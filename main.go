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

	// router.Initialize()
	router.Initialize()
	

}