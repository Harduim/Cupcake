package controllers

import (
	"cupcake/app/service"
	"github.com/gofiber/fiber/v2"
)

func AuthenticateSSO(sso *service.SSOClient) fiber.Handler {

	return func(ctx *fiber.Ctx) error {

		url, err := sso.AuthCodeURL(ctx.Context())

		if err != nil {
			return err
		}

		return ctx.Redirect(url, 302)
	}
}
