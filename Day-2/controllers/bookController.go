package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"Day-2/models"
)

var bookList = models.Books

func GetBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, book := range bookList {
		if book.ID == uint(id) {
			response := map[string]interface{}{
				"message": "success",
				"code":    http.StatusOK,
				"data":    book,
			}
			return c.JSON(http.StatusOK, response)
		}
	}

	response := map[string]interface{}{
		"message": "book not found",
		"code":    http.StatusBadRequest,
	}
	return c.JSON(http.StatusBadRequest, response)

}

func GetAllBooks(c echo.Context) error {
	response := map[string]interface{}{
		"message": "success",
		"code":    http.StatusOK,
		"data":    bookList,
	}
	return c.JSON(http.StatusOK, response)
}

func CreateNewBook(c echo.Context) error {
	bookPayload := models.Book{}
	err := json.NewDecoder(c.Request().Body).Decode(&bookPayload)
	if err != nil {
		response := map[string]interface{}{
			"message": "failed to create book",
			"code":    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	for _, book := range bookList {
		if book.ID == uint(bookPayload.ID) {
			response := map[string]interface{}{
				"message": "id already exist",
				"code":    http.StatusBadRequest,
			}
			return c.JSON(http.StatusOK, response)
		}
	}

	bookList = append(bookList, bookPayload)
	fmt.Println("Current Repo:", bookList)
	response := map[string]interface{}{
		"message": "success",
		"code":    http.StatusCreated,
	}
	return c.JSON(http.StatusCreated, response)
}

func UpdateBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	bookPayload := models.Book{}
	err := json.NewDecoder(c.Request().Body).Decode(&bookPayload)
	if err != nil {
		response := map[string]interface{}{
			"message": "failed to update book",
			"code":    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	for i, book := range bookList {
		if book.ID == uint(id) {
			bookList[i] = bookPayload
			response := map[string]interface{}{
				"message": "success",
				"code":    http.StatusOK,
			}
			fmt.Println("Current Repo:", bookList)
			return c.JSON(http.StatusOK, response)
		}
	}

	response := map[string]interface{}{
		"message": "failed to update book",
		"code":    http.StatusBadRequest,
	}
	return c.JSON(http.StatusBadRequest, response)
}

func DeleteBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, book := range bookList {
		if book.ID == uint(id) {
			bookList[i] = bookList[0]
			bookList = bookList[1:]
			response := map[string]interface{}{
				"message": "success",
				"code":    http.StatusOK,
			}
			fmt.Println("Current Repo:", bookList)
			return c.JSON(http.StatusOK, response)
		}
	}

	response := map[string]interface{}{
		"message": "failed to delete book",
		"code":    http.StatusBadRequest,
	}
	return c.JSON(http.StatusBadRequest, response)
}
