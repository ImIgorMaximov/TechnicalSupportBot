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

// Глобальная переменная для хранения введённых данных от пользователей
var userInputValues = make(map[int64][]string)
var previousStateSizingPrivateCloudStandalone = make(map[int64]string)

func HandleSizingPrivateCloudStandalone(bot *tgbotapi.BotAPI, chatID int64) {
	// Запрос данных у пользователя
	msg := tgbotapi.NewMessage(chatID, "Введите количество пользователей (Например, 50):")
	bot.Send(msg)
	previousStateSizingPrivateCloudStandalone[chatID] = "awaitingUserCountPrivateCloud"
	log.Printf("Cостояние previousStateSizingPrivateCloudStandalone: %s", previousStateSizingPrivateCloudStandalone[chatID])
}

// HandleUserInputPrivateCloud обрабатывает ввод пользователя
func HandleUserInputPrivateCloud(bot *tgbotapi.BotAPI, chatID int64, userInput string) {
	log.Println("Текущее состояние перед обработкой ввода:", previousStateSizingPrivateCloudStandalone[chatID])
	log.Println("Полученное значение от пользователя:", userInput)

	switch previousStateSizingPrivateCloudStandalone[chatID] {
	case "awaitingUserCountPrivateCloud":
		userInputValues[chatID] = []string{userInput} // Сохраняем первое значение
		msg := tgbotapi.NewMessage(chatID, "Введите количество одновременно активных пользователей (Например, 10):")
		bot.Send(msg)
		previousStateSizingPrivateCloudStandalone[chatID] = "awaitingActiveUserCountPrivateCloud" // Обновляем состояние
		log.Println("Состояние изменено на awaitingActiveUserCountPrivateCloud")

	case "awaitingActiveUserCountPrivateCloud":
		userInputValues[chatID] = append(userInputValues[chatID], userInput) // Сохраняем второе значение
		msg := tgbotapi.NewMessage(chatID, "Введите количество редактируемых документов одновременно (Например, 10):")
		bot.Send(msg)
		previousStateSizingPrivateCloudStandalone[chatID] = "awaitingDocumentCountPrivateCloud" // Обновляем состояние
		log.Println("Состояние изменено на awaitingDocumentCountPrivateCloud")

	case "awaitingDocumentCountPrivateCloud":
		userInputValues[chatID] = append(userInputValues[chatID], userInput) // Сохраняем третье значение
		msg := tgbotapi.NewMessage(chatID, "Введите дисковую квоту пользователей в хранилище (ГБ) (Например, 2):")
		bot.Send(msg)
		previousStateSizingPrivateCloudStandalone[chatID] = "awaitingStorageQuotaPrivateCloud" // Обновляем состояние
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
	previousStateSizingPrivateCloudStandalone[chatID] = nextState
}

// calculateAndSendSizing выполняет расчет и отправляет результат пользователю
func calculateAndSendSizing(bot *tgbotapi.BotAPI, chatID int64) {
	// Открытие файла Excel
	filePath := "/home/admin-msk/Documents/sizingPrivateCloudPSNStandalone.xlsx"
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Println("Ошибка открытия файла:", err)
		msg := tgbotapi.NewMessage(chatID, "Произошла ошибка при открытии файла.")
		bot.Send(msg)
		return
	}
	defer f.Close()

	// Заполнение ячеек данными
	err = f.SetCellValue("PSN", "D4", userInputValues[chatID][0])
	err = f.SetCellValue("PSN", "F6", userInputValues[chatID][1])
	err = f.SetCellValue("PSN", "F7", userInputValues[chatID][2])
	err = f.SetCellValue("PSN", "D8", userInputValues[chatID][3])
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

	value1, err := strconv.ParseFloat(userInputValues[chatID][0], 64)
	if err != nil {
		log.Println("Ошибка преобразования строки в число:", err)
		msg := tgbotapi.NewMessage(chatID, "Ошибка данных: невозможно преобразовать значение в число.")
		bot.Send(msg)
		return
	}

	value2, err := strconv.ParseFloat(userInputValues[chatID][3], 64)
	if err != nil {
		log.Println("Ошибка преобразования строки в число:", err)
		msg := tgbotapi.NewMessage(chatID, "Ошибка данных: невозможно преобразовать значение в число.")
		bot.Send(msg)
		return
	}

	ssdValue := 100 + value1*value2
	pgsSSD := int(math.Round(ssdValue))

	// Отправка результата пользователю
	resultMsg := fmt.Sprintf(
		"Результаты расчета сайзинга для продукта Частное Облако Standalone:\n\n"+
			"ВМ Operator: кол-во ВМ - %s, CPU - %s, RAM - %s ГБ, SSD - %s ГБ;\n"+
			"Компонент CO: кол-во ВМ - %s, CPU - %s, RAM - %s ГБ, SSD - %s ГБ;\n"+
			"Компонент PGS: кол-во ВМ - %s, CPU - %s, RAM - %s ГБ, SSD - %d ГБ.",
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
	previousStateSizingPrivateCloudStandalone[chatID] = "sizingResultProvided"
	userInputValues[chatID] = nil
}
