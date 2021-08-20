package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func writeView(w http.ResponseWriter, view string) {
	pathFile := path.Join("./views/", view)
	data, err := ioutil.ReadFile(pathFile)
	check(err)

	w.Write(data)
}

func home(w http.ResponseWriter, r *http.Request) {
	writeView(w, "index.html")
}

func about(w http.ResponseWriter, r *http.Request) {
	writeView(w, "about.html")
}

func faq(w http.ResponseWriter, r *http.Request) {
	writeView(w, "faq.html")
}

var Serve http.Handler

func init() {
	r := mux.NewRouter()

	r.Use(loggingMiddleware)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/about", about).Methods("GET")
	r.HandleFunc("/faq", faq).Methods("GET")

	Serve = r
}

func main() {
	log.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", Serve)

	check(err)
}
