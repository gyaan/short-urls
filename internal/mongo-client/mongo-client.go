package mongo_client

import (
	"context"
	"github.com/gyaan/short-urls/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoClient struct {
	conf *config.Config
}

func NewMongoClient(conf *config.Config) *MongoClient {
	return &MongoClient{
		conf: conf,
	}
}

func (m *MongoClient) GetClient() (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.conf.MongoDbConnectionUrl))

	if err != nil {
		log.Fatalf("unable to connect to db with error : %v", err)
		return nil, err
	}
	return client, nil
}
