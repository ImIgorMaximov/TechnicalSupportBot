package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("Переменная окружения /'TELEGRAM_BOT_TOKEN/' не объявлена. Добавьте токен бота")
	}

	//Создание бота

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	//Установка режима дебага

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			switch update.Message.Text {
			case "/start":
				welcomeMessage := "Добро пожаловать в чат бот тех. поддержки МойОфис! :) " +
					"Выберите необходимую функцию:\n" +
					"1. Инструкции по продуктам.\n" +
					"2. Связаться с инженером тех. поддержки.\n"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, welcomeMessage)
				keyboard := tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Инструкции по продуктам"),
					),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Связаться с инженером тех. поддержки"),
					),
				)
				msg.ReplyMarkup = keyboard
				bot.Send(msg)
			case "Инструкции по продуктам":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ссылка на инструкции по продуктам - https://support.myoffice.ru/products/\n")
				msg.ParseMode = "markdown"
				bot.Send(msg)
			case "Связаться с инженером тех. поддержки":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Технический пресейл Игорь - @IgorMaksimov2000 \n")
				msg.ParseMode = "markdown"
				bot.Send(msg)
			}
		}
	}
}

