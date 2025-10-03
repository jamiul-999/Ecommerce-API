package cmd

import (
	"ecommerce-api/global_router"
	"ecommerce-api/handlers"
	"fmt"
	"net/http"
)

func Serve() {

	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))

	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))

	mux.Handle("GET /products/{productID}", http.HandlerFunc(handlers.GetProductByID))

	fmt.Println("Server running on :8080")

	globalRouter := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
