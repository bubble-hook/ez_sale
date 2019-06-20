package api

import (
	"ezsale/db"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db.Init()
	newUser := `{
		"username" : "${uname}",
		"name" :"12345678"
	}`
	newUser = strings.Replace(newUser, "${uname}", uuid.New().String(), 1)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(newUser))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestGetUsers(t *testing.T) {
	db.Init()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, GetUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
