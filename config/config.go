package config

import (
	"flag"
	"fmt"
	"os"

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
	var port = os.Getenv("PORT")

	if flag.Lookup("test.v") == nil {
		fmt.Println("user c config")
		if port == "" {
			gonfig.GetConf("config.c.json", &configuration)
		} else {
			gonfig.GetConf("config.l.json", &configuration)
		}

	} else {
		fmt.Println("user l config")
		if port == "" {
			gonfig.GetConf("../config.c.json", &configuration)
		} else {
			gonfig.GetConf("../config.l.json", &configuration)
		}
	}

	return configuration
}
