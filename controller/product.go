package controller

import (
	"net/http"
	"pos/lib/database"
	"pos/model"

	"github.com/labstack/echo/v4"
)

func GetAllProductsController(c echo.Context) error {
	products := database.GetProducts()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetAllProductsController",
		"data":    products,
	})
}

func GetProductByIDController(c echo.Context) error {
	id := c.Param("id")
	product := database.GetProductByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetProductByIDController",
		"data":    product,
	})
}

func DeleteProductByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteProductByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeleteProductByIDController",
	})
}

func UpdateProductByIDController(c echo.Context) error {
	id := c.Param("id")

	var product model.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateProductController",
			"error":   err.Error(),
		})
	}
	database.UpdateProductByID(id, product)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetProductByIDController",
		"data":    product,
	})
}

func CreateProductController(c echo.Context) error {
	var newProduct model.Product
	if err := c.Bind(&newProduct); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateProductController",
			"error":   err.Error(),
		})
	}

	newProduct = database.CreateProduct(newProduct)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateProductController",
		"data":    newProduct,
	})
}
