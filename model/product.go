package model

type Product struct {
	ID        int    `json: "id" form:"id"`
	Name      string `json: "name" form:"name"`
	HargaBeli int    `json: "hargabeli" form:"hargabeli"`
	HargaJual int    `json: "hargajual" form:"hargajual"`
	Stock     int    `json: "stock" form: "stock"`
}