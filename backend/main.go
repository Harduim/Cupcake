package main

import (
	configuration "cupcake/app/config"
	"cupcake/app/database"
	"cupcake/app/database/fixtures"
	"cupcake/app/routes"
	"cupcake/app/service"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
)

type App struct {
	*fiber.App
	DB *database.Database
}

func getDB(config *configuration.Config) (*database.Database, error) {
	dbConfig := database.DatabaseConfig{
		Host:     config.GetString("DB_HOST"),
		Username: config.GetString("DB_USERNAME"),
		Database: config.GetString("DB_DATABASE"),
		Password: config.GetString("DB_PASSWORD"),
		Port:     config.GetInt("DB_PORT"),
	}

	db, err := database.New(&dbConfig)

	return db, err
}

func getSSO(config *configuration.Config) (*service.SSOClient, error) {
	ssoClient := service.NewSSO()
	ssoConfig := service.SSOConfig{
		Authority:    config.GetString("AUTHORITY"),
		ClientId:     config.GetString("CLIENT_ID"),
		ClientSecret: config.GetString("CLIENT_SECRET"),
		RedirectUrl:  config.GetString("REDIRECT"),
		Scopes:       []string{config.GetString("SCOPES")}}

	err, s := ssoClient.Init(&ssoConfig)

	if err != nil {
		return nil, err
	}
	return s, nil
}

func main() {
	config := configuration.New()

	db, err := getDB(config)

	if err != nil {
		log.Fatalf("Error connecting to DB")
	}

	sso, err := getSSO(config)

	if err != nil {
		log.Fatalf("Error getting sso service")
	}

	if config.GetString("ENV") == "dev" {
		err = fixtures.CreateFixtures(db)

		if err != nil {
			log.Fatalf("Error creating fixuters")
		}
	}

	app := App{
		App: fiber.New(),
		DB:  db,
	}

	api := app.Group("/api")
	routes.RegisterRoutes(api, app.DB, sso)

	// Custom 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		if err := c.SendStatus(fiber.StatusNotFound); err != nil {
			panic(err)
		}
		if err := c.Render("errors/404", fiber.Map{}); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return err
	})

	// Close any connections on interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		app.exit()
	}()

	// Start listening on the specified address
	err = app.Listen(config.GetString("APP_ADDR"))
	if err != nil {
		app.exit()
	}
}

func (app *App) exit() {
	_ = app.Shutdown()
}
