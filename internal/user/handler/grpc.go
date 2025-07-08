package handler

import (
	"context"

	"github.com/sxd0/SweetTweet/internal/user/service"
	"github.com/sxd0/SweetTweet/proto/userpb"
)

type GRPCHandler struct {
	userpb.UnimplementedUserServiceServer
	svc *service.ProfileService
}

func NewGRPCHandler(svc *service.ProfileService) *GRPCHandler {
	return &GRPCHandler{svc: svc}
}

func (h *GRPCHandler) CreateProfile(ctx context.Context, req *userpb.CreateProfileRequest) (*userpb.Profile, error) {
	return h.svc.CreateProfile(ctx, req)
}

func (h *GRPCHandler) GetProfile(ctx context.Context, req *userpb.ProfileRequest) (*userpb.Profile, error) {
	return h.svc.GetProfile(ctx, req)
}

func (h *GRPCHandler) UpdateProfile(ctx context.Context, req *userpb.UpdateProfileRequest) (*userpb.Profile, error) {
	return h.svc.UpdateProfile(ctx, req)
}
