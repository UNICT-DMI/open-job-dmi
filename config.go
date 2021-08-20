package main

import (
	"os"
	"strconv"
)

func loadConfig() {
	var err error
	TelegramToken = os.Getenv("TELEGRAM_TOKEN")

	if TelegramToken == "" {
		panic("Error: TELEGRAM_TOKEN is not set.")
	}

	channelIdStr := os.Getenv("CHANNEL_ID")

	if channelIdStr == "" {
		panic("Error: CHANNEL_ID is not set.")
	}

	ChannelId, err = strconv.ParseInt(channelIdStr, 10, 64)

	check(err)

	adminGroupIdStr := os.Getenv("ADMIN_GROUP_ID")

	if adminGroupIdStr == "" {
		panic("Error: CHANNEL_ID is not set.")
	}

	AdminGroupId, err = strconv.ParseInt(adminGroupIdStr, 10, 64)

	check(err)
}
