package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

var (
	config ApplicationConfig
)

type ApplicationConfig struct {
	Server struct {
		Fiber fiber.Config `json:"fiber"`
		Port  int          `json:"port"`
	} `json:"server"`
	DB struct {
		Dsn string `json:"dsn"`
	} `json:"db"`
}

func Load(configPath string) error {
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.Unmarshal(&config)
}

func GetConfig() *ApplicationConfig {
	return &config
}
