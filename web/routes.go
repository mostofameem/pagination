package web

import (
	"net/http"
	"products/web/handlers"
	"products/web/middlewares"
)

func InitRouts(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(handlers.List),
		),
	)
}
