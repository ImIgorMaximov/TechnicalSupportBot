package sizing

import (
	"bytes"
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

	sheetName = "Standalone"
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

	// Извлечение результатов
	sendSizingResultsPrivateCloudStandalone(bot, chatID, f, userInputValuesPrivateCloudStandalone)
}

// fillExcelFile заполняет Excel данными пользователя
func fillExcelFilePrivateCloudStandalone(f *excelize.File, userInputValuesPrivateCloudStandalone []string) error {
	if len(userInputValuesPrivateCloudStandalone) < 4 {
		errors.New("len is less than 4")
	}

	err := f.SetCellValue(sheetName, "D4", userInputValuesPrivateCloudStandalone[0]) // Макс. количество пользователей
	err = f.SetCellValue(sheetName, "F6", userInputValuesPrivateCloudStandalone[1])  // Кол-во активных пользователей
	err = f.SetCellValue(sheetName, "F7", userInputValuesPrivateCloudStandalone[2])  // Кол-во редактируемых документов
	err = f.SetCellValue(sheetName, "D8", userInputValuesPrivateCloudStandalone[3])  // Дисковая квота пользователя
	return err
}

// sendSizingResults извлекает результаты из Excel и отправляет их пользователю
func sendSizingResultsPrivateCloudStandalone(bot *tgbotapi.BotAPI, chatID int64, f *excelize.File, userInputValuesPrivateCloudStandalone []string) {
	//GET VALUE
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

	psnVM, _ := f.GetCellValue("Standalone", "C18")
	psnCPU, _ := f.GetCellValue("Standalone", "D18")
	psnRAM, _ := f.GetCellValue("Standalone", "E18")

	itogoVM, _ := f.GetCellValue("Standalone", "C19")
	itogoCPU, _ := f.GetCellValue("Standalone", "D19")
	itogoRAM, _ := f.GetCellValue("Standalone", "E19")
	itogoSSD, _ := f.GetCellValue("Standalone", "F19")
	// itogoHDD, _ := f.GetCellValue("Standalone", "G19")

	newFile, err := newExcelFile(sheetName)
	if err != nil {
		log.Println("creating new file err:", err)
		return
	}

	err = newFile.SetCellValue(sheetName, "C2", operatorVM)
	err = newFile.SetCellValue(sheetName, "D2", operatorCPU)
	err = newFile.SetCellValue(sheetName, "E2", operatorRAM)
	err = newFile.SetCellValue(sheetName, "F2", operatorSSD)

	err = newFile.SetCellValue(sheetName, "C3", coVM)
	err = newFile.SetCellValue(sheetName, "D3", coCPU)
	err = newFile.SetCellValue(sheetName, "E3", coRAM)
	err = newFile.SetCellValue(sheetName, "F3", coSSD)

	err = newFile.SetCellValue(sheetName, "C4", pgsVM)
	err = newFile.SetCellValue(sheetName, "D4", pgsCPU)
	err = newFile.SetCellValue(sheetName, "E4", pgsRAM)

	err = newFile.SetCellValue(sheetName, "C5", psnVM)
	err = newFile.SetCellValue(sheetName, "D5", psnCPU)
	err = newFile.SetCellValue(sheetName, "E5", psnRAM)

	err = f.SetCellValue(sheetName, "C6", itogoVM)
	err = f.SetCellValue(sheetName, "D6", itogoCPU)
	err = f.SetCellValue(sheetName, "E6", itogoRAM)
	err = f.SetCellValue(sheetName, "F6", itogoSSD)

	// Расчет значения для PGS SSD
	ssdValue := calculateSSD(userInputValuesPrivateCloudStandalone)

	// Создание буфера для хранения файла в памяти
	buf := new(bytes.Buffer)
	if err := newFile.Write(buf); err != nil {
		log.Fatalf("Ошибка при записи в буфер: %v", err)
	}

	fbytes := tgbotapi.FileBytes{
		Name:  "sizing.xlsx",
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

func newExcelFile(sheetName string) (*excelize.File, error) {
	f := excelize.NewFile()

	index, err := f.NewSheet(sheetName)
	if err != nil {
		log.Fatalf("Ошибка при создании sheet: %v", err)
		return nil, err
	}
	defer f.Close()

	f.SetActiveSheet(index)

	err = f.SetColWidth(sheetName, "A", "G", 12)
	if err != nil {
		log.Fatal(err)
	}

	err = f.SetCellDefault(sheetName, "A1", "Компонент")
	if err != nil {
		log.Fatal(err)
	}
	err = f.SetCellDefault(sheetName, "B1", "Роль\\Сервис")
	err = f.SetCellDefault(sheetName, "C1", "Кол-во VM")
	err = f.SetCellDefault(sheetName, "D1", "CPU, vCPU")
	err = f.SetCellDefault(sheetName, "E1", "RAM, GB")
	err = f.SetCellDefault(sheetName, "F1", "SSD, GB")
	err = f.SetCellDefault(sheetName, "G1", "HDD, GB")

	err = f.SetCellDefault(sheetName, "A3", "COS")
	err = f.SetCellDefault(sheetName, "A4", "PGS")
	err = f.SetCellDefault(sheetName, "A5", "PSN")

	err = f.SetCellDefault(sheetName, "B2", "оператор")
	err = f.SetCellDefault(sheetName, "B3", "все роли")
	err = f.SetCellDefault(sheetName, "B4", "все роли")
	err = f.SetCellDefault(sheetName, "B5", "все роли")

	top := excelize.Border{Type: "top", Style: 1, Color: "000000"}
	left := excelize.Border{Type: "left", Style: 1, Color: "000000"}
	right := excelize.Border{Type: "right", Style: 1, Color: "000000"}
	bottom := excelize.Border{Type: "bottom", Style: 1, Color: "000000"}

	err = f.MergeCell(sheetName, "A6", "B6")
	if err != nil {
		return nil, err
	}
	style, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{top, left, right, bottom},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
	})
	err = f.SetCellStyle(sheetName, "A5", "B5", style)
	if err != nil {
		return nil, err
	}

	rowStyle, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{top, left, right, bottom},
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"e6b690"},
			Shading: 1,
		},
		Font: &excelize.Font{
			Strike:       false,
			Color:        "000000",
			ColorIndexed: index,
		},
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})
	if err != nil {
		log.Fatal("rowstyle: ", err)
		return nil, err
	}
	err = f.SetCellStyle(sheetName, "A1", "G1", rowStyle)
	if err != nil {
		return nil, err
	}

	row1, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{top, left, right, bottom},
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"b4c7dc"},
			Shading: 1,
		},
		Font: &excelize.Font{
			Color:        "000000",
			ColorIndexed: index,
		},
	})
	row2, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{top, left, right, bottom},
		Font: &excelize.Font{
			Color:        "000000",
			ColorIndexed: index,
		},
	})

	if err != nil {
		log.Fatal("rowstyle: ", err)
		return nil, err
	}
	err = f.SetCellStyle(sheetName, "A2", "G2", row2)
	if err != nil {
		return nil, err
	}
	err = f.SetCellStyle(sheetName, "A3", "G3", row1)
	if err != nil {
		return nil, err
	}
	err = f.SetCellStyle(sheetName, "A4", "G4", row2)
	if err != nil {
		return nil, err
	}
	err = f.SetCellStyle(sheetName, "A5", "G5", row1)
	if err != nil {
		return nil, err
	}
	err = f.SetCellStyle(sheetName, "A6", "G6", row2)
	if err != nil {
		return nil, err
	}

	err = f.SetCellValue(sheetName, "A6", "Итого")

	// err = f.SetCellDefault(sheetName, "C2", " 1")
	// err = f.SetCellDefault(sheetName, "D2", " 1")
	// err = f.SetCellDefault(sheetName, "E2", " 4")
	// err = f.SetCellDefault(sheetName, "f2", " 50")
	// err = f.SetCellDefault(sheetName, "C3", " 1")
	// err = f.SetCellDefault(sheetName, "D3", " 8")
	// err = f.SetCellDefault(sheetName, "E3", " 20")
	// err = f.SetCellDefault(sheetName, "f3", " 100")
	// err = f.SetCellDefault(sheetName, "C4", " 1")
	// err = f.SetCellDefault(sheetName, "D4", " 8")
	// err = f.SetCellDefault(sheetName, "E4", " 20")
	// err = f.SetCellDefault(sheetName, "F4", " 140")
	// err = f.SetCellDefault(sheetName, "C5", " 1")
	// err = f.SetCellDefault(sheetName, "D5", " 8")
	// err = f.SetCellDefault(sheetName, "E5", " 20")
	// err = f.SetCellDefault(sheetName, "F5", " 140")

	if err != nil {
		log.Println("errs", err)
	}

	return f, nil
}
