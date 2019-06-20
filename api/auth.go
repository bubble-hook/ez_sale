package api

import (
	"ezsale/config"
	"ezsale/db"
	"ezsale/model"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func createJWT(u *model.User) (string, error) {

	configuration := config.GetConfig()

	sUID := strconv.Itoa(int(u.ID))

	claims := JwtClaims{
		u.Name,
		jwt.StandardClaims{
			Id:        sUID,
			ExpiresAt: time.Now().Add(48 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte(configuration.APP_SINGINGKEY))
	if err != nil {
		return "", err
	}
	return token, nil
}

func Login(c echo.Context) error {
	db := db.DbManager()
	loginRequest := LoginRequest{}
	err := JsonBodyTo(c, &loginRequest)
	if err != nil {
		return ErrorResponse(c, err)
	}

	u := model.User{}

	db.First(&u, model.User{Username: loginRequest.UserName})

	if &u == nil {
		return ErrorResponseMessage(c, http.StatusUnauthorized, "Not Found UserName ")
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(loginRequest.Password))

	if err != nil {

		return ErrorResponseMessage(c, http.StatusUnauthorized, "Password Not Match. , "+err.Error())
	}

	nAccessToken, err := createJWT(&u)

	db.Where("user_id = ?", u.ID).Unscoped().Delete(model.UserToken{})

	nUToken := model.UserToken{UserId: u.ID, AccessToken: nAccessToken}

	db.Create(&nUToken)

	return c.JSON(http.StatusOK, nUToken)

}
