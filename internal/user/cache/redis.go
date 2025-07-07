package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sxd0/SweetTweet/internal/user/model"
)

type Cache struct {
	client *redis.Client
}

func NewCache(addr string) *Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	return &Cache{client: rdb}
}

func (c *Cache) GetProfile(ctx context.Context, userID int64) (*model.UserProfile, error) {
	key := fmt.Sprintf("profile:%d", userID)
	data, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var profile model.UserProfile
	err = json.Unmarshal([]byte(data), &profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (c *Cache) SetProfile(ctx context.Context, profile *model.UserProfile) error {
	key := fmt.Sprintf("profile:%d", profile.UserID)
	data, err := json.Marshal(profile)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, key, data, 10*time.Minute).Err()
}
