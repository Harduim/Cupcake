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
	registerMatches(api, db)
	registerNationalTeam(api, db)
	registerUserPoints(api, db)
	registerBets(api, db)
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
	users.Get("/token", Controller.Token(sso))
}

func registerBrackets(api fiber.Router, db *database.Database) {
	brackets := api.Group("/brackets")
	brackets.Get("/", Controller.GetAllBrackets(db))
}

func registerMatches(api fiber.Router, db *database.Database) {
	matches := api.Group("/matches")
	matches.Get("/", Controller.GetAllMatches(db))
	matches.Post("/", Controller.CreateMatch(db))
	matches.Put("/", Controller.UpdateMatch(db))
	matches.Put("/", Controller.DeleteMatch(db))
}

func registerNationalTeam(api fiber.Router, db *database.Database) {
	nationalTeams := api.Group("/national_teams")
	nationalTeams.Get("/", Controller.GetAllNationalTeams(db))
}

func registerUserPoints(api fiber.Router, db *database.Database) {
	userPoints := api.Group("/points")
	userPoints.Get("/", Controller.GetAllUserPoints(db))
}

func registerBets(api fiber.Router, db *database.Database) {
	userPoints := api.Group("/bets")
	userPoints.Get("/", Controller.GetAllBets(db))
	userPoints.Post("/", Controller.CreateBet(db))
}
