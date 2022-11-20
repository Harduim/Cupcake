package controllers

import (
	"cupcake/app/database"
	"cupcake/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetGroups(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user_id := ctx.Locals("user_id").(string)
		var group models.Group

		db.Find(&group, "user_id = ?", user_id)

		if group.UserID != "" {
			response := ctx.JSON(group)
			return response

		}
		newGroups := &models.Group{UserID: user_id}
		db.Create(&newGroups)

		response := ctx.JSON(newGroups)
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
		db.Delete(&models.Group{}, "user_id = ?", user_id)
		db.Save(&groups.Groups)
		response := ctx.JSON(groups)
		return response

	}
}
