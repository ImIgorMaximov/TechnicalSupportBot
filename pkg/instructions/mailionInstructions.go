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
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/e16/9uzqg7r0zek2rfa4x63zfstxdwph3y1a/Mailion_1.9_Installation_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard() // Кнопка "Назад"
	bot.Send(msg)
}

// SendAdminGuideMailion отправляет ссылку на руководство по администрированию Mailion
func SendAdminGuideMailion(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/6a4/lht8xpd3zc13bzlbaz4rvj5u7cpr85v3/Mailion_1.9_Admin_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard() // Кнопка "Назад"
	bot.Send(msg)
}

// SendSystemRequirementsMailion отправляет ссылку на руководство по системным требованиям для Mailion
func SendSystemRequirementsMailion(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/f74/b753ow0h94l8gh6s59wg4n11ckhcyzbf/Mailion_1.9_System_Requirements.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard() // Кнопка "Назад"
	bot.Send(msg)
}
