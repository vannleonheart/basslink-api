package admin

import (
	"CRM/src/lib/basslink"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	App *basslink.App
}

func New(app *basslink.App) *Service {
	return &Service{
		App: app,
	}
}

func (s *Service) handleSession(c *fiber.Ctx) error {
	headers := strings.TrimSpace(c.Get("authorization", ""))

	if len(headers) > 0 {
		headerValue := strings.Split(headers, "Bearer ")

		if len(headerValue) == 2 {
			token := strings.TrimSpace(headerValue[1])

			if len(token) > 0 {
				claims := jwt.MapClaims{}

				jwtToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
					if jwt.GetSigningMethod("HS256") != token.Method {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
					}

					return []byte(s.App.Config.JwtKey), nil
				})

				if err == nil && jwtToken != nil {
					if iAs, asIsExist := claims["as"]; asIsExist {
						as := iAs.(string)

						if as == "admin" {
							if exp, exists := claims["exp"]; exists {
								if expiredAt := int64(exp.(float64)); expiredAt >= time.Now().Unix() {
									if adminId, ok := claims["user"]; ok {
										var admin basslink.Administrator

										if err = s.App.DB.Connection.Where("id = ?", adminId).First(&admin).Error; err == nil {
											if admin.IsEnable {
												c.Locals("admin", &admin)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return c.Next()
}

func (s *Service) shouldBeUser(roles []string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ctxAdmin := c.Locals("admin")

		if ctxAdmin != nil {
			admin := ctxAdmin.(*basslink.Administrator)

			if admin != nil {
				if len(roles) > 0 {
					allow := false

					for _, role := range roles {
						if admin.Role == role {
							allow = true
						}
					}

					if !allow {
						return errors.New("permission denied")
					}
				}

				return c.Next()
			}
		}

		return errors.New("authentication required")
	}
}

func (s *Service) shouldBeGuest(c *fiber.Ctx) error {
	if ctxAdmin := c.Locals("admin"); ctxAdmin == nil {
		return c.Next()
	}

	return errors.New("permission denied")
}
