package middlewares

import (
	"fees-ibd/internal/entity"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func RegisterCheckAuthToken(app *fiber.App, jwtSecret []byte) {
	app.Use(
		jwtware.New(jwtware.Config{
			SigningKey: jwtSecret,
			SuccessHandler: func(c *fiber.Ctx) error {
				tokenString := extractToken(c)

				token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}

					return jwtSecret, nil
				})

				if err != nil {
					return c.Status(fiber.StatusUnauthorized).JSON(
						entity.ErrorResp{
							Status:  false,
							Message: err.Error(),
						},
					)
				}

				claims, ok := token.Claims.(jwt.MapClaims)
				if !ok || !token.Valid {
					return c.Status(fiber.StatusUnauthorized).JSON(
						entity.ErrorResp{
							Status:  false,
							Message: "invalid token",
						},
					)
				}

				_, hasUserID := claims["user_id"]
				_, hasLogin := claims["login"]
				_, hasOgdCode := claims["ogd_code"]
				if !hasUserID || !hasLogin || !hasOgdCode {
					return c.Status(fiber.StatusUnauthorized).JSON(
						entity.ErrorResp{
							Status:  false,
							Message: "invalid token",
						},
					)
				}

				userID, validUserID := claims["user_id"].(float64) // почему то парсит не как float
				login, validLogin := claims["login"].(string)
				ogdCode, validOgdCode := claims["ogd_code"].(string)
				if !validUserID || !validLogin || !validOgdCode {
					return c.Status(fiber.StatusUnauthorized).JSON(
						entity.ErrorResp{
							Status:  false,
							Message: "invalid token",
						},
					)
				}

				user := entity.User{
					ID:      int(userID),
					Login:   login,
					OgdCode: ogdCode,
				}

				// include in fiber local
				c.Locals("user", user)

				return c.Next()
			},
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				code := fiber.StatusUnauthorized
				return c.Status(code).JSON(
					entity.ErrorResp{
						Status:  false,
						Message: err.Error(),
					},
				)
			},
		}),
	)
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}
