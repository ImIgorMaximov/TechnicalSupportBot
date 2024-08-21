package handlers

import (
	"technicalSupportBot/pkg/deployment"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleNextStep(bot *tgbotapi.BotAPI, chatID int64) {
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
