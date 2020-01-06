package router

import (
	"github.com/gorilla/mux"
	"github.com/gyaan/short-urls/internal/handler"
)

type router struct {
	h handler.Handler
	r *mux.Router
}

func RegisterRoutes(h handler.Handler, r *mux.Router) {
	routes := router{h: h, r: r}
	routes.r.HandleFunc("/", routes.h.HomeHandler).Methods("GET")
	routes.r.HandleFunc("/short-urls", routes.h.CreateShortUrl).Methods("POST")
	routes.r.HandleFunc("/short-urls", routes.h.GetAllShortUrl).Methods("GET")
	routes.r.HandleFunc("/short-urls/{short_url_id}", routes.h.GetAShortUrl).Methods("GET")
	routes.r.HandleFunc("/short-urls/{short_url_id}", routes.h.UpdateShortUrl).Methods("POST")
	routes.r.HandleFunc("/short-urls/{short_url_id}", routes.h.DeleteShortUrl).Methods("DELETE")
}
