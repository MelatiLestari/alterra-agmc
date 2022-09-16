package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"Day-2/lib/database"
	"Day-2/models"
)

func GetAllUsers(c echo.Context) error {
	users, err := database.GetAllUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := map[string]interface{}{
		"message": "success",
		"code":    http.StatusOK,
		"data":    users,
	}
	return c.JSON(http.StatusOK, response)
}

func GetUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := database.GetUserById(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := map[string]interface{}{
		"message": "success",
		"code":    http.StatusOK,
		"data":    user,
	}
	return c.JSON(http.StatusOK, response)
}

func CreateNewUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	err := database.CreateNewUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	message := fmt.Sprintf("succesfully created with ID: %d", user.ID)
	response := map[string]interface{}{
		"message": message,
		"code":    http.StatusCreated,
	}
	return c.JSON(http.StatusCreated, response)
}

func UpdateUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
	c.Bind(&user)

	err := database.UpdateUserById(user, uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	message := fmt.Sprintf("succesfully updated user with ID: %d", id)
	response := map[string]interface{}{
		"message": message,
		"code":    http.StatusOK,
	}
	return c.JSON(http.StatusOK, response)
}

func DeleteUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := database.DeleteUserById(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	message := fmt.Sprintf("succesfully deleted user with ID: %d", id)
	response := map[string]interface{}{
		"message": message,
		"code":    http.StatusOK,
	}
	return c.JSON(http.StatusOK, response)
}

func LoginUsers(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	token, err := database.LoginUsers(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := map[string]interface{}{
		"token": token,
	}
	return c.JSON(http.StatusOK, response)
}
