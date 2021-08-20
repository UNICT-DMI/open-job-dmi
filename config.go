package main

import "os"

func loadConfig() {
	token = os.Getenv("TELEGRAM_TOKEN")

	if token == "" {
		panic("Error: TELEGRAM_TOKEN is not set.")
	}

	channel_id_str = os.Getenv("CHANNEL_ID")

	if channel_id_str == "" {
		panic("Error: CHANNEL_ID is not set.")
	}
}
