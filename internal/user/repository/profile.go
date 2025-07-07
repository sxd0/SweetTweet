package repository

import (
	"context"
	"time"

	"github.com/sxd0/SweetTweet/internal/user/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileRepository struct {
	collection *mongo.Collection
}

func NewProfileRepository(client *mongo.Client) *ProfileRepository {
	coll := client.Database("sweetweet").Collection("profiles")
	return &ProfileRepository{collection: coll}
}

func (r *ProfileRepository) Create(ctx context.Context, profile *model.UserProfile) error {
	profile.CreatedAt = time.Now().Unix()
	profile.UpdatedAt = profile.CreatedAt
	_, err := r.collection.InsertOne(ctx, profile)
	return err
}

func (r *ProfileRepository) GetByUserID(ctx context.Context, userID int64) (*model.UserProfile, error) {
	var profile model.UserProfile
	err := r.collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *ProfileRepository) Update(ctx context.Context, userID int64, update bson.M) error {
	update["updated_at"] = time.Now().Unix()
	_, err := r.collection.UpdateOne(ctx, bson.M{"user_id": userID}, bson.M{"$set": update})
	return err
}
