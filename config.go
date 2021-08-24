package main

import (
	"log"
	"os"
	"strconv"
)

func loadConfig() {
	var err error
	TelegramToken = os.Getenv("TELEGRAM_TOKEN")

	if TelegramToken == "" {
		log.Panic("Error: TELEGRAM_TOKEN is not set.")
	}

	channelIdStr := os.Getenv("CHANNEL_ID")

	if channelIdStr == "" {
		log.Panic("Error: CHANNEL_ID is not set.")
	}

	ChannelId, err = strconv.ParseInt(channelIdStr, 10, 64)

	check(err)

	adminGroupIdStr := os.Getenv("ADMIN_GROUP_ID")

	if adminGroupIdStr == "" {
		log.Panic("Error: CHANNEL_ID is not set.")
	}

	AdminGroupId, err = strconv.ParseInt(adminGroupIdStr, 10, 64)

	check(err)

	ReCaptchaConf.SiteKey = os.Getenv("RECAPTCHA_SITE_KEY")

	if ReCaptchaConf.SiteKey == "" {
		log.Panic("Error: RECAPTCHA_SITE_KEY is not set.")
	}

	ReCaptchaConf.Secret = os.Getenv("RECAPTCHA_SECRET")

	if ReCaptchaConf.Secret == "" {
		log.Panic("Error: RECAPTCHA_SECRET is not set.")
	}
}
