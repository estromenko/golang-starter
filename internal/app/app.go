package app

import (
	"golang-starter/internal/server"
	"golang-starter/pkg/config"
	"golang-starter/pkg/database"
)

func Run(configPath string) error {
	if err := config.Load(configPath); err != nil {
		return err
	}

	if err := database.Init(); err != nil {
		return err
	}

	if err := server.Init(); err != nil {
		return err
	}

	return server.Run()
}
