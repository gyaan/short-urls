package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/gyaan/short-urls/config"
	"github.com/gyaan/short-urls/internal/handler"
	mclient "github.com/gyaan/short-urls/internal/mongo-client"
	"github.com/gyaan/short-urls/internal/repositories"
	"github.com/gyaan/short-urls/internal/router"
	"log"
	"net/http"
)

func main() {

	//create config
	conf := config.NewConfig(config.Config{
		ApplicationPort:           ":1334",
		MongoDbConnectionUrl:      "mongodb://localhost:27017",
		MongoDatabaseName:         "my_project",
		ShortUrlExpiryTime:        3 * 24, //three days
		BaseUrl:                   "http://localhsot:1334/",
		MinimumShortUrlIdentifier: 1,
		JWTSecret:                 "+,g~9Ywa8)7D<nbR",
		TokenExpiryTime:           5 * 24 * 60 * 60, //5 days
	})

	//mongodb repositories client
	mClient, err := mclient.NewMongoClient(conf).GetClient()

	if err != nil {
		log.Fatalf("not able to connect to repositories %v", err)
	}

	//get repositories
	shortUrlRepository := repositories.NewShortUrlRepository(mClient)
	userRepository := repositories.NewUserRepository(mClient)

	//get handler
	h := handler.NewHandler(shortUrlRepository, userRepository)

	//create routes
	r := chi.NewRouter()

	//register routes to handle request
	router.RegisterRoutes(h, r)

	//start server
	err = http.ListenAndServe(conf.ApplicationPort, r)

	if err != nil {
		fmt.Println("Error:", err.Error())
	}
}
