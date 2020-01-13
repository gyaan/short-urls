package repositories

import (
	"context"
	"github.com/gyaan/short-urls/config"
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
}

type Counters interface {
	UpdateAndGetCounter(ctx context.Context, counter string) (int64,error)
}

func NewCounterRepository(client *mongo.Client) Counters {
	return &counters{mongoClient: client}
}

//UpdateCounter increase sequence of a counter
func (c counters) UpdateAndGetCounter(ctx context.Context, counterStr string) (int64,error) {
	collection := c.mongoClient.Database(config.GetConf().MongoDatabaseName).Collection("counters")
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
		 counter.Sequence = int64(config.GetConf().MinimumShortUrlIdentifier)
		 counter.Name = counterStr
		 _, err := collection.InsertOne(ctx1, counter)
		 if err != nil {
			 log.Printf("Error setting initial count for %s", counterStr)
			 return 0, err
		 }
	}
	log.Printf("Successfully updated counter for %s", counterStr)
	return counter.Sequence,nil
}