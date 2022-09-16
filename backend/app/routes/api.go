package routes

import (
	Controller "cupcake/app/controllers"
	"cupcake/app/database"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router, db *database.Database) {
	registerUsers(api, db)
}

func registerUsers(api fiber.Router, db *database.Database) {
	users := api.Group("/users")

	users.Get("/", Controller.GetAllUsers(db))
	users.Get("/:id", Controller.GetUser(db))
	users.Post("/", Controller.AddUser(db))
	users.Put("/:id", Controller.EditUser(db))
	users.Delete("/:id", Controller.DeleteUser(db))
}
