package cmd

import (
	"ecommerce-api/config"
	"ecommerce-api/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Serve() {
	cnf := config.GetConfig()
	manager := middleware.NewManager()
	manager.Use(
		middleware.PreFlight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("Server running on port", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting server", err)
		os.Exit(1)
	}
}
