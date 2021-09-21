package route

import (
	"pos/controller"

	"github.com/labstack/echo/v4"
)

func NewProduct(app *echo.Echo) {
	app.GET("/products", controller.GetAllProductsController)
	app.POST("/products", controller.CreateProductController)
	app.GET("/products/:id", controller.GetProductByIDController)
	app.DELETE("/products/:id", controller.DeleteProductByIDController)
	app.PUT("/products/:id", controller.UpdateProductByIDController)
}