package main

import (
	"log"
	"strconv"
	"time"

	"gopkg.in/tucnak/telebot.v2"
)

func initTelegramBot() {
	var err error
	var channel_id int64
	channel_id, err = strconv.ParseInt(channel_id_str, 10, 64)

	check(err)

	Channel = &telebot.Chat{ID: channel_id}

	log.Println("[Bot] TELEGRAM_TOKEN: " + token)
	log.Println("[Bot] CHANNEL_ID: " + channel_id_str)

	Bot, err = telebot.NewBot(telebot.Settings{
		Token:     token,
		ParseMode: telebot.ModeMarkdown,
		Poller:    &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	check(err)
}

func startTelegramBot() {
	log.Println("[Bot] Starting Telegram Bot...")
	Bot.Start()
}
