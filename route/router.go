package route

import (
	"ezsale/api"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {

	userResouceName := "/users"
	productCategoryResouceName := "/productCategory"

	//configuration := config.GetConfig()

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.POST("/login", api.Login)
	e.POST(userResouceName, api.CreateUser)

	authGroup := e.Group("api")

	// authGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: []byte(configuration.APP_SINGINGKEY),
	// }))

	authGroup.GET(userResouceName, api.GetUsers)
	authGroup.GET(fmt.Sprintf("%s/:id", userResouceName), api.GetUserById)
	authGroup.DELETE(fmt.Sprintf("%s/:id", userResouceName), api.DeleteUser)

	authGroup.GET(productCategoryResouceName, api.GetProductCategory)
	authGroup.POST(productCategoryResouceName, api.CreateProductCategory)

	// e.GET("/", home)
	// e.GET("/info", getInfo)
	// e.POST("/info", createInfo)
	return e
}
