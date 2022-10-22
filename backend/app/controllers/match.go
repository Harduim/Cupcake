package controllers

import (
	"cupcake/app/database"
	"cupcake/app/repositories"
	"github.com/gofiber/fiber/v2"
)

// GetAllMatches Return all matches as JSON
func GetAllMatches(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		repo := repositories.MatchRepositoryDb{Db: db}
		brackets, err := repo.FindAll()

		if err != nil {
			panic("Error occurred while retrieving matches from the database: " + err.Error())
		}

		response := ctx.JSON(brackets)

		if err != nil {
			panic("Error occurred when returning JSON of matches: " + err.Error())
		}

		return response
	}
}
