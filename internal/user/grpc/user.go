package grpc

import (
	"context"

	"github.com/sxd0/SweetTweet/internal/user/service"
	pb "github.com/sxd0/SweetTweet/proto/userpb"
)

type GRPCServer struct {
	pb.UnimplementedUserServiceServer
	Service *service.ProfileService
}

func NewGRPCServer(s *service.ProfileService) *GRPCServer {
	return &GRPCServer{Service: s}
}

func (g *GRPCServer) CreateProfile(ctx context.Context, req *pb.CreateProfileRequest) (*pb.Profile, error) {
	return g.Service.CreateProfile(ctx, req)
}

func (g *GRPCServer) GetProfile(ctx context.Context, req *pb.ProfileRequest) (*pb.Profile, error) {
	return g.Service.GetProfile(ctx, req)
}

func (g *GRPCServer) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.Profile, error) {
	return g.Service.UpdateProfile(ctx, req)
}
