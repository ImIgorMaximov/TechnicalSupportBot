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
		sendInstructions(bot, chatID, "privateCloud")
		previousState[chatID] = "privateCloud"
	case "Squadus":
		sendInstructions(bot, chatID, "squadus")
		previousState[chatID] = "squadus"
	case "Mailion":
		sendInstructions(bot, chatID, "mailion")
		previousState[chatID] = "mailion"
	case "Системные требования":
		handleSystemRequirements(bot, chatID)
	case "Руководство по установке":
		handleInstallationGuide(bot, chatID)
	case "PGS":
		sendPGSInstallationGuide(bot, chatID)
		previousState[chatID] = "pgs"
	case "CO":
		sendCOInstallationGuide(bot, chatID)
		previousState[chatID] = "co"
	case "Руководство по администрированию":
		handleAdminGuide(bot, chatID)
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
	case "privateCloud", "squadus", "mailion":
		sendProduct(bot, chatID)
		previousState[chatID] = "product"
	case "requirementsPrivateCloud", "installationGuidePrivateCloud", "adminGuidePrivateCloud":
		sendInstructions(bot, chatID, "privateCloud")
		previousState[chatID] = "privateCloud"
	case "pgs", "co":
		sendInstallationGuideOptionsPrivateCloud(bot, chatID)
		previousState[chatID] = "installationGuidePrivateCloud"
	case "requirementsSquadus", "installationGuideSquadus", "adminGuideSquadus":
		sendInstructions(bot, chatID, "squadus")
		previousState[chatID] = "squadus"
	case "requirementsMailion", "installationGuideMailion", "adminGuideMailion":
		sendInstructions(bot, chatID, "mailion")
		previousState[chatID] = "mailion"
	default:
		sendWelcomeMessage(bot, chatID)
		previousState[chatID] = "start"
	}
}

func handleSystemRequirements(bot *tgbotapi.BotAPI, chatID int64) {
	if previousState[chatID] == "privateCloud" {
		sendSystemRequirementsPivateCloud(bot, chatID)
		previousState[chatID] = "requirementsPrivateCloud"
	} else if previousState[chatID] == "squadus" {
		sendSystemRequirementsSquadus(bot, chatID)
		previousState[chatID] = "requirementsSquadus"
	} else if previousState[chatID] == "mailion" {
		sendSystemRequirementsMailion(bot, chatID)
		previousState[chatID] = "requirementsMailion"
	}
}

func handleInstallationGuide(bot *tgbotapi.BotAPI, chatID int64) {
	if previousState[chatID] == "privateCloud" {
		sendInstallationGuideOptionsPrivateCloud(bot, chatID)
		previousState[chatID] = "installationGuidePrivateCloud"
	} else if previousState[chatID] == "squadus" {
		sendInstallationGuideSquadus(bot, chatID)
		previousState[chatID] = "installationGuideSquadus"
	} else if previousState[chatID] == "mailion" {
		sendInstallationGuideMailion(bot, chatID)
		previousState[chatID] = "installationGuideMailion"
	}
}

func handleAdminGuide(bot *tgbotapi.BotAPI, chatID int64) {
	if previousState[chatID] == "privateCloud" {
		sendAdminGuidePrivateCloud(bot, chatID)
		previousState[chatID] = "adminGuidePrivateCloud"
	} else if previousState[chatID] == "squadus" {
		sendAdminGuideSquadus(bot, chatID)
		previousState[chatID] = "adminGuideSquadus"
	} else if previousState[chatID] == "mailion" {
		sendAdminGuideMailion(bot, chatID)
		previousState[chatID] = "adminGuideMailion"
	}
}