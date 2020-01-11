package repositories

import (
	"context"
	"fmt"
	"github.com/gyaan/short-urls/config"
	"github.com/gyaan/short-urls/internal/models"
	"go.mongodb.org/mongo-driver/bson"
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
	UpdateUser(ctx context.Context, user models.User) (*models.User, error)
	DeleteUser(ctx context.Context, user models.User) (*models.User, error)
}

//NewUserRepository
func NewUserRepository(client *mongo.Client) Users {
	return &users{
		mongoClient: client,
	}
}

//CreateUser creates new user
func (u *users) CreateUser(ctx context.Context, user models.User) (*models.User, error) {

	collection := u.mongoClient.Database(config.GetConf().MongoDatabaseName).Collection("users")

	res, err := collection.InsertOne(ctx, bson.M{"name": user.Name, "email": user.Email, "password": user.Password, "status": user.Status})
	fmt.Print(res.InsertedID)

	if err != nil {
		log.Print("error creating short urls")
		return nil, err
	}

	return nil, nil
}

//UpdateUser updates a user
func (u *users) UpdateUser(ctx context.Context, user models.User) (*models.User, error) {

	collection := u.mongoClient.Database(config.GetConf().MongoDatabaseName).Collection("users")
	ctx1, _ := context.WithTimeout(context.Background(), 30*time.Second)

	filter := bson.D{{"email", user.Email}}
	_, err := collection.UpdateOne(ctx1, filter, user)

	return nil, err
}

//DeleteUser deletes a user
func (u *users) DeleteUser(ctx context.Context, user models.User) (*models.User, error) {
	panic("implement me")
}
