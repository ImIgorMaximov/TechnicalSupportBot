/*
Package instructions реализует функции для отправки ссылок на различные руководства и инструкции через Telegram-бота поддержки.

Функции пакета предназначены для того, чтобы предоставить пользователям ссылки на важные документы, такие как руководство по установке, администрированию и системным требованиям для продукта МойОфис Почта.

Каждая функция формирует сообщение с ссылкой на соответствующее руководство и отправляет его пользователю в Telegram, добавляя клавишу "Назад" для удобства навигации.

Функции:
- SendInstallationGuideMail: Отправляет ссылку на руководство по установке МойОфис Почта.
- SendAdminGuideMail: Отправляет ссылку на руководство по администрированию МойОфис Почта.
- SendSystemRequirementsMail: Отправляет ссылку на руководство по системным требованиям для МойОфис Почта.

Автор: Максимов Игорь
Email: imigormaximov@gmail.com
*/

package instructions

import (
	"technicalSupportBot/pkg/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// SendInstallationGuideMail отправляет ссылку на руководство по установке МойОфис Почта
func SendInstallationGuideMail(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/eb4/wfd8c8ndheus26cfk90dvjc4r53ssly1/MyOffice_Mail_Server_3.1_Installation_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}

// SendAdminGuideMail отправляет ссылку на руководство по администрированию МойОфис Почта
func SendAdminGuideMail(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/080/ipwd8vvczh7bw4ilmmwriewy28zkslv5/MyOffice_Mail_3.1_Admin_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}

// SendSystemRequirementsMail отправляет ссылку на руководство по системным требованиям для МойОфис Почта
func SendSystemRequirementsMail(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/3a3/4j4idjbivdp7y953tu4w93opnqmk1qe9/MyOffice_Mail_3.1_System_Requirements.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard()
	bot.Send(msg)
}
