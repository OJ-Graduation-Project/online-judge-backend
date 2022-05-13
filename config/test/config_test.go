package config_test

import (
	"testing"

	"github.com/OJ-Graduation-Project/online-judge-backend/config"
)

func Test(t *testing.T) {
	config.LoadConfig("../../config/res/config.json")
	if config.AppConfig.Server.Host == "" {
		t.Fail()
	}
}
