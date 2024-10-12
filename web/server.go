package web

import (
	"fmt"
	"log/slog"
	"net/http"
	"post-service/config"
	"post-service/web/handlers"
	"post-service/web/middlewares"
	"sync"
)

type Server struct {
	handlers *handlers.Handlers
	cnf      *config.Config
	Wg       sync.WaitGroup
}

func NewServer(cnf *config.Config, handlers *handlers.Handlers) *Server {
	server := &Server{
		cnf:      cnf,
		handlers: handlers,
	}
	return server
}
func (server *Server)Run(){
	server.Start()
}


func (server *Server) Start() {
	manager := middlewares.NewManager()

	mux := http.NewServeMux()

	server.initRoutes(mux, manager)

	handler := middlewares.EnableCors(mux)

	server.Wg.Add(1)

	go func() {
		defer server.Wg.Done()
		conf := config.GetConfig()

		addr := fmt.Sprintf(":%d", conf.HttpPort)

		slog.Info(fmt.Sprintf("Listening at %s", addr))

		if err := http.ListenAndServe(addr, handler); err != nil {
			slog.Error(err.Error())
		}
	}()

}
