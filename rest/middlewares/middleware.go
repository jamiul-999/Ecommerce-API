package middleware

import "ecommerce-api/config"

type Middlewares struct {
	cnf *config.Config
}

func NewMiddlewares(cnf *config.Config) *Middlewares {
	return &Middlewares{
		cnf: cnf,
	}
}
