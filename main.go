package main

import (
	"ezsale/config"
	"ezsale/db"
	"ezsale/route"
	"fmt"
)

// type InfoRequest struct {
// 	Name string `json:"name"`
// }

// type AuthRequest struct {
// 	UserName string `json:"userName"`
// 	Password string `json:"password"`
// }

// type JwtClaims struct {
// 	Name string `json:"name"`
// 	jwt.StandardClaims
// }

// func errResponse(c echo.Context, errorMessage string) error {
// 	return c.JSON(http.StatusInternalServerError, map[string]string{
// 		"errorMessage": errorMessage,
// 	})
// }

// func home(c echo.Context) error {
// 	token, err := createJWT()
// 	if err != nil {
// 		return errResponse(c, err.Error())
// 	}
// 	return c.String(http.StatusOK, token)
// }

// func getInfo(c echo.Context) error {
// 	q := c.QueryParam("q")
// 	s := c.QueryParam("s")
// 	return c.JSON(http.StatusOK, map[string]string{
// 		"q": q,
// 		"s": s,
// 	})
// }

// func createJWT() (string, error) {
// 	claims := JwtClaims{
// 		"Admin",
// 		jwt.StandardClaims{
// 			Id:        "uid",
// 			ExpiresAt: time.Now().Add(48 * time.Hour).Unix(),
// 		},
// 	}
// 	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
// 	token, err := rawToken.SignedString([]byte("sdaxc32rf"))
// 	if err != nil {
// 		return "", err
// 	}
// 	return token, nil
// }

// func createInfo(c echo.Context) error {
// 	infoRequest := InfoRequest{}
// 	defer c.Request().Body.Close()
// 	b, err := ioutil.ReadAll(c.Request().Body)
// 	if err != nil {
// 		return errResponse(c, err.Error())
// 	}

// 	err = json.Unmarshal(b, &infoRequest)

// 	if err != nil {
// 		return errResponse(c, err.Error())
// 	}

// 	return c.String(http.StatusOK, "infoRequest.Name :"+infoRequest.Name)
// }

func main() {
	fmt.Println("start server")
	configuration := config.GetConfig()
	db.Init()
	e := route.Init()

	e.Logger.Fatal(e.Start(":" + configuration.APP_PORT))
}
