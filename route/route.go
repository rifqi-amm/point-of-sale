package route

import (
	"pos/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()

	//  routing
	e.GET("/products", controller.GetProductController)
	e.POST("/products", controller.CreatProductController)

	return e
}
