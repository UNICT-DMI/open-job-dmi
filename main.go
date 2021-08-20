package main

import (
	"log"
	"net/http"

	"gopkg.in/tucnak/telebot.v2"
)

var Serve http.Handler
var Bot *telebot.Bot
var Channel *telebot.Chat
var token string
var channel_id_str string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	initRouter()
	loadConfig()
	initTelegramBot()
}

func main() {
	go startTelegramBot()

	log.Println("[HTTP] Listening on :8080...")
	err := http.ListenAndServe(":8080", Serve)

	check(err)
}
