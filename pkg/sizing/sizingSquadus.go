package sizing

import (
	"log"
	"os"
	"path/filepath"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// sizingSquadus обрабатывает выбор количества пользователей и отправку PDF-файла
func SizingSquadus(bot *tgbotapi.BotAPI, chatID int64) {
	// Определение кнопок для выбора количества пользователей
	// изменение
	buttons := [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("<50", "<50"),
			tgbotapi.NewInlineKeyboardButtonData("<500", "<500"),
		},
		{
			tgbotapi.NewInlineKeyboardButtonData("<1000", "<1000"),
			tgbotapi.NewInlineKeyboardButtonData("<2000", "<2000"),
		},
		{
			tgbotapi.NewInlineKeyboardButtonData("<3000", "<3000"),
			tgbotapi.NewInlineKeyboardButtonData("<5000", "<5000"),
		},
		{
			tgbotapi.NewInlineKeyboardButtonData("<10000", "<10000"),
			tgbotapi.NewInlineKeyboardButtonData("<20000", "<20000"),
		},
	}

	// Создание клавиатуры с кнопками
	keyboard := tgbotapi.NewInlineKeyboardMarkup(buttons...)

	// Сообщение с инструкциями по выбору
	msg := tgbotapi.NewMessage(chatID, "Пожалуйста, выберите количество пользователей:")
	msg.ReplyMarkup = keyboard

	// Отправляем сообщение пользователю
	if _, err := bot.Send(msg); err != nil {
		log.Println("Ошибка отправки сообщения:", err)
	}
}

// HandleUserSelection обрабатывает нажатие кнопки и отправляет PDF в зависимости от выбора пользователя
func HandleUserSelection(chatID int64, data string, bot *tgbotapi.BotAPI) {
	// Определяем путь к соответствующему PDF-файлу в зависимости от выбора
	var pdfFilePath string
	switch data {
	case "<50":
		pdfFilePath = "/home/admin-msk/MyOfficeConfig/sizingSquadus_50.pdf"
	case "<500":
		pdfFilePath = "/home/admin-msk/MyOfficeConfig/sizingSquadus_500.pdf"
	case "<1000":
		pdfFilePath = "/home/admin-msk/MyOfficeConfig/sizingSquadus_1000.pdf"
	case "<2000":
		pdfFilePath = "/home/admin-msk/MyOfficeConfig/sizingSquadus_2000.pdf"
	case "<3000":
		pdfFilePath = "/home/admin-msk/MyOfficeConfig/sizingSquadus_3000.pdf"
	case "<5000":
		pdfFilePath = "/home/admin-msk/MyOfficeConfig/sizingSquadus_5000.pdf"
	case "<10000":
		pdfFilePath = "/home/admin-msk/MyOfficeConfig/sizingSquadus_10000.pdf"
	case "<20000":
		pdfFilePath = "/home/admin-msk/MyOfficeConfig/sizingSquadus_20000.pdf"
	default:
		msg := tgbotapi.NewMessage(chatID, "Неверный выбор. Попробуйте еще раз.")
		bot.Send(msg)
		return
	}

	// Проверяем, что файл существует
	if _, err := os.Stat(pdfFilePath); os.IsNotExist(err) {
		msg := tgbotapi.NewMessage(chatID, "Извините, файл для выбранного количества пользователей не найден.")
		bot.Send(msg)
		return
	}

	// Отправляем предупреждение перед отправкой файла
	warningMessage := "Данный сайзинг является типовым. Для более детального расчета обратитесь к инженеру или в техническую поддержку."
	warningMsg := tgbotapi.NewMessage(chatID, warningMessage)
	if _, err := bot.Send(warningMsg); err != nil {
		log.Println("Ошибка отправки предупреждения:", err)
	}

	// Отправляем PDF-файл
	pdfFile := tgbotapi.NewDocument(chatID, tgbotapi.FilePath(filepath.Clean(pdfFilePath)))
	if _, err := bot.Send(pdfFile); err != nil {
		log.Println("Ошибка отправки PDF-файла:", err)
	}
}
