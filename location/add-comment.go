package location

import (
	"context"
	"post-service/mongodb"
	"time"
)

type Comment struct {
	Userid     int
	UserName   string
	Content    string
	Created_at time.Time
	Updated_at time.Time
	Vote       int
}

func (svc *service) AddReviews(ctx context.Context, locationId int, cmnt Comment) error {
	err := svc.mdblocationTypeRepo.AddReviews(ctx, locationId, mongodb.Comment{
		Userid:     cmnt.Userid,
		UserName:   cmnt.UserName,
		Content:    cmnt.Content,
		Created_at: time.Now(),
		Updated_at: time.Now(),
		Vote:       0,
	})
	return err
}
