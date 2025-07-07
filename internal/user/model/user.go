package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserProfile struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    UserID    int64              `bson:"user_id" json:"user_id"`
    Username  string             `bson:"username" json:"username"`
    Bio       string             `bson:"bio,omitempty" json:"bio,omitempty"`
    AvatarURL string             `bson:"avatar_url,omitempty" json:"avatar_url,omitempty"`
    CreatedAt int64              `bson:"created_at" json:"created_at"`
    UpdatedAt int64              `bson:"updated_at" json:"updated_at"`
}
