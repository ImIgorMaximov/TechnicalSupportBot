package handlers

import (
	"technicalSupportBot/pkg/deployment"
	"technicalSupportBot/pkg/instructions"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleBackButton(bot *tgbotapi.BotAPI, chatID int64) {
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
