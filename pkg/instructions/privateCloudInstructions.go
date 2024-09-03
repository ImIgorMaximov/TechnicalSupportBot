package instructions

import (
	"technicalSupportBot/pkg/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendInstallationGuideOptionsPrivateCloud(bot *tgbotapi.BotAPI, chatID int64) {

	chooseComponent := "Частное облако состоит из двух компонентов: PGS (Система хранения данных) и CO (Система редактирования и совместной работы)\n" +
		"Выберите компонент:\n" +
		"- PGS \n" +
		"- CO \n"
	msg := tgbotapi.NewMessage(chatID, chooseComponent)
	msg.ReplyMarkup = keyboards.GetInstallationGuideKeyboard()
	bot.Send(msg)
}

func SendPGSInstallationGuide(bot *tgbotapi.BotAPI, chatID int64) {

	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/7d2/nlof37tjayp6yhqedq8xkrour82m1wef/MyOffice_PGS_3.1_Installation_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}

func SendCOInstallationGuide(bot *tgbotapi.BotAPI, chatID int64) {

	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/05c/137yo7qojdz3hm5k46ngil1nf0opt52p/MyOffice_Private_Cloud_3.0_CO_Installation_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}

func SendAdminGuidePrivateCloud(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/7d5/pibun6pza2j67c2cf9oz6dbiqtmw0aly/MyOffice_CO_PGS_3.1_Admin_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}

func SendSystemRequirementsPivateCloud(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/8c0/aa0g1jhk9phcxh0229zq1fli11qlrtpb/MyOffice_CO_PGS_3.1_System_Requirements.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}
