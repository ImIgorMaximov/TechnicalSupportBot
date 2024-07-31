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
	case "Почта 3":
		sendInstructions(bot, chatID, "mail3")
		previousState[chatID] = "mail3"
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
	case "Standalone":
		sendStandaloneRequirements(bot, chatID, "Standalone")
		previousState[chatID] = "standaloneRequirements"
	case "Cluster":
		sendClusterDevelopmentMessage(bot, chatID)
		previousState[chatID] = "clusterDevelopment"
	case "Далее":
		if previousState[chatID] == "standaloneRequirements" {
			sendStandaloneDownloadPackages(bot, chatID)
			previousState[chatID] = "standaloneDownloadPackages"
		}
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
	case "requirementsMail3", "installationGuideMail3", "adminGuideMail3":
		sendInstructions(bot, chatID, "mail3")
		previousState[chatID] = "mail3"
	case "deploymentOptions":
		sendWelcomeMessage(bot, chatID)
		previousState[chatID] = "start"
	case "standaloneRequirements":
		sendDeploymentOptions(bot, chatID)
		previousState[chatID] = "deploymentOptions"
	case "standaloneDownloadPackages":
		sendStandaloneRequirements(bot, chatID, "Standalone")
		previousState[chatID] = "standaloneRequirements"
	case "clusterDevelopment":
		sendDeploymentOptions(bot, chatID)
		previousState[chatID] = "deploymentOptions"
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
	} else if previousState[chatID] == "mail3" {
		sendSystemRequirementsMail3(bot, chatID)
		previousState[chatID] = "requirementsMail3"
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
	} else if previousState[chatID] == "mail3" {
		sendInstallationGuideMail3(bot, chatID)
		previousState[chatID] = "installationGuideMail3"
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
	} else if previousState[chatID] == "mail3" {
		sendAdminGuideMailion(bot, chatID)
		previousState[chatID] = "adminGuideMail3"
	}
}