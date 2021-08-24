package main

import (
	"log"
	"strings"
	"time"

	"gopkg.in/tucnak/telebot.v2"
)

var selector *telebot.ReplyMarkup
var approve telebot.Btn
var disapprove telebot.Btn

func initTelegramBot() {
	var err error
	Channel = &telebot.Chat{ID: ChannelId}
	AdminGroup = &telebot.Chat{ID: AdminGroupId}

	Bot, err = telebot.NewBot(telebot.Settings{
		Token:     TelegramToken,
		ParseMode: telebot.ModeHTML,
		Poller:    &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	check(err)

	selector = &telebot.ReplyMarkup{}
	approve = selector.Data("👍🏻", "approve")
	disapprove = selector.Data("👎🏻", "disapprove")

	selector.Inline(
		selector.Row(approve, disapprove),
	)

	Bot.Handle(&approve, func(c *telebot.Callback) {
		c.Message.Text = strings.Replace(c.Message.Text, "Azienda:", "<b>Azienda</b>:", -1)
		c.Message.Text = strings.Replace(c.Message.Text, "Email:", "<b>Email</b>:", -1)
		c.Message.Text = strings.Replace(c.Message.Text, "Sede di Lavoro:", "<b>Sede di Lavoro</b>:", -1)
		c.Message.Text = strings.Replace(c.Message.Text, "Ruolo:", "<b>Ruolo</b>:", -1)
		c.Message.Text = strings.Replace(c.Message.Text, "Salario:", "<b>Salario</b>:", -1)
		c.Message.Text = strings.Replace(c.Message.Text, "Descrizione", "<b>Descrizione</b>", -1)
		c.Message.Text = strings.Replace(c.Message.Text, "Competenze Richieste", "<b>Competenze Richieste</b>", -1)
		c.Message.Text = strings.Replace(c.Message.Text, "Benefits", "<b>Benefits</b>", -1)
		c.Message.Text = strings.Replace(c.Message.Text, "Disponibilità:", "<b>Disponibilità</b>:", -1)

		sendMessageToChannel(c.Message.Text)

		Bot.Edit(c.Message, c.Message.Text)
		message := "Approvato da " + c.Sender.Username
		Bot.Reply(c.Message, message)
	})

	Bot.Handle(&disapprove, func(c *telebot.Callback) {
		Bot.Edit(c.Message, c.Message.Text)
		message := "Rifiutato da " + c.Sender.Username
		Bot.Reply(c.Message, message)
	})
}

func startTelegramBot() {
	log.Println("[Bot] Starting Telegram Bot...")
	Bot.Start()
}

func sendOfferToAdminGroup(message string) {
	Bot.Send(AdminGroup, message, selector)
}

func sendMessageToChannel(message string) {
	Bot.Send(Channel, message)
}
