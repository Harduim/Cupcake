package config

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

const (
	BRKT_FINAIS     = "5ef28a89-f697-4af2-931d-808c41cbd2d1"
	BRKT_TERCERIRO  = "5e87f7d4-ac15-4f8b-b82c-3f50f2d5371f"
	BRKT_SEMIFINAIS = "40e58268-0fc5-4dec-8fcb-b52b46006215"
	BRKT_QUARTAS    = "22ecea42-848e-43d1-a387-5de1bd468338"
	BRKT_OITAVAS    = "ef13e77f-b345-4f4d-b4a7-2d1cfb12fa48"
	BRKT_CORINGA    = "ffa8f4c1-1183-477a-991f-d4652ba2e227"

	MATCH_CORINGA = "602d4734-ed53-4f19-9afd-695c7bcdba6c"

	MATCH_OITAVAS_01 = "602d4734-ed53-4f19-9afd-695c7bcdba6c"
	MATCH_OITAVAS_02 = "132068de-6616-4dd8-9607-593f9703ea98"
	MATCH_OITAVAS_03 = "a7afd395-23c6-4d57-a04f-7360cdb0cd0b"
	MATCH_OITAVAS_04 = "c2caa666-3340-4a68-8dd6-57e779f2ad4c"
	MATCH_OITAVAS_05 = "4a343fb8-5b40-4b92-a324-7595430012d0"
	MATCH_OITAVAS_06 = "69613c5d-0b22-4069-8b16-4d342402c700"
	MATCH_OITAVAS_07 = "04660dfb-21af-4e68-bf63-2da682e70526"
	MATCH_OITAVAS_08 = "775957af-2909-4815-b0de-4e83272de49b"

	MATCH_QUARTAS_01 = "382b4798-4e88-4a79-92ec-860e11a2fd17"
	MATCH_QUARTAS_02 = "f2b8b4c0-74c6-400d-8348-c7b97938eabd"
	MATCH_QUARTAS_03 = "3437999d-c42a-40ab-ab2d-1d791d0b72d0"
	MATCH_QUARTAS_04 = "7bf5625f-84c0-4501-8ac8-dc3523450613"

	MATCH_SEMIFINAIS_01 = "5f055bc0-1eeb-4d54-8c5a-1db00a2836a9"
	MATCH_SEMIFINAIS_02 = "8c869654-d89d-49c0-b60d-5101de8be22b"

	MATCH_TERCEIRO = "3ed6c6bb-cc7d-44e3-b14e-0ab5ae6160bc"

	MATCH_FINAIS = "ed101cf9-7697-4ffa-ae19-f27ed2d1c091"

	NT_BRAZIL = "6d71278a-4eca-42a8-8ec2-fa51a31ef95c"
	NT_FRANCE = "4935b4e1-f422-41a7-9a22-051f429ff5e4"
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
