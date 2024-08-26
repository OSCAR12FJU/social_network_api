package follows

import (
	"api_red_social/internal/domain"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r Repository) Insert(follow domain.Follows) (id interface{}, err error) {
	collection := r.Client.Database("mydatabase").Collection("follows")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, follow)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error inseting follower: %w", err)
	}
	return result.InsertedID.(primitive.ObjectID), nil

}
