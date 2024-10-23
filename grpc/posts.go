package grpc

import (
	pb "post-service/grpc/posts"
	"post-service/location"
)

type PostsService struct {
	pb.UnimplementedPostServiceServer
	locSvc location.Service
}

func NewPostsService(locSvc location.Service) *PostsService {
	return &PostsService{
		locSvc: locSvc,
	}
}
