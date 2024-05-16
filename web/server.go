package web

import (
	"net/http"
	"products/web/middlewares"
)

func StartServer() *http.ServeMux {
	manager := middlewares.NewManager()
	mux := http.NewServeMux()
	InitRouts(mux, manager)
	return mux
}
