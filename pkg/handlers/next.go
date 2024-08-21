package handlers

import (
	"log"
	"technicalSupportBot/pkg/deployment"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleNextStep(bot *tgbotapi.BotAPI, chatID int64) {
	state, exists := PreviousState[chatID]
	if !exists {
		log.Printf("Предыдущего состояния для chatID %d не найдено", chatID)
		return
	}

	log.Printf("Обработка следующего шага для chatID %d, текущее состояние: %s", chatID, state)

	switch state {
	case "reqPsn", "reqPrivateCloud":
		log.Printf("Состояние: %s. Отправка пакетов для самостоятельной загрузки.", state)
		sendStandaloneDownloadPackages(bot, chatID)
		log.Printf("Обработка вставки приватного ключа.")
		handlePrivateKeyInsert(bot, chatID)

	case "standaloneDownloadPackages":
		log.Printf("Состояние: %s. Обработка вставки приватного ключа.", state)
		handlePrivateKeyInsert(bot, chatID)

	case "privateKeyInsert":
		log.Printf("Состояние: %s. Отправка DNS-опций PGS.", state)
		deployment.SendDNSOptionsPGS(bot, chatID)
		PreviousState[chatID] = "dnsPGS"

	case "privateKeyInsertPSN":
		log.Printf("Состояние: %s. Отправка DNS-опций PSN.", state)
		deployment.SendDNSOptionsPSN(bot, chatID)
		PreviousState[chatID] = "dnsPSN"

	case "dnsPSN":
		log.Printf("Состояние: %s. Отправка пакетов для самостоятельной загрузки PSN.", state)
		deployment.SendStandaloneDownloadDistributionPSN(bot, chatID)
		PreviousState[chatID] = "standaloneDownloadDistributionPSN"

	case "dnsPGS":
		log.Printf("Состояние: %s. Отправка пакетов для самостоятельной загрузки PGS.", state)
		deployment.SendStandaloneDownloadDistribution(bot, chatID)
		PreviousState[chatID] = "standaloneDownloadDistribution"

	case "standaloneDownloadDistributionPSN":
		log.Printf("Состояние: %s. Отправка сертификатов и ключей PSN.", state)
		deployment.SendCertificatesAndKeysPSN(bot, chatID)
		PreviousState[chatID] = "certificatesAndKeysPSN"

	case "standaloneDownloadDistribution":
		log.Printf("Состояние: %s. Отправка сертификатов и ключей PGS.", state)
		deployment.SendCertificatesAndKeysPGS(bot, chatID)
		PreviousState[chatID] = "certificatesAndKeysPGS"

	case "certificatesAndKeysPSN":
		log.Printf("Состояние: %s. Отправка конфигурации для PSN.", state)
		deployment.SendStandalonePSNConfigure(bot, chatID)
		PreviousState[chatID] = "psnConfigure"

	case "psnConfigure":
		log.Printf("Состояние: %s. Отправка развертывания PSN.", state)
		deployment.SendPSNDeploy(bot, chatID)
		PreviousState[chatID] = "psnDeploy"

	case "certificatesAndKeysPGS":
		log.Printf("Состояние: %s. Отправка конфигурации для PGS.", state)
		deployment.SendStandalonePGSConfigure(bot, chatID)
		PreviousState[chatID] = "pgsConfigure"

	case "pgsConfigure":
		log.Printf("Состояние: %s. Отправка развертывания PGS.", state)
		deployment.SendPGSDeploy(bot, chatID)
		PreviousState[chatID] = "pgsDeploy"

	case "pgsDeploy":
		log.Printf("Состояние: %s. Отправка DNS-опций CO.", state)
		deployment.SendDNSOptionsCO(bot, chatID)
		PreviousState[chatID] = "dnsCO"

	case "dnsCO":
		log.Printf("Состояние: %s. Отправка сертификатов и ключей CO.", state)
		deployment.SendCertificatesAndKeysCO(bot, chatID)
		PreviousState[chatID] = "certificatesAndKeysCO"

	case "certificatesAndKeysCO":
		log.Printf("Состояние: %s. Отправка установки CO.", state)
		deployment.SendCOInstallation(bot, chatID)
		PreviousState[chatID] = "coInstallation"

	case "coInstallation":
		log.Printf("Состояние: %s. Отправка конфигурации CO.", state)
		deployment.SendCOConfigure(bot, chatID)
		PreviousState[chatID] = "coConfigure"

	case "coConfigure":
		log.Printf("Состояние: %s. Отправка развертывания CO.", state)
		deployment.SendCODeploy(bot, chatID)
		PreviousState[chatID] = "coDeploy"

	default:
		log.Printf("Состояние: %s. Неизвестное состояние, действие не выполнено.", state)
	}
}
