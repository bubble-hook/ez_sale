package main

import (
	"ezsale/config"
	"ezsale/db"
	"ezsale/route"
	"fmt"
	"os"
)

func main() {
	fmt.Println("start server")

	db.Init()
	e := route.Init()

	e.Logger.Fatal(e.Start(":" + getPort()))
}

func getPort() string {
	configuration := config.GetConfig()
	var port = os.Getenv("PORT")
	if port == "" {
		port = configuration.APP_PORT
		fmt.Println("No Port In Heroku" + port)
	}
	return "" + port
}
