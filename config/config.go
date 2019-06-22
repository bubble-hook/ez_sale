package config

import (
	"flag"
	"fmt"

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
	DB_SSL         bool
}

func GetConfig() Configuration {
	// _, filename, _, _ := runtime.Caller(0)
	// fmt.Println("Current config filename: " + filename)
	configuration := Configuration{}

	if flag.Lookup("test.v") == nil {
		fmt.Println("user c config")
		gonfig.GetConf("config.c.json", &configuration)
	} else {
		fmt.Println("user l config")
		gonfig.GetConf("../config.c.json", &configuration)
	}

	return configuration
}
