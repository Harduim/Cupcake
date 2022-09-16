package config

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper

	errorHandler fiber.ErrorHandler
	fiber        *fiber.Config
}

func New() *Config {
	config := &Config{
		Viper: viper.New(),
	}

	// Select the .env file
	config.SetConfigName(".env")
	config.SetConfigType("dotenv")
	config.AddConfigPath(".")

	// Automatically refresh environment variables
	config.AutomaticEnv()

	// Read configuration
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("failed to read configuration:", err.Error())
			os.Exit(1)
		}
	}

	return config
}

func (config *Config) GetFiberConfig() *fiber.Config {
	return config.fiber
}
