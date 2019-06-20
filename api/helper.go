package api

import (
	"encoding/json"
	"ezsale/db"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

func JsonBodyTo(c echo.Context, object interface{}) error {

	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &object)
	if err != nil {
		return err
	}
	return nil
}

func ErrorResponse(c echo.Context, e error) error {
	return c.JSON(http.StatusInternalServerError, map[string]string{
		"message": e.Error(),
	})
}

func ErrorResponseMessage(c echo.Context, statusCode int, errorMessage string) error {
	return c.JSON(statusCode, map[string]string{
		"message": errorMessage,
	})
}

func SearchMaster(c echo.Context, masterModel interface{}) error {
	db := db.DbManager()
	sName := c.QueryParam("sName")
	sCode := c.QueryParam("sCode")
	tx := db

	if sName != "" {
		tx = db.Where("name like ?", fmt.Sprintf("%s%s%s", "%", sName, "%ss"))
	}

	if sCode != "" {
		tx = db.Where("code like ?", fmt.Sprintf("%s%s%s", "%", sCode, "%ss"))
	}

	tx.Find(masterModel).Debug()
	return c.JSON(http.StatusOK, masterModel)
}

func DeleteMaster(c echo.Context, masterModel interface{}) error {
	id := c.Param("id")
	db := db.DbManager()
	cModel := 0
	db.Model(masterModel).Where("id = ?", id).Count(&cModel)
	if cModel < 1 {
		return ErrorResponseMessage(c, http.StatusNotFound, "Not Found Item")
	}
	db.Where("id = ?", id).First(masterModel).Delete(masterModel)
	return c.JSON(http.StatusOK, &masterModel)
}
