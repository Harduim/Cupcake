package controllers

import (
	"cupcake/app/database"
	"cupcake/app/domain"
	"github.com/gofiber/fiber/v2"
)

// GetAllUserPoints Return all user points as JSON
func GetAllUserPoints(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var UserPoints []domain.UserPoints
		if response := db.Find(&UserPoints); response.Error != nil {
			panic("Error occurred while retrieving user points from the database: " + response.Error.Error())
		}
		err := ctx.JSON(UserPoints)
		if err != nil {
			panic("Error occurred when returning JSON of users points: " + err.Error())
		}
		return err
	}
}
