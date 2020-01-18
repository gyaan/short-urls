package repositories

import (
	"context"
	"fmt"
	"github.com/gyaan/short-urls/internal/config"
	"github.com/gyaan/short-urls/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type users struct {
	mongoClient *mongo.Client
}

//Users
type Users interface {
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	UpdateUser(ctx context.Context, userId string, user models.User) error
	GetUserDetailsByName(ctx context.Context, name string) (*models.User, error)
	GetUserDetailsById(ctx context.Context, name string) (*models.User, error)
}

//NewUserRepository
func NewUserRepository(client *mongo.Client) Users {
	return &users{
		mongoClient: client,
	}
}

//RegisterUser creates new user
func (u *users) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	collection := u.mongoClient.Database(config.GetConf().MongoDatabaseName).Collection("users")
	res, err := collection.InsertOne(ctx, user)

	if err != nil {
		log.Printf("Error in user creation")
		return nil, err
	}
	fmt.Printf("successfully user created with id %v", res.InsertedID)
	return &user, nil
}

//UpdateUser updates a user
func (u *users) UpdateUser(ctx context.Context, userId string, user models.User) error {

	collection := u.mongoClient.Database(config.GetConf().MongoDatabaseName).Collection("users")
	ctx1, _ := context.WithTimeout(context.Background(), 30*time.Second)

	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		fmt.Printf("Error getting mongo object id for user id %s", userId)
		return err
	}

	filter := bson.D{{"_id", id}}
	res, err := collection.UpdateOne(ctx1, filter, bson.D{{
		"$set", bson.D{
			{"email", user.Email},
			{"password", user.Password},
			{"status", user.Status},
		},
	}})
	if err != nil {
		fmt.Printf("Error updating user details for user id %s", userId)
		return err
	}

	log.Printf("Successfully update user details for user id %s, total updated records %d", userId, res.UpsertedCount)
	return nil
}

//GetUserDetails get user details for a name
func (u *users) GetUserDetailsByName(ctx context.Context, name string) (*models.User, error) {
	collection := u.mongoClient.Database(config.GetConf().MongoDatabaseName).Collection("users")
	ctx1, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var user models.User

	filter := bson.D{{"name", name}}
	err := collection.FindOne(ctx1, filter).Decode(&user)
	if err != nil {
		log.Printf("Error getting user details")
		log.Print(err)
		return nil, err
	}
	return &user, nil
}

//GetUserDetails get user details by id
func (u *users) GetUserDetailsById(ctx context.Context, userId string) (*models.User, error) {
	collection := u.mongoClient.Database(config.GetConf().MongoDatabaseName).Collection("users")
	ctx1, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var user models.User

	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", id}}
	err = collection.FindOne(ctx1, filter).Decode(&user)
	if err != nil {
		log.Printf("Error getting user details %v", err)
		return nil, err
	}
	return &user, nil
}
