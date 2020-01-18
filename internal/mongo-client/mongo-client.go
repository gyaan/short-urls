package mongo_client

import (
	"context"
	"github.com/gyaan/short-urls/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var client *mongo.Client

type MongoClient struct{}

//NewMongoClient
func New() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.GetConf().MongoContextTimeout)*time.Second)
	defer cancel()

	c, err := mongo.Connect(ctx, options.Client().ApplyURI(config.GetConf().MongoDbConnectionUrl))
	if err != nil {
		log.Fatalf("unable to connect to db with error : %v", err)
	}
	client = c
	return client
}

//GetClient returns mongo client
func GetClient() *mongo.Client {
	return client
}