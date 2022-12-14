package routes

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"Day-2/controllers"
)

func New() *echo.Echo {
	e := echo.New()

	v1 := e.Group("/v1")
	jwtSecret := os.Getenv("JWT_SECRET")
	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(jwtSecret),
	})

	v1.POST("/login", controllers.LoginUsers)

	v1Users := v1.Group("/users")
	v1Users.GET("", controllers.GetAllUsers, jwtMiddleware)
	v1Users.GET("/:id", controllers.GetUserById, jwtMiddleware)
	v1Users.POST("", controllers.CreateNewUser)
	v1Users.PUT("/:id", controllers.UpdateUserById, jwtMiddleware)
	v1Users.DELETE("/:id", controllers.DeleteUserById, jwtMiddleware)

	v1Books := v1.Group("/books")
	v1Books.GET("", controllers.GetAllBooks, jwtMiddleware)
	v1Books.GET("/:id", controllers.GetBookById)
	v1Books.POST("", controllers.CreateNewBook, jwtMiddleware)
	v1Books.PUT("/:id", controllers.UpdateBookById, jwtMiddleware)
	v1Books.DELETE("/:id", controllers.DeleteBookById, jwtMiddleware)

	return e
}
