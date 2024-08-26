package users

import (
	"api_red_social/internal/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (r Repository) ProfileId(userId string) (id interface{}, err error) {
	collection := r.Client.Database("mydatabase").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user domain.Users

	filter := bson.M{"_id": id}
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
