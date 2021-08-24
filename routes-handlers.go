package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

func renderView(view string) *template.Template {
	content := path.Join("./views/", view)
	header := "./views/template/header.html"
	footer := "./views/template/footer.html"
	tpl, err := template.ParseFiles(content, header, footer)

	if err != nil {
		panic("Error while rendering pages.")
	}

	return tpl
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func home(w http.ResponseWriter, r *http.Request) {
	tpl := renderView("index.html")
	tpl.ExecuteTemplate(w, "index", map[string]interface{}{"RecaptchaSiteKey": ReCaptchaConf.SiteKey})
}

func about(w http.ResponseWriter, r *http.Request) {
	tpl := renderView("about.html")
	tpl.ExecuteTemplate(w, "about", nil)
}

func faq(w http.ResponseWriter, r *http.Request) {
	tpl := renderView("faq.html")
	tpl.ExecuteTemplate(w, "faq", nil)
}

func offer(w http.ResponseWriter, r *http.Request) {
	var o Offer
	err := json.NewDecoder(r.Body).Decode(&o)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = CheckRecaptcha(ReCaptchaConf.Secret, o.ReCaptchaToken)

	if err != nil {
		log.Println(err)
		return
	}

	var message string

	message = "<b>Azienda</b>: " + o.Azienda
	message += "\n<b>Email</b>: " + o.Email
	message += "\n<b>Sede di lavoro</b>: " + o.Sede
	message += "\n<b>Ruolo</b>: " + o.Ruolo
	message += "\n<b>Salario</b>: " + o.Salario
	message += "\n<b>Disponibilità</b>: "

	if o.FullTime && o.PartTime {
		message += "Full-Time/Part-time"
	} else if o.FullTime {
		message += "Full-Time"
	} else {
		message += "Part-Time"
	}

	message += "\n\n<b>Descrizione</b>\n" + o.Descrizione
	message += "\n\n<b>Competenze Richieste</b>\n" + o.Competenze

	if o.Benefits != "" {
		message += "\n\n<b>Benefits</b>\n" + o.Benefits
	}

	sendOfferToAdminGroup(message)

	response := "{\"message\": \"success\"}"
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, response)
}
