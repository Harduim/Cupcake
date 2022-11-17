package routes

import (
	"cupcake/app/config"
	Controller "cupcake/app/controllers"
	"cupcake/app/database"
	"cupcake/app/middlewares"
	"cupcake/app/service"

	_ "cupcake/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func RegisterRoutes(api fiber.Router, db *database.Database, sso *service.SSOClient, config *config.Config) {
	tokenSecret := config.GetString("TOKEN_SECRET")
	authorizationMiddleware := middlewares.Authorization(tokenSecret)

	registerUsers(api, db, authorizationMiddleware)
	registerAuth(api, db, sso, tokenSecret)
	registerBrackets(api, db, authorizationMiddleware)
	registerMatches(api, db, authorizationMiddleware)
	registerNationalTeam(api, db, authorizationMiddleware)
	registerUserPoints(api, db, authorizationMiddleware)
	registerBets(api, db, authorizationMiddleware)
	registerJoker(api, db, authorizationMiddleware)
	registerDocs(api, db)
}

func registerUsers(api fiber.Router, db *database.Database, authorization func(c *fiber.Ctx) (err error)) {
	users := api.Group("/users", authorization)

	users.Get("/", Controller.GetAllUsers(db))
	users.Get("/me", Controller.GetMe(db))
	// users.Get("/:id", Controller.GetUser(db))
	// users.Post("/", Controller.AddUser(db))
	// users.Put("/:id", Controller.EditUser(db))
	// users.Delete("/:id", Controller.DeleteUser(db))
}

func registerAuth(api fiber.Router, db *database.Database, sso *service.SSOClient, secretKey string) {
	users := api.Group("/auth")
	users.Get("/sso", Controller.AuthenticateSSO(sso))
	users.Get("/token", Controller.Token(sso, db, secretKey))
}

func registerBrackets(api fiber.Router, db *database.Database, authorization func(c *fiber.Ctx) (err error)) {
	brackets := api.Group("/brackets", authorization)
	brackets.Get("/", Controller.GetAllBrackets(db))
}

func registerMatches(api fiber.Router, db *database.Database, authorization func(c *fiber.Ctx) (err error)) {
	matches := api.Group("/matches", authorization)
	//	@Summary      Get ALL Matches
	//	@Description  Get Matches
	//	@ID           get-all-matches
	//	@Produce      json
	//	@Success      200      {string}  string        "ok"
	//	@Failure      400      {object}  web.APIError  "We need ID!!"
	//	@Failure      404      {object}  web.APIError  "Can not find ID"
	//	@Router       /api/matches [get]
	matches.Get("/", Controller.GetAllMatches(db))
	matches.Post("/", Controller.CreateMatch(db))
	matches.Put("/", Controller.UpdateMatch(db))
	matches.Put("/", Controller.DeleteMatch(db))
}

func registerNationalTeam(api fiber.Router, db *database.Database, authorization func(c *fiber.Ctx) (err error)) {
	nationalTeams := api.Group("/national-teams", authorization)
	nationalTeams.Get("/", Controller.GetAllNationalTeams(db))
}

func registerUserPoints(api fiber.Router, db *database.Database, authorization func(c *fiber.Ctx) (err error)) {
	userPoints := api.Group("/points", authorization)
	userPoints.Get("/", Controller.GetAllUserPoints(db))
	userPoints.Patch("/", Controller.UpdateUserPoints(db))
}

func registerBets(api fiber.Router, db *database.Database, authorization func(c *fiber.Ctx) (err error)) {
	userPoints := api.Group("/bets", authorization)
	userPoints.Get("/", Controller.GetAllBets(db))
	userPoints.Post("/", Controller.CreateBet(db))
	userPoints.Put("/", Controller.UpdateBet(db))
	userPoints.Put("/", Controller.DeleteBet(db))
}

func registerJoker(api fiber.Router, db *database.Database, authorization func(c *fiber.Ctx) (err error)) {
	userPoints := api.Group("/joker", authorization)
	userPoints.Get("/", Controller.GetAllJokers(db))
	userPoints.Post("/", Controller.CreateJoker(db))
	userPoints.Put("/", Controller.UpdateJoker(db))
	userPoints.Put("/", Controller.DeleteJoker(db))
}

func registerDocs(api fiber.Router, db *database.Database) {
	docs := api.Group("/docs")
	docs.Get("/*", swagger.HandlerDefault)
}
