package controllers

import (
	"cupcake/app/database"
	"cupcake/app/models"
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
		bet := new(models.Bet)
		bet.Prepare()
		id := ctx.Locals("user_id").(string)
		bet.UserID = id
		err := ctx.BodyParser(bet)
		if err != nil {
			panic("Unable to parse body: " + err.Error())
		}
		err = bet.Validate()
		if err != nil {
			err_msg := "Unable to validate body: " + err.Error() + "\n"
			return fiber.NewError(500, err_msg)
		}

		repo := repositories.BetRepositoryDb{Db: db}
		n_bet, err := repo.Insert(bet)

		if err != nil {
			panic("Error occurred while creating bet from the database: " + err.Error())
		}

		response := ctx.JSON(n_bet)

		if err != nil {
			panic("Error occurred when returning JSON of bets: " + err.Error())
		}

		return response
	}
}

// UpdateBet Update bet
func UpdateBet(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		bet := new(models.Bet)
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

// DeleteBet Delete bet
func DeleteBet(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		bet := new(models.Bet)

		err := ctx.BodyParser(bet)

		if err != nil {
			return err
		}

		repo := repositories.BetRepositoryDb{Db: db}
		data, err := repo.Delete(bet)

		if err != nil {
			panic("Error occurred while deleting bet from the database: " + err.Error())
		}

		response := ctx.JSON(data)

		if err != nil {
			panic("Error occurred when returning JSON of bet: " + err.Error())
		}

		return response
	}
}
