package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/OJ-Graduation-Project/online-judge-backend/config"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/routes"
)

func main() {

	// env file path
	var envFilePath string

	flag.StringVar(
		&envFilePath,
		"env",
		"./config/res/.env-example",
		"Location of the environment file",
	)
	


	// config file flag
	var configFilePath string
	flag.StringVar(
		&configFilePath,
		"config",
		"./config/res/config.json",
		"Location of the config file",
	)

	flag.Parse()


	config.LoadEnv(envFilePath)
	config := config.LoadConfig(configFilePath)
	
	router:= routes.LoadRoutes()
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d",config.Server.Host, config.Server.Port), router))

	_ = config // will be removed later
}

