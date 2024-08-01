package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"technicalSupportBot/pkg/keyboards"
)

func sendWelcomeMessage(bot *tgbotapi.BotAPI, chatID int64) {
	previousState[chatID] = "main"

	welcomeMessage := "Добро пожаловать в чат бот тех. поддержки МойОфис! :) " +
		"Выберите необходимую функцию:\n" +
		"1. Инструкции по продуктам.\n" +
		"2. Развертывание продуктов. \n" +
		"3. Рассчет сайзинга продуктов. \n" +
		"4. Связаться с инженером тех. поддержки.\n"
	msg := tgbotapi.NewMessage(chatID, welcomeMessage)
	msg.ReplyMarkup = keyboards.GetMainKeyboard()
	bot.Send(msg)
}

func sendProduct(bot *tgbotapi.BotAPI, chatID int64) {
	previousState[chatID] = "production_instructions"

	chooseProductMessage := "Выберите продукт:"
	msg := tgbotapi.NewMessage(chatID, chooseProductMessage)
	msg.ReplyMarkup = keyboards.GetProductKeyboard()
	bot.Send(msg)
}

func sendDeploymentOptions(bot *tgbotapi.BotAPI, chatID int64) {
    previousState[chatID] = "deployment_options"

    deploymentMessage := "Выберите тип инсталляции:"
    msg := tgbotapi.NewMessage(chatID, deploymentMessage)
    msg.ReplyMarkup = keyboards.GetDeploymentOptionsKeyboard()
    bot.Send(msg)
}

func sendInstructions(bot *tgbotapi.BotAPI, chatID int64, product string) {
	previousState[chatID] = product
	chooseFunction := "Что подсказать? \n" +
		"- Cистемные требования \n" +
		"- Руководство по установке \n" +
		"- Руководство по администрированию \n"
	msg := tgbotapi.NewMessage(chatID, chooseFunction)
	msg.ReplyMarkup = keyboards.GetInstructionsKeyboard()
	bot.Send(msg)
}

func sendSupportEngineerContact(bot *tgbotapi.BotAPI, chatID int64) {
	errorMessage := "Направьте описание проблемы или скриншот ошибки инженеру \nТГ: @IgorMaksimov2000\nПочта: igor.maksimov@myoffice.team \n"
	msg := tgbotapi.NewMessage(chatID, errorMessage)
	bot.Send(msg)
}