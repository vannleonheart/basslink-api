package basslink

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const (
	TokenBindingHeader = "header"
	TokenBindingBody   = "body"
	TokenBindingQuery  = "query"
)

var (
	tokenBindings        = []string{TokenBindingHeader, TokenBindingQuery, TokenBindingBody}
	defaultTokenBindings = []string{TokenBindingHeader}
	defaultCorsOrigins   = "*"
	defaultCorsMethods   = "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS"
	defaultCorsMaxAge    = 0
)

type HttpServer struct {
	Config    HttpConfig
	Handler   *fiber.App
	serviceId *string
}

type HttpConfig struct {
	Host           string                    `json:"host"`
	Port           string                    `json:"port"`
	ReadTimeout    time.Duration             `json:"read_timeout"`
	WriteTimeout   time.Duration             `json:"write_timeout"`
	Authentication *HttpAuthenticationConfig `json:"authentication,omitempty"`
	Cors           *HttpCorsConfig           `json:"cors,omitempty"`
}

type HttpAuthenticationConfig struct {
	TokenBindings *[]string `json:"token_bindings,omitempty"`
	TokenName     string    `json:"token_name"`
	TokenValue    string    `json:"token_value"`
	TokenPrefix   *string   `json:"token_prefix,omitempty"`
	Excludes      *[]string `json:"excludes,omitempty"`
	Includes      *[]string `json:"includes,omitempty"`
}

type HttpCorsConfig struct {
	Origins *string `json:"origins,omitempty"`
	Methods *string `json:"methods,omitempty"`
	MaxAge  *int    `json:"max_age,omitempty"`
}

func NewHttpServer(cfg HttpConfig, serviceId string) *HttpServer {
	handler := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			var cErr *AppError
			if errors.As(err, &cErr) {
				return NewErrorResponse(ctx, cErr)
			}

			code := ErrInternalHost
			kind := ErrInternalHostUnknown
			message := err.Error()

			var e *fiber.Error
			if errors.As(err, &e) {
				code = fmt.Sprintf("%d", e.Code)
				if e.Code == fiber.StatusNotFound {
					kind = ErrNotFoundRouteNotExist
					message = "ERR_NOT_FOUND"
				} else if e.Code == fiber.StatusMethodNotAllowed {
					message = "ERR_METHOD_NOT_ALLOWED"
				}
			}

			return NewErrorResponse(ctx, NewError(serviceId, kind, code, message, "", nil))
		},
	})

	if cfg.Authentication != nil {
		useAuthenticationMiddleware(handler, cfg.Authentication)
	}

	if cfg.Cors != nil {
		useCorsMiddleware(handler, cfg.Cors)
	}

	return &HttpServer{
		Config:    cfg,
		Handler:   handler,
		serviceId: &serviceId,
	}
}

func (s *HttpServer) Start() error {
	if len(s.Config.Host) < 1 {
		return errors.New("http server host is not set")
	}

	if len(s.Config.Port) < 1 {
		return errors.New("http server port is not set")
	}

	addr := fmt.Sprintf("%s:%s", s.Config.Host, s.Config.Port)

	return s.Handler.Listen(addr)
}

func (s *HttpServer) Stop() error {
	if s.Handler != nil {
		if err := s.Handler.Shutdown(); err != nil {
			return err
		}
	}

	return nil
}

func useCorsMiddleware(ctx *fiber.App, c *HttpCorsConfig) {
	origins := defaultCorsOrigins
	if c.Origins != nil {
		origins = *c.Origins
	}

	methods := defaultCorsMethods
	if c.Methods != nil {
		methods = *c.Methods
	}

	maxAge := defaultCorsMaxAge
	if c.MaxAge != nil {
		maxAge = *c.MaxAge
	}

	ctx.Use(cors.New(cors.Config{
		AllowOrigins: origins,
		AllowMethods: methods,
		MaxAge:       maxAge,
	}))
}

func useAuthenticationMiddleware(ctx *fiber.App, c *HttpAuthenticationConfig) {
	if len(c.TokenName) < 1 || len(c.TokenValue) < 1 {
		return
	}

	tokenName := strings.ToLower(c.TokenName)

	tokenValue := c.TokenValue
	if c.TokenPrefix != nil && len(*c.TokenPrefix) > 0 {
		tokenValue = fmt.Sprintf("%s%s", *c.TokenPrefix, tokenValue)
	}

	currentTokenBindings := defaultTokenBindings
	if c.TokenBindings != nil && len(*c.TokenBindings) > 0 {
		currentTokenBindings = *c.TokenBindings
	}

	if len(currentTokenBindings) < 1 {
		return
	}

	ctx.Use(func(ctx *fiber.Ctx) error {
		if (c.Excludes != nil && !isCurrentPathInExclusionList(ctx, c.Excludes)) || (c.Includes != nil && isCurrentPathInInclusionList(ctx, c.Includes)) {
			bindingFound := false
			bindingFoundValue := ""

			for _, key := range currentTokenBindings {
				key = strings.ToLower(key)
				if !slices.Contains(tokenBindings, key) {
					continue
				}

				bindingFound = true
				if found := findTokenInBinding(ctx, key, tokenName); len(found) > 0 {
					bindingFoundValue = found
					break
				}
			}

			if bindingFound && bindingFoundValue != tokenValue {
				return ctx.SendStatus(fiber.StatusUnauthorized)
			}
		}

		return ctx.Next()
	})
}

func isCurrentPathInExclusionList(ctx *fiber.Ctx, exclusionList *[]string) bool {
	if len(*exclusionList) > 0 {
		curPath := strings.TrimLeft(strings.ToLower(strings.TrimSpace(ctx.Path())), "/")
		for _, exValue := range *exclusionList {
			exValueString := strings.TrimLeft(strings.ToLower(strings.TrimSpace(exValue)), "/")
			if curPath == exValueString {
				return true
			}
		}
	}

	return false
}

func isCurrentPathInInclusionList(ctx *fiber.Ctx, inclusionList *[]string) bool {
	if len(*inclusionList) > 0 {
		curPath := strings.TrimLeft(strings.ToLower(strings.TrimSpace(ctx.Path())), "/")
		for _, exValue := range *inclusionList {
			exValueString := strings.TrimLeft(strings.ToLower(strings.TrimSpace(exValue)), "/")
			if curPath == exValueString {
				return true
			}
		}
	}

	return false
}

func findTokenInBinding(ctx *fiber.Ctx, key, tokenName string) string {
	switch key {
	default:
		return ""
	case TokenBindingHeader:
		return findTokenInBindingHeader(ctx, tokenName)
	case TokenBindingQuery:
		return findTokenInBindingQuery(ctx, tokenName)
	case TokenBindingBody:
		return findTokenInBindingBody(ctx, tokenName)
	}
}

func findTokenInBindingHeader(ctx *fiber.Ctx, tokenName string) string {
	headers := ctx.GetReqHeaders()
	for headerName, headerValue := range headers {
		headerNameStr := strings.ToLower(headerName)
		if headerNameStr == tokenName {
			return headerValue[0]
		}
	}

	return ""
}

func findTokenInBindingQuery(ctx *fiber.Ctx, tokenName string) string {
	return ctx.Query(tokenName, "")
}

func findTokenInBindingBody(ctx *fiber.Ctx, tokenName string) string {
	reqBody := new(map[string]interface{})
	if err := ctx.BodyParser(reqBody); err == nil {
		for bodyKey, bodyValue := range *reqBody {
			bodyKeyString := strings.ToLower(strings.TrimSpace(bodyKey))
			if bodyKeyString == tokenName {
				return bodyValue.(string)
			}
		}
	}

	return ""
}
