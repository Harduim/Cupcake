package controllers

import (
	"cupcake/app/database"
	"cupcake/app/domain"
	"cupcake/app/repositories"

	"github.com/gofiber/fiber/v2"
)

// GetAllUsers Return all users as JSON
func GetAllUsers(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var Users []domain.User
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

// GetUser Return a single user as JSON
func GetUser(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		User := new(domain.User)
		id := ctx.Params("id")
		User.ID = id

		repo := repositories.UserRepositoryDb{Db: db}

		if _, err := repo.Find(id); err != nil {
			err := ctx.Status(fiber.StatusNotFound).SendString(err.Error())
			return err
		}

		if User.ID == "" {
			err := ctx.SendStatus(fiber.StatusNotFound)
			if err != nil {
				panic("Cannot return status not found: " + err.Error())
			}
			err = ctx.JSON(fiber.Map{
				"ID": id,
			})
			if err != nil {
				panic("Error occurred when returning JSON of a role: " + err.Error())
			}
			return err
		}
		err := ctx.JSON(User)
		if err != nil {
			panic("Error occurred when returning JSON of a user: " + err.Error())
		}
		return err
	}
}

// AddUser Add a single user to the database
func AddUser(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		User := new(domain.User)
		if err := ctx.BodyParser(User); err != nil {
			panic("An error occurred when parsing the new user: " + err.Error())
		}
		if response := db.Create(&User); response.Error != nil {
			panic("An error occurred when storing the new user: " + response.Error.Error())
		}
		err := ctx.JSON(User)
		if err != nil {
			panic("Error occurred when returning JSON of a user: " + err.Error())
		}
		return err
	}
}

// EditUser Edit a single user
func EditUser(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		EditUser := new(domain.User)
		User := new(domain.User)
		if err := ctx.BodyParser(EditUser); err != nil {
			panic("An error occurred when parsing the edited user: " + err.Error())
		}
		if response := db.Find(&User, id); response.Error != nil {
			panic("An error occurred when retrieving the existing user: " + response.Error.Error())
		}
		// User does not exist
		if User.ID == "" {
			err := ctx.SendStatus(fiber.StatusNotFound)
			if err != nil {
				panic("Cannot return status not found: " + err.Error())
			}
			err = ctx.JSON(fiber.Map{
				"ID": id,
			})
			if err != nil {
				panic("Error occurred when returning JSON of a user: " + err.Error())
			}
			return err
		}
		User.Name = EditUser.Name
		User.Email = EditUser.Email

		// Save user
		db.Save(&User)

		err := ctx.JSON(User)
		if err != nil {
			panic("Error occurred when returning JSON of a user: " + err.Error())
		}
		return err
	}
}

// DeleteUser Delete a single user
func DeleteUser(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		var User domain.User
		db.Find(&User, id)
		if response := db.Find(&User); response.Error != nil {
			panic("An error occurred when finding the user to be deleted" + response.Error.Error())
		}
		db.Delete(&User)

		err := ctx.JSON(fiber.Map{
			"ID":      id,
			"Deleted": true,
		})
		if err != nil {
			panic("Error occurred when returning JSON of a user: " + err.Error())
		}
		return err
	}
}

func GetMe(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		User := new(domain.User)
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
