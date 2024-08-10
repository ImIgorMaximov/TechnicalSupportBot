package instructions

import (
	"technicalSupportBot/pkg/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendInstallationGuideSquadus(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/e28/jm16tlpxh8ws3stx2ojdw0ru0fjhll3k/Squadus_Server_Web_1.5_Installation_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}

func SendAdminGuideSquadus(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/ea3/pjcqlet1o7mb75c2f7voy04dktbb01v3/Squadus_1.5_Admin_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}

func SendSystemRequirementsSquadus(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/3be/swmsya6ze6kf4cuhgx61b3dmfv1w8xyr/Squadus_1.5_System_Requirements.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}
