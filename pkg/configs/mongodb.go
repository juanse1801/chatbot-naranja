package configs

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, error) {
	credentials := options.Credential{
		Username: os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PASS"),
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetAuth(credentials))
	if err != nil {
		return nil, errors.New(err.Error())
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	// ping to database
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	fmt.Println("Connected to MongoDB")
	return client, nil
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("chatbot_naranja").Collection(collectionName)

	return collection
}
