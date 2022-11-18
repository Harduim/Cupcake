package controllers

import (
	"cupcake/app/database"
	domain "cupcake/app/models"
	"cupcake/app/repositories"

	"github.com/gofiber/fiber/v2"
)

// GetAllMatches Return all matches as JSON
func GetAllMatches(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		repo := repositories.MatchRepositoryDb{Db: db}
		matches, err := repo.FindAll()

		if err != nil {
			panic("Error occurred while retrieving matches from the database: " + err.Error())
		}

		response := ctx.JSON(matches)

		if err != nil {
			panic("Error occurred when returning JSON of matches: " + err.Error())
		}

		return response
	}
}

// CreateMatch Create match from body
func CreateMatch(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		match := new(domain.Match)

		err := ctx.BodyParser(match)

		if err != nil {
			return err
		}

		repo := repositories.MatchRepositoryDb{Db: db}
		match, err = repo.Insert(match)

		if err != nil {
			panic("Error occurred while creating match from the database: " + err.Error())
		}

		response := ctx.JSON(match)

		if err != nil {
			panic("Error occurred when returning JSON of match: " + err.Error())
		}

		return response
	}
}

// UpdateMatch Update match from body
func UpdateMatch(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		match := new(domain.Match)

		err := ctx.BodyParser(match)

		if err != nil {
			return err
		}

		repo := repositories.MatchRepositoryDb{Db: db}
		data, err := repo.Update(match)

		if err != nil {
			panic("Error occurred while updating match from the database: " + err.Error())
		}

		response := ctx.JSON(data)

		if err != nil {
			panic("Error occurred when returning JSON of match: " + err.Error())
		}

		return response
	}
}

// DeleteMatch Delete match from body
func DeleteMatch(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		match := new(domain.Match)

		err := ctx.BodyParser(match)

		if err != nil {
			return err
		}

		repo := repositories.MatchRepositoryDb{Db: db}
		data, err := repo.Delete(match)

		if err != nil {
			panic("Error occurred while deleting match from the database: " + err.Error())
		}

		response := ctx.JSON(data)

		if err != nil {
			panic("Error occurred when returning JSON of match: " + err.Error())
		}

		return response
	}
}
