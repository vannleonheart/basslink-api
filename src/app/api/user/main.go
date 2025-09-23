package user

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

						if as == "user" {
							if exp, exists := claims["exp"]; exists {
								if expiredAt := int64(exp.(float64)); expiredAt >= time.Now().Unix() {
									if userId, ok := claims["user"]; ok {
										var user basslink.User

										if err = s.App.DB.Connection.Where("id = ?", userId).First(&user).Error; err == nil {
											if user.IsEnable {
												c.Locals("user", &user)
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

func (s *Service) shouldBeUser(c *fiber.Ctx) error {
	ctxUser := c.Locals("user")

	if ctxUser != nil {
		user := ctxUser.(*basslink.User)

		if user != nil {
			return c.Next()
		}
	}

	return errors.New("authentication required")
}

func (s *Service) shouldBeGuest(c *fiber.Ctx) error {
	if ctxUser := c.Locals("user"); ctxUser == nil {
		return c.Next()
	}

	return errors.New("permission denied")
}
