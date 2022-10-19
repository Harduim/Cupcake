package controllers

import (
	"cupcake/app/database"
	"cupcake/app/domain"
	"github.com/gofiber/fiber/v2"
)

// GetAllMatches Return all matches as JSON
func GetAllMatches(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var Matches []domain.Match
		if response := db.Find(&Matches); response.Error != nil {
			panic("Error occurred while retrieving matches from the database: " + response.Error.Error())
		}
		err := ctx.JSON(Matches)
		if err != nil {
			panic("Error occurred when returning JSON of matches: " + err.Error())
		}
		return err
	}
}
