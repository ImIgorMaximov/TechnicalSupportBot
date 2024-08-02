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
	case "В главное меню":
		sendWelcomeMessage(bot, chatID)
		previousState[chatID] = "mainMenu"
	case "Инструкции по продуктам":
		sendProduct(bot, chatID)
		previousState[chatID] = "instr"
	case "Частное Облако":
		handlePrivateCloud(bot, chatID)
	case "Squadus":
		sendInstructions(bot, chatID, "squadus")
		previousState[chatID] = "squadus"
	case "Mailion":
		sendInstructions(bot, chatID, "mailion")
		previousState[chatID] = "mailion"
	case "Почта":
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
	case "Развертывание продуктов":
		sendProduct(bot, chatID)
		previousState[chatID] = "deploy"
	case "Standalone":
		sendStandaloneRequirements(bot, chatID, "Standalone")
		previousState[chatID] = "standaloneRequirements"
	case "Готово":
		handleNextStep(bot, chatID)
	case "Все Окей! Деплоим!":
		handleNextStep(bot, chatID)
	case "Проверить корректность сертификатов и ключа":
		sendIsCertificates(bot, chatID)
	case "Пример конфига hosts.yml":
		sendPGSConfig(bot, chatID)
	case "Далее":
		handleNextStep(bot, chatID)
	case "Установка CO":
		handleNextStep(bot, chatID)
	case "Распаковка ISO образа":
		sendUnzippingISO(bot, chatID)
	}
}

func handleNextStep(bot *tgbotapi.BotAPI, chatID int64) {
	switch previousState[chatID] {
	case "standaloneRequirements":
		sendStandaloneDownloadPackages(bot, chatID)
		previousState[chatID] = "standaloneDownloadPackages"
	case "privateKeyInsert":
		sendDNSOptions(bot, chatID)
		previousState[chatID] = "standaloneDNSOptions"
	case "standaloneDownloadPackages":
		sendPrivateKeyInsert(bot, chatID)
		previousState[chatID] = "privateKeyInsert"
	case "standaloneDNSOptions":
		sendStandaloneDownloadDistribution(bot, chatID)
		previousState[chatID] = "standaloneDownloadDistribution"
	case "standaloneDownloadDistribution":
		sendCertificatesAndKeys(bot, chatID)
		previousState[chatID] = "certificatesAndKeys"
	case "certificatesAndKeys":
		sendStandalonePGSConfigure(bot, chatID)
		previousState[chatID] = "pgsConfigure"
	case "pgsConfigure":
		sendPGSDeploy(bot, chatID)
		previousState[chatID] = "pgsDeploy"
	case "pgsDeploy":
		sendCOInstallation(bot, chatID)
		previousState[chatID] = "coInstallation"
	case "coInstallation":
		sendCOScripts(bot, chatID)
		previousState[chatID] = "coScripts"
	}
}

func handleBackButton(bot *tgbotapi.BotAPI, chatID int64) {
	currentMenu := previousState[chatID]
	switch currentMenu {
	case "instr":
		sendWelcomeMessage(bot, chatID)
		previousState[chatID] = "start"
	case "deploy":
		sendWelcomeMessage(bot, chatID)
		previousState[chatID] = "start"
	case "privateCloud", "squadus", "mailion":
		sendProduct(bot, chatID)
		previousState[chatID] = "instr"
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
	case "requirementsMail", "installationGuideMail", "adminGuideMail":
		sendInstructions(bot, chatID, "mail")
		previousState[chatID] = "mail"
	case "deploymentOptions":
		sendWelcomeMessage(bot, chatID)
		previousState[chatID] = "standaloneDownloadPackages"
	case "standaloneDownloadDistribution":
		sendDNSOptions(bot, chatID)
		previousState[chatID] = "standaloneDNSOptions"
	case "standaloneDNSOptions":
		sendPrivateKeyInsert(bot, chatID)
		previousState[chatID] = "privateKeyInsert"
	case "certificatesAndKeys":
		sendStandaloneDownloadDistribution(bot, chatID)
		previousState[chatID] = "standaloneDownloadDistribution"
	case "pgsConfigure":
		sendCertificatesAndKeys(bot, chatID)
		previousState[chatID] = "certificatesAndKeys"
	case "pgsDeploy":
		sendStandalonePGSConfigure(bot, chatID)
		previousState[chatID] = "pgsConfigure"
	case "coInstallation":
		sendPGSDeploy(bot, chatID)
		previousState[chatID] = "pgsDeploy"
	case "coScripts":
		sendCOInstallation(bot, chatID)
		previousState[chatID] = "coInstallation"
	case "standaloneRequirements":
		sendDeploymentOptions(bot, chatID)
		previousState[chatID] = "deploymentOptions"
	default:
		sendWelcomeMessage(bot, chatID)
		previousState[chatID] = "start"
	}
}

func handlePrivateCloud(bot *tgbotapi.BotAPI, chatID int64) {
    if previousState[chatID] == "instr" {
        sendInstructions(bot, chatID, "privateCloud")
        previousState[chatID] = "privateCloud"
    } else if previousState[chatID] == "deploy" {
        sendDeploymentOptions(bot, chatID)
        previousState[chatID] = "privateCloudDeploy"
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
		previousState[chatID] = "requirementsMail"
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
		previousState[chatID] = "installationGuideMail"
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
		sendAdminGuideMail3(bot, chatID)
		previousState[chatID] = "adminGuideMail"
	}
}