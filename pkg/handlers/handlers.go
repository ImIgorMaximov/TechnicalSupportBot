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

	log.Printf("Получено сообщение от chatID %d: %s", chatID, text)

	switch text {
	case "/start", "В главное меню":
		sendWelcomeMessage(bot, chatID)
		PreviousState[chatID] = "start"

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
		HandleBackButton(bot, chatID)

	case "Связаться с инженером тех. поддержки":
		sendSupportEngineerContact(bot, chatID)

	case "Standalone":
		handleStandalone(bot, chatID)

	case "Готово":
		HandleNextStep(bot, chatID)

	case "Запустить деплой":
		HandleNextStep(bot, chatID)

	case "Проверить корректность сертификатов и ключа":
		sendIsCertificates(bot, chatID)

	case "Описание ролей":
		sendRoleDescriptionsPrivateCloudCluster2k(bot, chatID)

	case "Пример конфига PGS - hosts.yml":
		sendConfigFile(bot, chatID, "/home/admin-msk/MyOfficeConfig/hostsPGS.yml", "hostsPGS.yml")

	case "Пример конфига PSN - hosts.yml":
		sendConfigFile(bot, chatID, "/home/admin-msk/MyOfficeConfig/hostsPSN.yml", "hostsPSN.yml")

	case "Пример конфига CO - main.yml":
		sendConfigFile(bot, chatID, "/home/admin-msk/MyOfficeConfig/mainCO.yml", "mainCO.yml")

	case "Пример конфига CO - hosts.yml":
		sendConfigFile(bot, chatID, "/home/admin-msk/MyOfficeConfig/hostsCO.yml", "hostsCO.yml")

	case "Далее":
		HandleNextStep(bot, chatID)

	case "Установка CO":
		HandleNextStep(bot, chatID)

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

	log.Printf("handleStandalone: chatID %d, action %s, previousState %s", chatID, action, PreviousState[chatID])

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

	log.Printf("handleCluster: chatID %d, action %s, previousState %s", chatID, action, PreviousState[chatID])

	if action == "sizing" {
		SendClusterRangeKeyboard(bot, chatID)
		PreviousState[chatID] = "clusterSelection"
	} else if action == "deploy" {
		msg := tgbotapi.NewMessage(chatID, "Извините, раздел находится в разработке 😢")
		bot.Send(msg)
	}
}

func handleClusterUserRange(bot *tgbotapi.BotAPI, chatID int64, userRange string) {
	log.Printf("handleClusterUserRange: chatID %d, userRange %s", chatID, userRange)

	switch userRange {
	case "<2k":
		PreviousState[chatID] = "awaitingClusterMoreThan2kInput"
		sizing.HandleClusterMoreThan2k(bot, chatID)
	default:
		msg := tgbotapi.NewMessage(chatID, "Выберите корректный диапазон пользователей.")
		bot.Send(msg)
	}
}

func handleDefaultState(bot *tgbotapi.BotAPI, chatID int64, text string) {
	log.Printf("handleDefaultState: chatID %d, previousState %s, text %s", chatID, PreviousState[chatID], text)

	if PreviousState[chatID] == "awaitingUserCountPrivateCloud" {
		sizing.HandleUserInputPrivateCloud(bot, chatID, text)
	} else if PreviousState[chatID] == "awaitingUserCountMail" {
		sizing.HandleUserInputMail(bot, chatID, text)
	} else if PreviousState[chatID] == "awaitingClusterMoreThan2kInput" {
		sizing.HandleClusterMoreThan2kInput(bot, chatID, text)
	}
}

func handlePrivateKeyInsert(bot *tgbotapi.BotAPI, chatID int64) {
	log.Printf("handlePrivateKeyInsert: chatID %d, previousState %s", chatID, PreviousState[chatID])

	if PreviousState[chatID] == "reqPrivateCloud" {
		deployment.SendPrivateKeyInsert(bot, chatID)
		PreviousState[chatID] = "privateKeyInsert"
	} else if PreviousState[chatID] == "reqPsn" {
		deployment.SendPrivateKeyInsertPSN(bot, chatID)
		PreviousState[chatID] = "privateKeyInsertPSN"
	}
}

func handlePrivateCloud(bot *tgbotapi.BotAPI, chatID int64) {
	log.Printf("handlePrivateCloud: chatID %d, previousState %s", chatID, PreviousState[chatID])

	if PreviousState[chatID] == "instr" {
		sendInstructions(bot, chatID)
		PreviousState[chatID] = "privateCloud"
	} else if PreviousState[chatID] == "deploy" || PreviousState[chatID] == "sizing" {
		sendDeploymentOptions(bot, chatID)
		PreviousState[chatID] = "privateCloud"
	}
}

func handleMail(bot *tgbotapi.BotAPI, chatID int64) {
	log.Printf("handleMail: chatID %d, previousState %s", chatID, PreviousState[chatID])

	if PreviousState[chatID] == "instr" {
		sendInstructions(bot, chatID)
		PreviousState[chatID] = "mail"
	} else if PreviousState[chatID] == "deploy" || PreviousState[chatID] == "sizing" {
		sendDeploymentOptions(bot, chatID)
		PreviousState[chatID] = "mail"
	}
}

func handleSystemRequirements(bot *tgbotapi.BotAPI, chatID int64) {
	log.Printf("handleSystemRequirements: chatID %d, previousState %s", chatID, PreviousState[chatID])

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
	log.Printf("handleInstallationGuide: chatID %d, previousState %s", chatID, PreviousState[chatID])

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
	log.Printf("handleAdminGuide: chatID %d, previousState %s", chatID, PreviousState[chatID])

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
