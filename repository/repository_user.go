package repository

import (
	"api_red_social/model"
	"api_red_social/utils/password"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	client *mongo.Client
}

func NewUserRepository(client *mongo.Client) *UserRepository {

	return &UserRepository{
		client: client,
	}
}

func (repo *UserRepository) FindUserByID(id string) (*model.Users, error) {
	collection := repo.client.Database("mydatabase").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user model.Users
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) CreateNewUser(user *model.Users) (primitive.ObjectID, error) {
	hashedPassword, err := password.Hash(user.Password)
	if err != nil {
		return primitive.NilObjectID, err
	}
	user.Password = string(hashedPassword)
	user.User_id = primitive.NewObjectID()
	user.CreatedAt = time.Now()

	collection := repo.client.Database("mydatabase").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("error al insertar el usuario: %w", err)
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (repo *UserRepository) FindUserByEmail(email string) (*model.Users, error) {
	collection := repo.client.Database("mydatabase").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user model.Users
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) CloseConnection() error {
	if err := repo.client.Disconnect(context.Background()); err != nil {
		return fmt.Errorf("error al desconectar de MongoDB: %w", err)
	}
	return nil
}
