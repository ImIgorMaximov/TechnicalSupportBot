package handlers

import (
	"log"

	"technicalSupportBot/pkg/deployment"
	"technicalSupportBot/pkg/instructions"
	"technicalSupportBot/pkg/sizing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var PreviousState = make(map[int64]string)
var sizingOrDeployment = make(map[int64]string)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	text := update.Message.Text

	switch text {
	case "/start":
		sendWelcomeMessage(bot, chatID)
		PreviousState[chatID] = "start"
	case "В главное меню":
		sendWelcomeMessage(bot, chatID)
		PreviousState[chatID] = "mainMenu"
	case "Инструкции по продуктам":
		sendProduct(bot, chatID)
		PreviousState[chatID] = "instr"
	case "Развертывание продуктов":
		sendProduct(bot, chatID)
		PreviousState[chatID] = "deploy"
		sizingOrDeployment[chatID] = "deploy"
	case "Расчет сайзинга продуктов":
		sendProduct(bot, chatID)
		PreviousState[chatID] = "sizing"
		sizingOrDeployment[chatID] = "sizing"
	case "Частное Облако":
		handlePrivateCloud(bot, chatID)
	case "Squadus":
		sendInstructions(bot, chatID)
		PreviousState[chatID] = "squadus"
	case "Mailion":
		sendInstructions(bot, chatID)
		PreviousState[chatID] = "mailion"
	case "Почта":
		handleMail(bot, chatID)
	case "Системные требования":
		handleSystemRequirements(bot, chatID)
	case "Руководство по установке":
		handleInstallationGuide(bot, chatID)
	case "PGS":
		instructions.SendPGSInstallationGuide(bot, chatID)
		PreviousState[chatID] = "pgs"
	case "CO":
		instructions.SendCOInstallationGuide(bot, chatID)
		PreviousState[chatID] = "co"
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
	case "Cluster":
		handleCluster(bot, chatID)
	case "<2k":
		handleClusterUserRange(bot, chatID, text)
	default:
		handleDefaultState(bot, chatID, text)
	}
}

func handleStandalone(bot *tgbotapi.BotAPI, chatID int64) {
	action := sizingOrDeployment[chatID]

	if action == "sizing" {
		if PreviousState[chatID] == "privateCloud" {
			PreviousState[chatID] = "awaitingUserCountPrivateCloud"
			sizing.HandleSizingPrivateCloudStandalone(bot, chatID)
		} else if PreviousState[chatID] == "mail" {
			PreviousState[chatID] = "awaitingUserCountMail"
			sizing.HandleSizingMailStandalone(bot, chatID)
		}
	} else if action == "deploy" {
		if PreviousState[chatID] == "privateCloud" {
			deployment.SendStandaloneRequirementsPrivateCloud(bot, chatID)
			PreviousState[chatID] = "reqPrivateCloud"
		} else if PreviousState[chatID] == "mail" {
			deployment.SendStandaloneRequirementsPSN(bot, chatID)
			PreviousState[chatID] = "reqPsn"
		}
	}
}

func handleCluster(bot *tgbotapi.BotAPI, chatID int64) {
	action := sizingOrDeployment[chatID]

	if action == "sizing" {
		SendClusterRangeKeyboard(bot, chatID)
		PreviousState[chatID] = "clusterSelection"
	} else if action == "deploy" {
		msg := tgbotapi.NewMessage(chatID, "Извините, раздел находится в разработке 😢")
		bot.Send(msg)
	}
}

func handleClusterUserRange(bot *tgbotapi.BotAPI, chatID int64, userRange string) {
	switch userRange {
	case "<2k":
		// Обработка ввода для диапазона >2k
		PreviousState[chatID] = "awaitingClusterMoreThan2kInput"
		sizing.HandleClusterMoreThan2k(bot, chatID)
	default:
		msg := tgbotapi.NewMessage(chatID, "Выберите корректный диапазон пользователей.")
		bot.Send(msg)
	}
}

func handleDefaultState(bot *tgbotapi.BotAPI, chatID int64, text string) {
	log.Printf("handleDefaultState: %s, %s", PreviousState[chatID], text)

	// Обработка ввода для частного облака
	if PreviousState[chatID] == "awaitingUserCountPrivateCloud" {
		sizing.HandleUserInputPrivateCloud(bot, chatID, text)
	}

	// Обработка ввода данных для почты
	if PreviousState[chatID] == "awaitingUserCountMail" {
		sizing.HandleUserInputMail(bot, chatID, text)
	}

	// Обработка ввода для диапазона <2k
	if PreviousState[chatID] == "awaitingClusterMoreThan2kInput" {
		sizing.HandleClusterMoreThan2kInput(bot, chatID, text)
	}
}

func handleNextStep(bot *tgbotapi.BotAPI, chatID int64) {
	switch PreviousState[chatID] {
	case "reqPsn", "reqPrivateCloud":
		sendStandaloneDownloadPackages(bot, chatID)
		handlePrivateKeyInsert(bot, chatID)
	case "standaloneDownloadPackages":
		handlePrivateKeyInsert(bot, chatID)
	case "privateKeyInsert":
		deployment.SendDNSOptionsPGS(bot, chatID)
		PreviousState[chatID] = "dnsPGS"
	case "privateKeyInsertPSN":
		deployment.SendDNSOptionsPSN(bot, chatID)
		PreviousState[chatID] = "dnsPSN"
	case "dnsPSN":
		deployment.SendStandaloneDownloadDistributionPSN(bot, chatID)
		PreviousState[chatID] = "standaloneDownloadDistributionPSN"
	case "dnsPGS":
		deployment.SendStandaloneDownloadDistribution(bot, chatID)
		PreviousState[chatID] = "standaloneDownloadDistribution"
	case "standaloneDownloadDistributionPSN":
		deployment.SendCertificatesAndKeysPSN(bot, chatID)
		PreviousState[chatID] = "certificatesAndKeysPSN"
	case "standaloneDownloadDistribution":
		deployment.SendCertificatesAndKeysPGS(bot, chatID)
		PreviousState[chatID] = "certificatesAndKeysPGS"
	case "certificatesAndKeysPSN":
		deployment.SendStandalonePSNConfigure(bot, chatID)
		PreviousState[chatID] = "psnConfigure"
	case "psnConfigure":
		deployment.SendPSNDeploy(bot, chatID)
		PreviousState[chatID] = "psnDeploy"
	case "certificatesAndKeysPGS":
		deployment.SendStandalonePGSConfigure(bot, chatID)
		PreviousState[chatID] = "pgsConfigure"
	case "pgsConfigure":
		deployment.SendPGSDeploy(bot, chatID)
		PreviousState[chatID] = "pgsDeploy"
	case "pgsDeploy":
		deployment.SendDNSOptionsCO(bot, chatID)
		PreviousState[chatID] = "dnsCO"
	case "dnsCO":
		deployment.SendCertificatesAndKeysCO(bot, chatID)
		PreviousState[chatID] = "certificatesAndKeysCO"
	case "certificatesAndKeysCO":
		deployment.SendCOInstallation(bot, chatID)
		PreviousState[chatID] = "coInstallation"
	case "coInstallation":
		deployment.SendCOConfigure(bot, chatID)
		PreviousState[chatID] = "coConfigure"
	case "coConfigure":
		deployment.SendCODeploy(bot, chatID)
		PreviousState[chatID] = "coDeploy"
	}

}

func handleBackButton(bot *tgbotapi.BotAPI, chatID int64) {
	currentMenu := PreviousState[chatID]
	switch currentMenu {
	case "instr":
		sendWelcomeMessage(bot, chatID)
		PreviousState[chatID] = "start"
	case "deploy":
		sendWelcomeMessage(bot, chatID)
		PreviousState[chatID] = "start"
	case "privateCloud", "squadus", "mailion":
		sendProduct(bot, chatID)
		PreviousState[chatID] = "instr"
	case "requirementsPrivateCloud", "installationGuidePrivateCloud", "adminGuidePrivateCloud":
		sendInstructions(bot, chatID)
		PreviousState[chatID] = "privateCloud"
	case "pgs", "co":
		instructions.SendInstallationGuideOptionsPrivateCloud(bot, chatID)
		PreviousState[chatID] = "installationGuidePrivateCloud"
	case "requirementsSquadus", "installationGuideSquadus", "adminGuideSquadus":
		sendInstructions(bot, chatID)
		PreviousState[chatID] = "squadus"
	case "requirementsMailion", "installationGuideMailion", "adminGuideMailion":
		sendInstructions(bot, chatID)
		PreviousState[chatID] = "mailion"
	case "requirementsMail", "installationGuideMail", "adminGuideMail":
		sendInstructions(bot, chatID)
		PreviousState[chatID] = "mail"
	case "Standalone":
		sendProduct(bot, chatID)
		PreviousState[chatID] = "deploy"
	case "standaloneDownloadDistribution":
		deployment.SendDNSOptionsPGS(bot, chatID)
		PreviousState[chatID] = "dnsPGS"
	case "standaloneDownloadDistributionPSN":
		deployment.SendDNSOptionsPSN(bot, chatID)
		PreviousState[chatID] = "dnsPSN"
	case "dnsPGS":
		deployment.SendPrivateKeyInsert(bot, chatID)
		PreviousState[chatID] = "privateKeyInsert"
	case "standaloneDownloadPackages":
		deployment.SendStandaloneRequirementsPrivateCloud(bot, chatID)
		PreviousState[chatID] = "requirements"
	case "privateKeyInsert", "privateKeyInsertPSN":
		deployment.SendStandaloneDownloadPackages(bot, chatID)
		PreviousState[chatID] = "standaloneDownloadPackages"
	case "certificatesAndKeysPGS":
		deployment.SendStandaloneDownloadDistribution(bot, chatID)
		PreviousState[chatID] = "standaloneDownloadDistribution"
	case "certificatesAndKeysPSN":
		deployment.SendStandaloneDownloadDistributionPSN(bot, chatID)
		PreviousState[chatID] = "standaloneDownloadDistributionPSN"
	case "psnConfigure":
		deployment.SendCertificatesAndKeysPSN(bot, chatID)
		PreviousState[chatID] = "certificatesAndKeysPSN"
	case "pgsConfigure":
		deployment.SendCertificatesAndKeysPGS(bot, chatID)
		PreviousState[chatID] = "certificatesAndKeysPGS"
	case "pgsDeploy":
		deployment.SendStandalonePGSConfigure(bot, chatID)
		PreviousState[chatID] = "pgsConfigure"
	case "psnDeploy":
		deployment.SendStandalonePSNConfigure(bot, chatID)
		PreviousState[chatID] = "psnConfigure"
	case "coInstallation":
		deployment.SendPGSDeploy(bot, chatID)
		PreviousState[chatID] = "pgsDeploy"
	case "coDeploy":
		deployment.SendCOConfigure(bot, chatID)
		PreviousState[chatID] = "coConfigure"
	default:
		sendWelcomeMessage(bot, chatID)
		PreviousState[chatID] = "start"
	}
}

func handlePrivateKeyInsert(bot *tgbotapi.BotAPI, chatID int64) {
	if PreviousState[chatID] == "reqPrivateCloud" {
		deployment.SendPrivateKeyInsert(bot, chatID)
		PreviousState[chatID] = "privateKeyInsert"
	} else if PreviousState[chatID] == "reqPsn" {
		deployment.SendPrivateKeyInsertPSN(bot, chatID)
		PreviousState[chatID] = "privateKeyInsertPSN"
	}
}

func handlePrivateCloud(bot *tgbotapi.BotAPI, chatID int64) {
	if PreviousState[chatID] == "instr" {
		sendInstructions(bot, chatID)
		PreviousState[chatID] = "privateCloud"
	} else if PreviousState[chatID] == "deploy" || PreviousState[chatID] == "sizing" {
		sendDeploymentOptions(bot, chatID)
		PreviousState[chatID] = "privateCloud"
	}
}

func handleMail(bot *tgbotapi.BotAPI, chatID int64) {
	if PreviousState[chatID] == "instr" {
		sendInstructions(bot, chatID)
		PreviousState[chatID] = "mail"
	} else if PreviousState[chatID] == "deploy" || PreviousState[chatID] == "sizing" {
		sendDeploymentOptions(bot, chatID)
		PreviousState[chatID] = "mail"
	}
}

func handleSystemRequirements(bot *tgbotapi.BotAPI, chatID int64) {
	if PreviousState[chatID] == "privateCloud" {
		instructions.SendSystemRequirementsPivateCloud(bot, chatID)
		PreviousState[chatID] = "requirementsPrivateCloud"
	} else if PreviousState[chatID] == "squadus" {
		instructions.SendSystemRequirementsSquadus(bot, chatID)
		PreviousState[chatID] = "requirementsSquadus"
	} else if PreviousState[chatID] == "mailion" {
		instructions.SendSystemRequirementsMailion(bot, chatID)
		PreviousState[chatID] = "requirementsMailion"
	} else if PreviousState[chatID] == "mail" {
		instructions.SendSystemRequirementsMail(bot, chatID)
		PreviousState[chatID] = "requirementsMail"
	}
}

func handleInstallationGuide(bot *tgbotapi.BotAPI, chatID int64) {
	if PreviousState[chatID] == "privateCloud" {
		instructions.SendInstallationGuideOptionsPrivateCloud(bot, chatID)
		PreviousState[chatID] = "installationGuidePrivateCloud"
	} else if PreviousState[chatID] == "squadus" {
		instructions.SendInstallationGuideSquadus(bot, chatID)
		PreviousState[chatID] = "installationGuideSquadus"
	} else if PreviousState[chatID] == "mailion" {
		instructions.SendInstallationGuideMailion(bot, chatID)
		PreviousState[chatID] = "installationGuideMailion"
	} else if PreviousState[chatID] == "mail" {
		instructions.SendInstallationGuideMail(bot, chatID)
		PreviousState[chatID] = "installationGuideMail"
	}
}

func handleAdminGuide(bot *tgbotapi.BotAPI, chatID int64) {
	if PreviousState[chatID] == "privateCloud" {
		instructions.SendAdminGuidePrivateCloud(bot, chatID)
		PreviousState[chatID] = "adminGuidePrivateCloud"
	} else if PreviousState[chatID] == "squadus" {
		instructions.SendAdminGuideSquadus(bot, chatID)
		PreviousState[chatID] = "adminGuideSquadus"
	} else if PreviousState[chatID] == "mailion" {
		instructions.SendAdminGuideMailion(bot, chatID)
		PreviousState[chatID] = "adminGuideMailion"
	} else if PreviousState[chatID] == "mail" {
		instructions.SendAdminGuideMail(bot, chatID)
		PreviousState[chatID] = "adminGuideMail"
	}
}
