package controllers

import (
	"cupcake/app/database"
	"cupcake/app/models"
	"cupcake/app/repositories"

	"github.com/gofiber/fiber/v2"
)

func Get(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user_id := ctx.Locals("user_id").(string)
		var joker models.Joker

		db.Find(&joker, "user_id = ?", user_id)

		if joker.ID != "" {
			response := ctx.JSON(joker)
			return response

		}
		newJoker := &models.Joker{UserID: user_id}
		newJoker.Prepare()
		db.Create(&newJoker)

		response := ctx.JSON(newJoker)
		return response

	}
}

func UpdateJoker(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		joker := new(models.Joker)
		user_id := ctx.Locals("user_id").(string)
		err := ctx.BodyParser(joker)
		if err != nil {
			panic("Unable to parse body: " + err.Error())
		}

		repo := repositories.JokerRepositoryDb{Db: db}
		joker.UserID = user_id
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
