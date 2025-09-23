package agent

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

						if as == "agent" {
							if exp, exists := claims["exp"]; exists {
								if expiredAt := int64(exp.(float64)); expiredAt >= time.Now().Unix() {
									if agentUserId, ok := claims["user"]; ok {
										var agentUser basslink.AgentUser

										if err = s.App.DB.Connection.Preload("Agent").Where("id = ?", agentUserId).First(&agentUser).Error; err == nil {
											if agentUser.Agent != nil && agentUser.IsEnable {
												c.Locals("agent", &agentUser)
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
		ctxAgent := c.Locals("agent")

		if ctxAgent != nil {
			agent := ctxAgent.(*basslink.AgentUser)

			if agent != nil {
				if len(roles) > 0 {
					allow := false

					for _, role := range roles {
						if agent.Role == role {
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
	if ctxAgent := c.Locals("agent"); ctxAgent == nil {
		return c.Next()
	}

	return errors.New("permission denied")
}
