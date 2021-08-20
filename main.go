package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/tucnak/telebot.v2"
)

type Offer struct {
	Azienda     string `json:"azienda"`
	Email       string `json:"email"`
	Ruolo       string `json:"ruolo"`
	Descrizione string `json:"descrizione"`
	Competenze  string `json:"competenze"`
	Benefits    string `json:"benefits"`
	FullTime    bool   `json:"fulltime"`
	PartTime    bool   `json:"parttime"`
}

var Serve http.Handler
var Bot *telebot.Bot
var Channel *telebot.Chat

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

func offer(w http.ResponseWriter, r *http.Request) {
	var o Offer
	err := json.NewDecoder(r.Body).Decode(&o)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var message string

	message = "Azienda: " + o.Azienda + "\n"
	message += "Email: " + o.Email + "\n"
	message += "Ruolo: " + o.Ruolo + "\n\n"
	message += "Descrizione\n" + o.Descrizione + "\n"
	message += "Competenze Richieste\n" + o.Competenze + "\n"

	if o.Benefits != "" {
		message += "Benefits\n" + o.Benefits + "\n"
	}

	message += "Disponibilità: "

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

func init() {
	r := mux.NewRouter()

	r.Use(loggingMiddleware)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/about", about).Methods("GET")
	r.HandleFunc("/faq", faq).Methods("GET")
	r.HandleFunc("/api/offer", offer).Methods("POST")

	Serve = r
}

func init() {
	var err error

	token := os.Getenv("TELEGRAM_TOKEN")

	if token == "" {
		panic("Error: TELEGRAM_TOKEN is not set.")
	}

	channel_id_str := os.Getenv("CHANNEL_ID")

	if channel_id_str == "" {
		panic("Error: CHANNEL_ID is not set.")
	}

	var channel_id int64
	channel_id, err = strconv.ParseInt(channel_id_str, 10, 64)

	check(err)

	Channel = &telebot.Chat{ID: channel_id}

	log.Println("[Bot] TELEGRAM_TOKEN: " + token)
	log.Println("[Bot] CHANNEL_ID: " + channel_id_str)

	Bot, err = telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	check(err)

	Bot.Handle("/hello", func(m *telebot.Message) {
		Bot.Send(Channel, "Hello World!")
	})

}

func startBot() {
	log.Println("[Bot] Starting Telegram Bot...")
	Bot.Start()
}

func main() {
	go startBot()

	log.Println("[HTTP] Listening on :8080...")
	err := http.ListenAndServe(":8080", Serve)

	check(err)
}
