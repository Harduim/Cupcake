package controllers

import (
	"cupcake/app/database"
	"cupcake/app/domain"
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

// CreateBet Create bet
func CreateBet(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		bet := new(domain.Bet)
		err := ctx.BodyParser(bet)

		repo := repositories.BetRepositoryDb{Db: db}
		brackets, err := repo.Insert(bet)

		if err != nil {
			panic("Error occurred while creating bet from the database: " + err.Error())
		}

		response := ctx.JSON(brackets)

		if err != nil {
			panic("Error occurred when returning JSON of bets: " + err.Error())
		}

		return response
	}
}

// UpdateBet Update bet
func UpdateBet(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		bet := new(domain.Bet)
		err := ctx.BodyParser(bet)

		repo := repositories.BetRepositoryDb{Db: db}
		brackets, err := repo.Update(bet)

		if err != nil {
			panic("Error occurred while updating bet from the database: " + err.Error())
		}

		response := ctx.JSON(brackets)

		if err != nil {
			panic("Error occurred when returning JSON of bets: " + err.Error())
		}

		return response
	}
}
