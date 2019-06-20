package api

import (
	"ezsale/db"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	login = `{
		"username" : "a11",
		"password" :"12345678"
	}`
	loginFaild = `{
		"username" : "a11123xfm",
		"password" :"12345678"
	}`
)

func TestLoginSuccess(t *testing.T) {
	db.Init()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(login))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestLoginFailed(t *testing.T) {
	db.Init()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(loginFaild))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, Login(c)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	}
}
