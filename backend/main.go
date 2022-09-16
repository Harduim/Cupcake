package main

import (
	configuration "cupcake/app/config"
	"cupcake/app/database"
	"cupcake/app/models"
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	*fiber.App

	DB *database.Database
}

func main() {
	config := configuration.New()

	app := App{
		App: fiber.New(*config.GetFiberConfig()),
	}

	// Initialize database
	db, err := database.New(&database.DatabaseConfig{
		Host:     config.GetString("DB_HOST"),
		Username: config.GetString("DB_USERNAME"),
		Password: config.GetString("DB_PASSWORD"),
		Port:     config.GetInt("DB_PORT"),
		Database: config.GetString("DB_DATABASE"),
	})

	// Auto-migrate database models
	if err != nil {
		fmt.Println("failed to connect to database:", err.Error())
	} else {
		if db == nil {
			fmt.Println("failed to connect to database: db variable is nil")
		} else {
			app.DB = db

			err = app.DB.AutoMigrate(&models.User{})
			if err != nil {
				fmt.Println("failed to auto migrate user model:", err.Error())
				return
			}
		}
	}

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
