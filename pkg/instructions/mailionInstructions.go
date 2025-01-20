/*
Package instructions предоставляет функции для отправки ссылок на различные руководства и инструкции для продукта Mailion через Telegram-бота поддержки.

Функции пакета помогают пользователям легко находить важные документы, такие как руководство по установке, руководство администратора и системные требования для Mailion версии 1.9.
Каждая функция отправляет пользователю в Telegram сообщение со ссылкой на соответствующее руководство и включает кнопку "Назад" для возврата в предыдущее меню.

Функции:
- SendInstallationGuideMailion: Отправляет ссылку на руководство по установке Mailion.
- SendAdminGuideMailion: Отправляет ссылку на руководство по администрированию Mailion.
- SendSystemRequirementsMailion: Отправляет ссылку на руководство по системным требованиям для Mailion.

Автор: Максимов Игорь
Email: imigormaximov@gmail.com
*/

package instructions

import (
	"technicalSupportBot/pkg/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// SendInstallationGuideMailion отправляет ссылку на руководство по установке Mailion
func SendInstallationGuideMailion(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/b56/elzc5aji5b1ggizqw32mf2hkdj61537y/Mailion_2.0_Installation_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard() // Кнопка "Назад"
	bot.Send(msg)
}

// SendAdminGuideMailion отправляет ссылку на руководство по администрированию Mailion
func SendAdminGuideMailion(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/cc3/ydr4czrrxc2335sn15pqsvlqbkft0fuq/Mailion_2.0_Admin_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard() // Кнопка "Назад"
	bot.Send(msg)
}

// SendSystemRequirementsMailion отправляет ссылку на руководство по системным требованиям для Mailion
func SendSystemRequirementsMailion(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/f38/jy5nm1ehb709mh2hu79qrtd7jy230k83/Mailion_2.0_System_Requirements.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard() // Кнопка "Назад"
	bot.Send(msg)
}
