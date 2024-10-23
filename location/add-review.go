package location

import (
	"context"
	"log/slog"
	"post-service/grpc/users"
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

func (svc *service) AddReviews(ctx context.Context, locationId int, cmnt Comment, id int) error {
	userInfo, err := svc.grpcUserClient.GetUserName(ctx, &users.GetUserNameReq{
		UserId: int32(id),
	})
	if err != nil {
		slog.Error("failed to get username from grpc")
		return err
	}

	err = svc.mdblocationTypeRepo.AddReviews(ctx, locationId, mongodb.Comment{
		Userid:     cmnt.Userid,
		UserName:   userInfo.Name,
		Content:    cmnt.Content,
		Created_at: time.Now(),
		Updated_at: time.Now(),
		Vote:       0,
	})
	return err
}
