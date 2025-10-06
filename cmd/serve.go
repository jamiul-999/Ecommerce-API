package cmd

import (
	"ecommerce-api/config"
	"ecommerce-api/rest"
)

func Serve() {
	cnf := config.GetConfig()
	rest.Start(cnf)
}
