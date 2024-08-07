package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateConnection() (*mongo.Client, error) {

	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error al cargar el archivo .env: %w", err)
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		return nil, fmt.Errorf("MONGO_URI no esta definido en el archivo .env")
	}
	fmt.Println("URI de MongoDB", mongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error al conectar a MongoDB: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("error al hacer ping a MongoDB: %w", err)
	}

	fmt.Println("Conectado a MongoDB!")
	return client, nil
}
