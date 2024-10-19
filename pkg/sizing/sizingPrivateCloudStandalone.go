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

	err := f.SetCellValue(sheetName, "D4", userInputValuesPrivateCloudStandalone[0]) // Макс. количество пользователей
	err = f.SetCellValue(sheetName, "F6", userInputValuesPrivateCloudStandalone[1])  // Кол-во активных пользователей
	err = f.SetCellValue(sheetName, "F7", userInputValuesPrivateCloudStandalone[2])  // Кол-во редактируемых документов
	err = f.SetCellValue(sheetName, "D8", userInputValuesPrivateCloudStandalone[3])  // Дисковая квота пользователя

	// Расчет значения для PGS SSD
	pgsSSD := calculateSSD(userInputValuesPrivateCloudStandalone)
	err = f.SetCellValue(sheetName, "F17", pgsSSD)

	return err
}

// sendSizingResults извлекает результаты из Excel и отправляет их пользователю
func sendSizingResultsPrivateCloudStandalone(bot *tgbotapi.BotAPI, chatID int64, f *excelize.File, userInputValuesPrivateCloudStandalone []string) {
	//GET VALUE
	operatorVM, _ := f.GetCellValue("Standalone", "C15")
	operatorCPU, _ := f.GetCellValue("Standalone", "D15")
	operatorRAM, _ := f.GetCellValue("Standalone", "E15")
	operatorSSD, _ := f.GetCellValue("Standalone", "F15")
	operatorHDD, _ := f.GetCellValue("Standalone", "G15")

	coVM, _ := f.GetCellValue("Standalone", "C16")
	coCPU, _ := f.GetCellValue("Standalone", "D16")
	coRAM, _ := f.GetCellValue("Standalone", "E16")
	coSSD, _ := f.GetCellValue("Standalone", "F16")
	coHDD, _ := f.GetCellValue("Standalone", "G16")

	pgsVM, _ := f.GetCellValue("Standalone", "C17")
	pgsCPU, _ := f.GetCellValue("Standalone", "D17")
	pgsRAM, _ := f.GetCellValue("Standalone", "E17")
	pgsSSD, _ := f.GetCellValue("Standalone", "F17")
	pgsHDD, _ := f.GetCellValue("Standalone", "G17")

	// resultVM, _ := f.GetCellValue("Standalone", "C19")
	// resultCPU, _ := f.GetCellValue("Standalone", "D19")
	// resultRAM, _ := f.GetCellValue("Standalone", "E19")
	// resultSSD, _ := f.GetCellValue("Standalone", "F19")
	// resultHDD, _ := f.GetCellValue("Standalone", "G19")

	// calculate cells value manually
	i_operatorVM, _ := strconv.Atoi(operatorVM)
	i_operatorCPU, _ := strconv.Atoi(operatorCPU)
	i_operatorRAM, _ := strconv.Atoi(operatorRAM)
	i_operatorSSD, _ := strconv.Atoi(operatorSSD)
	i_operatorHDD, _ := strconv.Atoi(operatorHDD)

	i_coVM, _ := strconv.Atoi(coVM)
	i_coCPU, _ := strconv.Atoi(coCPU)
	i_coRAM, _ := strconv.Atoi(coRAM)
	i_coSSD, _ := strconv.Atoi(coSSD)
	i_coHDD, _ := strconv.Atoi(coHDD)

	i_pgsVM, _ := strconv.Atoi(pgsVM)
	i_pgsCPU, _ := strconv.Atoi(pgsCPU)
	i_pgsRAM, _ := strconv.Atoi(pgsRAM)
	i_pgsSSD, _ := strconv.Atoi(pgsSSD)
	i_pgsHDD, _ := strconv.Atoi(pgsHDD)

	resultVM := strconv.Itoa(i_operatorVM + i_coVM + i_pgsVM)
	resultCPU := strconv.Itoa(i_operatorCPU + i_coCPU + i_pgsCPU)
	resultRAM := strconv.Itoa(i_operatorRAM + i_coRAM + i_pgsRAM)
	resultSSD := strconv.Itoa(i_operatorSSD + i_coSSD + i_pgsSSD)
	resultHDD := strconv.Itoa(i_operatorHDD + i_coHDD + i_pgsHDD)

	newFile, err := newExcelFile()
	if err != nil {
		log.Println("creating new file err:", err)
		return
	}

	err = configurePCS(newFile, sheetName)
	if err != nil {
		log.Println("configurePCS", err)
		return
	}

	err = newFile.SetCellValue(sheetName, "B2", operatorVM)
	err = newFile.SetCellValue(sheetName, "C2", operatorCPU)
	err = newFile.SetCellValue(sheetName, "D2", operatorRAM)
	err = newFile.SetCellValue(sheetName, "E2", operatorSSD)
	err = newFile.SetCellValue(sheetName, "F2", operatorHDD)

	err = newFile.SetCellValue(sheetName, "B3", coVM)
	err = newFile.SetCellValue(sheetName, "C3", coCPU)
	err = newFile.SetCellValue(sheetName, "D3", coRAM)
	err = newFile.SetCellValue(sheetName, "E3", coSSD)
	err = newFile.SetCellValue(sheetName, "F3", coHDD)

	err = newFile.SetCellValue(sheetName, "B4", pgsVM)
	err = newFile.SetCellValue(sheetName, "C4", pgsCPU)
	err = newFile.SetCellValue(sheetName, "D4", pgsRAM)
	err = newFile.SetCellValue(sheetName, "E4", pgsSSD)
	err = newFile.SetCellValue(sheetName, "F4", pgsHDD)

	err = newFile.SetCellValue(sheetName, "B5", resultVM)
	err = newFile.SetCellValue(sheetName, "C5", resultCPU)
	err = newFile.SetCellValue(sheetName, "D5", resultRAM)
	err = newFile.SetCellValue(sheetName, "E5", resultSSD)
	err = newFile.SetCellValue(sheetName, "F5", resultHDD)

	// Создание буфера для хранения файла в памяти
	buf := new(bytes.Buffer)
	if err := newFile.Write(buf); err != nil {
		log.Fatalf("Ошибка при записи в буфер: %v", err)
	}

	fbytes := tgbotapi.FileBytes{
		Name:  "sizingPrivateCloud.xlsx",
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
		pgsVM, pgsCPU, pgsRAM, pgsSSD,
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

func newExcelFile() (*excelize.File, error) {
	return excelize.NewFile(), nil
}

func configurePCS(f *excelize.File, sheetname string) error {
	index, err := f.NewSheet(sheetname)
	if err != nil {
		log.Fatalf("Ошибка при создании sheet: %v", err)
		return err
	}

	f.SetActiveSheet(index)

	err = f.SetColWidth(sheetname, "A", "G", 12)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = f.SetCellDefault(sheetname, "A1", "Компонент")
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = f.SetCellDefault(sheetname, "B1", "Кол-во VM")
	err = f.SetCellDefault(sheetname, "C1", "CPU, vCPU")
	err = f.SetCellDefault(sheetname, "D1", "RAM, GB")
	err = f.SetCellDefault(sheetname, "E1", "SSD, GB")
	err = f.SetCellDefault(sheetname, "F1", "HDD, GB")

	err = f.SetCellDefault(sheetname, "A2", "Operator")
	err = f.SetCellDefault(sheetname, "A3", "CO")
	err = f.SetCellDefault(sheetname, "A4", "PGS")

	top := excelize.Border{Type: "top", Style: 1, Color: "000000"}
	left := excelize.Border{Type: "left", Style: 1, Color: "000000"}
	right := excelize.Border{Type: "right", Style: 1, Color: "000000"}
	bottom := excelize.Border{Type: "bottom", Style: 1, Color: "000000"}

	// err = f.MergeCell(sheetname, "A6", "B6")
	// if err != nil {
	// 	return nil, err
	// }
	style, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{top, left, right, bottom},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
	})
	err = f.SetCellStyle(sheetname, "A5", "B5", style)
	if err != nil {
		log.Println("set style:", err)
		return err
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
		return err
	}
	err = f.SetCellStyle(sheetname, "A1", "F1", rowStyle)
	if err != nil {
		log.Println(err)
		return err
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
		Alignment: &excelize.Alignment{
			Horizontal: "left",
		},
	})
	if err != nil {
		return err
	}
	row2, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{top, left, right, bottom},
		Font: &excelize.Font{
			Color:        "000000",
			ColorIndexed: index,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "left",
		},
	})
	if err != nil {
		log.Fatal("rowstyle: ", err)
		return err
	}

	err = f.SetCellStyle(sheetname, "A2", "F2", row2)
	if err != nil {
		log.Println(err)
		return err
	}
	err = f.SetCellStyle(sheetname, "A3", "F3", row1)
	if err != nil {
		log.Println(err)
		return err
	}
	err = f.SetCellStyle(sheetname, "A4", "F4", row2)
	if err != nil {
		log.Println(err)
		return err
	}
	err = f.SetCellStyle(sheetname, "A5", "F5", row1)
	if err != nil {
		log.Println(err)
		return err
	}

	err = f.SetCellValue(sheetname, "A5", "Итого")
	if err != nil {
		log.Println("errs", err)
		return err
	}

	return nil
}

func configurePSN(f *excelize.File, sheetname string) error {
	index, err := f.NewSheet(sheetname)
	if err != nil {
		log.Fatalf("Ошибка при создании sheet: %v", err)
		return err
	}

	f.SetActiveSheet(index)

	err = f.SetColWidth(sheetname, "A", "G", 12)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = f.SetCellDefault(sheetname, "A1", "Компонент")
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = f.SetCellDefault(sheetname, "B1", "Кол-во VM")
	err = f.SetCellDefault(sheetname, "C1", "CPU, vCPU")
	err = f.SetCellDefault(sheetname, "D1", "RAM, GB")
	err = f.SetCellDefault(sheetname, "E1", "SSD, GB")
	err = f.SetCellDefault(sheetname, "F1", "HDD, GB")

	err = f.SetCellDefault(sheetname, "A2", "PSN")

	top := excelize.Border{Type: "top", Style: 1, Color: "000000"}
	left := excelize.Border{Type: "left", Style: 1, Color: "000000"}
	right := excelize.Border{Type: "right", Style: 1, Color: "000000"}
	bottom := excelize.Border{Type: "bottom", Style: 1, Color: "000000"}

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
		return err
	}
	err = f.SetCellStyle(sheetname, "A1", "F1", rowStyle)
	if err != nil {
		log.Println(err)
		return err
	}
	row2, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{top, left, right, bottom},
		Font: &excelize.Font{
			Color:        "000000",
			ColorIndexed: index,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "left",
		},
	})
	if err != nil {
		log.Fatal("rowstyle: ", err)
		return err
	}

	err = f.SetCellStyle(sheetname, "A2", "F2", row2)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
