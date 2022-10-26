package controllers

import (
	"cupcake/app/database"
	"cupcake/app/repositories"
	"github.com/gofiber/fiber/v2"
)

// GetAllBrackets Return all brackets as JSON
func GetAllBrackets(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		repo := repositories.BracketRepositoryDb{Db: db}
		brackets, err := repo.FindAll()

		if err != nil {
			panic("Error occurred while retrieving brackets from the database: " + err.Error())
		}

		response := ctx.JSON(brackets)

		if err != nil {
			panic("Error occurred when returning JSON of brackets: " + err.Error())
		}

		return response
	}
}
