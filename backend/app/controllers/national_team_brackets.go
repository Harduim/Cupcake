package controllers

import (
	"cupcake/app/database"
	"cupcake/app/repositories"
	"github.com/gofiber/fiber/v2"
)

// GetAllNationalTeamsBrackets Return all national teams brackets as JSON
func GetAllNationalTeamsBrackets(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		repo := repositories.NationalTeamBracketRepositoryDb{Db: db}
		brackets, err := repo.FindAll()

		if err != nil {
			panic("Error occurred while retrieving national teams brackets from the database: " + err.Error())
		}

		response := ctx.JSON(brackets)

		if err != nil {
			panic("Error occurred when returning JSON of national teams brackets: " + err.Error())
		}

		return response
	}
}
