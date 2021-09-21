package database

import (
	"pos/config"
	"pos/model"
)

func GetProducts() []model.Product {
	var products []model.Product
	config.DB.Find(&products)
	return products
}

func GetProductByID(id string) model.Product {
	var product model.Product
	config.DB.Where("id = ?", id).Find(&product)
	return product
}

func CreateProduct(product model.Product) model.Product {
	config.DB.Create(&product)
	return product
}

func DeleteProductByID(id string) {
	var product model.Product
	config.DB.Where("id = ?", id).Delete(&product)
}

func UpdateProductByID(id string, product model.Product) {
	config.DB.Where("id = ?", id).Updates(&product)
}
