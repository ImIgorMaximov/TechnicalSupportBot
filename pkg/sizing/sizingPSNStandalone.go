/*

Package sizing содержит функции для расчета сайзинга с учетом введенных данных пользователем для продуктов: Частное облако, Mailion, Почта и Squadus.
Отправляет соответствующий xlsx-файл с сайзингом в конфигурации Standalone для продукта Почта в чат Telegram.


Автор: Максимов Игорь
Email: imigormaximov@gmail.com
*/

package sizing

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xuri/excelize/v2"
)

const (
	psnMaxUser    = 1000
	psnQuota      = 1000
	psnMaxLetters = 10000
	psnSpam       = 10000
)

// Определение глобальной переменной для хранения пользовательских вводов по состояниям
var userInputValuesPSNStandalone = make(map[int64][]string)

// Глобальная переменная для хранения текущего состояния
var currentStatePSNStandalone string

// Обработка ввода пользователя и управление состоянием
func HandleUserInputPSNStandalone(bot *tgbotapi.BotAPI, chatID int64, state *string, text string) {
	currentStatePSNStandalone = *state

	log.Printf("Функция HandleUserInput. Текущее состояние: %s", currentStatePSNStandalone)

	switch currentStatePSNStandalone {
	case "standalone":
		log.Printf("Обработка состояния: %s.", currentStatePSNStandalone)
		currentStatePSNStandalone = "awaitingMaxUserCountPSN"
		userInputValuesPSNStandalone[chatID] = []string{} // Инициализация мапы для пользователя
		HandleNextInputPSNStandalone(bot, chatID, "", "Введите максимальное количество пользователей (например, 50):", "awaitingMaxUserCountPSN")

	case "awaitingMaxUserCountPSN":
		log.Printf("Обработка состояния: %s.", currentStatePSNStandalone)

		if ok := validateInput(text, psnMaxUser); !ok {
			msg := tgbotapi.NewMessage(chatID, "Некорректный ввод. Пожалуйста, введите числа в диапазоне от 1 до 1000.")
			bot.Send(msg)
			return
		}

		userInputValuesPSNStandalone[chatID] = append(userInputValuesPSNStandalone[chatID], text) // Сохраняем ввод
		currentStatePSNStandalone = "awaitingDiskQuotaMail"
		HandleNextInputPSNStandalone(bot, chatID, text, "Введите дисковую квоту пользователей в почте (ГБ) (Например, 2):", "awaitingDiskQuotaMail")

	case "awaitingDiskQuotaMail":
		log.Printf("Обработка состояния: %s.", currentStatePSNStandalone)

		if ok := validateInput(text, psnQuota); !ok {
			msg := tgbotapi.NewMessage(chatID, "Некорректный ввод. Пожалуйста, введите числа в диапазоне от 1 до 10000.")
			bot.Send(msg)
			return
		}

		userInputValuesPSNStandalone[chatID] = append(userInputValuesPSNStandalone[chatID], text) // Сохраняем ввод
		currentStatePSNStandalone = "awaitingEmailsPerDayMail"
		HandleNextInputPSNStandalone(bot, chatID, text, "Введите количество писем в сутки на пользователя (Например, 100):", "awaitingEmailsPerDayMail")

	case "awaitingEmailsPerDayMail":
		log.Printf("Обработка состояния: %s.", currentStatePSNStandalone)

		if ok := validateInput(text, psnMaxLetters); !ok {
			msg := tgbotapi.NewMessage(chatID, "Некорректный ввод. Пожалуйста, введите числа в диапазоне от 1 до 10000.")
			bot.Send(msg)
			return
		}

		userInputValuesPSNStandalone[chatID] = append(userInputValuesPSNStandalone[chatID], text) // Сохраняем ввод
		currentStatePSNStandalone = "awaitingSpamCoefficientMail"
		HandleNextInputPSNStandalone(bot, chatID, text, "Введите коэффициент спама (Например, 0.1):", "awaitingSpamCoefficientMail")

	case "awaitingSpamCoefficientMail":
		log.Printf("Обработка состояния: %s.", currentStatePSNStandalone)

		if ok := validateInputSpam(text, psnSpam); !ok {
			msg := tgbotapi.NewMessage(chatID, "Некорректный ввод. Пожалуйста, введите числа в диапазоне от 0.1 до 0.9.")
			bot.Send(msg)
			return
		}

		userInputValuesPSNStandalone[chatID] = append(userInputValuesPSNStandalone[chatID], text) // Сохраняем ввод
		log.Println("Все данные от пользователя получены:", userInputValuesPSNStandalone[chatID])

		// После получения всех значений выполняем расчет
		calculateAndSendMailSizingPSNStandalone(bot, chatID, userInputValuesPSNStandalone[chatID])
		log.Println("Результаты расчета отправлены пользователю")
		currentStatePSNStandalone = "calculationDone"
	default:
		log.Printf("Ошибка: Неизвестное состояние или некорректный ввод. Состояние: %s", currentStatePSNStandalone)
		sendErrorMessage(bot, chatID, "Ошибка: некорректный ввод. Введите кнопку /start для выхода в Главное меню.")
	}

	// Обновление состояния после обработки
	*state = currentStatePSNStandalone
}

// HandleNextInput помогает запрашивать следующие данные и обновлять состояние
func HandleNextInputPSNStandalone(bot *tgbotapi.BotAPI, chatID int64, userInput string, nextMessage string, nextState string) {
	currentStatePSNStandalone = nextState
	// userInputValuesPSNStandalone[chatID] = append(userInputValuesPSNStandalone[chatID], userInput) // Сохраняем текущий ввод
	msg := tgbotapi.NewMessage(chatID, nextMessage) // Создаём сообщение для пользователя
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Ошибка при отправке сообщения: %s", err) // Логируем ошибку
	}
}

// calculateAndSendMailSizingPSNStandalone выполняет расчет и отправляет результат пользователю
func calculateAndSendMailSizingPSNStandalone(bot *tgbotapi.BotAPI, chatID int64, userInputValuesPSNStandalone []string) {
	// Открытие файла Excel
	filePath := "/home/admin-msk/MyOfficeConfig/sizingPSNStandalone.xlsx"
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		sendErrorMessage(bot, chatID, "Произошла ошибка при открытии файла.")
		log.Println("Ошибка открытия файла:", err)
		return
	}
	defer f.Close()

	// Заполнение ячеек данными
	err = fillExcelFilePSN(f, userInputValuesPSNStandalone)
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
	sendSizingResultsPSNStandalone(bot, chatID, f, userInputValuesPSNStandalone)
}

// fillExcelFilePSN заполняет Excel данными пользователя
func fillExcelFilePSN(f *excelize.File, userInputValuesPSNStandalone []string) error {
	if len(userInputValuesPSNStandalone) < 4 {
		errors.New("len is less than 4")
	}

	err := f.SetCellValue("Standalone", "D4", userInputValuesPSNStandalone[0]) // Макс. количество пользователей
	err = f.SetCellValue("Standalone", "D9", userInputValuesPSNStandalone[1])  // Дисковая квота пользователя
	err = f.SetCellValue("Standalone", "D10", userInputValuesPSNStandalone[2]) // Кол-во писем в сутки
	err = f.SetCellValue("Standalone", "D11", userInputValuesPSNStandalone[3]) // Коэффициент Спама
	return err
}

// sendSizingResults извлекает результаты из Excel и отправляет их пользователю
func sendSizingResultsPSNStandalone(bot *tgbotapi.BotAPI, chatID int64, f *excelize.File, userInputValuesPSNStandalone []string) {

	psnVM, _ := f.GetCellValue("Standalone", "C18")
	psnCPU, _ := f.GetCellValue("Standalone", "D18")
	psnRAM, _ := f.GetCellValue("Standalone", "E18")
	// Расчет значения для SSD
	psnSSD := calculateSSDPSN(userInputValuesPSNStandalone)
	psnHDD, _ := f.GetCellValue("Standalone", "G18")

	// Расчет значения для SSD
	// ssdValue := calculateSSD(userInputValuesPSNStandalone)

	newFile, err := newExcelFile()
	if err != nil {
		log.Println(err)
		return
	}

	sheetPSN := "PSN"
	err = configurePSN(newFile, sheetPSN)
	if err != nil {
		log.Println(err)
		return
	}

	err = newFile.SetCellValue(sheetPSN, "B2", psnVM)
	err = newFile.SetCellValue(sheetPSN, "C2", psnCPU)
	err = newFile.SetCellValue(sheetPSN, "D2", psnRAM)
	err = newFile.SetCellValue(sheetPSN, "E2", psnSSD)
	err = newFile.SetCellValue(sheetPSN, "F2", psnHDD)

	// Создание буфера для хранения файла в памяти
	buf := new(bytes.Buffer)
	if err := newFile.Write(buf); err != nil {
		log.Fatalf("Ошибка при записи в буфер: %v", err)
	}

	fbytes := tgbotapi.FileBytes{
		Name:  "sizingPSN.xlsx",
		Bytes: buf.Bytes(),
	}

	// отправка файла в чат
	doc := tgbotapi.NewDocument(chatID, fbytes)
	if _, err := bot.Send(doc); err != nil {
		log.Printf("Ошибка отправки %s файла, err: %v", fbytes.Name, err)
		return
	}
	// Отправка результата пользователю
	resultMsg := fmt.Sprintf(
		"Результаты расчета сайзинга для продукта Почта Standalone:\n\n"+
			"Компонент PSN: кол-во ВМ - %s, CPU - %s, RAM - %s ГБ, SSD - %d ГБ;\n",
		psnVM, psnCPU, psnRAM, psnSSD,
	)
	msg := tgbotapi.NewMessage(chatID, resultMsg)
	bot.Send(msg)

	// Отправка клавиатуры с основным меню
	showMainMenu(bot, chatID)
}

func calculateSSDPSN(userInputValuesPrivateCloudStandalone []string) int {
	value1, err := strconv.ParseFloat(userInputValuesPrivateCloudStandalone[0], 64) // Количество пользователей
	if err != nil {
		log.Println("Ошибка преобразования строки в число:", err)
		return 0
	}

	value2, err := strconv.ParseFloat(userInputValuesPrivateCloudStandalone[1], 64) // Дисковая квота
	if err != nil {
		log.Println("Ошибка преобразования строки в число:", err)
		return 0
	}

	ssdValue := 50 + value1*value2*1.3
	return int(math.Round(ssdValue))
}

func validateInputSpam(input string, max float64) bool {
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return false
	}

	return num > 0 && num < 1
}
