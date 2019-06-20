package main

import (
	"ezsale/config"
	"ezsale/db"
	"ezsale/route"
	"fmt"
)

func main() {
	fmt.Println("start server")
	configuration := config.GetConfig()
	db.Init()
	e := route.Init()

	e.Logger.Fatal(e.Start(":" + configuration.APP_PORT))
}
