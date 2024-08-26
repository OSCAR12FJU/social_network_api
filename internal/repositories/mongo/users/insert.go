package users

import (
	"api_red_social/internal/domain"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r Repository) Insert(user domain.Users) (id interface{}, err error) {

	collection := r.Client.Database("mydatabase").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error inseting follower: %w", err)
	}
	return result.InsertedID.(primitive.ObjectID), nil

}

// func (repo *UserRepository) CreateNewUser(user *model.Users) (primitive.ObjectID, error) {
// 	hashedPassword, err := password.Hash(user.Password)
// 	if err != nil {
// 		return primitive.NilObjectID, err
// 	}
// 	user.Password = string(hashedPassword)
// 	user.User_id = primitive.NewObjectID()
// 	user.CreatedAt = time.Now()

// 	collection := repo.client.Database("mydatabase").Collection("users")
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	result, err := collection.InsertOne(ctx, user)
// 	if err != nil {
// 		return primitive.NilObjectID, fmt.Errorf("error al insertar el usuario: %w", err)
// 	}

// 	return result.InsertedID.(primitive.ObjectID), nil
// }
