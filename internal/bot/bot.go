package bot

import (
	"database/sql"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	db *sql.DB
}

func NewBot(db *sql.DB) *Bot {
	return &Bot{db: db}
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

		switch update.Message.Text {
		case "/start":
			findOrCreateChatByID(r.db, update.Message.Chat.ID)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.Text = "Привет, напишите название продукта который вы хотите искать:"
			bot.Send(msg)
		// case "yes":
		// 	msg.Text = "Great! Do you want to search for another product?"
		// case "no":
		// 	if len(userProductMap[update.Message.Chat.ID]) > 0 {
		// 		msg.Text = "Here are the products you are tracking: " + strings.Join(userProductMap[update.Message.Chat.ID], ", ")
		// 	} else {
		// 		msg.Text = "Do you want to search for a product?"
		// 	}
		default:
			// 	userProductMap[update.Message.Chat.ID] = append(userProductMap[update.Message.Chat.ID], update.Message.Text)
			// 	msg.Text = "Do you want to track when new " + update.Message.Text + " products appear?"
		}
	}
}

type Chat struct {
	ID int64
}

func findOrCreateChatByID(db *sql.DB, chatID int64) (*Chat, error) {
	chat := &Chat{}
	err := db.QueryRow("SELECT chat_id FROM chats WHERE chat_id = $1", chatID).Scan(&chat.ID)

	if err == sql.ErrNoRows {
		err = db.QueryRow("INSERT INTO chats (chat_id) VALUES ($1) RETURNING chat_id", chatID).Scan(&chat.ID)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	return chat, nil
}
