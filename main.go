package main

import (
	"pos/config"
	"pos/route"
)

func main() {
	config.InitDB()

	e := route.New()
	e.Start(":8000")
}