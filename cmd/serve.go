package cmd

import (
	"ecommerce-api/config"
	"ecommerce-api/infra/db"
	"ecommerce-api/repo"
	"ecommerce-api/rest"
	prdctHandler "ecommerce-api/rest/handlers/product"
	usrHandler "ecommerce-api/rest/handlers/user"
	middleware "ecommerce-api/rest/middlewares"
	"ecommerce-api/user"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	//domains
	usrSvc := user.NewService(userRepo)
	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := prdctHandler.NewHandler(middlewares, productRepo)
	userHandler := usrHandler.NewHandler(cnf, usrSvc)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
