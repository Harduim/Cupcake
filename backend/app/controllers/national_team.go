package controllers

import (
	"cupcake/app/database"
	"cupcake/app/repositories"
	"github.com/gofiber/fiber/v2"
)

// GetAllNationalTeams Return all national teams as JSON
func GetAllNationalTeams(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		repo := repositories.NationalTeamRepositoryDb{Db: db}
		brackets, err := repo.FindAll()

		if err != nil {
			panic("Error occurred while retrieving national teams from the database: " + err.Error())
		}

		response := ctx.JSON(brackets)

		if err != nil {
			panic("Error occurred when returning JSON of national teams: " + err.Error())
		}

		return response
	}
}
