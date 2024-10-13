package location

import (
	"context"
	"post-service/mongodb"
)

func (svc *service) GetLocations(ctx context.Context) (*[]mongodb.Location, error) {
	locations, err := svc.mdblocationTypeRepo.GetLocations(ctx)
	if err != nil {
		return nil, err
	}
	
	return locations, nil
}
