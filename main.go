package main

import (
	"log"
	"net/http"

	"gopkg.in/tucnak/telebot.v2"
)

var Serve http.Handler
var Bot *telebot.Bot
var TelegramToken string
var ChannelId int64
var AdminGroupId int64
var Channel *telebot.Chat
var AdminGroup *telebot.Chat

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
