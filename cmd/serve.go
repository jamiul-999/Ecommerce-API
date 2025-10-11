package cmd

import (
	"ecommerce-api/config"
	"ecommerce-api/rest"
	"ecommerce-api/rest/handlers/product"
	"ecommerce-api/rest/handlers/user"
)

func Serve() {
	cnf := config.GetConfig()
	productHandler := product.NewHandler()
	userHandler := user.NewHandler()
	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
