package main

import (
	"github.com/MarcosViniicius/gopportunities/config"
	_ "github.com/MarcosViniicius/gopportunities/docs"
	"github.com/MarcosViniicius/gopportunities/router"
)

var (
	logger *config.Logger
)

// @title           Gopportunities API
// @version         1.0
// @description     API para gerenciar oportunidades de emprego
// @host            localhost:8080
// @basePath        /api/v1
// @schemes         http
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
