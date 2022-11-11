package controllers

import (
	"cupcake/app/database"
	"cupcake/app/repositories"
	"github.com/gofiber/fiber/v2"
)

// GetAllUserPoints Return all user points as JSON
func GetAllUserPoints(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		repo := repositories.UserPointsRepositoryDb{Db: db}
		brackets, err := repo.FindAll()

		if err != nil {
			panic("Error occurred while retrieving user points from the database: " + err.Error())
		}

		response := ctx.JSON(brackets)

		if err != nil {
			panic("Error occurred when returning JSON of user points: " + err.Error())
		}

		return response
	}
}

// UpdateUserPoints Run procedure to update user points
func UpdateUserPoints(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		repo := repositories.UserPointsRepositoryDb{Db: db}
		err := repo.RunScoreUpdate()

		if err != nil {
			panic("Error occurred while updating user points from the database: " + err.Error())
		}

		response := ctx.Status(fiber.StatusAccepted).SendString("Running table update")

		return response
	}
}
