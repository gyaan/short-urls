package database

import (
	"context"
	"fmt"
	"github.com/gyaan/short-urls/internal/models"
	"github.com/gyaan/short-urls/internal/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type database struct {
	MongoClient *mongo.Client
}

type Database interface {
	CreateShortUrl(ctx context.Context, url models.ShortUrl) error
	GetAllShortUrls(ctx context.Context) ([]models.ShortUrl, error)
	UpdateShortUrls(ctx context.Context, shortUrl models.ShortUrl) error
}

func NewDatabase(client *mongo.Client) Database {
	return &database{
		MongoClient: client,
	}
}

func (db *database) CreateShortUrl(ctx context.Context, url models.ShortUrl) error {

	collection := db.MongoClient.Database("my_project").Collection("short_urls")

	//todo move this to some where else before creating it
	shortUrlService := services.NewShortUrlService()
	shortUrl := shortUrlService.GetShortUrl(url.Url)
	res, err := collection.InsertOne(ctx, bson.M{"actual_url": url.Url, "short_url": shortUrl})
	fmt.Print(res.InsertedID)

	if err != nil {
		log.Print("error creating short urls")
		return err
	}

	return nil
}

func (db *database) GetAllShortUrls(ctx context.Context) ([]models.ShortUrl, error) {
	collection := db.MongoClient.Database("my_project").Collection("short_urls")
	ctx1, _ := context.WithTimeout(context.Background(), 30*time.Second)

	var res []models.ShortUrl

	cur, err := collection.Find(ctx1, bson.D{})

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, models.ShortUrl{
			Url:    result["actual_url"].(string),
			NewUrl: result["short_url"].(string),
		})
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return res, nil
}

func (db *database) UpdateShortUrls(ctx context.Context, shortUrl models.ShortUrl) error {
	collection := db.MongoClient.Database("my_project").Collection("short_urls")
	ctx1, _ := context.WithTimeout(context.Background(), 30*time.Second)

	filter := bson.D{{"actual_url", shortUrl.Url}}
	_, err := collection.UpdateOne(ctx1, filter, shortUrl)

	return err

}
