package handlers

import (
	"post-service/config"
	"post-service/location"
)

type Handlers struct {
	cnf    *config.Config
	locSvc location.Service
}

func NewHandlers(cnf *config.Config, locSvc location.Service) *Handlers {
	return &Handlers{
		cnf:    cnf,
		locSvc: locSvc,
	}
}
