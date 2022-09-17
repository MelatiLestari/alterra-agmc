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
	userJSON = `{"Email": "george@falcon.com", "password": "123"}`

// anotherBookJSON = `{"id": 2, "title": "Mocking Jay", "author": "George"}`
// otherBookJSON   = `{"id": 5, "title": "Mocking Jay", "author": "George"}`
)

func TestGetAllUsers(t *testing.T) {
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetAllUsers(c)) {
		// fmt.Println("isinya", rec.Body.String())
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
		// fmt.Println("isinya", rec.Body.String())
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
		// fmt.Println("isinya", rec.Body.String())
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.True(t, strings.Contains(rec.Body.String(), "{\"code\":500,\"message\":\"Record not found\"}"))
	}
}

// func TestCreateNewBook_Success(t *testing.T) {
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(bookJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	if assert.NoError(t, CreateNewBook(c)) {
// 		fmt.Println("isinya", rec.Body.String())
// 		assert.Equal(t, http.StatusCreated, rec.Code)
// 		assert.True(t, strings.Contains(rec.Body.String(), "{\"code\":201,\"message\":\"success\"}"))
// 	}
// }

// func TestCreateNewBook_Failed(t *testing.T) {
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(anotherBookJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	if assert.NoError(t, CreateNewBook(c)) {
// 		assert.Equal(t, http.StatusBadRequest, rec.Code)
// 		assert.True(t, strings.Contains(rec.Body.String(), "{\"code\":400,\"message\":\"id already exist\"}"))
// 	}
// }

// func TestUpdateBookById_BookFound(t *testing.T) {
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(anotherBookJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetParamNames("id")
// 	c.SetParamValues("2")

// 	if assert.NoError(t, UpdateBookById(c)) {
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		assert.True(t, strings.Contains(rec.Body.String(), "{\"code\":200,\"message\":\"success\"}"))
// 	}
// }

// func TestUpdateBookById_BookNotFound(t *testing.T) {
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(otherBookJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetParamNames("id")
// 	c.SetParamValues("5")

// 	if assert.NoError(t, UpdateBookById(c)) {
// 		assert.Equal(t, http.StatusBadRequest, rec.Code)
// 		assert.True(t, strings.Contains(rec.Body.String(), "{\"code\":400,\"message\":\"failed to update book\"}"))
// 	}
// }

// func TestDeleteBookById_BookFound(t *testing.T) {
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodDelete, "/", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetParamNames("id")
// 	c.SetParamValues("2")

// 	if assert.NoError(t, DeleteBookById(c)) {
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		assert.True(t, strings.Contains(rec.Body.String(), "{\"code\":200,\"message\":\"success\"}"))
// 	}
// }

// func TestDeleteBookById_BookNotFound(t *testing.T) {
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodDelete, "/", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetParamNames("id")
// 	c.SetParamValues("5")

// 	if assert.NoError(t, DeleteBookById(c)) {
// 		assert.Equal(t, http.StatusBadRequest, rec.Code)
// 		assert.True(t, strings.Contains(rec.Body.String(), "{\"code\":400,\"message\":\"failed to delete book\"}"))
// 	}
// }
