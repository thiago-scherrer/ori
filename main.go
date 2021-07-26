package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	token := os.Getenv("TG_TOKEN")

	if token == "" {
		log.Fatalln("TG_TOKEN required!")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalln(err)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	start(bot)
}

func start(bot *tgbotapi.BotAPI) (bool, error) {
	u := tgbotapi.NewUpdate(5)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalln(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		chatId := update.Message.Chat.ID

		switch update.Message.Text {
		case "/classes":
			classes(bot, chatId)
		case "/lista-de-magia":
			magic(bot, chatId)
		default:
			help(bot, chatId)
		}

	}
	return true, nil
}
