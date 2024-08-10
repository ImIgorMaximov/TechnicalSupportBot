package handlers

import (
	"fmt"
	"log"

	"technicalSupportBot/pkg/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xuri/excelize/v2"
)

// Глобальная переменная для хранения введённых данных от пользователей
var mailInputValues = make(map[int64][]string)

func HandleSizingMailStandalone(bot *tgbotapi.BotAPI, chatID int64) {
	// Запрос данных у пользователя
	msg := tgbotapi.NewMessage(chatID, "Введите количество пользователей (Например, 50):")
	bot.Send(msg)
	previousState[chatID] = "awaitingUserCountMail"
}

// HandleUserInputMail обрабатывает ввод пользователя
func HandleUserInputMail(bot *tgbotapi.BotAPI, chatID int64, userInput string) {
	log.Println("Текущее состояние перед обработкой ввода:", previousState[chatID])
	log.Println("Полученное значение от пользователя:", userInput)

	switch previousState[chatID] {
	case "awaitingUserCountMail":
		mailInputValues[chatID] = []string{userInput} // Сохраняем первое значение
		msg := tgbotapi.NewMessage(chatID, "Введите дисковую квоту пользователей в почте (ГБ) (Например, 2):")
		bot.Send(msg)
		previousState[chatID] = "awaitingDiskQuotaMail" // Обновляем состояние
		log.Println("Состояние изменено на awaitingDiskQuotaMail")

	case "awaitingDiskQuotaMail":
		mailInputValues[chatID] = append(mailInputValues[chatID], userInput) // Сохраняем второе значение
		msg := tgbotapi.NewMessage(chatID, "Введите количество писем в сутки на пользователя (Например, 100):")
		bot.Send(msg)
		previousState[chatID] = "awaitingEmailsPerDayMail" // Обновляем состояние
		log.Println("Состояние изменено на awaitingEmailsPerDayMail")

	case "awaitingEmailsPerDayMail":
		mailInputValues[chatID] = append(mailInputValues[chatID], userInput) // Сохраняем третье значение
		msg := tgbotapi.NewMessage(chatID, "Введите коэффициент спама (Например, 0.1):")
		bot.Send(msg)
		previousState[chatID] = "awaitingSpamCoefficientMail" // Обновляем состояние
		log.Println("Состояние изменено на awaitingSpamCoefficientMail")

	case "awaitingSpamCoefficientMail":
		mailInputValues[chatID] = append(mailInputValues[chatID], userInput) // Сохраняем четвертое значение
		log.Println("Все данные от пользователя получены:", mailInputValues[chatID])

		// После получения всех значений выполняем расчет
		calculateAndSendMailSizing(bot, chatID)
		log.Println("Результаты расчета отправлены пользователю")

	default:
		msg := tgbotapi.NewMessage(chatID, "Ошибка: неизвестное состояние.")
		bot.Send(msg)
		log.Println("Ошибка: состояние неизвестно")
	}
}

// calculateAndSendMailSizing выполняет расчет и отправляет результат пользователю
func calculateAndSendMailSizing(bot *tgbotapi.BotAPI, chatID int64) {
	// Открытие файла Excel
	filePath := "/home/admin-msk/Documents/sizingPSNStandalone.xlsx"
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Println("Ошибка открытия файла:", err)
		msg := tgbotapi.NewMessage(chatID, "Произошла ошибка при открытии файла.")
		bot.Send(msg)
		return
	}
	defer f.Close()

	// Заполнение ячеек данными
	err = f.SetCellValue("PSN", "D2", mailInputValues[chatID][0])
	err = f.SetCellValue("PSN", "D7", mailInputValues[chatID][1])
	err = f.SetCellValue("PSN", "D8", mailInputValues[chatID][2])
	err = f.SetCellValue("PSN", "D9", mailInputValues[chatID][3])
	if err != nil {
		log.Println("Ошибка записи в файл:", err)
		msg := tgbotapi.NewMessage(chatID, "Произошла ошибка при записи в файл.")
		bot.Send(msg)
		return
	}

	// Сохранение изменений
	if err := f.Save(); err != nil {
		log.Println("Ошибка сохранения файла:", err)
		msg := tgbotapi.NewMessage(chatID, "Произошла ошибка при сохранении файла.")
		bot.Send(msg)
		return
	}

	// Извлечение результатов
	vmCount, _ := f.GetCellValue("PSN", "C16")
	cpu, _ := f.GetCellValue("PSN", "D16")
	ram, _ := f.GetCellValue("PSN", "E16")
	ssd, _ := f.GetCellValue("PSN", "F16")

	// Отправка результата пользователю
	resultMsg := fmt.Sprintf(
		"Результаты расчета сайзинга для Почты (PSN) Standalone:\n\n"+
			"Количество VM: %s\n"+
			"CPU: %s\n"+
			"RAM: %s GB\n"+
			"SSD: %s GB",
		vmCount, cpu, ram, ssd,
	)
	msg := tgbotapi.NewMessage(chatID, resultMsg)
	bot.Send(msg)

	keyboard := keyboards.GetMainMenuKeyboard()
	msgWithKeyboard := tgbotapi.NewMessage(chatID, "Выберите следующую опцию:")
	msgWithKeyboard.ReplyMarkup = keyboard
	if _, err := bot.Send(msgWithKeyboard); err != nil {
		log.Printf("Ошибка отправки клавиатуры: %v", err)
	}
	// Очистка состояния
	previousState[chatID] = "sizingResultProvided"
	mailInputValues[chatID] = nil
}
