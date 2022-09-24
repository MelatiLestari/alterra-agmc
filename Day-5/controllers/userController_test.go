package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"Day-2/config"
)

var (
	userJSON      = `{"Email": "george@falcon.com", "password": "123"}`
	otherUserJSON = `{"Email": "george@falcon.com", "password": "1245"}`
)

func TestGetAllUsers(t *testing.T) {
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetAllUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.True(t, strings.Contains(rec.Body.String(), "{\"code\":200,\"data\":[{\"ID\":1"))
	}
}

func TestGetUserById_UserFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, GetUserById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.True(t, strings.Contains(rec.Body.String(), "{\"code\":200,\"data\":{\"ID\":1"))
	}
}

func TestGetUserById_UserNotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("5")

	if assert.NoError(t, GetUserById(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.True(t, strings.Contains(rec.Body.String(), "{\"code\":500,\"message\":\"Record not found\"}"))
	}
}

func TestCreateNewUser_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateNewUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.True(t, strings.Contains(rec.Body.String(), "{\"code\":201,\"message\":\"succesfully created with ID: 0\"}"))
	}
}

func TestUpdateUserById_UserFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(otherUserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	if assert.NoError(t, UpdateUserById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.True(t, strings.Contains(rec.Body.String(), "{\"code\":200,\"message\":\"succesfully updated user with ID: 2\"}"))
	}
}

func TestUpdateUserById_UserNotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(otherBookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("100")

	if assert.NoError(t, UpdateUserById(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.True(t, strings.Contains(rec.Body.String(), "{\"code\":500,\"message\":\"Record not found\"}"))
	}
}

func TestDeleteUserById_UserFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("3")

	if assert.NoError(t, DeleteUserById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.True(t, strings.Contains(rec.Body.String(), "{\"code\":200,\"message\":\"succesfully deleted user with ID: 3\"}"))
	}
}
