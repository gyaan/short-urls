package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	config2 "github.com/gyaan/short-urls/config"
	dao2 "github.com/gyaan/short-urls/dao"
	"github.com/gyaan/short-urls/model"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

var dao = dao2.ShortUrlsDao{}
var config = config2.Config{}
var userDao = dao2.UsersDao{}

func AllShortUrls(w http.ResponseWriter, r *http.Request) {

	shortUrls, err := dao.FindAll()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, shortUrls)
}

func GetShortUrl(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	shortUrl, err := dao.FindById(params["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Short url id")
		return
	}

	respondWithJson(w, http.StatusOK, shortUrl)
}

func CreateShortUrl(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	var shortUrl model.ShortUrl

	if err := json.NewDecoder(r.Body).Decode(&shortUrl); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	shortUrl.Id = bson.NewObjectId()

	if err := dao.Insert(shortUrl);

		err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusCreated, shortUrl)
}

func UpdateShortUrl(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var shortUrl model.ShortUrl

	if err := json.NewDecoder(r.Body).Decode(&shortUrl); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := dao.Update(shortUrl); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteShortUrl(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var shortUrl model.ShortUrl

	if err := json.NewDecoder(r.Body).Decode(&shortUrl); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := dao.Delete(shortUrl); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

/**
  Function for users
 */
func GetUser(writer http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	shortUrl, err := userDao.FindById(params["id"])
	if err != nil {
		respondWithError(writer, http.StatusBadRequest, "Invalid user id")
		return
	}
	respondWithJson(writer, http.StatusOK, shortUrl)
}

func CreateUser(writer http.ResponseWriter, request *http.Request) {

	defer request.Body.Close()
	var user model.User

	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		respondWithError(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user.Id = bson.NewObjectId()
	if err := userDao.Insert(user);
		err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(writer, http.StatusCreated, user)
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	var user model.User
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		respondWithError(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := userDao.Update(user); err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(writer, http.StatusOK, map[string]string{"result": "success"})
}

func GetAllUsers(writer http.ResponseWriter, request *http.Request) {

	users, err := userDao.FindAll()

	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(writer, http.StatusOK, users)
}



func SingUp(writer http.ResponseWriter, request *http.Request) {

	defer request.Body.Close()
	var user model.User

	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		respondWithError(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user.Id = bson.NewObjectId()
	if err := userDao.Insert(user);
		err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	//todo create the jwt token and send it back in the response

	respondWithJson(writer, http.StatusCreated, user)

}

//similar to get user details
func SignIn(writer http.ResponseWriter, request *http.Request) {

	//here will get the customer details using email and password
	params := mux.Vars(request)
	shortUrl, err := userDao.FindById(params["id"])
	if err != nil {
		respondWithError(writer, http.StatusBadRequest, "Invalid user id")
		return
	}
	respondWithJson(writer, http.StatusOK, shortUrl)

}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {

	r := mux.NewRouter()

	/**
	 Routes for short urls
	 */
	r.HandleFunc("/short-urls", AllShortUrls).Methods("GET")
	r.HandleFunc("/short-urls", CreateShortUrl).Methods("POST")
	r.HandleFunc("/short-urls", UpdateShortUrl).Methods("PUT")
	r.HandleFunc("/short-urls/{id}", GetShortUrl).Methods("GET")
	r.HandleFunc("/short-urls", DeleteShortUrl).Methods("Delete")

	/**
	Routes for user
	 */

	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users",GetAllUsers).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users", UpdateUser).Methods("PUT")

	/**
	Routes for authentication
	 */

	//simillar to user create
	r.HandleFunc("/sign-up", SingUp).Methods("POST")
	r.HandleFunc("/sign-in", SignIn).Methods("POST")

	if err := http.ListenAndServe(":1334", r);

		err != nil {
		log.Fatal(err)
	}

}

