package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var previousState = make(map[int64]string)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	switch update.Message.Text {
	case "/start":
		sendWelcomeMessage(bot, chatID)
		previousState[chatID] = "start"
	case "Инструкции по продуктам":
		sendProduct(bot, chatID)
		previousState[chatID] = "product"
	case "Частное Облако":
		sendInstructions(bot, chatID)
		previousState[chatID] = "privateCloud"
	case "Системные требования":
		sendSystemRequirementsPivateCloud(bot, chatID)
		previousState[chatID] = "requirementsPrivateCloud"
	case "Руководство по установке":
		sendInstallationGuideOptionsPrivateCloud(bot, chatID)
		previousState[chatID] = "installationGuidePrivateCloud"
	case "PGS":
		sendPGSInstallationGuide(bot, chatID)
		previousState[chatID] = "pgs"
	case "CO":
		sendCOInstallationGuide(bot, chatID)
		previousState[chatID] = "co"
	case "Руководство по администрированию":
		sendAdminGuidePrivateCloud(bot, chatID)
		previousState[chatID] = "adminGuide"
	case "Назад":
		handleBackButton(bot, chatID)
	case "Связаться с инженером тех. поддержки":
		sendSupportEngineerContact(bot, chatID)
		previousState[chatID] = "supportContact"
	}
}

func handleBackButton(bot *tgbotapi.BotAPI, chatID int64) {
    currentMenu := previousState[chatID]
    switch currentMenu {
    case "product":
        sendWelcomeMessage(bot, chatID)
        previousState[chatID] = "start"
    case "privateCloud":
        sendProduct(bot, chatID)
        previousState[chatID] = "product"
    case "requirementsPrivateCloud":
        sendInstructions(bot, chatID)
        previousState[chatID] = "privateCloud"
    case "installationGuidePrivateCloud":
        sendInstructions(bot, chatID)
        previousState[chatID] = "privateCloud"
    case "pgs", "co":
        sendInstallationGuideOptionsPrivateCloud(bot, chatID)
        previousState[chatID] = "installationGuidePrivateCloud"
    default:
        sendWelcomeMessage(bot, chatID)
        previousState[chatID] = "start"
    }
}