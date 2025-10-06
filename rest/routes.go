package rest

import (
	"ecommerce-api/rest/handlers"
	middleware "ecommerce-api/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(handlers.GetProducts),
	),
	)

	mux.Handle("POST /products", manager.With(
		http.HandlerFunc(handlers.CreateProduct),
	),
	)

	mux.Handle("GET /products/{productID}", manager.With(
		http.HandlerFunc(handlers.GetProductByID),
	),
	)

}
