package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"technicalSupportBot/pkg/keyboards" 
)

func sendInstallationGuideMailion(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/61a/ycy96cdtr5s4n74y9pqj4uj0jtzvv8er/Mailion_1.9_Installation_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}

func sendAdminGuideMailion(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/dd6/ectn8biypiful7glx1dg7iqnfkatjha8/Mailion_1.9_Admin_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}

func sendSystemRequirementsMailion(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/f74/b753ow0h94l8gh6s59wg4n11ckhcyzbf/Mailion_1.9_System_Requirements.pdf \n")
	bot.Send(msg)
}