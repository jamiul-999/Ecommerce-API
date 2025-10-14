package cmd

import (
	"ecommerce-api/config"
	"ecommerce-api/repo"
	"ecommerce-api/rest"
	"ecommerce-api/rest/handlers/product"
	"ecommerce-api/rest/handlers/user"
	middleware "ecommerce-api/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
