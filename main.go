package main

import (
	"net/http"
	"pos/config"
	"pos/routes"

	"github.com/labstack/echo/v4"
)

const(
	username = "rifqi"
	password = "1233"
)

 func BasicAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	 return func(c echo.Context) error{
		usernameFromHeader := c.Request().Header.Get("username")
		passwordFromHeader := c.Request().Header.Get("password")
		if username == usernameFromHeader && password == passwordFromHeader{
			return next(c)
		}
		return c.String(http.StatusForbidden, "username atau pasword salah")
	 }
 }
 
func main() {
	config.InitDB()
	config.InitMigration()

	app := echo.New()
	// app.Use(middleware.Logger())
	
	routes.NewProduct(app)
	app.Start(":8080")
}