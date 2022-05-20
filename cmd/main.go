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
		"./config/res/localhost_config.json",
		"Location of the config file",
	)

	flag.Parse()

	config.LoadEnv(envFilePath)
	config.LoadConfig(configFilePath)
	//Test contest and test submissions.
	/*
		all := contest.GetInstance()
		all.GetContestAndStart(131458055240186)
	*/
	router := routes.LoadRoutes()
	router.Use(routes.Middleware)
	server_uri := fmt.Sprintf("%s:%s", config.AppConfig.Server.Host, config.AppConfig.Server.Port)
	log.Fatal(http.ListenAndServe(server_uri, router))

}
