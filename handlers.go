package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var previousState = make(map[int64]string)

func handleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.Message.Text {
	case "/start":
		sendWelcomeMessage(bot, update.Message.Chat.ID)
	case "Инструкции по продуктам":
		sendProductInstructions(bot, update.Message.Chat.ID)
	case "Частное Облако":
		sendPrivateCloudOptions(bot, update.Message.Chat.ID)
	case "Системные требования":
		sendSystemRequirements(bot, update.Message.Chat.ID)
	case "Руководство по установке":
		sendInstallationGuideOptions(bot, update.Message.Chat.ID)
	case "PGS":
		sendPGSInstallationGuide(bot, update.Message.Chat.ID)
	case "CO":
		sendCOInstallationGuide(bot, update.Message.Chat.ID)
	case "Руководство по администрированию":
		sendAdminGuide(bot, update.Message.Chat.ID)
	case "Назад":
		handleBackButton(bot, update.Message.Chat.ID)
	case "Связаться с инженером тех. поддержки":
		sendSupportEngineerContact(bot, update.Message.Chat.ID)
	}
}

func handleBackButton(bot *tgbotapi.BotAPI, chatID int64) {
	currentMenu := previousState[chatID]
	switch currentMenu {
	case "production_instructions":
		sendWelcomeMessage(bot, chatID)
	case "private_cloud":
		sendProductInstructions(bot, chatID)
	case "installation_guide":
		sendPrivateCloudOptions(bot, chatID)
	case "pgs_guide", "co_guide":
		sendInstallationGuideOptions(bot, chatID)
	default:
		sendWelcomeMessage(bot, chatID)
	}
}

func sendWelcomeMessage(bot *tgbotapi.BotAPI, chatID int64) {
	previousState[chatID] = "main"

	welcomeMessage := "Добро пожаловать в чат бот тех. поддержки МойОфис! :) " +
		"Выберите необходимую функцию:\n" +
		"1. Инструкции по продуктам.\n" +
		"2. Развертывание продуктов. \n" +
		"3. Связаться с инженером тех. поддержки.\n"
	msg := tgbotapi.NewMessage(chatID, welcomeMessage)
	msg.ReplyMarkup = getMainKeyboard()
	bot.Send(msg)
}

func sendProductInstructions(bot *tgbotapi.BotAPI, chatID int64) {
	previousState[chatID] = "production_instructions"

	chooseProductMessage := "Выберите продукт:"
	msg := tgbotapi.NewMessage(chatID, chooseProductMessage)
	msg.ReplyMarkup = getProductInstructionsKeyboard()
	bot.Send(msg)
}

func sendPrivateCloudOptions(bot *tgbotapi.BotAPI, chatID int64) {
	previousState[chatID] = "private_cloud"
	chooseFunction := "Что подсказать? \n" +
		"- Cистемные требования \n" +
		"- Руководство по установке \n" +
		"- Руководство по администрированию \n"
	msg := tgbotapi.NewMessage(chatID, chooseFunction)
	msg.ReplyMarkup = getPrivateCloudKeyboard()
	bot.Send(msg)
}

func sendSystemRequirements(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/e09/ejjo29n32sj1f93bwyoa5y0upppfbuux/MyOffice_Private_Cloud_3.0_System_Requirements.pdf \n")
	bot.Send(msg)
}

func sendInstallationGuideOptions(bot *tgbotapi.BotAPI, chatID int64) {
	previousState[chatID] = "installation_guide"
	chooseComponent := "Частное облако состоит из двух компонентов: PGS (Система хранения данных) и CO (Система редактирования и совместной работы)\n" +
		"Выберите компонент:\n" +
		"- PGS \n" +
		"- CO \n"
	msg := tgbotapi.NewMessage(chatID, chooseComponent)
	msg.ReplyMarkup = getInstallationGuideKeyboard()
	bot.Send(msg)
}

func sendPGSInstallationGuide(bot *tgbotapi.BotAPI, chatID int64) {
	previousState[chatID] = "pgs_guide"
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/171/x2vqm7n0zp0jg1qjwsieymtp5pemtcg1/MyOffice_Private_Cloud_3.0_PGS_Installation_Guide.pdf \n")
	msg.ReplyMarkup = getBackKeyboard()
	bot.Send(msg)
}

func sendCOInstallationGuide(bot *tgbotapi.BotAPI, chatID int64) {
	previousState[chatID] = "co_guide"
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/05c/137yo7qojdz3hm5k46ngil1nf0opt52p/MyOffice_Private_Cloud_3.0_CO_Installation_Guide.pdf \n")
	msg.ReplyMarkup = getBackKeyboard()
	bot.Send(msg)
}

func sendAdminGuide(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/bc7/a30e4keqfke8h5m7r8mgx4asur3oedlf/MyOffice_Private_Cloud_3.0_Admin_Guide.pdf \n")
	msg.ReplyMarkup = getBackKeyboard()
	bot.Send(msg)
}

func sendSupportEngineerContact(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Технический пресейл Игорь - ТГ: @IgorMaksimov2000 / Почта: igor.maksimov@myoffice.team \n")
	msg.ReplyMarkup = getBackKeyboard()
	bot.Send(msg)
}
