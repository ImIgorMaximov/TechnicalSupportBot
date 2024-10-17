package sizing

import (
	"errors"
	"log"
	"strconv"
	"technicalSupportBot/pkg/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	mailionMaxUser = 140000
	mailionQuota   = 200
)

// Определение глобальной переменной для хранения пользовательских вводов по состояниям
var userInputValuesMailion = make(map[int64][]string)

// Глобальная переменная для хранения текущего состояния
var currentStateMailion string

// Обработка ввода пользователя и управление состоянием
func HandleUserInputMailion(bot *tgbotapi.BotAPI, chatID int64, state *string, text string) {
	currentStateMailion = *state

	log.Printf("Функция HandleUserInput. Текущее состояние: %s", currentStateMailion)

	switch currentStateMailion {
	case "mailion":
		log.Printf("Обработка состояния: %s.", currentStateMailion)
		currentStateMailion = "awaitingMaxUserMailion"
		userInputValuesMailion[chatID] = []string{} // Инициализация мапы для пользователя
		HandleNextInputMailion(bot, chatID, "", "Введите максимальное количество пользователей (например, 500):", "awaitingMaxUserMailion")

	case "awaitingMaxUserMailion":
		log.Printf("Обработка состояния: %s.", currentStateMailion)

		if ok := validateInput(text, mailionMaxUser); !ok {
			msg := tgbotapi.NewMessage(chatID, "Некорректный ввод. Пожалуйста, введите числа в диапазоне от 1 до 1400000.")
			bot.Send(msg)
			return
		}

		userInputValuesMailion[chatID] = append(userInputValuesMailion[chatID], text) // Сохраняем ввод
		currentStateMailion = "awaitingDiskQuotaMailion"
		HandleNextInputMailion(bot, chatID, text, "Введите дисковую квоту пользователей в почте (ГБ) (например, 3):", "awaitingDiskQuotaMailion")

	case "awaitingDiskQuotaMailion":
		log.Printf("Обработка состояния: %s.", currentStateMailion)

		if ok := validateInput(text, mailionQuota); !ok {
			msg := tgbotapi.NewMessage(chatID, "Некорректный ввод. Пожалуйста, введите числа в диапазоне от 1 до 10000.")
			bot.Send(msg)
			return
		}

		userInputValuesMailion[chatID] = append(userInputValuesMailion[chatID], text) // Сохраняем ввод

		// Все данные собраны, теперь определяем какой PDF отправить
		maxUsers, _ := strconv.Atoi(userInputValuesMailion[chatID][0])
		quota, _ := strconv.Atoi(userInputValuesMailion[chatID][1])

		// Определяем, какой файл PDF отправить
		pdfFilePath, err := determinePDFFileMailion(maxUsers, quota)
		if err != nil {
			msg := tgbotapi.NewMessage(chatID, "Произошла ошибка при выборе файла. Попробуйте снова.")
			bot.Send(msg)
			return
		}

		// Отправляем PDF файл пользователю
		err = sendPDFToUserMailion(bot, chatID, pdfFilePath)
		if err != nil {
			msg := tgbotapi.NewMessage(chatID, "Ошибка при отправке файла. Попробуйте снова.")
			bot.Send(msg)
			return
		}

		// Сбрасываем состояние после завершения
		currentStateMailion = "calculation done"
		delete(userInputValuesMailion, chatID)

		showMainMenuForMailion(bot, chatID)

	default:
		msg := tgbotapi.NewMessage(chatID, "Неизвестное состояние. Начните заново.")
		bot.Send(msg)
		currentStateMailion = "mailion"
		delete(userInputValuesMailion, chatID)
	}

	*state = currentStateMailion
}

// HandleNextInputMailion помогает запрашивать следующие данные и обновлять состояние
func HandleNextInputMailion(bot *tgbotapi.BotAPI, chatID int64, userInput string, nextMessage string, nextState string) {
	currentStateMailion = nextState
	msg := tgbotapi.NewMessage(chatID, nextMessage) // Создаём сообщение для пользователя
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Ошибка при отправке сообщения: %s", err) // Логируем ошибку
	}
}

// Функция для определения пути к PDF файлу на основе введённых данных
func determinePDFFileMailion(maxUsers, quota int) (string, error) {
	// Пример логики выбора PDF файла
	if (maxUsers <= 500) && (quota <= 15) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_500_15GB.pdf", nil
	} else if ((maxUsers > 500) && (maxUsers <= 700)) && (quota <= 8) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_700_15GB.pdf", nil
	} else if ((maxUsers <= 1000) && (maxUsers > 700)) && (quota == 1) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_1000_1GB.pdf", nil
	} else if ((maxUsers <= 1000) && (maxUsers > 700)) && (quota <= 3) && (quota > 1) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_1000_3GB.pdf", nil
	} else if ((maxUsers <= 1000) && (maxUsers > 700)) && ((quota <= 5) && (quota > 3)) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_1000_5GB.pdf", nil
	} else if ((maxUsers > 1000) && (maxUsers <= 1500)) && (quota <= 3) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_1500_3GB.pdf", nil
	} else if ((maxUsers > 1000) && (maxUsers <= 1500)) && ((quota > 3) && (quota <= 10)) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_1500_10GB.pdf", nil
	} else if ((maxUsers > 1000) && (maxUsers <= 1500)) && ((quota > 10) && (quota <= 50)) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_1500_50GB.pdf", nil
	} else if ((maxUsers > 1000) && (maxUsers <= 1500)) && ((quota > 50) && (quota <= 200)) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_1500_200GB.pdf", nil
	} else if ((maxUsers > 1500) && (maxUsers <= 2000)) && (quota <= 3) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_2000_3GB.pdf", nil
	} else if ((maxUsers > 1500) && (maxUsers <= 2000)) && ((quota <= 5) && (quota > 3)) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_2000_5GB.pdf", nil
	} else if ((maxUsers > 1500) && (maxUsers <= 2000)) && ((quota <= 10) && (quota > 5)) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_2000_10GB.pdf", nil
	} else if ((maxUsers > 2000) && (maxUsers <= 3000)) && (quota <= 3) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_3000_3GB.pdf", nil
	} else if ((maxUsers > 3000) && (maxUsers <= 5000)) && (quota <= 3) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_5000_3GB.pdf", nil
	} else if ((maxUsers > 3000) && (maxUsers <= 5000)) && ((quota <= 5) && (quota > 3)) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_5000_5GB.pdf", nil
	} else if ((maxUsers > 5000) && (maxUsers <= 6000)) && (quota == 1) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_6000_1GB.pdf", nil
	} else if ((maxUsers > 6000) && (maxUsers <= 7000)) && (quota == 1) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_7000_1GB.pdf", nil
	} else if ((maxUsers > 7000) && (maxUsers <= 10000)) && (quota == 1) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_10000_1GB.pdf", nil
	} else if ((maxUsers > 7000) && (maxUsers <= 10000)) && ((quota <= 5) && (quota > 1)) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_10000_5GB.pdf", nil
	} else if ((maxUsers > 7000) && (maxUsers <= 10000)) && ((quota > 5) && (quota <= 10)) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_10000_10GB.pdf", nil
	} else if ((maxUsers > 10000) && (maxUsers <= 30000)) && (quota == 1) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_30000_1GB.pdf", nil
	} else if ((maxUsers > 10000) && (maxUsers <= 30000)) && (quota == 2) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_30000_2GB.pdf", nil
	} else if ((maxUsers > 10000) && (maxUsers <= 30000)) && ((quota > 2) && (quota <= 5)) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_30000_5GB.pdf", nil
	} else if ((maxUsers > 10000) && (maxUsers <= 30000)) && ((quota > 5) && (quota <= 20)) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_30000_20GB.pdf", nil
	} else if ((maxUsers > 30000) && (maxUsers <= 140000)) && (quota <= 20) {
		return "/home/admin-msk/MyOfficeConfig/sizing_mailion_140000_20GB.pdf", nil
	}

	return "", errors.New("не удалось подобрать нужный сайзинг для введённых данных. Обратитесь к инженеру или технической поддежке.")
}

// Функция для отправки PDF файла пользователю
func sendPDFToUserMailion(bot *tgbotapi.BotAPI, chatID int64, filePath string) error {
	if filePath == "" {
		log.Printf("Путь к файлу пустой, файл не найден для отправки")
		return errors.New("путь к файлу пустой")
	}

	log.Printf("Отправка PDF файла пользователю: %s", filePath)

	// Загружаем файл
	file := tgbotapi.NewDocument(chatID, tgbotapi.FilePath(filePath))
	_, err := bot.Send(file)
	if err != nil {
		log.Printf("Ошибка при отправке PDF файла: %v", err)
		return err
	}

	log.Printf("Файл успешно отправлен пользователю: %d", chatID)

	// Отправляем сообщение о том, что сайзинг является типовым
	msg := tgbotapi.NewMessage(chatID, "Данный сайзинг является типовым. Для более детального расчета и оптимизации обратитесь к инженеру или техническую поддержку.")
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Ошибка при отправке сообщения: %v", err)
		return err
	}

	return nil
}

// showMainMenu отправляет клавиатуру с главным меню
func showMainMenuForMailion(bot *tgbotapi.BotAPI, chatID int64) {
	keyboard := keyboards.GetMainMenuKeyboardForMailion()
	msgWithKeyboard := tgbotapi.NewMessage(chatID, "Выберите следующую опцию:")
	msgWithKeyboard.ReplyMarkup = keyboard
	bot.Send(msgWithKeyboard)
}
