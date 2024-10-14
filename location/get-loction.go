package location

import (
	"context"
	"post-service/db"
	"post-service/mongodb"
	"post-service/web/utils"
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

func (svc *service) GetLocations(ctx context.Context, filter utils.PaginationParams) (*[]db.Location, error) {
	locations, err := svc.dblocationTypeRepo.GetLocations(ctx, filter)
	if err != nil {
		return nil, err
	}

	return locations, nil
}
