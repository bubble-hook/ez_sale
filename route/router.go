package route

import (
	"ezsale/api"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {

	// userResouceName := "/users"
	// productCategoryResouceName := "/productCategory"
	// unitQuantityResouceName := "/unitQuantity"
	// productResouceName := ""

	resouceNames := map[string]string{
		"users":           "/users",
		"productCategory": "/productCategory",
		"sellingUnit":     "/sellingUnit",
		"product":         "/product",
	}

	//configuration := config.GetConfig()

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.POST("/login", api.Login)
	e.POST(resouceNames["users"], api.CreateUser)

	authGroup := e.Group("api")

	// authGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: []byte(configuration.APP_SINGINGKEY),
	// }))

	authGroup.GET(resouceNames["users"], api.GetUsers)
	authGroup.GET(fmt.Sprintf("%s/:id", resouceNames["users"]), api.GetUserById)
	authGroup.DELETE(fmt.Sprintf("%s/:id", resouceNames["users"]), api.DeleteUser)

	authGroup.GET(resouceNames["productCategory"], api.GetProductCategory)
	authGroup.POST(resouceNames["productCategory"], api.CreateProductCategory)
	authGroup.PUT(resouceNames["productCategory"], api.UpdateProductCategory)
	authGroup.DELETE(fmt.Sprintf("%s/:id", resouceNames["productCategory"]), api.DeleteProductCategory)

	authGroup.GET(resouceNames["sellingUnit"], api.GetSellingUnit)
	authGroup.POST(resouceNames["sellingUnit"], api.CreateSellingUnit)
	authGroup.PUT(resouceNames["sellingUnit"], api.UpdateSellingUnit)
	authGroup.DELETE(fmt.Sprintf("%s/:id", resouceNames["sellingUnit"]), api.DeleteSellingUnit)

	authGroup.POST(resouceNames["product"], api.CreateProduct)

	authGroup.POST("/order/process", api.ProcessOrder)

	// e.GET("/", home)
	// e.GET("/info", getInfo)
	// e.POST("/info", createInfo)
	return e
}
