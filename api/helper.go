package api

import (
	"encoding/json"
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
