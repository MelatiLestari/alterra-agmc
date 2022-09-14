package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// func main() {
// 	// dsn := "host=127.0.0.1 user=postgres password=admin dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
// 	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// type Recipe struct {
// 	// 	gorm.Model
// 	// 	recipe_id   int
// 	// 	recipe_name string
// 	// }

// 	// var recipes []Recipe

// 	// if result := db.Raw("SELECT * FROM recipes").Scan(&recipes); result.Error != nil {
// 	// 	fmt.Println(result.Error)
// 	// }

// 	// fmt.Println("result", result)

// 	e := echo.New()
// 	e.GET("/", func(c echo.Context) error {
// 		return c.JSON(http.StatusOK, map[string]interface{}{
// 			"message": "success get all recipes",
// 			"recipes": recipes,
// 		})
// 	})
// 	e.Logger.Fatal(e.Start(":8080"))
// }

func main() {
	e := echo.New()

	v1 := e.Group("/v1")

	v1Books := v1.Group("/books")
	v1Users := v1.Group("/users")

	v1Users.GET("/:id", getUserById)

	v1Books.GET("", getAllBooks)
	v1Books.GET("/:id", getBookById)
	v1Books.POST("", createNewBook)
	v1Books.PUT("/:id", updateBookById)
	v1Books.DELETE("/:id", deleteBookById)

	e.Logger.Fatal(e.Start(":8080"))
}

type User struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Email: "melati@falcon.com"},
	{ID: 2, Email: "jason@falcon.com"},
}

func getUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, user := range users {
		if user.ID == uint(id) {
			response := map[string]interface{}{
				"message": "success",
				"code":    http.StatusOK,
				"data":    user,
			}
			return c.JSON(http.StatusOK, response)
		}
	}

	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "user not found",
		"code":    http.StatusBadRequest,
	})

}

type Book struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: 1, Title: "Into The Woods", Author: "Jack"},
	{ID: 2, Title: "Breaking Back", Author: "Simon"},
}

func getBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, book := range books {
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

func getAllBooks(c echo.Context) error {
	response := map[string]interface{}{
		"message": "success",
		"code":    http.StatusOK,
		"data":    books,
	}
	return c.JSON(http.StatusOK, response)
}

func createNewBook(c echo.Context) error {
	book := Book{}
	err := json.NewDecoder(c.Request().Body).Decode(&book)
	if err != nil {
		response := map[string]interface{}{
			"message": "failed to create book",
			"code":    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	books = append(books, book)
	fmt.Println("Current Repo:", books)
	response := map[string]interface{}{
		"message": "success",
		"code":    http.StatusCreated,
	}
	return c.JSON(http.StatusCreated, response)
}

func updateBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	bookPayload := Book{}
	err := json.NewDecoder(c.Request().Body).Decode(&bookPayload)
	if err != nil {
		response := map[string]interface{}{
			"message": "failed to update book",
			"code":    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	for i, book := range books {
		if book.ID == uint(id) {
			books[i] = bookPayload
			response := map[string]interface{}{
				"message": "success",
				"code":    http.StatusOK,
			}
			fmt.Println("Current Repo:", books)
			return c.JSON(http.StatusOK, response)
		}
	}

	response := map[string]interface{}{
		"message": "failed to update book",
		"code":    http.StatusBadRequest,
	}
	return c.JSON(http.StatusBadRequest, response)
}

func deleteBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, book := range books {
		if book.ID == uint(id) {
			books[i] = books[0]
			books = books[1:]
			response := map[string]interface{}{
				"message": "success",
				"code":    http.StatusOK,
			}
			fmt.Println("Current Repo:", books)
			return c.JSON(http.StatusOK, response)
		}
	}

	response := map[string]interface{}{
		"message": "failed to delete book",
		"code":    http.StatusBadRequest,
	}
	return c.JSON(http.StatusBadRequest, response)
}
