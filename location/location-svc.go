package location

import (
	"post-service/config"
	"post-service/db"
	"post-service/grpc/clients"
	"post-service/mongodb"
	"sync"
)

type service struct {
	cnf                 *config.Config
	dblocationTypeRepo  *db.LocationTypeRepo
	mdblocationTypeRepo *mongodb.LocationTypeRepo
	grpcUserClient      *clients.UserClients
}

var location *service
var locationCnt = sync.Once{}

func NewService(cnf *config.Config) Service {
	locationCnt.Do(func() {
		locationTypeRepo := db.NewLocationTypeRepo(&cnf.DB)
		mongoSvc := mongodb.NewLocationTypeRepo(&cnf.MongoDB)
		grpcUserClients := clients.NewUserClient(cnf)
		location = &service{
			cnf:                 cnf,
			dblocationTypeRepo:  locationTypeRepo,
			mdblocationTypeRepo: mongoSvc,
			grpcUserClient:      grpcUserClients,
		}
	})
	return location
}
