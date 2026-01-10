// @title           Gopportunities API
// @version         1.0
// @description     API para gerenciamento de vagas
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

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
