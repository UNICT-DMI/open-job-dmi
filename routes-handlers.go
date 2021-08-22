package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
)

func renderView(w http.ResponseWriter, r *http.Request, view string) {
	pathFile := path.Join("./views/", view)
	http.ServeFile(w, r, pathFile)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func home(w http.ResponseWriter, r *http.Request) {
	renderView(w, r, "index.html")
}

func about(w http.ResponseWriter, r *http.Request) {
	renderView(w, r, "about.html")
}

func faq(w http.ResponseWriter, r *http.Request) {
	renderView(w, r, "faq.html")
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

	message = "*Azienda*: " + o.Azienda
	message += "\n*Email*: " + o.Email
	message += "\n*Ruolo*: " + o.Ruolo
	message += "\n*Disponibilità*: "

	if o.FullTime && o.PartTime {
		message += "Full-Time/Part-time"
	} else if o.FullTime {
		message += "Full-Time"
	} else {
		message += "Part-Time"
	}

	message += "\n\n*Descrizione*\n" + o.Descrizione
	message += "\n\n*Competenze Richieste*\n" + o.Competenze

	if o.Benefits != "" {
		message += "\n\n**Benefits**\n" + o.Benefits
	}

	sendOfferToAdminGroup(message)

	response := "{\"message\": \"success\"}"
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, response)
}
