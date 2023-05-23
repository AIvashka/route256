package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string `yaml:"server_address"`
}

func LoadConfig() (Config, error) {
	viper.AddConfigPath("/app/config")
	viper.SetConfigName("app")
	viper.SetConfigType("yml")

	if err := godotenv.Load("/app/config/.env"); err != nil {
		fmt.Printf("No .env file found")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	return Config{
		ServerAddress: viper.GetString("server_address"),
	}, err
}
