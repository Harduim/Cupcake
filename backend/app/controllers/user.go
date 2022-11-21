package controllers

import (
	"cupcake/app/database"
	models "cupcake/app/models"
	"cupcake/app/repositories"

	"github.com/gofiber/fiber/v2"
)

// GetAllUsers Return all users as JSON
func GetAllUsers(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var Users []models.User
		if response := db.Find(&Users); response.Error != nil {
			panic("Error occurred while retrieving users from the database: " + response.Error.Error())
		}
		err := ctx.JSON(Users)
		if err != nil {
			panic("Error occurred when returning JSON of users: " + err.Error())
		}
		return err
	}
}

func GetMe(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		User := new(models.User)
		id := ctx.Locals("user_id").(string)
		User.ID = id

		repo := repositories.UserRepositoryDb{Db: db}

		user, err := repo.Find(id)
		if err == nil {
			ctx.JSON(user)
		}
		return err
	}
}
