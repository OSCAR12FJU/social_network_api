package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Publication struct {
	Publication_id primitive.ObjectID `bson:"publication_id,omitempty"`
	Content        string             `bson:"content"`
	Follower       string             `bson:"follower"`
	CreatedAt      time.Time          `bson:"created_at"`
}
