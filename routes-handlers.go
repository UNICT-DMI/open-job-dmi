package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

func writeView(w http.ResponseWriter, view string) {
	pathFile := path.Join("./views/", view)
	data, err := ioutil.ReadFile(pathFile)
	check(err)

	w.Write(data)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
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

func offer(w http.ResponseWriter, r *http.Request) {
	var o Offer
	err := json.NewDecoder(r.Body).Decode(&o)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var message string

	message = "*Azienda*: " + o.Azienda + "\n"
	message += "*Email*: " + o.Email + "\n"
	message += "*Ruolo*: " + o.Ruolo + "\n\n"
	message += "*Descrizione*\n" + o.Descrizione + "\n"
	message += "*Competenze Richieste*\n" + o.Competenze + "\n"

	if o.Benefits != "" {
		message += "**Benefits**\n" + o.Benefits + "\n"
	}

	message += "*Disponibilità*: "

	if o.FullTime && o.PartTime {
		message += "Full-Time/Part-time"
	} else if o.FullTime {
		message += "Full-Time"
	} else {
		message += "Part-Time"
	}

	Bot.Send(Channel, message)

	response := "{\"message\": \"success\"}"
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, response)
}
