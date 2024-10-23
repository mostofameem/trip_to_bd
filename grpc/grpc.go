package grpc

import (
	"post-service/config"
	"post-service/location"
	"sync"
)

type grpc struct {
	cnf      *config.Config
	locSvc location.Service
	Wg       sync.WaitGroup
}

func NewGRPC(cnf *config.Config, locSvc location.Service) *grpc {
	return &grpc{
		cnf:      cnf,
		locSvc: locSvc,
	}
}
