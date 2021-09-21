package main

import (
	"pos/config"
	"pos/route"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()
	config.InitMigration()

	app := echo.New()
	route.NewProduct(app)
	app.Start(":8080")
}