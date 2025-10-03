package cmd

import (
	"ecommerce-api/global_router"
	"ecommerce-api/handlers"
	"ecommerce-api/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

	mux := http.NewServeMux()

	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(handlers.GetProducts),
		middleware.Logger,
	))

	mux.Handle("POST /products", manager.With(
		http.HandlerFunc(handlers.CreateProduct),
		middleware.Logger,
	))

	mux.Handle("GET /products/{productID}", manager.With(
		http.HandlerFunc(handlers.GetProductByID),
		middleware.Logger,
	))

	fmt.Println("Server running on :8080")

	globalRouter := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
