package route

import (
	"ezsale/api"
	"ezsale/config"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {

	configuration := config.GetConfig()

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.POST("/login", api.Login)

	authGroup := e.Group("/api")

	authGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(configuration.APP_SINGINGKEY),
	}))

	authGroup.GET("/users", api.GetUsers)
	authGroup.GET("/users/:id", api.GetUserById)
	authGroup.DELETE("/users/:id", api.DeleteUser)
	authGroup.POST("/users", api.CreateUser)

	// e.GET("/", home)
	// e.GET("/info", getInfo)
	// e.POST("/info", createInfo)
	return e
}
