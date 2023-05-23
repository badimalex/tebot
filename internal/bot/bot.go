package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	events *Events
}

func NewBot(events *Events) *Bot {
	return &Bot{events: events}
}

func (r *Bot) Start() {
	bot, err := tgbotapi.NewBotAPI("5903480848:AAHL46d6qYLy7N90jt2DECNIVXOHadTL8Js")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		response := r.events.HandleMessage(update.Message.Text, update.Message.Chat.ID)
		msg.Text = response
		bot.Send(msg)
	}
}
