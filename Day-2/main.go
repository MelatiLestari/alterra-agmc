package main

import (
	"Day-2/config"
	"Day-2/routes"
)

// func main() {
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

func init() {
	config.InitDB()
	config.InitialMigration()
}

func main() {
	e := routes.New()

	e.Logger.Fatal(e.Start(":8080"))
}

// type User struct {
// 	ID    uint   `json:"id"`
// 	Email string `json:"email"`
// }

// var users = []User{
// 	{ID: 1, Email: "melati@falcon.com"},
// 	{ID: 2, Email: "jason@falcon.com"},
// }
