package repository

import (
	"api_red_social/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FollowRepository struct {
	client *mongo.Client
}

func NewFollowRepository(client *mongo.Client) *FollowRepository {
	return &FollowRepository{
		client: client,
	}
}

func (repo *FollowRepository) SaveFollow(follow *model.Follow) (*model.Follow, error) {
	collection := repo.client.Database("mydatabase").Collection("follows")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, follow)
	if err != nil {
		return nil, err
	}

	follow.ID = result.InsertedID.(primitive.ObjectID)
	return follow, nil
}

func (repo *FollowRepository) DeleteFollow(userID, followedID string) error {
	collection := repo.client.Database("mydatabase").Collection("follows")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user": userID, "followed": followedID}

	_, err := collection.DeleteOne(ctx, filter)

	return err
}

func (repo *FollowRepository) FindFollowing(userID string, page int, itemsPerPage int) ([]model.Follow, int, error) {
	collection := repo.client.Database("mydatabase").Collection("follows")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user": userID}
	opts := options.Find().SetSkip(int64((page - 1) * itemsPerPage)).SetLimit(int64(itemsPerPage))

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var follows []model.Follow
	if err = cursor.All(ctx, &follows); err != nil {
		return nil, 0, err
	}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return follows, int(count), nil
}
