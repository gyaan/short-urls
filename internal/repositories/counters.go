package repositories

import (
	"context"
	"github.com/gyaan/short-urls/internal/config"
	"github.com/gyaan/short-urls/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type counters struct {
	mongoClient *mongo.Client
	conf        *config.Config
}

//Counters
type Counters interface {
	UpdateAndGetCounter(ctx context.Context, counter string) (int64, error)
}

//NewCounterRepository
func NewCounterRepository(client *mongo.Client, config2 *config.Config) Counters {
	return &counters{mongoClient: client, conf: config2}
}

//UpdateCounter increase sequence of a counter
func (c counters) UpdateAndGetCounter(ctx context.Context, counterStr string) (int64, error) {
	collection := c.mongoClient.Database(c.conf.MongoDatabaseName).Collection("counters")
	ctx1, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var counter models.Counter

	opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.D{{"name", counterStr}}
	update := primitive.D{{"$inc", primitive.D{{"sequence", 1}}}}

	err := collection.FindOneAndUpdate(ctx1, filter, update, opts).Decode(&counter)

	if err != nil {
		//still no row in the counter collection
		//set initial sequence
		counter.Sequence = int64(c.conf.MinimumShortUrlIdentifier)
		counter.Name = counterStr
		counter.ID = primitive.NewObjectIDFromTimestamp(time.Now())
		_, err := collection.InsertOne(ctx1, counter)
		if err != nil {
			log.Printf("Error setting initial count for %s", counterStr)
			return 0, err
		}
	}
	log.Printf("Successfully updated counter for %s", counterStr)
	return counter.Sequence, nil
}