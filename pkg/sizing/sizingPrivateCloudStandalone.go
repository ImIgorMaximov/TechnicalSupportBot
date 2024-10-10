package sizing

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"technicalSupportBot/pkg/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xuri/excelize/v2"
)

// Определение глобальной переменной для хранения пользовательских вводов по состояниям
var userInputValues = make(map[int64][]string)

// Глобальная переменная для хранения текущего состояния
var currentState string

// Обработка ввода пользователя и управление состоянием
func HandleUserInput(bot *tgbotapi.BotAPI, chatID int64, state *string, text string) {
	currentState = *state
	// var userInput string // Переменная для хранения ввода пользователя

	log.Printf("Получен ввод от пользователя. Текущее состояние: %s", currentState)

	switch currentState {
	case "privateCloud":
		log.Printf("Обработка состояния: %s.", currentState)
		currentState = "awaitingMaxUserCountPrivateCloud"
		userInputValues[chatID] = []string{} // Инициализация мапы для пользователя
		HandleNextInput(bot, chatID, "", "Введите максимальное количество пользователей (например, 50):", "awaitingMaxUserCountPrivateCloud")

	case "awaitingMaxUserCountPrivateCloud":
		log.Println("Обработка состояния: awaitingMaxUserCountPrivateCloud")
		userInputValues[chatID] = append(userInputValues[chatID], text) // Сохраняем ввод
		currentState = "awaitingActiveUserCountPrivateCloud"
		HandleNextInput(bot, chatID, text, "Введите количество одновременно активных пользователей (например, 10):", "awaitingActiveUserCountPrivateCloud")

	case "awaitingActiveUserCountPrivateCloud":
		log.Println("Обработка состояния: awaitingActiveUserCountPrivateCloud")
		userInputValues[chatID] = append(userInputValues[chatID], text) // Сохраняем ввод
		currentState = "awaitingDocumentCountPrivateCloud"
		HandleNextInput(bot, chatID, text, "Введите количество редактируемых документов (например, 200):", "awaitingDocumentCountPrivateCloud")

	case "awaitingDocumentCountPrivateCloud":
		log.Println("Обработка состояния: awaitingDocumentCountPrivateCloud")
		userInputValues[chatID] = append(userInputValues[chatID], text) // Сохраняем ввод
		currentState = "awaitingStorageQuotaPrivateCloud"
		HandleNextInput(bot, chatID, text, "Введите дисковую квоту пользователей в хранилище (ГБ) (например, 2):", "awaitingStorageQuotaPrivateCloud")

	case "awaitingStorageQuotaPrivateCloud":
		log.Println("Обработка состояния: awaitingStorageQuotaPrivateCloud")
		userInputValues[chatID] = append(userInputValues[chatID], text) // Сохраняем ввод
		log.Println("Все данные от пользователя получены:", userInputValues[chatID])

		// После получения всех значений выполняем расчет
		calculateAndSendSizing(bot, chatID, userInputValues[chatID])
		log.Println("Результаты расчета отправлены пользователю")

	default:
		log.Printf("Ошибка: Неизвестное состояние или некорректный ввод. Состояние: %s", currentState)
		sendErrorMessage(bot, chatID, "Ошибка: неизвестное состояние или некорректный ввод.")
	}

	// Обновление состояния после обработки
	*state = currentState
}

// HandleNextInput помогает запрашивать следующие данные и обновлять состояние
func HandleNextInput(bot *tgbotapi.BotAPI, chatID int64, userInput string, nextMessage string, nextState string) {
	currentState = nextState
	// userInputValues[chatID] = append(userInputValues[chatID], userInput) // Сохраняем текущий ввод
	msg := tgbotapi.NewMessage(chatID, nextMessage) // Создаём сообщение для пользователя
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Ошибка при отправке сообщения: %s", err) // Логируем ошибку
	}
}

// calculateAndSendSizing выполняет расчет и отправляет результат пользователю
func calculateAndSendSizing(bot *tgbotapi.BotAPI, chatID int64, userInputValues []string) {
	// Открытие файла Excel
	filePath := "/home/admin-msk/Documents/sizingPrivateCloudPSNStandalone.xlsx"
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		sendErrorMessage(bot, chatID, "Произошла ошибка при открытии файла.")
		log.Println("Ошибка открытия файла:", err)
		return
	}
	defer f.Close()

	// Заполнение ячеек данными
	err = fillExcelFile(f, userInputValues)
	if err != nil {
		sendErrorMessage(bot, chatID, "Произошла ошибка при записи в файл.")
		log.Println("Ошибка записи в файл:", err)
		return
	}

	// Сохранение изменений
	if err := f.Save(); err != nil {
		sendErrorMessage(bot, chatID, "Произошла ошибка при сохранении файла.")
		log.Println("Ошибка сохранения файла:", err)
		return
	}

	// Извлечение результатов
	sendSizingResults(bot, chatID, f)
}

// fillExcelFile заполняет Excel данными пользователя
func fillExcelFile(f *excelize.File, userInputValues []string) error {
	err := f.SetCellValue("PSN", "D4", userInputValues[0]) // Макс. количество пользователей
	err = f.SetCellValue("PSN", "F6", userInputValues[1])  // Кол-во активных пользователей
	err = f.SetCellValue("PSN", "F7", userInputValues[2])  // Кол-во редактируемых документов
	err = f.SetCellValue("PSN", "D8", userInputValues[3])  // Дисковая квота пользователя
	return err
}

// sendSizingResults извлекает результаты из Excel и отправляет их пользователю
func sendSizingResults(bot *tgbotapi.BotAPI, chatID int64, f *excelize.File) {
	operatorVM, _ := f.GetCellValue("PSN", "C15")
	operatorCPU, _ := f.GetCellValue("PSN", "D15")
	operatorRAM, _ := f.GetCellValue("PSN", "E15")
	operatorSSD, _ := f.GetCellValue("PSN", "F15")

	coVM, _ := f.GetCellValue("PSN", "C16")
	coCPU, _ := f.GetCellValue("PSN", "D16")
	coRAM, _ := f.GetCellValue("PSN", "E16")
	coSSD, _ := f.GetCellValue("PSN", "F16")

	pgsVM, _ := f.GetCellValue("PSN", "C17")
	pgsCPU, _ := f.GetCellValue("PSN", "D17")
	pgsRAM, _ := f.GetCellValue("PSN", "E17")

	// Расчет значения для PGS SSD
	ssdValue := calculatePGSSSD(userInputValues[0])

	// Отправка результата пользователю
	resultMsg := fmt.Sprintf(
		"Результаты расчета сайзинга для продукта Частное Облако Standalone:\n\n"+
			"ВМ Operator: кол-во ВМ - %s, CPU - %s, RAM - %s ГБ, SSD - %s ГБ;\n"+
			"Компонент CO: кол-во ВМ - %s, CPU - %s, RAM - %s ГБ, SSD - %s ГБ;\n"+
			"Компонент PGS: кол-во ВМ - %s, CPU - %s, RAM - %s ГБ, SSD - %d ГБ.",
		operatorVM, operatorCPU, operatorRAM, operatorSSD,
		coVM, coCPU, coRAM, coSSD,
		pgsVM, pgsCPU, pgsRAM, ssdValue,
	)
	msg := tgbotapi.NewMessage(chatID, resultMsg)
	bot.Send(msg)

	// Отправка клавиатуры с основным меню
	showMainMenu(bot, chatID)
}

// calculatePGSSSD вычисляет значение для PGS SSD
func calculatePGSSSD(userInputValues []string) int {
	value1, err := strconv.ParseFloat(userInputValues[0], 64) // Количество пользователей
	if err != nil {
		log.Println("Ошибка преобразования строки в число:", err)
		return 0
	}

	value2, err := strconv.ParseFloat(userInputValues[3], 64) // Дисковая квота
	if err != nil {
		log.Println("Ошибка преобразования строки в число:", err)
		return 0
	}

	ssdValue := 100 + value1*value2
	return int(math.Round(ssdValue))
}

// showMainMenu отправляет клавиатуру с главным меню
func showMainMenu(bot *tgbotapi.BotAPI, chatID int64) {
	keyboard := keyboards.GetMainMenuKeyboard()
	msgWithKeyboard := tgbotapi.NewMessage(chatID, "Выберите следующую опцию:")
	msgWithKeyboard.ReplyMarkup = keyboard
	bot.Send(msgWithKeyboard)
}

// sendErrorMessage отправляет сообщение об ошибке
func sendErrorMessage(bot *tgbotapi.BotAPI, chatID int64, errorMessage string) {
	msg := tgbotapi.NewMessage(chatID, errorMessage)
	bot.Send(msg)
	log.Println(errorMessage)
}
