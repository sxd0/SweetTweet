package service

import (
	"context"
	"time"

	"github.com/sxd0/SweetTweet/internal/user/model"
	"github.com/sxd0/SweetTweet/internal/user/repository"
	pb "github.com/sxd0/SweetTweet/proto/userpb"
	"go.mongodb.org/mongo-driver/bson"
)

type ProfileService struct {
	Repo *repository.ProfileRepository
}

func NewProfileService(repo *repository.ProfileRepository) *ProfileService {
	return &ProfileService{Repo: repo}
}

func (s *ProfileService) CreateProfile(ctx context.Context, req *pb.CreateProfileRequest) (*pb.Profile, error) {
	now := time.Now().Unix()
	profile := model.UserProfile{
		UserID:    req.UserId,
		Username:  req.Username,
		Bio:       req.Bio,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.Repo.Create(ctx, &profile); err != nil {
		return nil, err
	}

	return toProto(&profile), nil
}

func (s *ProfileService) GetProfile(ctx context.Context, req *pb.ProfileRequest) (*pb.Profile, error) {
	profile, err := s.Repo.GetByUserID(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return toProto(profile), nil
}

func (s *ProfileService) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.Profile, error) {
	now := time.Now().Unix()

	profile, err := s.Repo.GetByUserID(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	profile.Bio = req.Bio
	profile.UpdatedAt = now

	if err := s.Repo.Update(ctx, profile.UserID, bson.M{
		"bio":        profile.Bio,
		"updated_at": profile.UpdatedAt,
	}); err != nil {
		return nil, err
	}
	return toProto(profile), nil
}

func toProto(p *model.UserProfile) *pb.Profile {
	return &pb.Profile{
		Id:        p.ID.Hex(),
		UserId:    p.UserID,
		Username:  p.Username,
		Bio:       p.Bio,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

