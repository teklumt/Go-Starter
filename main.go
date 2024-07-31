package main

import (
	"crud/config"
	"crud/routes"
)

func main() {
	config.InitiEnvConfigs() 
	routes.IntializeRoutes()

	routes.Router.Run(config.EnvConfigs.LocalServerPort)
	 
} 