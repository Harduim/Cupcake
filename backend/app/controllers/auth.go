package controllers

import (
	"cupcake/app/service"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type AuthResponse struct {
	Token     string `json:"access_token"`
	TokenType string `json:"token_type"`
}

func AuthenticateSSO(sso *service.SSOClient) fiber.Handler {

	return func(ctx *fiber.Ctx) error {

		url, err := sso.AuthCodeURL(ctx.Context())

		if err != nil {
			return err
		}

		return ctx.Redirect(url, 302)
	}
}

func Token(sso *service.SSOClient) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		code := ctx.Query("code")
		result, err := sso.App.AcquireTokenByAuthCode(ctx.Context(), code, sso.RedirectUrl, sso.Scopes)
		if err != nil {
			err := ctx.Status(fiber.StatusUnauthorized).SendString(err.Error())
			if err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}

		response := AuthResponse{
			Token:     result.AccessToken,
			TokenType: "bearer",
		}

		u, err := json.Marshal(response)

		if err != nil {
			return err
		}

		return ctx.Status(fiber.StatusOK).SendString(string(u))
	}
}
