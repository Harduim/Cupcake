package controllers

import (
	"cupcake/app/database"
	"cupcake/app/domain"
	"github.com/gofiber/fiber/v2"
)

// GetAllNationalTeamsBrackets Return all national teams brackets as JSON
func GetAllNationalTeamsBrackets(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var NationalTeamsBrackets []domain.NationalTeamBracket
		if response := db.Find(&NationalTeamsBrackets); response.Error != nil {
			panic("Error occurred while retrieving national teams brackets from the database: " + response.Error.Error())
		}
		err := ctx.JSON(NationalTeamsBrackets)
		if err != nil {
			panic("Error occurred when returning JSON of National Teams Brackets: " + err.Error())
		}
		return err
	}
}
