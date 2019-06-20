package config

import (
	"flag"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DB_USERNAME    string
	DB_PASSWORD    string
	DB_PORT        string
	DB_HOST        string
	DB_NAME        string
	APP_PORT       string
	APP_SINGINGKEY string
}

func GetConfig() Configuration {
	// _, filename, _, _ := runtime.Caller(0)
	// fmt.Println("Current config filename: " + filename)
	configuration := Configuration{}

	if flag.Lookup("test.v") == nil {
		gonfig.GetConf("config.json", &configuration)
	} else {
		gonfig.GetConf("../config.json", &configuration)
	}

	return configuration
}
