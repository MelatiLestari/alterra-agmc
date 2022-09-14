package routes

import (
	"github.com/labstack/echo/v4"

	"Day-2/controllers"
)

func New() *echo.Echo {
	e := echo.New()

	v1 := e.Group("/v1")

	v1Books := v1.Group("/books")
	v1Users := v1.Group("/users")

	v1Users.GET("", controllers.GetAllUsers)
	v1Users.GET("/:id", controllers.GetUserById)
	v1Users.POST("", controllers.CreateNewUser)
	v1Users.PUT("/:id", controllers.UpdateUserById)
	v1Users.DELETE("/:id", controllers.DeleteUserById)

	v1Books.GET("", controllers.GetAllBooks)
	v1Books.GET("/:id", controllers.GetBookById)
	v1Books.POST("", controllers.CreateNewBook)
	v1Books.PUT("/:id", controllers.UpdateBookById)
	v1Books.DELETE("/:id", controllers.DeleteBookById)

	return e
}
