package repositories

import (
	"context"
	"fmt"
	"github.com/gyaan/short-urls/internal/config"
	"github.com/gyaan/short-urls/internal/models"
	"github.com/gyaan/short-urls/pkg/url_shortner"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type shortUrls struct {
	mongoClient       *mongo.Client
	counterRepository Counters
	conf              *config.Config
}

//ShortUrls
type ShortUrls interface {
	CreateShortUrl(ctx context.Context, urlString string) (*models.ShortUrl, error)
	GetAShortUrl(ctx context.Context, shortUrlId string) (*models.ShortUrl, error)
	GetAllShortUrls(ctx context.Context, page int, limit int) ([]models.ShortUrl, error)
	UpdateShortUrls(ctx context.Context, shortUrlId string, url models.ShortUrl) error
	DeleteShortUrl(ctx context.Context, shortUrlId string) error
	GetActualUrlOfAShortUrl(ctx context.Context, shortUrl string) (*models.ShortUrl, error)
	IncrementClickCountOfShortUrl(ctx context.Context, shortUrl string) error
	GetTotalShortUrlsCount(ctx context.Context) (int64, error)
}

// NewShortUrlRepository creates new repositories for short urls
func NewShortUrlRepository(client *mongo.Client, counterRepository Counters, config2 *config.Config) ShortUrls {
	return &shortUrls{
		mongoClient:       client,
		counterRepository: counterRepository,
		conf:              config2,
	}
}

//CreateShortUrl creates new short urls
func (s *shortUrls) CreateShortUrl(ctx context.Context, urlString string) (*models.ShortUrl, error) {

	var srtUrl models.ShortUrl
	collection := s.mongoClient.Database(s.conf.MongoDatabaseName).Collection("short_urls")
	urlIdentifier, err := s.counterRepository.UpdateAndGetCounter(ctx, "url_identifier")

	log.Printf("New url identifier %d", urlIdentifier)
	if err != nil {
		log.Printf("Error with getting current sequence of url_identifier")
		return nil, err
	}

	shortUrl := url_shortner.New().GetShortUrl(urlIdentifier)

	srtUrl.ID = primitive.NewObjectIDFromTimestamp(time.Now())
	srtUrl.Url = urlString
	srtUrl.NewUrl = shortUrl
	srtUrl.UrlIdentifier = urlIdentifier
	srtUrl.Status = 1 //default status active
	srtUrl.CreatedAt = time.Now()
	srtUrl.ExpireTime = time.Now().Add(time.Duration(s.conf.ShortUrlExpiryTime) * time.Hour)
	srtUrl.CreatedBy = fmt.Sprintf("%v", ctx.Value("user_id"))

	res, err := collection.InsertOne(ctx, srtUrl)
	if err != nil {
		log.Println("Error in creating short url", srtUrl)
		return nil, err
	}
	log.Println("short url created successfully with id ", res.InsertedID)
	return &srtUrl, nil
}

//GetAllShortUrls returns all short urls
func (s *shortUrls) GetAllShortUrls(ctx context.Context, offset int, limit int) ([]models.ShortUrl, error) {
	log.Println("Get all short urls")

	var res []models.ShortUrl
	collection := s.mongoClient.Database(s.conf.MongoDatabaseName).Collection("short_urls")
	ctx1, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(offset))
	cur, err := collection.Find(ctx1, bson.D{{"created_by", ctx.Value("user_id").(string)}}, findOptions)
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
	return res, nil
}

//UpdateShortUrls update existing short url
func (s *shortUrls) UpdateShortUrls(ctx context.Context, shortUrlId string, shortUrl models.ShortUrl) error {
	collection := s.mongoClient.Database(s.conf.MongoDatabaseName).Collection("short_urls")
	ctx1, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(shortUrlId)
	if err != nil {
		log.Printf("Error creating mongo object id for %s", shortUrlId)
	}

	//updating status only as of now
	filter := bson.D{{"_id", objectId}, {"created_by", ctx.Value("user_id")}}
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
	collection := s.mongoClient.Database(s.conf.MongoDatabaseName).Collection("short_urls")
	ctx1, cancel := context.WithTimeout(ctx, 10*time.Second)
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

//DeleteShortUrl delete a short url
func (s *shortUrls) DeleteShortUrl(ctx context.Context, srtUrlId string) error {
	collection := s.mongoClient.Database(s.conf.MongoDatabaseName).Collection("short_urls")
	ctx1, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(srtUrlId)

	if err != nil {
		log.Printf("Error creating mongo object id for %s", srtUrlId)
		return err
	}

	filter := bson.D{{"_id", id}, {"created_by", ctx.Value("user_id")}}
	res, err := collection.DeleteOne(ctx1, filter)
	if err != nil {
		log.Printf("Error in deleting short url details for short url id %s", srtUrlId)
		return err
	}

	log.Printf("successfully deleted short url count %d", res.DeletedCount)
	return nil
}

//GetActualUrlOfAShortUrl get short url details from url identifier
func (s *shortUrls) GetActualUrlOfAShortUrl(ctx context.Context, shortUrl string) (*models.ShortUrl, error) {

	var srtUrl models.ShortUrl
	collection := s.mongoClient.Database(s.conf.MongoDatabaseName).Collection("short_urls")
	ctx1, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	//get url identifier
	identifierNumber := url_shortner.New().GetIdentifierNumberFromShortUrl(shortUrl)
	filter := bson.D{{"url_identifier", identifierNumber}}
	err := collection.FindOne(ctx1, filter).Decode(&srtUrl)

	if err != nil {
		log.Printf("Error in fecthing short url details of url identifier %d", identifierNumber)
		return nil, err
	}

	return &srtUrl, nil
}

//IncrementClickCountOfShortUrl increase clicks count for a url after clicking it
func (s *shortUrls) IncrementClickCountOfShortUrl(ctx context.Context, shortUrl string) error {
	collection := s.mongoClient.Database(s.conf.MongoDatabaseName).Collection("short_urls")
	ctx1, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	//updating status only as of now
	filter := bson.D{{"new_url", shortUrl}}
	res, err := collection.UpdateOne(ctx1, filter, bson.D{{"$inc", bson.D{{"clicks_count", 1}}}})

	if err != nil {
		log.Printf("Error increasing click count for short url id %s", shortUrl)
		return err
	}

	log.Printf("successfully increased short urls clicks count %d", res.MatchedCount)
	return nil
}

//GetTotalShortUrlsCount returns count for a user
func (s *shortUrls) GetTotalShortUrlsCount(ctx context.Context) (int64, error) {

	log.Println("Get short urls count for a user")
	collection := s.mongoClient.Database(s.conf.MongoDatabaseName).Collection("short_urls")
	ctx1, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(ctx1, bson.D{{"created_by", ctx.Value("user_id").(string)}})

	if err != nil {
		return 0, err
	}

	return count, nil
}
