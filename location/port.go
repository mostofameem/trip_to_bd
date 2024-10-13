package location

import (
	"context"
	"post-service/mongodb"
)

type Service interface {
	AddLocation(ctx context.Context, location *Location) error
	GetLocations(ctx context.Context) (*[]mongodb.Location, error)
}
