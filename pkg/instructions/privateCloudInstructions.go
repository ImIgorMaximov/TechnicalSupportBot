/*
Package instructions предоставляет функции для отправки пользователям ссылок на руководства по установке и администрированию компонентов Частного облака.

Функции пакета позволяют пользователям выбрать интересующий компонент Частного облака — систему хранения данных (PGS) или систему редактирования и совместной работы (CO) — и получить соответствующее руководство.
Также предоставляются ссылки на руководство администратора и на системные требования.

Функции:
- SendInstallationGuideOptionsPrivateCloud: Отправляет сообщение с выбором компонентов Частного облака.
- SendPGSInstallationGuide: Отправляет ссылку на руководство по установке для компонента PGS.
- SendCOInstallationGuide: Отправляет ссылку на руководство по установке для компонента CO.
- SendAdminGuidePrivateCloud: Отправляет ссылку на руководство администратора для Частного облака.
- SendSystemRequirementsPrivateCloud: Отправляет ссылку на руководство по системным требованиям для Частного облака.

Автор: Максимов Игорь
Email: imigormaximov@gmail.com
*/

package instructions

import (
	"technicalSupportBot/pkg/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// SendInstallationGuideOptionsPrivateCloud отправляет сообщение с выбором компонентов Частного облака
func SendInstallationGuideOptionsPrivateCloud(bot *tgbotapi.BotAPI, chatID int64) {

	chooseComponent := "Частное облако состоит из двух компонентов: PGS (Система хранения данных) и CO (Система редактирования и совместной работы)\n" +
		"Выберите компонент:\n" +
		"- PGS \n" +
		"- CO \n"
	msg := tgbotapi.NewMessage(chatID, chooseComponent)
	msg.ReplyMarkup = keyboards.GetInstallationGuideKeyboard() // Клавиатура с кнопками для выбора компонента
	bot.Send(msg)
}

// SendPGSInstallationGuide отправляет ссылку на руководство по установке для компонента PGS
func SendPGSInstallationGuide(bot *tgbotapi.BotAPI, chatID int64) {

	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/e55/4zbxzq7p4g4zeb3et5k793ouzvksdmqi/MyOffice_PGS_3.1_Installation_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard() // Кнопка "Назад"
	bot.Send(msg)
}

// SendCOInstallationGuide отправляет ссылку на руководство по установке для компонента CO
func SendCOInstallationGuide(bot *tgbotapi.BotAPI, chatID int64) {

	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/afc/8f0559sw11uoy033r8tuwvfimimnrlt9/MyOffice_CO_3.1_Installation_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard() // Кнопка "Назад"
	bot.Send(msg)
}

// SendAdminGuidePrivateCloud отправляет ссылку на руководство администратора для Частного облака
func SendAdminGuidePrivateCloud(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/06a/s042042xdt7la4pn4wa17hukf13l4ohr/MyOffice_CO_PGS_3.1_Admin_Guide.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard() // Кнопка "Назад"
	bot.Send(msg)
}

// SendSystemRequirementsPrivateCloud отправляет ссылку на руководство по системным требованиям для Частного облака
func SendSystemRequirementsPrivateCloud(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "https://support.myoffice.ru/upload/iblock/8c0/aa0g1jhk9phcxh0229zq1fli11qlrtpb/MyOffice_CO_PGS_3.1_System_Requirements.pdf \n")
	msg.ReplyMarkup = keyboards.GetBackKeyboard() // Кнопка "Назад"
	bot.Send(msg)
}
