package routes

import (
	Controller "cupcake/app/controllers"
	"cupcake/app/database"
	"cupcake/app/service"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router, db *database.Database, sso *service.SSOClient) {
	registerUsers(api, db)
	registerAuth(api, db, sso)
	registerBrackets(api, db)
}

func registerUsers(api fiber.Router, db *database.Database) {
	users := api.Group("/users")

	users.Get("/", Controller.GetAllUsers(db))
	users.Get("/:id", Controller.GetUser(db))
	users.Post("/", Controller.AddUser(db))
	users.Put("/:id", Controller.EditUser(db))
	users.Delete("/:id", Controller.DeleteUser(db))
}

func registerAuth(api fiber.Router, db *database.Database, sso *service.SSOClient) {
	users := api.Group("/auth")
	users.Get("/sso", Controller.AuthenticateSSO(sso))
}

func registerBrackets(api fiber.Router, db *database.Database) {
	brackets := api.Group("/brackets")
	brackets.Get("/", Controller.GetAllBrackets(db))
}
