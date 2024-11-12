/*
Package handlers предоставляет функции для обработки различных команд и кнопок, используемых в техническом боте поддержки.

Функции пакета предназначены для взаимодействия с пользователями в Telegram, помогая им управлять процессами установки и конфигурации различных продуктов.
В частности, функция HandleBackButton обрабатывает нажатие кнопки "Назад", изменяя состояние в зависимости от текущего шага пользователя в процессе установки.

Функция динамически реагирует на текущее состояние пользователя и возвращает его на предыдущий шаг.
Это позволяет пользователю легко вернуться к предыдущим шагам, сохраняя контекст установки и необходимые инструкции.

Автор: Максимов Игорь
Email: imigormaximov@gmail.com
*/

package handlers

import (
	"log"
	"technicalSupportBot/pkg/deployment"
	"technicalSupportBot/pkg/instructions"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleBackButton обрабатывает нажатие кнопки "Назад" и меняет состояние в зависимости от текущего состояния пользователя.
func HandleBackButton(bot *tgbotapi.BotAPI, chatID int64, sm *StateManager) {

	// Получаем текущее состояние пользователя
	state := sm.GetState(chatID)
	if state == nil {
		log.Printf("Состояние для chatID %d не найдено", chatID)
		return
	}

	log.Printf("Нажата кнопка \"Назад\" для chatID %d, текущее состояние: %s", chatID, state.Current)

	switch state.Current {

	// Если текущее состояние связано с продуктом, возвращаемся в Главное меню
	case "privateCloud", "squadus", "mailion", "mail":
		if state.Previous == "standalone" {
			sendWelcomeMessage(bot, chatID)
			sm.SetState(chatID, state.Current, "start")
			sm.SetType(chatID, "")
			updatedState := sm.GetState(chatID)
			log.Printf("После выполнения кнопки Назад sendWelcomeMessage. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)
		} else {
			sendProduct(bot, chatID)
			sm.SetState(chatID, state.Current, state.Action)
			updatedState := sm.GetState(chatID)
			log.Printf("После выполнения кнопки Назад sendProduct. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)
		}

	// Если текущее состояние - тип инсталляции, то возвращаемся к выбору продуктов
	case "standalone", "cluster":
		sendProduct(bot, chatID)
		sm.SetState(chatID, state.Current, state.Product)
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад sendProduct. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	// Если текущее состояние - инструкции, то возвращаемся к типу инсталляции
	case "reqPrivateCloud", "reqPsn", "reqSquadus", "reqMailion":
		sendDeploymentOptions(bot, chatID)
		sm.SetState(chatID, state.Current, state.Type)
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад sendDeploymentOptions. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "standaloneDownloadPackages":

		if state.Product == "privateCloud" {
			deployment.SendStandaloneRequirementsPrivateCloud(bot, chatID)
			sm.SetState(chatID, state.Current, "reqPrivateCloud")
			updatedState := sm.GetState(chatID)
			log.Printf("После выполнения кнопки Назад SendStandaloneRequirementsPrivateCloud. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)
		} else if state.Product == "mail" {
			deployment.SendStandaloneRequirementsPSN(bot, chatID)
			sm.SetState(chatID, state.Current, "reqPsn")
			updatedState := sm.GetState(chatID)
			log.Printf("После выполнения кнопки Назад SendStandaloneRequirementsPSN. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)
		} else if state.Product == "squadus" {
			deployment.SendStandaloneRequirementsSquadus(bot, chatID)
			sm.SetState(chatID, state.Current, "reqSquadus")
			updatedState := sm.GetState(chatID)
			log.Printf("После выполнения кнопки Назад SendStandaloneRequirementsSquadus. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)
		} else if state.Product == "mailion" {
			deployment.SendStandaloneRequirementsMailion(bot, chatID)
			sm.SetState(chatID, state.Current, "reqMailion")
			updatedState := sm.GetState(chatID)
			log.Printf("После выполнения кнопки Назад SendStandaloneRequirementsMailion. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)
		}

	case "requirementsPrivateCloud", "installationGuidePrivateCloud", "adminGuidePrivateCloud":
		sendInstructions(bot, chatID)
		sm.SetState(chatID, state.Action, state.Product)
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад sendInstructions. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "pgs", "co":
		log.Printf("Состояние: %s. Отправка опций установки PrivateCloud.", state.Current)
		instructions.SendInstallationGuideOptionsPrivateCloud(bot, chatID)
		sm.SetState(chatID, state.Product, "installationGuidePrivateCloud")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendInstallationGuideOptionsPrivateCloud. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "requirementsSquadus", "installationGuideSquadus", "adminGuideSquadus":
		sendInstructions(bot, chatID)
		sm.SetState(chatID, state.Action, state.Product)
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад sendInstruction. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "requirementsMailion", "installationGuideMailion", "adminGuideMailion":
		sendInstructions(bot, chatID)
		sm.SetState(chatID, state.Action, state.Product)
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад sendInstruction. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "requirementsMail", "installationGuideMail", "adminGuideMail":
		sendInstructions(bot, chatID)
		sm.SetState(chatID, state.Action, state.Product)
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад sendInstruction. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "standaloneDownloadDistributionPrivateCloud":
		deployment.SendDNSOptionsPGS(bot, chatID)
		sm.SetState(chatID, state.Current, "dnsPGS")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendDNSOptionsPGS. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "standaloneDownloadDistributionPSN":
		deployment.SendDNSOptionsPSN(bot, chatID)
		sm.SetState(chatID, state.Current, "dnsPSN")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendDNSOptionsPSN. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "standaloneDownloadDistributionSquadus":
		deployment.SendDNSOptionsSquadus(bot, chatID)
		sm.SetState(chatID, state.Current, "dnsSquadus")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendDNSOptionsSquadus. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "standaloneDownloadDistributionMailion":
		deployment.SendDNSOptionsMailion(bot, chatID)
		sm.SetState(chatID, state.Current, "dnsMailion")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendDNSOptionsMailion. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "dnsPGS":
		deployment.SendPrivateKeyInsertPrivateCloud(bot, chatID)
		sm.SetState(chatID, state.Current, "privateKeyInsertPrivateCloud")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendPrivateKeyInsertPrivateCloud. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "dnsPSN":
		deployment.SendPrivateKeyInsertPSN(bot, chatID)
		sm.SetState(chatID, state.Current, "privateKeyInsertPSN")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendPrivateKeyInsertPSN. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "dnsSquadus":
		deployment.SendPrivateKeyInsertSquadus(bot, chatID)
		sm.SetState(chatID, state.Current, "privateKeyInsertSquadus")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendPrivateKeyInsertPSN. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "dnsMailion":
		deployment.SendPrivateKeyInsertMailion(bot, chatID)
		sm.SetState(chatID, state.Current, "privateKeyInsertMailion")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendPrivateKeyInsertMailion. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "privateKeyInsertPrivateCloud", "privateKeyInsertPSN", "privateKeyInsertSquadus", "privateKeyInsertMailion":
		log.Printf("Состояние: %s. Отправка пакетов для самостоятельной загрузки.", state.Current)
		deployment.SendStandaloneDownloadPackages(bot, chatID)
		sm.SetState(chatID, state.Current, "standaloneDownloadPackages")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendStandaloneDownloadPackages. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "certificatesAndKeysPGS":
		deployment.SendStandaloneDownloadDistributionPrivateCloud(bot, chatID)
		sm.SetState(chatID, state.Current, "standaloneDownloadDistributionPrivateCloud")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendStandaloneDownloadDistributionPrivateCloud. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "certificatesAndKeysPSN":
		deployment.SendStandaloneDownloadDistributionPSN(bot, chatID)
		sm.SetState(chatID, state.Current, "standaloneDownloadDistributionPSN")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendStandaloneDownloadDistributionPSN. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "certificatesAndKeysSquadus":
		deployment.SendStandaloneDownloadDistributionSquadus(bot, chatID)
		sm.SetState(chatID, state.Current, "standaloneDownloadDistributionSquadus")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendStandaloneDownloadDistributionSquadus. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "certificatesAndKeysMailion":
		deployment.SendStandaloneDownloadDistributionMailion(bot, chatID)
		sm.SetState(chatID, state.Current, "standaloneDownloadDistributionMailion")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendStandaloneDownloadDistributionMailion. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "psnConfigure":
		deployment.SendCertificatesAndKeysPSN(bot, chatID)
		sm.SetState(chatID, state.Current, "certificatesAndKeysPSN")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendCertificatesAndKeysPSN. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "squadusConfigure":
		deployment.SendCertificatesAndKeysSquadus(bot, chatID)
		sm.SetState(chatID, state.Current, "certificatesAndKeysSquadus")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendCertificatesAndKeysSquadus. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "mailionConfigure":
		deployment.SendCertificatesAndKeysMailion(bot, chatID)
		sm.SetState(chatID, state.Current, "certificatesAndKeysMailion")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendCertificatesAndKeysMailion. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "pgsConfigure":
		deployment.SendCertificatesAndKeysPGS(bot, chatID)
		sm.SetState(chatID, state.Current, "certificatesAndKeysPGS")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendCertificatesAndKeysPGS. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "pgsDeploy":
		deployment.SendStandalonePGSConfigure(bot, chatID)
		sm.SetState(chatID, state.Current, "pgsConfigure")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendStandalonePGSConfigure. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "dnsCO":
		deployment.SendPGSDeploy(bot, chatID)
		sm.SetState(chatID, state.Current, "pgsDeploy")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendPGSDeploy. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "certificatesAndKeysCO":
		deployment.SendDNSOptionsCO(bot, chatID)
		sm.SetState(chatID, state.Current, "dnsCO")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendDNSOptionsCO. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "squadusDeploy":
		deployment.SendStandaloneSquadusConfigure(bot, chatID)
		sm.SetState(chatID, state.Current, "squadusConfigure")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendStandaloneSquadusConfigure. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "mailionDeploy":
		deployment.SendStandaloneMailionConfigure(bot, chatID)
		sm.SetState(chatID, state.Current, "mailionConfigure")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendStandaloneMailionConfigure. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "psnDeploy":
		deployment.SendStandalonePSNConfigure(bot, chatID)
		sm.SetState(chatID, state.Current, "psnConfigure")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendStandalonePSNConfigure. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "coInstallation":
		deployment.SendCertificatesAndKeysCO(bot, chatID)
		sm.SetState(chatID, state.Current, "certificatesAndKeysCO")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendPGSDeploy. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "coConfigure":
		deployment.SendCOInstallation(bot, chatID)
		sm.SetState(chatID, state.Current, "coInstallation")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendPGSDeploy. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "coDeploy":
		deployment.SendCOConfigure(bot, chatID)
		sm.SetState(chatID, state.Current, "coConfigure")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад SendCOConfigure. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	case "awaitingMaxUserMailion":
		sendProduct(bot, chatID)
		sm.SetState(chatID, "start", "sizing")
		sm.SetType(chatID, "")
		updatedState := sm.GetState(chatID)
		log.Printf("После выполнения кнопки Назад sendProduct. Текущее состояние: %s, Предыдущее состояние: %s.", updatedState.Current, updatedState.Previous)

	default:
		log.Printf("Состояние: %s. Неизвестное состояние, отправка приветственного сообщения и переход на начальный экран.", state.Current)
		sendWelcomeMessage(bot, chatID)
		sm.SetState(chatID, state.Current, "start")
	}
}
