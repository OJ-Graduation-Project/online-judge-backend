package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/joho/godotenv"
)

var AppConfig ApplicationConfig

func LoadConfig(configFilePath string) {

	file, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(file, &AppConfig)
	if err != nil {
		fmt.Println(err)
	}

}

func LoadEnv(envPath string) {
	// load .env file
	err := godotenv.Load(envPath)
	if err != nil {
		fmt.Println(err)
		return
	}

}
