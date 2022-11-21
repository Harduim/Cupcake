package controllers

import (
	"cupcake/app/config"
	"cupcake/app/database"
	"cupcake/app/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetGroups(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user_id := ctx.Locals("user_id").(string)
		var group []models.Group

		err := db.Model(&models.Group{}).Find(&group, "user_id = ?", user_id).Error

		if err != nil {
			return ctx.SendStatus(fiber.ErrBadGateway.Code)
		}

		response := ctx.JSON(group)
		return response

	}
}

func UpdateGroup(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user_id := ctx.Locals("user_id").(string)
		groups := new(models.Groups)

		err := ctx.BodyParser(groups)
		if err != nil {
			return ctx.SendStatus(fiber.ErrUnprocessableEntity.Code)
		}
		for _, g := range groups.Groups {
			if g.UserID != user_id {
				return ctx.SendStatus(fiber.ErrUnauthorized.Code)
			}
		}

		datetime_now := time.Now().UTC()

		if datetime_now.After(config.FASE_GRUPOS_CLOSE_DATE) {
			return ctx.SendStatus(fiber.ErrConflict.Code)
		}

		db.Delete(&models.Group{}, "user_id = ?", user_id)
		db.Save(&groups.Groups)
		response := ctx.JSON(groups)
		return response

	}
}
