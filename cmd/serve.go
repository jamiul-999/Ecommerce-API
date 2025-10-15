package cmd

import (
	"ecommerce-api/config"
	"ecommerce-api/infra/db"
	"ecommerce-api/repo"
	"ecommerce-api/rest"
	"ecommerce-api/rest/handlers/product"
	"ecommerce-api/rest/handlers/user"
	middleware "ecommerce-api/rest/middlewares"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo(dbCon)

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
