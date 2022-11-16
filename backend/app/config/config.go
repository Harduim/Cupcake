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

	config.setDefaults()
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

func (config *Config) setDefaults() {
	// Set default App configuration
	config.SetDefault("APP_ADDR", ":3000")
	config.SetDefault("APP_ENV", "local")

	// Set default database configuration
	config.SetDefault("DB_DRIVER", "mysql")
	config.SetDefault("DB_HOST", "localhost")
	config.SetDefault("DB_USERNAME", "fiber")
	config.SetDefault("DB_PASSWORD", "password")
	config.SetDefault("DB_PORT", 3306)
	config.SetDefault("DB_DATABASE", "boilerplate")
}
