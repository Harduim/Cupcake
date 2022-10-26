package controllers

import (
	"cupcake/app/database"
	"cupcake/app/repositories"
	"github.com/gofiber/fiber/v2"
)

// GetAllBets Return all bets as JSON
func GetAllBets(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		repo := repositories.BetRepositoryDb{Db: db}
		brackets, err := repo.FindAll()

		if err != nil {
			panic("Error occurred while retrieving bets from the database: " + err.Error())
		}

		response := ctx.JSON(brackets)

		if err != nil {
			panic("Error occurred when returning JSON of bets: " + err.Error())
		}

		return response
	}
}
