package instructions

import (
	"technicalSupportBot/pkg/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendInstallationGuideSquadus(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/bc1/z5yney5pfjweu2txsq9djfstpwmcalbb/Squadus_Server_Web_1.6.2_Installation_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}

func SendAdminGuideSquadus(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/a79/blk4bkyttb9kussakcp3oos668k18ht0/Squadus_1.6_Admin_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}

func SendSystemRequirementsSquadus(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/02d/a825sft1g3vyrw8b8nsiqahnbx3gx2iu/Squadus_1.6_System_Requirements.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}
