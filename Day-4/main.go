package main

import (
	"Day-2/config"
	m "Day-2/middleware"
	"Day-2/routes"
)

func init() {
	config.InitDB()
	config.InitialMigration()
}

func main() {
	e := routes.New()

	m.LogMiddleware(e)

	e.Start(":8080")
	e.Logger.Fatal(e.Start(":8080"))
}
