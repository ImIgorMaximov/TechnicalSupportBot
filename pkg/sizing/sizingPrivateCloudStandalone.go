package sizing

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"

	"technicalSupportBot/pkg/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xuri/excelize/v2"
)

const (
	privateCloudMaxUser     = 1000
	privateCloudActiveUser  = 1000
	privateCloudDocument    = 10000
	privateCloudStorageInGB = 10000
)

// Определение глобальной переменной для хранения пользовательских вводов по состояниям
var userInputValuesPrivateCloudStandalone = make(map[int64][]string)

// Глобальная переменная для хранения текущего состояния
var currentStatePrivateCloudStandalone string

// Обработка ввода пользователя и управление состоянием
func HandleUserInputPrivateCloudStandalone(bot *tgbotapi.BotAPI, chatID int64, state *string, text string) {
	currentStatePrivateCloudStandalone = *state

	log.Printf("Функция HandleUserInput. Текущее состояние: %s", currentStatePrivateCloudStandalone)

	switch currentStatePrivateCloudStandalone {
	case "standalone":
		log.Printf("Обработка состояния: %s.", currentStatePrivateCloudStandalone)
		currentStatePrivateCloudStandalone = "awaitingMaxUserCountPrivateCloud"
		userInputValuesPrivateCloudStandalone[chatID] = []string{} // Инициализация мапы для пользователя
		HandleNextInputPrivateCloudStandalone(bot, chatID, "", "Введите максимальное количество пользователей (например, 50):", "awaitingMaxUserCountPrivateCloud")

	case "awaitingMaxUserCountPrivateCloud":
		log.Printf("Обработка состояния: %s.", currentStatePrivateCloudStandalone)

		if ok := validateInput(text, privateCloudMaxUser); !ok {
			msg := tgbotapi.NewMessage(chatID, "Некорректный ввод. Пожалуйста, введите числа в диапазоне от 1 до 1000.")
			bot.Send(msg)
			return
		}

		userInputValuesPrivateCloudStandalone[chatID] = append(userInputValuesPrivateCloudStandalone[chatID], text) // Сохраняем ввод
		currentStatePrivateCloudStandalone = "awaitingActiveUserCountPrivateCloud"
		HandleNextInputPrivateCloudStandalone(bot, chatID, text, "Введите количество одновременно активных пользователей (например, 10):", "awaitingActiveUserCountPrivateCloud")

	case "awaitingActiveUserCountPrivateCloud":
		log.Printf("Обработка состояния: %s.", currentStatePrivateCloudStandalone)

		if ok := validateInput(text, privateCloudActiveUser); !ok {
			msg := tgbotapi.NewMessage(chatID, "Некорректный ввод. Пожалуйста, введите числа в диапазоне от 1 до 10000.")
			bot.Send(msg)
			return
		}

		userInputValuesPrivateCloudStandalone[chatID] = append(userInputValuesPrivateCloudStandalone[chatID], text) // Сохраняем ввод
		currentStatePrivateCloudStandalone = "awaitingDocumentCountPrivateCloud"
		HandleNextInputPrivateCloudStandalone(bot, chatID, text, "Введите количество редактируемых документов (например, 200):", "awaitingDocumentCountPrivateCloud")

	case "awaitingDocumentCountPrivateCloud":
		log.Printf("Обработка состояния: %s.", currentStatePrivateCloudStandalone)

		if ok := validateInput(text, privateCloudDocument); !ok {
			msg := tgbotapi.NewMessage(chatID, "Некорректный ввод. Пожалуйста, введите числа в диапазоне от 1 до 10000.")
			bot.Send(msg)
			return
		}

		userInputValuesPrivateCloudStandalone[chatID] = append(userInputValuesPrivateCloudStandalone[chatID], text) // Сохраняем ввод
		currentStatePrivateCloudStandalone = "awaitingStorageQuotaPrivateCloud"
		HandleNextInputPrivateCloudStandalone(bot, chatID, text, "Введите дисковую квоту пользователей в хранилище (ГБ) (например, 2):", "awaitingStorageQuotaPrivateCloud")

	case "awaitingStorageQuotaPrivateCloud":
		log.Printf("Обработка состояния: %s.", currentStatePrivateCloudStandalone)

		if ok := validateInput(text, privateCloudStorageInGB); !ok {
			msg := tgbotapi.NewMessage(chatID, "Некорректный ввод. Пожалуйста, введите числа в диапазоне от 1 до 1000.")
			bot.Send(msg)
			return
		}

		userInputValuesPrivateCloudStandalone[chatID] = append(userInputValuesPrivateCloudStandalone[chatID], text) // Сохраняем ввод
		log.Println("Все данные от пользователя получены:", userInputValuesPrivateCloudStandalone[chatID])

		// После получения всех значений выполняем расчет
		calculateAndSendSizingPrivateCloudStandalone(bot, chatID, userInputValuesPrivateCloudStandalone[chatID])
		log.Println("Результаты расчета отправлены пользователю")
		currentStatePrivateCloudStandalone = "calculationDone"
	default:
		log.Printf("Ошибка: Неизвестное состояние или некорректный ввод. Состояние: %s", currentStatePrivateCloudStandalone)
		sendErrorMessage(bot, chatID, "Ошибка: некорректный ввод. Введите кнопку /start для выхода в Главное меню.")
	}

	// Обновление состояния после обработки
	*state = currentStatePrivateCloudStandalone
}

// HandleNextInputPrivateCloudStandalone помогает запрашивать следующие данные и обновлять состояние
func HandleNextInputPrivateCloudStandalone(bot *tgbotapi.BotAPI, chatID int64, userInput string, nextMessage string, nextState string) {
	currentStatePrivateCloudStandalone = nextState
	// userInputValuesPrivateCloudStandalone[chatID] = append(userInputValuesPrivateCloudStandalone[chatID], userInput) // Сохраняем текущий ввод
	msg := tgbotapi.NewMessage(chatID, nextMessage) // Создаём сообщение для пользователя
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Ошибка при отправке сообщения: %s", err) // Логируем ошибку
	}
}

// calculateAndSendSizing выполняет расчет и отправляет результат пользователю
func calculateAndSendSizingPrivateCloudStandalone(bot *tgbotapi.BotAPI, chatID int64, userInputValuesPrivateCloudStandalone []string) {
	// Открытие файла Excel
	filePath := "/home/admin-msk/MyOfficeConfig/sizingPrivateCloudStandalone.xlsx"
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		sendErrorMessage(bot, chatID, "Произошла ошибка при открытии файла.")
		log.Println("Ошибка открытия файла:", err)
		return
	}
	defer f.Close()

	// Заполнение ячеек данными
	err = fillExcelFilePrivateCloudStandalone(f, userInputValuesPrivateCloudStandalone)
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
	sendSizingResultsPrivateCloudStandalone(bot, chatID, f, userInputValuesPrivateCloudStandalone)
}

// fillExcelFile заполняет Excel данными пользователя
func fillExcelFilePrivateCloudStandalone(f *excelize.File, userInputValuesPrivateCloudStandalone []string) error {
	if len(userInputValuesPrivateCloudStandalone) < 4 {
		errors.New("len is less than 4")
	}

	err := f.SetCellValue("Standalone", "D4", userInputValuesPrivateCloudStandalone[0]) // Макс. количество пользователей
	err = f.SetCellValue("Standalone", "F6", userInputValuesPrivateCloudStandalone[1])  // Кол-во активных пользователей
	err = f.SetCellValue("Standalone", "F7", userInputValuesPrivateCloudStandalone[2])  // Кол-во редактируемых документов
	err = f.SetCellValue("Standalone", "D8", userInputValuesPrivateCloudStandalone[3])  // Дисковая квота пользователя
	return err
}

// sendSizingResults извлекает результаты из Excel и отправляет их пользователю
func sendSizingResultsPrivateCloudStandalone(bot *tgbotapi.BotAPI, chatID int64, f *excelize.File, userInputValuesPrivateCloudStandalone []string) {
	operatorVM, _ := f.GetCellValue("Standalone", "C15")
	operatorCPU, _ := f.GetCellValue("Standalone", "D15")
	operatorRAM, _ := f.GetCellValue("Standalone", "E15")
	operatorSSD, _ := f.GetCellValue("Standalone", "F15")

	coVM, _ := f.GetCellValue("Standalone", "C16")
	coCPU, _ := f.GetCellValue("Standalone", "D16")
	coRAM, _ := f.GetCellValue("Standalone", "E16")
	coSSD, _ := f.GetCellValue("Standalone", "F16")

	pgsVM, _ := f.GetCellValue("Standalone", "C17")
	pgsCPU, _ := f.GetCellValue("Standalone", "D17")
	pgsRAM, _ := f.GetCellValue("Standalone", "E17")

	// Расчет значения для PGS SSD
	ssdValue := calculateSSD(userInputValuesPrivateCloudStandalone)

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

// calculateSSD вычисляет значение для SSD
func calculateSSD(userInputValuesPrivateCloudStandalone []string) int {
	value1, err := strconv.ParseFloat(userInputValuesPrivateCloudStandalone[0], 64) // Количество пользователей
	if err != nil {
		log.Println("Ошибка преобразования строки в число:", err)
		return 0
	}

	value2, err := strconv.ParseFloat(userInputValuesPrivateCloudStandalone[3], 64) // Дисковая квота
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

func validateInput(input string, max int) bool {
	num, err := strconv.Atoi(input)
	if err != nil {
		return false
	}

	return num > 0 && num <= max
}
