package mongo_client

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var client *mongo.Client

type MongoClient struct{}

// NewMongoClient
func New(mongoDbConnectionUrl string, mongoContextTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoContextTimeout)*time.Second)
	defer cancel()

	c, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDbConnectionUrl))
	if err != nil {
		log.Printf("unable to connect to db with error : %v", err)
		return nil, err

	}
	client = c
	return client, nil
}

// GetClient returns mongo client
func GetClient() *mongo.Client {
	return client
}
