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

	MATCH_CORINGA       = "21c2c52b-c0c9-4f9f-a371-7f1a0ce99d7e"
	MATCH_OITAVAS_01    = "602d4734-ed53-4f19-9afd-695c7bcdba6c"
	MATCH_OITAVAS_02    = "132068de-6616-4dd8-9607-593f9703ea98"
	MATCH_OITAVAS_03    = "a7afd395-23c6-4d57-a04f-7360cdb0cd0b"
	MATCH_OITAVAS_04    = "c2caa666-3340-4a68-8dd6-57e779f2ad4c"
	MATCH_OITAVAS_05    = "4a343fb8-5b40-4b92-a324-7595430012d0"
	MATCH_OITAVAS_06    = "69613c5d-0b22-4069-8b16-4d342402c700"
	MATCH_OITAVAS_07    = "04660dfb-21af-4e68-bf63-2da682e70526"
	MATCH_OITAVAS_08    = "775957af-2909-4815-b0de-4e83272de49b"
	MATCH_QUARTAS_01    = "382b4798-4e88-4a79-92ec-860e11a2fd17"
	MATCH_QUARTAS_02    = "f2b8b4c0-74c6-400d-8348-c7b97938eabd"
	MATCH_QUARTAS_03    = "3437999d-c42a-40ab-ab2d-1d791d0b72d0"
	MATCH_QUARTAS_04    = "7bf5625f-84c0-4501-8ac8-dc3523450613"
	MATCH_SEMIFINAIS_01 = "5f055bc0-1eeb-4d54-8c5a-1db00a2836a9"
	MATCH_SEMIFINAIS_02 = "8c869654-d89d-49c0-b60d-5101de8be22b"
	MATCH_TERCEIRO      = "3ed6c6bb-cc7d-44e3-b14e-0ab5ae6160bc"
	MATCH_FINAIS        = "ed101cf9-7697-4ffa-ae19-f27ed2d1c091"

	NT_ALEMANHA       = "b6fb68cc-e1ad-4b5e-9255-a9733ed154e7"
	NT_ARABIA_SAUDITA = "339bc00e-f8d0-4f33-bfd8-95b4d192dbef"
	NT_ARGENTINA      = "cef4796c-5fb4-4f2d-809b-92f1d927c388"
	NT_AUSTRALIA      = "49a25e66-5ba7-4368-85c3-48403c95443d"
	NT_BELGICA        = "55f82c5d-4b28-48da-a92d-07bf8b939395"
	NT_BRASIL         = "fe0eb3c4-b1bb-4f72-ac4c-e025b2143fd6"
	NT_CAMAROES       = "a9761e12-2b6d-4c3d-96ec-a61a5496e7af"
	NT_CANADA         = "d34ecd53-ff91-4497-ba6f-39f2fdc96387"
	NT_CATAR          = "9f4a3073-e7d0-42c5-b616-e8af267db970"
	NT_COREIA_DO_SUL  = "994c3c5d-f235-402c-9f18-c9c9e448b7df"
	NT_COSTA_RICA     = "23729113-6189-42bf-bc3a-6d35d87ce9c5"
	NT_CROACIA        = "fc94209e-4058-445d-b151-d983e7e72b47"
	NT_DINAMARCA      = "d8f1ced1-bb7d-4dd1-b770-2211d5e835bd"
	NT_EQUADOR        = "f8419460-d30e-47a2-91c9-3e7a91688642"
	NT_ESPANHA        = "960b77ed-eb3e-4be2-bdc4-244fa01516d2"
	NT_ESTADOS_UNIDOS = "122f703d-86fe-4509-b55d-f30d4072dbfd"
	NT_FRANÇA         = "97301b75-9a85-413c-8221-dd1f7cc20c83"
	NT_GANA           = "ce0dc619-97bd-4784-93e5-5c46113b2f1a"
	NT_HOLANDA        = "8859a52c-ce9f-4ce5-b15a-afbe42a9e61c"
	NT_INGLATERRA     = "238b01dd-8380-43c1-831f-a35c3dac71d6"
	NT_IRÃ            = "06ce8b21-af8e-4ee8-ae42-39ecd7d24f41"
	NT_JAPAO          = "5109b641-26da-44c8-b80e-cdac7a72ddaa"
	NT_MARROCOS       = "e9e6f2ff-bb53-4bf6-bceb-2f795f671769"
	NT_MEXICO         = "d3834b17-6ea8-40bb-9227-ac277a890167"
	NT_PAIS_DE_GALES  = "e32f08ae-cc35-43b0-b89b-87a1a2cfb7b7"
	NT_POLONIA        = "d1ced66c-6135-435b-9f8b-1e74a5bbe62b"
	NT_PORTUGAL       = "2e89008d-8f13-413f-9772-9351c8decb91"
	NT_SENEGAL        = "cb3ebe1d-6852-4ae1-a0a4-7535b0c62ad0"
	NT_SERVIA         = "e58312cd-4698-4cc3-b642-a234df49b047"
	NT_SUIÇA          = "604e3940-2ae2-463f-841a-11088078b345"
	NT_TUNISIA        = "4ddcfd7e-4310-4a00-a8ac-f4de440dacd4"
	NT_URUGUAI        = "2c269272-d3b9-4dfb-8683-9d68f05b22e3"
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
