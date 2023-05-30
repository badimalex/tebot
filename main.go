package main

import (
	"fmt"
	"log"

	"github.com/badimalex/goshop/config"
	"github.com/badimalex/goshop/pkg/database"

	"github.com/badimalex/tebot/internal/bot"
	"github.com/badimalex/tebot/pkg/searches"
)

func main() {
	qwe := searches.SearchOnEbay("retro", 55)
	for _, items := range qwe {
		fmt.Println(items)
	}

	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	store := bot.New(db)
	events := bot.NewEvents(store)
	bot := bot.NewBot(events)

	bot.Start()

}
