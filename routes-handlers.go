package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

func renderView(w http.ResponseWriter, view string) {
	viewPath := view + ".html"
	content := path.Join("./views/", viewPath)
	header := "./views/template/header.html"
	footer := "./views/template/footer.html"
	tpl, err := template.ParseFiles(content, header, footer)

	if err != nil {
		panic("Error while rendering pages.")
	}

	tpl.ExecuteTemplate(w, view, nil)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func home(w http.ResponseWriter, r *http.Request) {
	renderView(w, "index")
}

func about(w http.ResponseWriter, r *http.Request) {
	renderView(w, "about")
}

func faq(w http.ResponseWriter, r *http.Request) {
	renderView(w, "faq")
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
