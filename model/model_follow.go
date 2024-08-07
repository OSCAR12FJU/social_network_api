package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Follow struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    string             `bson:"user_id"`
	FollowID  string             `bson:"follow_id"`
	CreatedAt time.Time          `bson:"created_at"`
}
