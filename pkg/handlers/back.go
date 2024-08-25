package handlers

import (
	"log"
	"technicalSupportBot/pkg/deployment"
	"technicalSupportBot/pkg/instructions"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleBackButton обрабатывает нажатие кнопки "Назад"
func HandleBackButton(bot *tgbotapi.BotAPI, chatID int64, sm *StateManager) {
	state := sm.GetState(chatID)
	if state == nil {
		log.Printf("Состояние для chatID %d не найдено", chatID)
		return
	}

	log.Printf("Нажата кнопка \"Назад\" для chatID %d, текущее состояние: %s", chatID, state.Current)

	switch state.Current {
	case "instr", "deploy", "sizing":
		log.Printf("Состояние: %s. Отправка приветственного сообщения и переход на начальный экран.", state.Current)
		sendWelcomeMessage(bot, chatID)
		sm.SetState(chatID, state.Current, "start")

	case "privateCloud", "squadus", "mailion", "mail":
		log.Printf("Состояние: %s. Отправка списка продуктов.", state.Current)
		sendProduct(bot, chatID)
		sm.SetState(chatID, state.Current, "instr")

	case "standalone", "cluster":
		log.Printf("Состояние: %s. Отправка списка продуктов.", state.Current)
		sendProduct(bot, chatID)
		sm.SetState(chatID, state.Current, state.Product)

	case "standaloneDownloadPackages":
		log.Printf("Состояние: %s. Отправка типа инсталляции продуктов.", state.Current)
		sendDeploymentOptions(bot, chatID)
		sm.SetState(chatID, state.Current, "standalone")

	case "requirementsPrivateCloud", "installationGuidePrivateCloud", "adminGuidePrivateCloud":
		log.Printf("Состояние: %s. Отправка инструкций для PrivateCloud.", state.Current)
		sendInstructions(bot, chatID)
		sm.SetState(chatID, "instr", "privateCloud")

	case "pgs", "co":
		log.Printf("Состояние: %s. Отправка опций установки PrivateCloud.", state.Current)
		instructions.SendInstallationGuideOptionsPrivateCloud(bot, chatID)
		sm.SetState(chatID, "privateCloud", "installationGuidePrivateCloud")

	case "requirementsSquadus", "installationGuideSquadus", "adminGuideSquadus":
		log.Printf("Состояние: %s. Отправка инструкций для Squadus.", state.Current)
		sendInstructions(bot, chatID)
		sm.SetState(chatID, "instr", "squadus")

	case "requirementsMailion", "installationGuideMailion", "adminGuideMailion":
		log.Printf("Состояние: %s. Отправка инструкций для Mailion.", state.Current)
		sendInstructions(bot, chatID)
		sm.SetState(chatID, "instr", "mailion")

	case "requirementsMail", "installationGuideMail", "adminGuideMail":
		log.Printf("Состояние: %s. Отправка инструкций для Mail.", state.Current)
		sendInstructions(bot, chatID)
		sm.SetState(chatID, "instr", "mail")

	case "standaloneDownloadDistribution":
		log.Printf("Состояние: %s. Отправка DNS-опций PGS.", state.Current)
		deployment.SendDNSOptionsPGS(bot, chatID)
		sm.SetState(chatID, state.Current, "dnsPGS")

	case "standaloneDownloadDistributionPSN":
		log.Printf("Состояние: %s. Отправка DNS-опций PSN.", state.Current)
		deployment.SendDNSOptionsPSN(bot, chatID)
		sm.SetState(chatID, state.Current, "dnsPSN")

	case "dnsPGS":
		log.Printf("Состояние: %s. Отправка вставки приватного ключа.", state.Current)
		deployment.SendPrivateKeyInsertPrivateCloud(bot, chatID)
		sm.SetState(chatID, state.Current, "privateKeyInsertPrivateCloud")

	case "privateKeyInsertPrivateCloud", "privateKeyInsertPSN":
		log.Printf("Состояние: %s. Отправка пакетов для самостоятельной загрузки.", state.Current)
		deployment.SendStandaloneDownloadPackages(bot, chatID)
		sm.SetState(chatID, state.Current, "standaloneDownloadPackages")

	case "certificatesAndKeysPGS":
		log.Printf("Состояние: %s. Отправка пакетов для самостоятельной загрузки.", state.Current)
		deployment.SendStandaloneDownloadDistribution(bot, chatID)
		sm.SetState(chatID, state.Current, "standaloneDownloadDistribution")

	case "certificatesAndKeysPSN":
		log.Printf("Состояние: %s. Отправка пакетов для самостоятельной загрузки PSN.", state.Current)
		deployment.SendStandaloneDownloadDistributionPSN(bot, chatID)
		sm.SetState(chatID, state.Current, "standaloneDownloadDistributionPSN")

	case "psnConfigure":
		log.Printf("Состояние: %s. Отправка сертификатов и ключей PSN.", state.Current)
		deployment.SendCertificatesAndKeysPSN(bot, chatID)
		sm.SetState(chatID, state.Current, "certificatesAndKeysPSN")

	case "pgsConfigure":
		log.Printf("Состояние: %s. Отправка сертификатов и ключей PGS.", state.Current)
		deployment.SendCertificatesAndKeysPGS(bot, chatID)
		sm.SetState(chatID, state.Current, "certificatesAndKeysPGS")

	case "pgsDeploy":
		log.Printf("Состояние: %s. Отправка конфигурации для PGS.", state.Current)
		deployment.SendStandalonePGSConfigure(bot, chatID)
		sm.SetState(chatID, state.Current, "pgsConfigure")

	case "psnDeploy":
		log.Printf("Состояние: %s. Отправка конфигурации для PSN.", state.Current)
		deployment.SendStandalonePSNConfigure(bot, chatID)
		sm.SetState(chatID, state.Current, "psnConfigure")

	case "coInstallation":
		log.Printf("Состояние: %s. Отправка развертывания PGS.", state.Current)
		deployment.SendPGSDeploy(bot, chatID)
		sm.SetState(chatID, state.Current, "pgsDeploy")

	case "coDeploy":
		log.Printf("Состояние: %s. Отправка конфигурации CO.", state.Current)
		deployment.SendCOConfigure(bot, chatID)
		sm.SetState(chatID, state.Current, "coConfigure")

	default:
		log.Printf("Состояние: %s. Неизвестное состояние, отправка приветственного сообщения и переход на начальный экран.", state.Current)
		sendWelcomeMessage(bot, chatID)
		sm.SetState(chatID, state.Current, "start")
	}
}
