package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	LomsUrl             string `yaml:"lomsUrl"`
	ProductServiceUrl   string `yaml:"productServiceUrl"`
	ProductServiceToken string
}

func LoadConfig(dotEnvRequired bool) (Config, error) {
	viper.AddConfigPath("/app/config")
	viper.AddConfigPath("./cmd/app/config")
	viper.SetConfigName("app")
	viper.SetConfigType("yml")

	dotEnvPaths := []string{"./cmd/app/config/.env", "/app/config/.env"}
	dotEnvFound := false
	for _, path := range dotEnvPaths {
		if err := godotenv.Load(path); err != nil {
			//fmt.Printf(".env file not found in %s\n", path)
			continue
		}
		dotEnvFound = true
	}
	if !dotEnvFound && dotEnvRequired {
		return Config{}, fmt.Errorf(".env file required but not found in paths %v", dotEnvPaths)
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	return Config{
		LomsUrl:             viper.GetString("lomsUrl"),
		ProductServiceUrl:   viper.GetString("productServiceUrl"),
		ProductServiceToken: viper.GetString("PRODUCT_SERVICE_TOKEN"),
	}, err
}
