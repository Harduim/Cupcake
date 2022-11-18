package controllers

import (
	"cupcake/app/database"
	domain "cupcake/app/models"
	"cupcake/app/repositories"

	"github.com/gofiber/fiber/v2"
)

// GetAllJokers Return all jokers as JSON
func GetAllJokers(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		repo := repositories.JokerRepositoryDb{Db: db}
		joker, err := repo.FindAll()

		if err != nil {
			panic("Error occurred while retrieving jokers from the database: " + err.Error())
		}

		response := ctx.JSON(joker)

		if err != nil {
			panic("Error occurred when returning JSON of jokers: " + err.Error())
		}

		return response
	}
}

// CreateJoker Create joker
func CreateJoker(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		joker := new(domain.Joker)
		err := ctx.BodyParser(joker)

		repo := repositories.JokerRepositoryDb{Db: db}
		joker, err = repo.Insert(joker)

		if err != nil {
			panic("Error occurred while creating joker from the database: " + err.Error())
		}

		response := ctx.JSON(joker)

		if err != nil {
			panic("Error occurred when returning JSON of jokers: " + err.Error())
		}

		return response
	}
}

// UpdateJoker Update joker
func UpdateJoker(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		joker := new(domain.Joker)
		err := ctx.BodyParser(joker)

		repo := repositories.JokerRepositoryDb{Db: db}
		joker, err = repo.Update(joker)

		if err != nil {
			panic("Error occurred while updating joker from the database: " + err.Error())
		}

		response := ctx.JSON(joker)

		if err != nil {
			panic("Error occurred when returning JSON of jokers: " + err.Error())
		}

		return response
	}
}

// DeleteJoker Delete joker
func DeleteJoker(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		joker := new(domain.Joker)

		err := ctx.BodyParser(joker)

		if err != nil {
			return err
		}

		repo := repositories.JokerRepositoryDb{Db: db}
		data, err := repo.Delete(joker)

		if err != nil {
			panic("Error occurred while deleting joker from the database: " + err.Error())
		}

		response := ctx.JSON(data)

		if err != nil {
			panic("Error occurred when returning JSON of joker: " + err.Error())
		}

		return response
	}
}
