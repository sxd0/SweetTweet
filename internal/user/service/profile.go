package service

import (
	"context"
	"errors"

	"github.com/sxd0/SweetTweet/internal/user/cache"
	"github.com/sxd0/SweetTweet/internal/user/model"
	"github.com/sxd0/SweetTweet/internal/user/repository"
)

type ProfileService struct {
	repo  *repository.ProfileRepository
	cache *cache.Cache
}

func NewProfileService(repo *repository.ProfileRepository, cache *cache.Cache) *ProfileService {
	return &ProfileService{repo: repo, cache: cache}
}

func (s *ProfileService) Create(ctx context.Context, profile *model.UserProfile) error {
	err := s.repo.Create(ctx, profile)
	if err != nil {
		return err
	}
	_ = s.cache.SetProfile(ctx, profile)
	return nil
}

func (s *ProfileService) GetByUserID(ctx context.Context, userID int64) (*model.UserProfile, error) {
	profile, err := s.cache.GetProfile(ctx, userID)
	if err == nil {
		return profile, nil
	}
	profile, err = s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	_ = s.cache.SetProfile(ctx, profile)
	return profile, nil
}

func (s *ProfileService) Update(ctx context.Context, userID int64, update map[string]interface{}) error {
	if len(update) == 0 {
		return errors.New("empty update")
	}
	err := s.repo.Update(ctx, userID, update)
	if err != nil {
		return err
	}
	profile, err := s.repo.GetByUserID(ctx, userID)
	if err == nil {
		_ = s.cache.SetProfile(ctx, profile)
	}
	return nil
}
