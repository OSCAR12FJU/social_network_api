package publications

import (
	"api_red_social/internal/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (r Repository) FindByID(publiId string) (id interface{}, err error) {
	collection := r.Client.Database("mydatabase").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var publication domain.Publication

	filter := bson.M{"_id": publiId}
	err = collection.FindOne(ctx, filter).Decode(&publication)
	if err != nil {
		return nil, err
	}
	return &publication, nil
}

// func (repo *PublicationRepository) FindByID(id string) (*model.Publication, error) {
// 	collection := repo.client.Database("mydatabase").Collection("publication")
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	objectId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var publication model.Publication
// 	filter := bson.M{"_id": objectId}
// 	err = collection.FindOne(ctx, filter).Decode(&publication)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}

// 	return &publication, nil
// }
