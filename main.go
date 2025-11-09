package main

import (
	"github.com/nathan/gopportunitiesbb/configs"
	"github.com/nathan/gopportunitiesbb/router"
)

var (
	logger *config.Logger
)

// @title           Go + Gin + Swagger API Demo
// @version         1.0
// @description     This is a sample API server using Gin, sqlite and documented with Swagger
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Failed to initialize configs: %v", err)
		return
		// panic(err)
	}

	// Database init
	// _, err = config.InitializeSQLite()
	// _, err = config.GetSQLite()
	// if err != nil {
	// 	logger.Errorf("Failed to initialize database: %v", err)
	// 	return
	// }


	// router.Initialize()
	router.Initialize()
	

}