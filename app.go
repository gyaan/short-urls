package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"short-urls/model"
	dao2 "short-urls/dao"
	config2 "short-urls/config"
)

var dao = dao2.ShortUrlsDao{}
var config = config2.Config{}

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

	if err:= json.NewDecoder(r.Body).Decode(&shortUrl); err != nil{
		respondWithError(w,http.StatusBadRequest,"Invalid request payload")
	    return
	}

	if err:= dao.Delete(shortUrl); err!=nil{
		respondWithError(w,http.StatusInternalServerError,err.Error())
        return
	}

	respondWithJson(w, http.StatusOK,map[string]string{"result":"success"})
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
	r.HandleFunc("/short-urls", AllShortUrls).Methods("GET")
	r.HandleFunc("/short-urls", CreateShortUrl).Methods("POST")
	r.HandleFunc("/short-urls", UpdateShortUrl).Methods("PUT")
	r.HandleFunc("/short-urls/{id}", GetShortUrl).Methods("GET")
	r.HandleFunc("/short-urls", DeleteShortUrl).Methods("Delete")

	if err := http.ListenAndServe(":1334", r);

		err != nil {
		log.Fatal(err)
	}

}
