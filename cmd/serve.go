package cmd

import (
	"ecommerce-api/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

	mux := http.NewServeMux()

	//globalRouter := middleware.CorsWithPreflight(mux)
	// wrappedMux := manager.WrapMux(mux,
	// 	middleware.Logger,
	// 	middleware.CorsWithPreflight,
	// )
	manager.Use(
		middleware.CorsWithPreflight,
		middleware.Logger,
	)

	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
