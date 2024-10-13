package location

import (
	"context"
	"post-service/mongodb"
)

type Service interface {
	AddLocation(ctx context.Context, location *Location) error
	GetLocation(ctx context.Context, title string) (*mongodb.Location, error)
}
