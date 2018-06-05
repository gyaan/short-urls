package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

func AllShortUrls(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "still need to implement")
}

func GetShortUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "still need to implement")
}

func CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "still need to implement")
}

func UpdateShortUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "still need to implement")
}

func DeleteShortUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "still need to implement")
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/short-urls", AllShortUrls).Methods("GET")
	r.HandleFunc("/short-urls", CreateShortUrl).Methods("POST")
	r.HandleFunc("/short-urls/{id}", UpdateShortUrl).Methods("PUT")
	r.HandleFunc("/short-urls/{id}", GetShortUrl).Methods("GET")
	r.HandleFunc("/short-urls/{id}", DeleteShortUrl).Methods("Delete")

	if err := http.ListenAndServe(":1334", r);

		err != nil {
		log.Fatal(err)
	}

}
