package handlers

import (
	"fmt"
	"log"

	"technicalSupportBot/pkg/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xuri/excelize/v2"
)

// Глобальная переменная для хранения введённых данных от пользователей
var userInputValues = make(map[int64][]string)

func HandleSizingPrivateCloudStandalone(bot *tgbotapi.BotAPI, chatID int64) {
	// Запрос данных у пользователя
	msg := tgbotapi.NewMessage(chatID, "Введите количество пользователей (Например, 50):")
	bot.Send(msg)
	previousState[chatID] = "awaitingUserCountPrivateCloud"
}

// HandleUserInputPrivateCloud обрабатывает ввод пользователя
func HandleUserInputPrivateCloud(bot *tgbotapi.BotAPI, chatID int64, userInput string) {
	log.Println("Текущее состояние перед обработкой ввода:", previousState[chatID])
	log.Println("Полученное значение от пользователя:", userInput)

	switch previousState[chatID] {
	case "awaitingUserCountPrivateCloud":
		userInputValues[chatID] = []string{userInput} // Сохраняем первое значение
		msg := tgbotapi.NewMessage(chatID, "Введите количество одновременно активных пользователей (Например, 10):")
		bot.Send(msg)
		previousState[chatID] = "awaitingActiveUserCountPrivateCloud" // Обновляем состояние
		log.Println("Состояние изменено на awaitingActiveUserCountPrivateCloud")

	case "awaitingActiveUserCountPrivateCloud":
		userInputValues[chatID] = append(userInputValues[chatID], userInput) // Сохраняем второе значение
		msg := tgbotapi.NewMessage(chatID, "Введите количество редактируемых документов одновременно (Например, 10):")
		bot.Send(msg)
		previousState[chatID] = "awaitingDocumentCountPrivateCloud" // Обновляем состояние
		log.Println("Состояние изменено на awaitingDocumentCountPrivateCloud")

	case "awaitingDocumentCountPrivateCloud":
		userInputValues[chatID] = append(userInputValues[chatID], userInput) // Сохраняем третье значение
		msg := tgbotapi.NewMessage(chatID, "Введите дисковую квоту пользователей в хранилище (ГБ) (Например, 2):")
		bot.Send(msg)
		previousState[chatID] = "awaitingStorageQuotaPrivateCloud" // Обновляем состояние
		log.Println("Состояние изменено на awaitingStorageQuotaPrivateCloud")

	case "awaitingStorageQuotaPrivateCloud":
		userInputValues[chatID] = append(userInputValues[chatID], userInput) // Сохраняем четвертое значение
		log.Println("Все данные от пользователя получены:", userInputValues[chatID])

		// После получения всех значений выполняем расчет
		calculateAndSendSizing(bot, chatID)
		log.Println("Результаты расчета отправлены пользователю")

	default:
		msg := tgbotapi.NewMessage(chatID, "Ошибка: неизвестное состояние.")
		bot.Send(msg)
		log.Println("Ошибка: состояние неизвестно")
	}
}

// HandleNextInput помогает запрашивать следующие данные
func HandleNextInput(bot *tgbotapi.BotAPI, chatID int64, userInput string, nextMessage string, nextState string) {
	userInputValues[chatID] = append(userInputValues[chatID], userInput)
	msg := tgbotapi.NewMessage(chatID, nextMessage)
	bot.Send(msg)
	previousState[chatID] = nextState
}

// calculateAndSendSizing выполняет расчет и отправляет результат пользователю
func calculateAndSendSizing(bot *tgbotapi.BotAPI, chatID int64) {
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
	err = f.SetCellValue("PSN", "D2", userInputValues[chatID][0])
	err = f.SetCellValue("PSN", "F4", userInputValues[chatID][1])
	err = f.SetCellValue("PSN", "F5", userInputValues[chatID][2])
	err = f.SetCellValue("PSN", "D6", userInputValues[chatID][3])
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
	operatorVM, _ := f.GetCellValue("PSN", "C13")
	operatorCPU, _ := f.GetCellValue("PSN", "D13")
	operatorRAM, _ := f.GetCellValue("PSN", "E13")
	operatorSSD, _ := f.GetCellValue("PSN", "F13")

	coVM, _ := f.GetCellValue("PSN", "C14")
	coCPU, _ := f.GetCellValue("PSN", "D14")
	coRAM, _ := f.GetCellValue("PSN", "E14")
	coSSD, _ := f.GetCellValue("PSN", "F14")

	pgsVM, _ := f.GetCellValue("PSN", "C15")
	pgsCPU, _ := f.GetCellValue("PSN", "D15")
	pgsRAM, _ := f.GetCellValue("PSN", "E15")
	pgsSSD, _ := f.GetCellValue("PSN", "F15")

	// Отправка результата пользователю
	resultMsg := fmt.Sprintf(
		"Результаты расчета сайзинга для продукта Частное Облако Standalone:\n\n"+
			"ВМ Operator: кол-во ВМ - %s, CPU - %s, RAM - %s ГБ, SSD - %s ГБ\n"+
			"Компонент CO: кол-во ВМ - %s, CPU - %s, RAM - %s ГБ, SSD - %s ГБ\n"+
			"Компонент PGS: кол-во ВМ - %s, CPU - %s, RAM - %s ГБ, SSD - %s ГБ",
		operatorVM, operatorCPU, operatorRAM, operatorSSD,
		coVM, coCPU, coRAM, coSSD,
		pgsVM, pgsCPU, pgsRAM, pgsSSD,
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
	userInputValues[chatID] = nil
}
