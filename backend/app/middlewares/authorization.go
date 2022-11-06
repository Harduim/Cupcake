package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"log"
)

func Authorization(tokenSecret string) func(c *fiber.Ctx) (err error) {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("access_token")

		secretFunc := func(token *jwt.Token) (interface{}, error) {
			return []byte(tokenSecret), nil
		}

		parsed, err := jwt.Parse(token, secretFunc)
		if err != nil {
			log.Printf("Failed to parse JWT.\nError: %s\n", err.Error())
			err := ctx.Status(fiber.StatusUnauthorized).SendString(err.Error())
			return err
		}

		if !parsed.Valid {
			log.Println("Token is not valid.")
			err := ctx.Status(fiber.StatusUnauthorized).SendString(err.Error())
			return err
		}

		log.Println("Token is valid.")
		return ctx.Next()
	}
}
