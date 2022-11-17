package middlewares

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Authorization(tokenSecret string) func(c *fiber.Ctx) (err error) {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization")
		token = strings.Replace(token, "Bearer ", "", 1)

		secretFunc := func(token *jwt.Token) (interface{}, error) {
			return []byte(tokenSecret), nil
		}

		parsed, err := jwt.Parse(token, secretFunc)
		if err != nil {
			err_msg := "Failed to parse JWT.\nError: " + err.Error()
			log.Println(err_msg)
			err := ctx.Status(fiber.StatusUnauthorized).SendString(err_msg)
			return err
		}

		if !parsed.Valid {
			err_msg := "Token is not valid"
			log.Println(err_msg)
			err := ctx.Status(fiber.StatusUnauthorized).SendString(err_msg)
			return err
		}

		claims := parsed.Claims.(jwt.MapClaims)
		if claims == nil {
			err_msg := "Unable to find claims"
			log.Println(err_msg)
			err := ctx.Status(fiber.StatusUnauthorized).SendString(err_msg)
			return err
		}

		user_id := claims["user"]
		if user_id == nil {
			err_msg := "Unable to find user UUID"
			log.Println(err_msg)
			err := ctx.Status(fiber.StatusUnauthorized).SendString(err_msg)
			return err
		}
		ctx.Locals("user_id", user_id)

		is_admin := claims["admin"]
		if is_admin == nil {
			err_msg := "Unable to find user permissions"
			log.Println(err_msg)
			err := ctx.Status(fiber.StatusUnauthorized).SendString(err_msg)
			return err
		}
		ctx.Locals("is_admin", is_admin)

		return ctx.Next()
	}
}
