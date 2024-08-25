package handlers

import (
	"log"
	"technicalSupportBot/pkg/deployment"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleNextStep обрабатывает следующий шаг на основе текущего состояния пользователя
func HandleNextStep(bot *tgbotapi.BotAPI, chatID int64, sm *StateManager) {
	state := sm.GetState(chatID)
	currentState := state.Current

	log.Printf("Обработка следующего шага для chatID %d, текущее состояние: %s; предыдущее состояние: %s", chatID, currentState, state.Previous)

	switch currentState {
	case "reqPsn", "reqPrivateCloud":
		deployment.SendStandaloneDownloadPackages(bot, chatID)
		sm.SetState(chatID, state.Current, "standaloneDownloadPackages")
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова SendStandaloneDownloadPackages. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "standaloneDownloadPackages":
		handlePrivateKeyInsert(bot, chatID, sm)
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова handlePrivateKeyInsert. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "privateKeyInsertPrivateCloud":
		log.Printf("Состояние: %s. Отправка DNS-опций PGS.", currentState)
		deployment.SendDNSOptionsPGS(bot, chatID)
		sm.SetState(chatID, currentState, "dnsPGS")
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова SendDNSOptionsPGS. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "privateKeyInsertPSN":
		log.Printf("Состояние: %s. Отправка DNS-опций PSN.", currentState)
		deployment.SendDNSOptionsPSN(bot, chatID)
		sm.SetState(chatID, currentState, "dnsPSN")
		log.Printf("Текущее состояние: %s, Предыдущее состояние: %s.", currentState, state.Previous)

	case "dnsPSN":
		log.Printf("Состояние: %s. Отправка пакетов для самостоятельной загрузки PSN.", currentState)
		deployment.SendStandaloneDownloadDistributionPSN(bot, chatID)
		sm.SetState(chatID, currentState, "standaloneDownloadDistributionPSN")

	case "dnsPGS":
		log.Printf("Состояние: %s. Отправка пакетов для самостоятельной загрузки PGS.", currentState)
		deployment.SendStandaloneDownloadDistribution(bot, chatID)
		sm.SetState(chatID, currentState, "standaloneDownloadDistribution")

	case "standaloneDownloadDistributionPSN":
		log.Printf("Состояние: %s. Отправка сертификатов и ключей PSN.", currentState)
		deployment.SendCertificatesAndKeysPSN(bot, chatID)
		sm.SetState(chatID, currentState, "certificatesAndKeysPSN")

	case "standaloneDownloadDistribution":
		log.Printf("Состояние: %s. Отправка сертификатов и ключей PGS.", currentState)
		deployment.SendCertificatesAndKeysPGS(bot, chatID)
		sm.SetState(chatID, currentState, "certificatesAndKeysPGS")

	case "certificatesAndKeysPSN":
		log.Printf("Состояние: %s. Отправка конфигурации для PSN.", currentState)
		deployment.SendStandalonePSNConfigure(bot, chatID)
		sm.SetState(chatID, currentState, "psnConfigure")

	case "psnConfigure":
		log.Printf("Состояние: %s. Отправка развертывания PSN.", currentState)
		deployment.SendPSNDeploy(bot, chatID)
		sm.SetState(chatID, currentState, "psnDeploy")

	case "certificatesAndKeysPGS":
		log.Printf("Состояние: %s. Отправка конфигурации для PGS.", currentState)
		deployment.SendStandalonePGSConfigure(bot, chatID)
		sm.SetState(chatID, currentState, "pgsConfigure")

	case "pgsConfigure":
		log.Printf("Состояние: %s. Отправка развертывания PGS.", currentState)
		deployment.SendPGSDeploy(bot, chatID)
		sm.SetState(chatID, currentState, "pgsDeploy")

	case "pgsDeploy":
		log.Printf("Состояние: %s. Отправка DNS-опций CO.", currentState)
		deployment.SendDNSOptionsCO(bot, chatID)
		sm.SetState(chatID, currentState, "dnsCO")

	case "dnsCO":
		log.Printf("Состояние: %s. Отправка сертификатов и ключей CO.", currentState)
		deployment.SendCertificatesAndKeysCO(bot, chatID)
		sm.SetState(chatID, currentState, "certificatesAndKeysCO")

	case "certificatesAndKeysCO":
		log.Printf("Состояние: %s. Отправка установки CO.", currentState)
		deployment.SendCOInstallation(bot, chatID)
		sm.SetState(chatID, currentState, "coInstallation")

	case "coInstallation":
		log.Printf("Состояние: %s. Отправка конфигурации CO.", currentState)
		deployment.SendCOConfigure(bot, chatID)
		sm.SetState(chatID, currentState, "coConfigure")

	case "coConfigure":
		log.Printf("Состояние: %s. Отправка развертывания CO.", currentState)
		deployment.SendCODeploy(bot, chatID)
		sm.SetState(chatID, currentState, "coDeploy")

	default:
		log.Printf("Состояние: %s. Неизвестное состояние, действие не выполнено.", currentState)
	}
}
