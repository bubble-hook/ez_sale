package api

import (
	"ezsale/db"
	"ezsale/model"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c echo.Context) error {
	return SearchMaster(c, &[]model.User{})
}

func CreateUser(c echo.Context) error {
	db := db.DbManager()
	u := model.User{}

	err := JsonBodyTo(c, &u)
	if err != nil {
		return ErrorResponse(c, err)
	}

	cUser := 0

	db.Model(&model.User{}).Where("username = ?", u.Username).Count(&cUser)
	if cUser > 0 {
		return ErrorResponseMessage(c, http.StatusBadRequest, "Duplicated userName")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.MinCost)
	if err != nil {
		return ErrorResponse(c, err)
	}
	u.Password = string(hash)
	db.Create(&u)
	//db.Save(&u)
	return c.JSON(http.StatusCreated, u)
}

func GetUserById(c echo.Context) error {
	id := c.Param("id")
	db := db.DbManager()
	u := model.User{}
	db.Where("id = ?", id).First(&u)
	if &u == nil {
		return c.String(http.StatusNotFound, "")
	}
	return c.JSON(http.StatusOK, u)
}

func DeleteUser(c echo.Context) error {
	return DeleteMaster(c, &model.User{})
}
