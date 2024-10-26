package instructions

import (
	"technicalSupportBot/pkg/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendInstallationGuideMail(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/eb4/wfd8c8ndheus26cfk90dvjc4r53ssly1/MyOffice_Mail_Server_3.1_Installation_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}

func SendAdminGuideMail(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/080/ipwd8vvczh7bw4ilmmwriewy28zkslv5/MyOffice_Mail_3.1_Admin_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}

func SendSystemRequirementsMail(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/3a3/4j4idjbivdp7y953tu4w93opnqmk1qe9/MyOffice_Mail_3.1_System_Requirements.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}
