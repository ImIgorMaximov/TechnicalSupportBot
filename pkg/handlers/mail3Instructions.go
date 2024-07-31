package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"technicalSupportBot/pkg/keyboards" 
)

func sendInstallationGuideMail3(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/b9e/rrjck7c57hzsic80ov3wpgl2dc64wl08/MyOffice_Mail_Server_3.0_Installation_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}

func sendAdminGuideMail3(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/c31/io6dlh8692093yaxmewgvw5p3wil2qed/MyOffice_Mail_3.0_Admin_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}

func sendSystemRequirementsMail3(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/73b/jw9hnfqltqw3qmjg47lu33r1su0wdlst/MyOffice_Mail_3.0_System_Requirements.pdf \n")
	bot.Send(msg)
}