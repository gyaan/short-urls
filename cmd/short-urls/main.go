package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/gyaan/short-urls/internal/config"
	"github.com/gyaan/short-urls/internal/handler"
	mclient "github.com/gyaan/short-urls/internal/mongo_client"
	"github.com/gyaan/short-urls/internal/repositories"
	"github.com/gyaan/short-urls/internal/router"
	"log"
	"net/http"
)

func main() {

	//create config
	conf := config.New()

	//mongodb client
	mClient, err := mclient.New(conf.MongoDbConnectionUrl, conf.MongoContextTimeout)

	if err != nil {
		log.Fatalf("Error connecting mongodb")
	}

	//get repositories
	counterRepository := repositories.NewCounterRepository(mClient, conf)
	shortUrlRepository := repositories.NewShortUrlRepository(mClient, counterRepository, conf)
	userRepository := repositories.NewUserRepository(mClient, conf)

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

	fmt.Println("short url application started and served at port:", conf.ApplicationPort)
}
