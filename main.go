package main

import (
	"log"

	"github.com/badimalex/goshop/config"
	"github.com/badimalex/goshop/pkg/database"

	"github.com/badimalex/tebot/internal/bot"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	bot := bot.NewBot(db)

	bot.Start()
}
