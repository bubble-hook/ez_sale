package api

import (
	"encoding/json"
	"errors"
	"ezsale/db"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

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

func ValidateMasterData(masterModel interface{}, isCreate bool) error {
	db := db.DbManager()
	r := reflect.ValueOf(masterModel)
	fCode := reflect.Indirect(r).FieldByName("Code")
	fID := reflect.Indirect(r).FieldByName("ID")
	cModel := 0
	code := fCode.String()
	id := fID.Uint()
	tableName := db.NewScope(masterModel).GetModelStruct().TableName(db)
	if isCreate {
		db.Table(tableName).Where("code = ?", code).Count(&cModel)
	} else {
		db.Table(tableName).Where("code = ?", code).Where("id <> ?", id).Count(&cModel)
	}

	if cModel > 0 {
		return errors.New("Duppicate Master Field Code")
	}
	return nil
}

func CreateMasterData(c echo.Context, masterModel interface{}) error {
	db := db.DbManager()
	err := JsonBodyTo(c, masterModel)
	if err != nil {
		return ErrorResponse(c, err)
	}
	err = ValidateMasterData(masterModel, true)
	if err != nil {
		return ErrorResponse(c, err)
	}
	db.Create(masterModel)
	return c.JSON(http.StatusCreated, masterModel)
}

func UpdateMasterData(c echo.Context, masterModel interface{}) error {
	err := JsonBodyTo(c, masterModel)
	if err != nil {
		return ErrorResponse(c, err)
	}
	err = ValidateMasterData(masterModel, false)
	if err != nil {
		return ErrorResponse(c, err)
	}
	///fmt.Println(spew.Sdump(&masterModel))
	db := db.DbManager()
	db.Debug().Model(masterModel).Updates(masterModel)
	return c.JSON(http.StatusOK, &masterModel)
}
