package main

import (
	"log"
	"os"
	"technicalSupportBot/pkg/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("Переменная окружения 'TELEGRAM_BOT_TOKEN' не объявлена. Добавьте токен бота")
	}

	// Создание бота
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	// Создание StateManager
	sm := handlers.NewStateManager()

	// Установка режима дебага
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			handlers.HandleUpdate(bot, update, sm)
		}
	}
}
