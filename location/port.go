package location

import "context"


type Service interface {
	AddLocation(ctx context.Context, location *Location) error
}
