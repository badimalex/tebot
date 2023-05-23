package bot

import (
	"database/sql"
)

type Storage struct {
	db *sql.DB
}

func New(db *sql.DB) *Storage {
	return &Storage{db: db}
}

type Chat struct {
	ID int64
}

func (c *Storage) FindOrCreateChatByID(chatID int64) (*Chat, error) {
	chat := &Chat{}
	err := c.db.QueryRow("SELECT chat_id FROM chats WHERE chat_id = $1", chatID).Scan(&chat.ID)

	if err == sql.ErrNoRows {
		err = c.db.QueryRow("INSERT INTO chats (chat_id) VALUES ($1) RETURNING chat_id", chatID).Scan(&chat.ID)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	return chat, nil
}

func (c *Storage) AddItemSearchInTable(message string, chatID int64) error {
	_, err := c.db.Exec("INSERT INTO searches (name, track, chat_id) VALUES ($1, $2, $3) ", message, false, chatID)
	return err
}

// Следующее задание перепесать plain sql на orm https://github.com/go-pg/pg

func (c *Storage) Subscribe(chatID int64) error {
	_, err := c.db.Exec("update searches set track = $1 where id in (select id from searches where chat_id = $2 order by id desc limit 1)", true, chatID)
	return err
}
