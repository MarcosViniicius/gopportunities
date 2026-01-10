package main

import (
	"github.com/MarcosViniicius/gopportunities/config"
	_ "github.com/MarcosViniicius/gopportunities/docs"
	"github.com/MarcosViniicius/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
