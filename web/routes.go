package web

import (
	"net/http"
	"post-service/web/middlewares"
)

func (server *Server) initRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"GET /hello-world",
		manager.With(
			http.HandlerFunc(server.handlers.Hello),
		),
	)

	mux.Handle(
		"POST /add-location",
		manager.With(
			http.HandlerFunc(server.handlers.AddLocation),
		),
	)

	mux.Handle(
		"GET /get-locations",
		manager.With(
			http.HandlerFunc(server.handlers.GetLocations),
		),
	)
}
