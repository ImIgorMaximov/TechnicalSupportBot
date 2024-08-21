package handlers

import (
	"log"
	"technicalSupportBot/pkg/deployment"
	"technicalSupportBot/pkg/instructions"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleBackButton(bot *tgbotapi.BotAPI, chatID int64) {
	currentMenu, exists := PreviousState[chatID]
	if !exists {
		log.Printf("Состояние для chatID %d не найдено", chatID)
		return
	}

	log.Printf("Нажата кнопка \"Назад\" для chatID %d, текущее состояние: %s", chatID, currentMenu)

	switch currentMenu {
	case "instr":
		log.Printf("Состояние: %s. Отправка приветственного сообщения и переход на начальный экран.", currentMenu)
		sendWelcomeMessage(bot, chatID)
		PreviousState[chatID] = "start"

	case "deploy":
		log.Printf("Состояние: %s. Отправка приветственного сообщения и переход на начальный экран.", currentMenu)
		sendWelcomeMessage(bot, chatID)
		PreviousState[chatID] = "start"

	case "privateCloud", "squadus", "mailion":
		log.Printf("Состояние: %s. Отправка информации о продукте.", currentMenu)
		sendProduct(bot, chatID)
		PreviousState[chatID] = "instr"

	case "requirementsPrivateCloud", "installationGuidePrivateCloud", "adminGuidePrivateCloud":
		log.Printf("Состояние: %s. Отправка инструкций для PrivateCloud.", currentMenu)
		sendInstructions(bot, chatID)
		PreviousState[chatID] = "privateCloud"

	case "pgs", "co":
		log.Printf("Состояние: %s. Отправка опций установки PrivateCloud.", currentMenu)
		instructions.SendInstallationGuideOptionsPrivateCloud(bot, chatID)
		PreviousState[chatID] = "installationGuidePrivateCloud"

	case "requirementsSquadus", "installationGuideSquadus", "adminGuideSquadus":
		log.Printf("Состояние: %s. Отправка инструкций для Squadus.", currentMenu)
		sendInstructions(bot, chatID)
		PreviousState[chatID] = "squadus"

	case "requirementsMailion", "installationGuideMailion", "adminGuideMailion":
		log.Printf("Состояние: %s. Отправка инструкций для Mailion.", currentMenu)
		sendInstructions(bot, chatID)
		PreviousState[chatID] = "mailion"

	case "requirementsMail", "installationGuideMail", "adminGuideMail":
		log.Printf("Состояние: %s. Отправка инструкций для Mail.", currentMenu)
		sendInstructions(bot, chatID)
		PreviousState[chatID] = "mail"

	case "Standalone":
		log.Printf("Состояние: %s. Отправка информации о продукте и переход к развертыванию.", currentMenu)
		sendProduct(bot, chatID)
		PreviousState[chatID] = "deploy"

	case "standaloneDownloadDistribution":
		log.Printf("Состояние: %s. Отправка DNS-опций PGS.", currentMenu)
		deployment.SendDNSOptionsPGS(bot, chatID)
		PreviousState[chatID] = "dnsPGS"

	case "standaloneDownloadDistributionPSN":
		log.Printf("Состояние: %s. Отправка DNS-опций PSN.", currentMenu)
		deployment.SendDNSOptionsPSN(bot, chatID)
		PreviousState[chatID] = "dnsPSN"

	case "dnsPGS":
		log.Printf("Состояние: %s. Отправка вставки приватного ключа.", currentMenu)
		deployment.SendPrivateKeyInsert(bot, chatID)
		PreviousState[chatID] = "privateKeyInsert"

	case "standaloneDownloadPackages":
		log.Printf("Состояние: %s. Отправка требований для PrivateCloud.", currentMenu)
		deployment.SendStandaloneRequirementsPrivateCloud(bot, chatID)
		PreviousState[chatID] = "requirements"

	case "privateKeyInsert", "privateKeyInsertPSN":
		log.Printf("Состояние: %s. Отправка пакетов для самостоятельной загрузки.", currentMenu)
		deployment.SendStandaloneDownloadPackages(bot, chatID)
		PreviousState[chatID] = "standaloneDownloadPackages"

	case "certificatesAndKeysPGS":
		log.Printf("Состояние: %s. Отправка пакетов для самостоятельной загрузки.", currentMenu)
		deployment.SendStandaloneDownloadDistribution(bot, chatID)
		PreviousState[chatID] = "standaloneDownloadDistribution"

	case "certificatesAndKeysPSN":
		log.Printf("Состояние: %s. Отправка пакетов для самостоятельной загрузки PSN.", currentMenu)
		deployment.SendStandaloneDownloadDistributionPSN(bot, chatID)
		PreviousState[chatID] = "standaloneDownloadDistributionPSN"

	case "psnConfigure":
		log.Printf("Состояние: %s. Отправка сертификатов и ключей PSN.", currentMenu)
		deployment.SendCertificatesAndKeysPSN(bot, chatID)
		PreviousState[chatID] = "certificatesAndKeysPSN"

	case "pgsConfigure":
		log.Printf("Состояние: %s. Отправка сертификатов и ключей PGS.", currentMenu)
		deployment.SendCertificatesAndKeysPGS(bot, chatID)
		PreviousState[chatID] = "certificatesAndKeysPGS"

	case "pgsDeploy":
		log.Printf("Состояние: %s. Отправка конфигурации для PGS.", currentMenu)
		deployment.SendStandalonePGSConfigure(bot, chatID)
		PreviousState[chatID] = "pgsConfigure"

	case "psnDeploy":
		log.Printf("Состояние: %s. Отправка конфигурации для PSN.", currentMenu)
		deployment.SendStandalonePSNConfigure(bot, chatID)
		PreviousState[chatID] = "psnConfigure"

	case "coInstallation":
		log.Printf("Состояние: %s. Отправка развертывания PGS.", currentMenu)
		deployment.SendPGSDeploy(bot, chatID)
		PreviousState[chatID] = "pgsDeploy"

	case "coDeploy":
		log.Printf("Состояние: %s. Отправка конфигурации CO.", currentMenu)
		deployment.SendCOConfigure(bot, chatID)
		PreviousState[chatID] = "coConfigure"

	default:
		log.Printf("Состояние: %s. Неизвестное состояние, отправка приветственного сообщения и переход на начальный экран.", currentMenu)
		sendWelcomeMessage(bot, chatID)
		PreviousState[chatID] = "start"
	}
}
