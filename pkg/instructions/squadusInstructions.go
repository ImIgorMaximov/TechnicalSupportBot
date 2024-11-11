/*
Package instructions предоставляет функции для отправки пользователям ссылок на руководства по установке, администрированию и системным требованиям для сервера Squadus.

Функции пакета позволяют отправлять пользователям в Telegram ссылки на соответствующие документы для помощи в установке и настройке сервера Squadus.
Включает ссылки на руководство по установке, руководство администратора и системные требования.

Функции:
- SendInstallationGuideSquadus: Отправляет ссылку на руководство по установке для Squadus.
- SendAdminGuideSquadus: Отправляет ссылку на руководство администратора для Squadus.
- SendSystemRequirementsSquadus: Отправляет ссылку на системные требования для Squadus.

Автор: Максимов Игорь
Email: imigormaximov@gmail.com
*/

package instructions

import (
	"technicalSupportBot/pkg/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// SendInstallationGuideSquadus отправляет ссылку на руководство по установке для Squadus.
func SendInstallationGuideSquadus(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/bc1/z5yney5pfjweu2txsq9djfstpwmcalbb/Squadus_Server_Web_1.6.2_Installation_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard() // Кнопка "Назад"
	bot.Send(msg)
}

// SendAdminGuideSquadus отправляет ссылку на руководство администратора для Squadus.
func SendAdminGuideSquadus(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/a79/blk4bkyttb9kussakcp3oos668k18ht0/Squadus_1.6_Admin_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard() // Кнопка "Назад"
	bot.Send(msg)
}

// SendSystemRequirementsSquadus отправляет ссылку на системные требования для Squadus.
func SendSystemRequirementsSquadus(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/02d/a825sft1g3vyrw8b8nsiqahnbx3gx2iu/Squadus_1.6_System_Requirements.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard() // Кнопка "Назад"
	bot.Send(msg)
}
