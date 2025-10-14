package product

import (
	"ecommerce-api/repo"
	middleware "ecommerce-api/rest/middlewares"
)

type Handler struct {
	middlewares *middleware.Middlewares
	productRepo repo.ProductRepo
}

func NewHandler(
	middlewares *middleware.Middlewares,
	productRepo repo.ProductRepo,
) *Handler {
	return &Handler{
		middlewares: middlewares,
		productRepo: productRepo,
	}
}
