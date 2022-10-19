package controllers

import (
	"cupcake/app/database"
	"cupcake/app/domain"
	"github.com/gofiber/fiber/v2"
)

// GetAllBrackets Return all brackets as JSON
func GetAllBrackets(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var Brackets []domain.Bracket
		if response := db.Find(&Brackets); response.Error != nil {
			panic("Error occurred while retrieving brackets from the database: " + response.Error.Error())
		}
		err := ctx.JSON(Brackets)
		if err != nil {
			panic("Error occurred when returning JSON of users: " + err.Error())
		}
		return err
	}
}
