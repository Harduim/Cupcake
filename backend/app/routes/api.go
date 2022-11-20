package routes

import (
	"cupcake/app/config"
	"cupcake/app/controllers"
	"cupcake/app/database"
	"cupcake/app/middleware"
	"cupcake/app/service"

	_ "cupcake/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func RegisterRoutes(api fiber.Router, db *database.Database, sso *service.SSOClient, config *config.Config) {
	tokenSecret := config.GetString("TOKEN_SECRET")
	authorizationMiddleware := middleware.Authorization(tokenSecret)

	registerUsers(api, db, authorizationMiddleware)
	registerAuth(api, db, sso, tokenSecret)
	registerBrackets(api, db, authorizationMiddleware)
	registerMatches(api, db, authorizationMiddleware)
	registerNationalTeam(api, db, authorizationMiddleware)
	registerBets(api, db, authorizationMiddleware)
	registerJoker(api, db, authorizationMiddleware)
	registerDocs(api, db)
}

func registerUsers(api fiber.Router, db *database.Database, authorization func(c *fiber.Ctx) (err error)) {
	users := api.Group("/users", authorization)

	users.Get("/", controllers.GetAllUsers(db))
	users.Get("/me", controllers.GetMe(db))
}

func registerAuth(api fiber.Router, db *database.Database, sso *service.SSOClient, secretKey string) {
	users := api.Group("/auth")
	users.Get("/sso", controllers.AuthenticateSSO(sso))
	users.Get("/token", controllers.Token(sso, db, secretKey))
}

func registerBrackets(api fiber.Router, db *database.Database, authorization func(c *fiber.Ctx) (err error)) {
	brackets := api.Group("/brackets", authorization)
	brackets.Get("/", controllers.GetAllBrackets(db))
}

func registerMatches(api fiber.Router, db *database.Database, authorization func(c *fiber.Ctx) (err error)) {
	matches := api.Group("/matches", authorization)
	matches.Get("/", controllers.GetAllMatches(db))
	matches.Post("/", controllers.CreateMatch(db))
	matches.Put("/", controllers.UpdateMatch(db))
}

func registerNationalTeam(api fiber.Router, db *database.Database, authorization func(c *fiber.Ctx) (err error)) {
	nationalTeams := api.Group("/national-teams", authorization)
	nationalTeams.Get("/", controllers.GetAllNationalTeams(db))
}

func registerBets(api fiber.Router, db *database.Database, authorization func(c *fiber.Ctx) (err error)) {
	userPoints := api.Group("/bets", authorization)
	userPoints.Get("/", controllers.GetAllBets(db))
	userPoints.Post("/", controllers.CreateBet(db))
	userPoints.Put("/", controllers.UpdateBet(db))
}

func registerJoker(api fiber.Router, db *database.Database, authorization func(c *fiber.Ctx) (err error)) {
	userPoints := api.Group("/joker", authorization)
	userPoints.Get("/", controllers.Get(db))
	userPoints.Put("/", controllers.UpdateJoker(db))
}

func registerDocs(api fiber.Router, db *database.Database) {
	docs := api.Group("/docs")
	docs.Get("/*", swagger.HandlerDefault)
}
