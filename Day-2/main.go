package main

import (
	"Day-2/config"
	"Day-2/routes"
)

func init() {
	config.InitDB()
	config.InitialMigration()
}

func main() {
	e := routes.New()

	e.Logger.Fatal(e.Start(":8080"))
}
