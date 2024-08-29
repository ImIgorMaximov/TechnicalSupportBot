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
	case "reqPsn", "reqPrivateCloud", "reqSquadus":
		deployment.SendStandaloneDownloadPackages(bot, chatID)
		sm.SetState(chatID, state.Current, "standaloneDownloadPackages")
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова SendStandaloneDownloadPackages. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "standaloneDownloadPackages":
		handlePrivateKeyInsert(bot, chatID, sm)
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова handlePrivateKeyInsert. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "privateKeyInsertPrivateCloud":
		deployment.SendDNSOptionsPGS(bot, chatID)
		sm.SetState(chatID, currentState, "dnsPGS")
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова SendDNSOptionsPGS. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "privateKeyInsertPSN":
		deployment.SendDNSOptionsPSN(bot, chatID)
		sm.SetState(chatID, currentState, "dnsPSN")
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова SendDNSOptionsPSN. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "privateKeyInsertSquadus":
		deployment.SendDNSOptionsSquadus(bot, chatID)
		sm.SetState(chatID, currentState, "dnsSquadus")
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова SendDNSOptionsSquadus. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "dnsSquadus":
		deployment.SendStandaloneDownloadDistributionSquadus(bot, chatID)
		sm.SetState(chatID, currentState, "standaloneDownloadDistributionSquadus")
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова SendStandaloneDownloadDistributionSquadus. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "dnsPGS":
		deployment.SendStandaloneDownloadDistributionPrivateCloud(bot, chatID)
		sm.SetState(chatID, currentState, "standaloneDownloadDistributionPrivateCloud")
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова SendStandaloneDownloadDistributionPrivateCloud. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "dnsPSN":
		deployment.SendStandaloneDownloadDistributionPSN(bot, chatID)
		sm.SetState(chatID, currentState, "standaloneDownloadDistributionPSN")
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова SendStandaloneDownloadDistributionPSN. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "standaloneDownloadDistributionPrivateCloud":
		deployment.SendCertificatesAndKeysPGS(bot, chatID)
		sm.SetState(chatID, currentState, "certificatesAndKeysPGS")
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова SendCertificatesAndKeysPGS. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "standaloneDownloadDistributionSquadus":
		deployment.SendCertificatesAndKeysSquadus(bot, chatID)
		sm.SetState(chatID, currentState, "certificatesAndKeysSquadus")
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова SendCertificatesAndKeysSquadus. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "standaloneDownloadDistributionPSN":
		deployment.SendCertificatesAndKeysPSN(bot, chatID)
		sm.SetState(chatID, currentState, "certificatesAndKeysPSN")
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова SendCertificatesAndKeysPSN. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "certificatesAndKeysPGS":
		deployment.SendStandalonePGSConfigure(bot, chatID)
		sm.SetState(chatID, currentState, "pgsConfigure")
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова SendStandalonePGSConfigure. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "certificatesAndKeysSquadus":
		deployment.SendSquadusConfigure(bot, chatID)
		sm.SetState(chatID, currentState, "squadusConfigure")
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова SendSquadusConfigure. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "certificatesAndKeysPSN":
		deployment.SendStandalonePSNConfigure(bot, chatID)
		sm.SetState(chatID, currentState, "psnConfigure")
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова SendStandalonePSNConfigure. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "psnConfigure":
		deployment.SendPSNDeploy(bot, chatID)
		sm.SetState(chatID, currentState, "psnDeploy")
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова SendPSNDeploy. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "pgsConfigure":
		deployment.SendPGSDeploy(bot, chatID)
		sm.SetState(chatID, currentState, "pgsDeploy")
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова SendPGSDeploy. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "squadusConfigure":
		deployment.SendSquadusDeploy(bot, chatID)
		sm.SetState(chatID, currentState, "squadusDeploy")
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова SendSquadusDeploy. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "pgsDeploy":
		deployment.SendDNSOptionsCO(bot, chatID)
		sm.SetState(chatID, currentState, "dnsCO")
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова SendDNSOptionsCO. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "dnsCO":
		deployment.SendCertificatesAndKeysCO(bot, chatID)
		sm.SetState(chatID, currentState, "certificatesAndKeysCO")
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова SendCertificatesAndKeysCO. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "certificatesAndKeysCO":
		deployment.SendCOInstallation(bot, chatID)
		sm.SetState(chatID, currentState, "coInstallation")
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова SendCOInstallation. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "coInstallation":
		deployment.SendCOConfigure(bot, chatID)
		sm.SetState(chatID, currentState, "coConfigure")
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова SendCOConfigure. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "coConfigure":
		deployment.SendCODeploy(bot, chatID)
		sm.SetState(chatID, currentState, "coDeploy")
		updatedState := sm.GetState(chatID)
		log.Printf("После вызова SendCODeploy. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	default:
		log.Printf("Состояние: %s. Неизвестное состояние, действие не выполнено.", currentState)
	}
}
