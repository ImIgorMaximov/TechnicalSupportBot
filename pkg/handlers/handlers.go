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
		handleMail(bot, chatID)
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
		handleStandaloneRequirements(bot, chatID)
	case "Готово":
		handleNextStep(bot, chatID)
	case "Запустить деплой":
		handleNextStep(bot, chatID)
	case "Проверить корректность сертификатов и ключа":
		sendIsCertificates(bot, chatID)
	case "Пример конфига PGS - hosts.yml":
		sendConfigFile(bot, chatID, "/home/admin-msk/MyOfficeConfig/hostsPGS.yml", "hostsPGS.yml")
	case "Пример конфига CO - main.yml":
		sendConfigFile(bot, chatID, "/home/admin-msk/MyOfficeConfig/mainCO.yml", "mainCO.yml")
	case "Пример конфига CO - hosts.yml":
		sendConfigFile(bot, chatID, "/home/admin-msk/MyOfficeConfig/hostsCO.yml", "hostsCO.yml")
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
	case "requirements":
		sendStandaloneDownloadPackages(bot, chatID)
		previousState[chatID] = "standaloneDownloadPackages"
	case "standaloneDownloadPackages":
		sendPrivateKeyInsert(bot, chatID)
		previousState[chatID] = "privateKeyInsert"
	case "privateKeyInsert":
		sendDNSOptionsPGS(bot, chatID)
		previousState[chatID] = "dnsPGS"
	case "dnsPGS":
		sendStandaloneDownloadDistribution(bot, chatID)
		previousState[chatID] = "standaloneDownloadDistribution"
	case "standaloneDownloadDistribution":
		sendCertificatesAndKeysPGS(bot, chatID)
		previousState[chatID] = "certificatesAndKeysPGS"
	case "certificatesAndKeysPGS":
		sendStandalonePGSConfigure(bot, chatID)
		previousState[chatID] = "pgsConfigure"
	case "pgsConfigure":
		sendPGSDeploy(bot, chatID)
		previousState[chatID] = "pgsDeploy"
	case "pgsDeploy":
		sendDNSOptionsCO(bot, chatID)
		previousState[chatID] = "dnsCO"
	case "dnsCO":
		sendCertificatesAndKeysCO(bot, chatID)
		previousState[chatID] = "certificatesAndKeysCO"
	case "certificatesAndKeysCO":
		sendCOInstallation(bot, chatID)
		previousState[chatID] = "coInstallation"
	case "coInstallation":
		sendCOConfigure(bot, chatID)
		previousState[chatID] = "coConfigure"
	case "coConfigure":
		sendCODeploy(bot, chatID)
		previousState[chatID] = "coDeploy"
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
	case "Standalone":
		sendProduct(bot, chatID)
		previousState[chatID] = "deploy"
	case "standaloneDownloadDistribution":
		sendDNSOptionsPGS(bot, chatID)
		previousState[chatID] = "dnsPGS"
	case "dnsPGS":
		sendPrivateKeyInsert(bot, chatID)
		previousState[chatID] = "privateKeyInsert"
	case "standaloneDownloadPackages":
		sendStandaloneRequirementsCO(bot, chatID)
		previousState[chatID] = "requirements"
	case "privateKeyInsert":
		sendStandaloneDownloadPackages(bot, chatID)
		previousState[chatID] = "standaloneDownloadPackages"
	case "certificatesAndKeysPGS":
		sendStandaloneDownloadDistribution(bot, chatID)
		previousState[chatID] = "standaloneDownloadDistribution"
	case "pgsConfigure":
		sendCertificatesAndKeysPGS(bot, chatID)
		previousState[chatID] = "certificatesAndKeysPGS"
	case "pgsDeploy":
		sendStandalonePGSConfigure(bot, chatID)
		previousState[chatID] = "pgsConfigure"
	case "coInstallation":
		sendPGSDeploy(bot, chatID)
		previousState[chatID] = "pgsDeploy"
	case "coConfigure":
		sendCOInstallation(bot, chatID)
		previousState[chatID] = "coInstallation"
	case "requirements":
		sendDeploymentOptions(bot, chatID)
		previousState[chatID] = "Standalone"
	case "coDeploy":
		sendCOConfigure(bot, chatID)
		previousState[chatID] = "coConfigure"
	default:
		sendWelcomeMessage(bot, chatID)
		previousState[chatID] = "start"
	}
}

func handleStandaloneRequirements(bot *tgbotapi.BotAPI, chatID int64) {
	if previousState[chatID] == "privateCloud" {
		sendStandaloneRequirementsCO(bot, chatID)
		previousState[chatID] = "requirements"
	} else if previousState[chatID] == "mail3" {
		sendStandaloneRequirementsPSN(bot, chatID)
		previousState[chatID] = "requirements"
	}
}

func handlePrivateCloud(bot *tgbotapi.BotAPI, chatID int64) {
	if previousState[chatID] == "instr" {
		sendInstructions(bot, chatID, "privateCloud")
		previousState[chatID] = "privateCloud"
	} else if previousState[chatID] == "deploy" {
		sendDeploymentOptions(bot, chatID)
		previousState[chatID] = "privateCloud"
	}
}

func handleMail(bot *tgbotapi.BotAPI, chatID int64) {
	if previousState[chatID] == "instr" {
		sendInstructions(bot, chatID, "mail3")
		previousState[chatID] = "mail3"
	} else if previousState[chatID] == "deploy" {
		sendDeploymentOptions(bot, chatID)
		previousState[chatID] = "mail3"
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
