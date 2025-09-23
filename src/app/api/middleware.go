package main

import (
	"CRM/src/lib/basslink"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"slices"
	"strings"
	"time"
)

func agentSessionHandler(c *fiber.Ctx) error {
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

					return []byte(config.JwtKey), nil
				})

				if err == nil && jwtToken != nil {
					if userType, isUserTypeExist := claims["user_type"]; isUserTypeExist {
						if userType != nil && strings.ToLower(userType.(string)) == "agent" {
							if exp, isExpExist := claims["exp"]; isExpExist {
								if exp != nil {
									if expiredAt := int64(exp.(float64)); expiredAt >= time.Now().Unix() {
										if userId, isUserIdExist := claims["user"]; isUserIdExist {
											if userId != nil {
												var user basslink.AgentUser

												if err = dbcon.Connection.Preload("Agent").Where("id = ?", userId.(string)).First(&user).Error; err == nil {
													if user.IsEnabled {
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
		}
	}

	return c.Next()
}

func agentAuthHandler(c *fiber.Ctx) error {
	return c.Next()
}

func clientSessionHandler(c *fiber.Ctx) error {
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

					return []byte(config.JwtKey), nil
				})

				if err == nil && jwtToken != nil {
					if userType, isUserTypeExist := claims["user_type"]; isUserTypeExist {
						if userType != nil && strings.ToLower(userType.(string)) == "client" {
							if exp, isExpExist := claims["exp"]; isExpExist {
								if exp != nil {
									if expiredAt := int64(exp.(float64)); expiredAt >= time.Now().Unix() {
										if userId, isUserIdExist := claims["user"]; isUserIdExist {
											if userId != nil {
												var user basslink.ClientUser

												if dbtx := dbcon.Connection.Preload("Client").Where("id = ?", userId.(string)).First(&user); dbtx.Error == nil {
													if user.IsEnabled {
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
		}
	}

	return c.Next()
}

func clientAuthHandler(c *fiber.Ctx) error {
	return c.Next()
}

func shouldBeGuest(c *fiber.Ctx) error {
	if ctxUser := c.Locals("user"); ctxUser == nil {
		return c.Next()
	}

	return basslink.NewError(serviceId, basslink.ErrForbiddenNotPermitted, basslink.ErrForbidden, "ERR_NOT_PERMITTED", "", nil)
}

func shouldBeAgentUser(roles *[]string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ctxUser := c.Locals("user")

		if ctxUser != nil {
			user := ctxUser.(*basslink.AgentUser)
			if user != nil {
				isAllowed := true

				if roles != nil && len(*roles) > 0 {
					if !slices.Contains(*roles, user.Role) {
						isAllowed = false
					}
				}

				if isAllowed {
					return c.Next()
				}
			}
		}

		return basslink.NewError(serviceId, basslink.ErrForbiddenNotPermitted, basslink.ErrUnauthorizedNeedAuthentication, "ERR_AUTH_REQUIRED", "", nil)
	}
}

func shouldBeClientUser(roles *[]string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ctxUser := c.Locals("user")

		if ctxUser != nil {
			user := ctxUser.(*basslink.ClientUser)
			if user != nil {
				isAllowed := true

				if roles != nil && len(*roles) > 0 {
					if !slices.Contains(*roles, user.Role) {
						isAllowed = false
					}
				}

				if isAllowed {
					return c.Next()
				}
			}
		}

		return basslink.NewError(serviceId, basslink.ErrForbiddenNotPermitted, basslink.ErrUnauthorizedNeedAuthentication, "ERR_AUTH_REQUIRED", "", nil)
	}
}

func adminSessionHandler(c *fiber.Ctx) error {
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

					return []byte(config.JwtKey), nil
				})

				if err == nil && jwtToken != nil {
					if userType, isUserTypeExist := claims["user_type"]; isUserTypeExist {
						if userType != nil && strings.ToLower(userType.(string)) == "admin" {
							if exp, isExpExist := claims["exp"]; isExpExist {
								if exp != nil {
									if expiredAt := int64(exp.(float64)); expiredAt >= time.Now().Unix() {
										if userId, isUserIdExist := claims["user"]; isUserIdExist {
											if userId != nil {
												var user basslink.AdminUser

												if err = dbcon.Connection.Where("id = ?", userId.(string)).First(&user).Error; err == nil {
													if user.IsEnabled {
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
		}
	}

	return c.Next()
}

func shouldBeAdmin(roles *[]string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ctxUser := c.Locals("user")

		if ctxUser != nil {
			user := ctxUser.(*basslink.AdminUser)
			if user != nil {
				isAllowed := true

				if roles != nil && len(*roles) > 0 {
					if !slices.Contains(*roles, user.Role) {
						isAllowed = false
					}
				}

				if isAllowed {
					return c.Next()
				}
			}
		}

		return basslink.NewError(serviceId, basslink.ErrForbiddenNotPermitted, basslink.ErrUnauthorizedNeedAuthentication, "ERR_AUTH_REQUIRED", "", nil)
	}
}

func adminAuthHandler(c *fiber.Ctx) error {
	return c.Next()
}
