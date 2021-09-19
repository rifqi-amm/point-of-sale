package controller

import (
	"net/http"
	"pos/config"
	"pos/model"

	"github.com/labstack/echo/v4"
)

func GetProductController(c echo.Context) error {
	var products []model.Product

	err := config.DB.Find(&products).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil menampilkan",
		"data":    products,
	})

}

func CreatProductController(c echo.Context) error {
	// email := c.FormValue("email")
	// name := c.FormValue("name")

	product := model.Product{}
	c.Bind(&product)

	err := config.DB.Save(&product).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "product berhasil ditambahkan",
		"product":    product,
	})
}