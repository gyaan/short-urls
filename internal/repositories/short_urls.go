package repositories

import (
"context"
"github.com/gyaan/short-urls/config"
"github.com/gyaan/short-urls/internal/models"
"github.com/gyaan/short-urls/pkg/url_shortner"
"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/bson/primitive"
"go.mongodb.org/mongo-driver/mongo"
"log"
"math/rand"
"time"
)

type shortUrls struct {
	mongoClient *mongo.Client
}

type ShortUrls interface {
	CreateShortUrl(ctx context.Context, urlString string) (*models.ShortUrl, error)
	GetAShortUrl(ctx context.Context, shortUrlId string) (*models.ShortUrl, error)
	GetAllShortUrls(ctx context.Context) ([]models.ShortUrl, error)
	UpdateShortUrls(ctx context.Context, shortUrlId string, url models.ShortUrl) error
	DeleteShortUrl(ctx context.Context, shortUrlId string) error
}

// NewShortUrlRepository creates new repositories for short urls
func NewShortUrlRepository(client *mongo.Client) ShortUrls {
	return &shortUrls{
		mongoClient: client,
	}
}

//CreateShortUrl creates new short urls
func (s *shortUrls) CreateShortUrl(ctx context.Context, urlString string) (*models.ShortUrl, error) {

	var srtUrl models.ShortUrl
	collection := s.mongoClient.Database(config.GetConf().MongoDatabaseName).Collection("short_urls")
	randomNumber := rand.Int() //todo get this number with better way
	shortUrl := url_shortner.New().GetShortUrl(randomNumber)

	srtUrl.Url = urlString
	srtUrl.NewUrl = shortUrl
	srtUrl.Status = 1 //default status active
	srtUrl.CreatedAt = time.Now()
	srtUrl.ExpireTime = time.Now().Add(time.Duration(config.GetConf().ShortUrlExpiryTime) * time.Hour)
	srtUrl.CreatedBy = "" //todo get the user id from token

	res, err := collection.InsertOne(ctx, srtUrl)
	if err != nil {
		log.Println("Error in creating short url", srtUrl)
		return nil, err
	}
	log.Println("short url created successfully with id ", res.InsertedID)

	return &srtUrl, nil
}

//GetAllShortUrls returns all short urls
func (s *shortUrls) GetAllShortUrls(ctx context.Context) ([]models.ShortUrl, error) {
	var res []models.ShortUrl
	collection := s.mongoClient.Database(config.GetConf().MongoDatabaseName).Collection("short_urls")
	ctx1, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//todo add created by user where condition
	cur, err := collection.Find(ctx1, bson.D{})
	if err != nil {
		log.Println("Error fetching all short urls")
		return res, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result models.ShortUrl
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Error iterating cursor for all short urls")
			return res, err
		}
		res = append(res, result)
	}
	if err := cur.Err(); err != nil {
		log.Println("Error get all short urls cursor")
		return res, err
	}

	//todo add paginated response
	return res, nil
}

//UpdateShortUrls update existing short url
func (s *shortUrls) UpdateShortUrls(ctx context.Context, shortUrlId string, shortUrl models.ShortUrl) error {
	collection := s.mongoClient.Database(config.GetConf().MongoDatabaseName).Collection("short_urls")
	ctx1, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(shortUrlId)
	if err != nil {
		log.Printf("Error creating mongo object id for %s", shortUrlId)
	}

	//updating status only as of now
	filter := bson.D{{"_id", objectId}}
	res, err := collection.UpdateOne(ctx1, filter, bson.D{{"$set", bson.D{{"status", shortUrl.Status}}}})

	if err != nil {
		log.Printf("Error updating short url for short url id %s", shortUrlId)
		return err
	}

	log.Printf("successfully updated short urls count %d", res.MatchedCount)
	return nil
}

//GetAShortUrl get single existing short url
func (s *shortUrls) GetAShortUrl(ctx context.Context, srtUrlId string) (*models.ShortUrl, error) {
	collection := s.mongoClient.Database(config.GetConf().MongoDatabaseName).Collection("short_urls")
	ctx1, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var result models.ShortUrl
	id, err := primitive.ObjectIDFromHex(srtUrlId)

	if err != nil {
		log.Printf("Error creating mongo object id for %s", srtUrlId)
		return nil, err
	}

	filter := bson.D{{"_id", id}}
	err = collection.FindOne(ctx1, filter).Decode(&result)

	if err != nil {
		log.Printf("Error in fetching short url details for short url id %s", srtUrlId)
		return nil, err
	}

	return &result, err
}

func (s *shortUrls) DeleteShortUrl(ctx context.Context, srtUrlId string) error {
	collection := s.mongoClient.Database(config.GetConf().MongoDatabaseName).Collection("short_urls")
	ctx1, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(srtUrlId)

	if err != nil {
		log.Printf("Error creating mongo object id for %s", srtUrlId)
		return err
	}

	filter := bson.D{{"_id", id}}
	res, err := collection.DeleteOne(ctx1, filter)
	if err != nil {
		log.Printf("Error in deleting short url details for short url id %s", srtUrlId)
		return err
	}

	log.Printf("successfully deleted short url count %d", res.DeletedCount)
	return nil
}
