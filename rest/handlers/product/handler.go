package product

import middleware "ecommerce-api/rest/middlewares"

type Handler struct {
	middlewares *middleware.Middlewares
}

func NewHandler(middlewares *middleware.Middlewares) *Handler {
	return &Handler{
		middlewares: middlewares,
	}
}
