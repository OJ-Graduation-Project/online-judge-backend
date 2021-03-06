package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/OJ-Graduation-Project/online-judge-backend/config"

	// "github.com/OJ-Graduation-Project/online-judge-backend/internal/contest"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/routes"
)

func main() {

	// env file path
	var envFilePath string
	os.Stdout = nil
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

	// all := contest.GetInstance()
	// all.GetContestAndStart(142995735221753)

	var err error
	db.DbConn, err = db.CreateDbConn()
	if err != nil {
		fmt.Print(err)
	}
	router := routes.LoadRoutes()
	router.Use(routes.Middleware)
	server_uri := fmt.Sprintf("%s:%s", config.AppConfig.Server.Host, config.AppConfig.Server.Port)
	log.Fatal(http.ListenAndServe(server_uri, router))

}
