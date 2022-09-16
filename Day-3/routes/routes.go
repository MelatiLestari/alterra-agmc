package routes

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"Day-2/controllers"
)

func New() *echo.Echo {
	e := echo.New()

	v1 := e.Group("/v1")
	jwtSecret := os.Getenv("JWT_SECRET")

	v1.POST("/login", controllers.LoginUsers)

	v1Users := v1.Group("/users")
	v1Users.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		Skipper: func(c echo.Context) bool {
			fmt.Println("echo url", c.Request().Method)
			return c.Request().Method == "POST"
		},
	}))
	v1Users.GET("", controllers.GetAllUsers)
	v1Users.GET("/:id", controllers.GetUserById)
	v1Users.POST("", controllers.CreateNewUser)
	v1Users.PUT("/:id", controllers.UpdateUserById)
	v1Users.DELETE("/:id", controllers.DeleteUserById)

	v1Books := v1.Group("/books")
	v1Books.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		Skipper: func(c echo.Context) bool {
			fmt.Println("echo url", c.Request().Method)
			return c.Request().Method == "GET"
		},
	}))
	v1Books.GET("", controllers.GetAllBooks)
	v1Books.GET("/:id", controllers.GetBookById)
	v1Books.POST("", controllers.CreateNewBook)
	v1Books.PUT("/:id", controllers.UpdateBookById)
	v1Books.DELETE("/:id", controllers.DeleteBookById)

	return e
}
