package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func initRouter() {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/about", about).Methods("GET")
	r.HandleFunc("/faq", faq).Methods("GET")
	r.HandleFunc("/api/offer", offer).Methods("POST")

	Serve = r
}
