package instructions

import (
	"technicalSupportBot/pkg/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendInstallationGuideMailion(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/e16/9uzqg7r0zek2rfa4x63zfstxdwph3y1a/Mailion_1.9_Installation_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}

func SendAdminGuideMailion(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/6a4/lht8xpd3zc13bzlbaz4rvj5u7cpr85v3/Mailion_1.9_Admin_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}

func SendSystemRequirementsMailion(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/f74/b753ow0h94l8gh6s59wg4n11ckhcyzbf/Mailion_1.9_System_Requirements.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}
