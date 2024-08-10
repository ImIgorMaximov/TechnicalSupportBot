package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var previousState = make(map[int64]string)
var sizingOrDeployment = make(map[int64]string)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	text := update.Message.Text

	switch text {
	case "/start":
		sendWelcomeMessage(bot, chatID)
		previousState[chatID] = "start"
	case "В главное меню":
		sendWelcomeMessage(bot, chatID)
		previousState[chatID] = "mainMenu"
	case "Инструкции по продуктам":
		sendProduct(bot, chatID)
		previousState[chatID] = "instr"
	case "Развертывание продуктов":
		sendProduct(bot, chatID)
		previousState[chatID] = "deploy"
		sizingOrDeployment[chatID] = "deploy"
	case "Расчет сайзинга продуктов":
		sendProduct(bot, chatID)
		previousState[chatID] = "sizing"
		sizingOrDeployment[chatID] = "sizing"
	case "Частное Облако":
		handlePrivateCloud(bot, chatID)
	case "Squadus":
		sendInstructions(bot, chatID)
		previousState[chatID] = "squadus"
	case "Mailion":
		sendInstructions(bot, chatID)
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
	case "Standalone":
		handleStandalone(bot, chatID)
	case "Готово":
		handleNextStep(bot, chatID)
	case "Запустить деплой":
		handleNextStep(bot, chatID)
	case "Проверить корректность сертификатов и ключа":
		sendIsCertificates(bot, chatID)
	case "Пример конфига PGS - hosts.yml":
		sendConfigFile(bot, chatID, "/home/admin-msk/MyOfficeConfig/hostsPGS.yml", "hostsPGS.yml")
	case "Пример конфига PSN - hosts.yml":
		sendConfigFile(bot, chatID, "/home/admin-msk/MyOfficeConfig/hostsPSN.yml", "hostsPSN.yml")
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
	default:
		// В этом блоке добавьте обработку для состояния, когда ожидается ввод пользователя
		if previousState[chatID] == "awaitingUserCountPrivateCloud" ||
			previousState[chatID] == "awaitingActiveUserCountPrivateCloud" ||
			previousState[chatID] == "awaitingDocumentCountPrivateCloud" ||
			previousState[chatID] == "awaitingStorageQuotaPrivateCloud" {
			HandleUserInputPrivateCloud(bot, chatID, text)
		}

		// Обработка ввода данных для почты
		if previousState[chatID] == "awaitingUserCountMail" ||
			previousState[chatID] == "awaitingDiskQuotaMail" ||
			previousState[chatID] == "awaitingEmailsPerDayMail" ||
			previousState[chatID] == "awaitingSpamCoefficientMail" {
			HandleUserInputMail(bot, chatID, text)
		}
	}
}

func handleNextStep(bot *tgbotapi.BotAPI, chatID int64) {
	switch previousState[chatID] {
	case "reqPsn", "reqPrivateCloud":
		sendStandaloneDownloadPackages(bot, chatID)
		handlePrivateKeyInsert(bot, chatID)
	case "standaloneDownloadPackages":
		handlePrivateKeyInsert(bot, chatID)
	case "privateKeyInsert":
		sendDNSOptionsPGS(bot, chatID)
		previousState[chatID] = "dnsPGS"
	case "privateKeyInsertPSN":
		sendDNSOptionsPSN(bot, chatID)
		previousState[chatID] = "dnsPSN"
	case "dnsPSN":
		sendStandaloneDownloadDistributionPSN(bot, chatID)
		previousState[chatID] = "standaloneDownloadDistributionPSN"
	case "dnsPGS":
		sendStandaloneDownloadDistribution(bot, chatID)
		previousState[chatID] = "standaloneDownloadDistribution"
	case "standaloneDownloadDistribution":
		sendCertificatesAndKeysPGS(bot, chatID)
		previousState[chatID] = "certificatesAndKeysPGS"
		if sizingOrDeployment[chatID] == "sizing" && previousState[chatID] == "privateCloud" {
			HandleSizingPrivateCloudStandalone(bot, chatID)
			previousState[chatID] = "awaitingUserCountPrivateCloud"
		}
		sendCertificatesAndKeysPSN(bot, chatID)
		previousState[chatID] = "certificatesAndKeysPSN"
	case "certificatesAndKeysPSN":
		sendStandalonePSNConfigure(bot, chatID)
		previousState[chatID] = "psnConfigure"
	case "psnConfigure":
		sendPSNDeploy(bot, chatID)
		previousState[chatID] = "psnDeploy"
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
		sendInstructions(bot, chatID)
		previousState[chatID] = "privateCloud"
	case "pgs", "co":
		sendInstallationGuideOptionsPrivateCloud(bot, chatID)
		previousState[chatID] = "installationGuidePrivateCloud"
	case "requirementsSquadus", "installationGuideSquadus", "adminGuideSquadus":
		sendInstructions(bot, chatID)
		previousState[chatID] = "squadus"
	case "requirementsMailion", "installationGuideMailion", "adminGuideMailion":
		sendInstructions(bot, chatID)
		previousState[chatID] = "mailion"
	case "requirementsMail", "installationGuideMail", "adminGuideMail":
		sendInstructions(bot, chatID)
		previousState[chatID] = "mail"
	case "Standalone":
		sendProduct(bot, chatID)
		previousState[chatID] = "deploy"
	case "standaloneDownloadDistribution":
		sendDNSOptionsPGS(bot, chatID)
		previousState[chatID] = "dnsPGS"
	case "standaloneDownloadDistributionPSN":
		sendDNSOptionsPSN(bot, chatID)
		previousState[chatID] = "dnsPSN"
	case "dnsPGS":
		sendPrivateKeyInsert(bot, chatID)
		previousState[chatID] = "privateKeyInsert"
	case "standaloneDownloadPackages":
		sendStandaloneRequirementsPrivateCloud(bot, chatID)
		previousState[chatID] = "requirements"
	case "privateKeyInsert", "privateKeyInsertPSN":
		sendStandaloneDownloadPackages(bot, chatID)
		previousState[chatID] = "standaloneDownloadPackages"
	case "certificatesAndKeysPGS":
		sendStandaloneDownloadDistribution(bot, chatID)
		previousState[chatID] = "standaloneDownloadDistribution"
	case "certificatesAndKeysPSN":
		sendStandaloneDownloadDistributionPSN(bot, chatID)
		previousState[chatID] = "standaloneDownloadDistributionPSN"
	case "psnConfigure":
		sendCertificatesAndKeysPSN(bot, chatID)
		previousState[chatID] = "certificatesAndKeysPSN"
	case "pgsConfigure":
		sendCertificatesAndKeysPGS(bot, chatID)
		previousState[chatID] = "certificatesAndKeysPGS"
	case "pgsDeploy":
		sendStandalonePGSConfigure(bot, chatID)
		previousState[chatID] = "pgsConfigure"
	case "psnDeploy":
		sendStandalonePSNConfigure(bot, chatID)
		previousState[chatID] = "psnConfigure"
	case "coInstallation":
		sendPGSDeploy(bot, chatID)
		previousState[chatID] = "pgsDeploy"
	case "coDeploy":
		sendCOConfigure(bot, chatID)
		previousState[chatID] = "coConfigure"
	default:
		sendWelcomeMessage(bot, chatID)
		previousState[chatID] = "start"
	}
}

func handlePrivateKeyInsert(bot *tgbotapi.BotAPI, chatID int64) {
	if previousState[chatID] == "reqPrivateCloud" {
		sendPrivateKeyInsert(bot, chatID)
		previousState[chatID] = "privateKeyInsert"
	} else if previousState[chatID] == "reqPsn" {
		sendPrivateKeyInsertPSN(bot, chatID)
		previousState[chatID] = "privateKeyInsertPSN"
	}
}

func handleStandalone(bot *tgbotapi.BotAPI, chatID int64) {
	action := sizingOrDeployment[chatID]

	if action == "sizing" {
		if previousState[chatID] == "privateCloud" {
			previousState[chatID] = "awaitingUserCountPrivateCloud"
			HandleSizingPrivateCloudStandalone(bot, chatID)
		} else if previousState[chatID] == "mail" {
			previousState[chatID] = "awaitingUserCountMail"
			HandleSizingMailStandalone(bot, chatID)
		}
	} else if action == "deploy" {
		if previousState[chatID] == "privateCloud" {
			sendStandaloneRequirementsPrivateCloud(bot, chatID)
			previousState[chatID] = "reqPrivateCloud"
		} else if previousState[chatID] == "mail" {
			sendStandaloneRequirementsPSN(bot, chatID)
			previousState[chatID] = "reqPsn"
		}
	}
}

func handlePrivateCloud(bot *tgbotapi.BotAPI, chatID int64) {
	if previousState[chatID] == "instr" {
		sendInstructions(bot, chatID)
		previousState[chatID] = "privateCloud"
	} else if previousState[chatID] == "deploy" || previousState[chatID] == "sizing" {
		sendDeploymentOptions(bot, chatID)
		previousState[chatID] = "privateCloud"
	}
}

func handleMail(bot *tgbotapi.BotAPI, chatID int64) {
	if previousState[chatID] == "instr" {
		sendInstructions(bot, chatID)
		previousState[chatID] = "mail"
	} else if previousState[chatID] == "deploy" || previousState[chatID] == "sizing" {
		sendDeploymentOptions(bot, chatID)
		previousState[chatID] = "mail"
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
	} else if previousState[chatID] == "mail" {
		sendSystemRequirementsMail(bot, chatID)
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
	} else if previousState[chatID] == "mail" {
		sendInstallationGuideMail(bot, chatID)
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
	} else if previousState[chatID] == "mail" {
		sendAdminGuideMail(bot, chatID)
		previousState[chatID] = "adminGuideMail"
	}
}
