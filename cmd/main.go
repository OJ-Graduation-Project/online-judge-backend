package main

import (
	"flag"
	"github.com/OJ-Graduation-Project/online-judge-backend/config"
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


	_ = config // will be remove later
}

