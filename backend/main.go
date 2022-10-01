package main

import (
	configuration "cupcake/app/config"
	"cupcake/app/database"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
)

type App struct {
	*fiber.App

	DB *database.Database
}

func main() {
	config := configuration.New()
	dbConfig := database.DatabaseConfig{
		Host:     config.GetString("DB_HOST"),
		Username: config.GetString("DB_USERNAME"),
		Database: config.GetString("DB_DATABASE"),
		Password: config.GetString("DB_PASSWORD"),
		Port:     config.GetInt("DB_PORT"),
	}

	_, err := database.New(&dbConfig)

	if err != nil {
		log.Fatalf("Error connecting to DB")
	}

	app := App{
		App: fiber.New(),
	}

	//api := app.Group("/api")
	//routes.RegisterRoutes(api, app.DB)

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
