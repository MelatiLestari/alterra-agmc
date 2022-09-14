package controllers

import (
	"Day-2/config"
	"Day-2/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// for _, user := range users {
	// 	if user.ID == uint(id) {
	// 		response := map[string]interface{}{
	// 			"message": "success",
	// 			"code":    http.StatusOK,
	// 			"data":    user,
	// 		}
	// 		return c.JSON(http.StatusOK, response)
	// 	}
	// }
	user := models.User{}
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := map[string]interface{}{
		"message": "user not found",
		"code":    http.StatusBadRequest,
	}
	return c.JSON(http.StatusBadRequest, response)

}
