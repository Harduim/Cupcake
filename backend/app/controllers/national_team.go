package controllers

import (
	"cupcake/app/database"
	"cupcake/app/domain"
	"github.com/gofiber/fiber/v2"
)

// GetAllNationalTeams Return all national teams as JSON
func GetAllNationalTeams(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var NationalTeams []domain.NationalTeam
		if response := db.Find(&NationalTeams); response.Error != nil {
			panic("Error occurred while retrieving national teams from the database: " + response.Error.Error())
		}
		err := ctx.JSON(NationalTeams)
		if err != nil {
			panic("Error occurred when returning JSON of national teams: " + err.Error())
		}
		return err
	}
}
