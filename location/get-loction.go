package location

import (
	"context"
	"post-service/mongodb"
)

func (svc *service) GetLocation(ctx context.Context, title string) (*mongodb.Location, error) {
	locationId, err := svc.dblocationTypeRepo.GetLocationID(ctx, title)
	if err != nil {
		return nil, err
	}

	locations, err := svc.mdblocationTypeRepo.GetLocation(ctx, locationId)
	if err != nil {
		return nil, err
	}

	return locations, nil
}
