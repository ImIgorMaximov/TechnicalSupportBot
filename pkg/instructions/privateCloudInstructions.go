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

	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/171/x2vqm7n0zp0jg1qjwsieymtp5pemtcg1/MyOffice_Private_Cloud_3.0_PGS_Installation_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}

func SendCOInstallationGuide(bot *tgbotapi.BotAPI, chatID int64) {

	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/05c/137yo7qojdz3hm5k46ngil1nf0opt52p/MyOffice_Private_Cloud_3.0_CO_Installation_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}

func SendAdminGuidePrivateCloud(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/bc7/a30e4keqfke8h5m7r8mgx4asur3oedlf/MyOffice_Private_Cloud_3.0_Admin_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}

func SendSystemRequirementsPivateCloud(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/e09/ejjo29n32sj1f93bwyoa5y0upppfbuux/MyOffice_Private_Cloud_3.0_System_Requirements.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}
