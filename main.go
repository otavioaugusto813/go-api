package main

import (
	"github.com/otavioaugusto813/go-api/config"
	"github.com/otavioaugusto813/go-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)

		// panic(err)
		return
	}
	//Inicializando os routers (a api)
	router.Initialize()
	// router.Run() // listen and serve on 0.0.0.0:8080
}
