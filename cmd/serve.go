package cmd

import (
	"ecommerce-api/config"
	"ecommerce-api/rest"
	"ecommerce-api/rest/handlers/product"
	"ecommerce-api/rest/handlers/review"
	"ecommerce-api/rest/handlers/user"
	middleware "ecommerce-api/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()
	middlewares := middleware.NewMiddlewares(cnf)
	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler()
	reviewHandler := review.NewHandler()
	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
		reviewHandler,
	)
	server.Start()
}
