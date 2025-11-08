package main

import (
	"github.com/nathan/gopportunitiesbb/config"
	"github.com/nathan/gopportunitiesbb/router"
)

var (
	logger *config.Logger
)

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
	_, err = config.GetSQLite()
	if err != nil {
		logger.Errorf("Failed to initialize database: %v", err)
		return
	}

	// router.Initialize()
	router.Initialize()
	

}