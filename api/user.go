package api

import (
	"ezsale/db"
	"ezsale/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c echo.Context) error {
	return SearchMaster(c, &[]model.User{})
}

func CreateUser(c echo.Context) error {
	db := db.DbManager()

	uRequest := model.CreateUserRequest{}

	err := JsonBodyTo(c, &uRequest)
	if err != nil {
		return ErrorResponse(c, err)
	}

	u := model.User(uRequest.User)

	cUser := 0

	db.Model(&model.User{}).Where("username = ?", u.Username).Count(&cUser)
	if cUser > 0 {
		return ErrorResponseMessage(c, http.StatusBadRequest, "Duplicated userName")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(uRequest.Password), bcrypt.MinCost)
	if err != nil {
		return ErrorResponse(c, err)
	}
	u.Password = string(hash)

	tx := db.Begin()

	store := model.Store{}
	storeUser := model.StoreUser{}

	if err := tx.Create(&u).Error; err != nil {
		tx.Rollback()
		return ErrorResponse(c, err)
	}
	store.OwnerUserID = u.ID
	store.StoreID = uuid.New().String()
	store.Name = uRequest.StoreName
	if err := tx.Create(&store).Error; err != nil {
		tx.Rollback()
		return ErrorResponse(c, err)
	}
	storeUser.StoreID = store.ID
	storeUser.UserID = u.ID
	if err := tx.Create(&storeUser).Error; err != nil {
		tx.Rollback()
		return ErrorResponse(c, err)
	}
	tx.Commit()

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
