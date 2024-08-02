package handlers

import (
	"log"
	"os"
	"technicalSupportBot/pkg/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func sendWelcomeMessage(bot *tgbotapi.BotAPI, chatID int64) {
	previousState[chatID] = "main"

	welcomeMessage := "Добро пожаловать в чат бот тех. поддержки МойОфис! :) " +
		"Выберите необходимую функцию:\n" +
		"1. Инструкции по продуктам.\n" +
		"2. Развертывание продуктов. \n" +
		"3. Расчет сайзинга продуктов. \n" +
		"4. Связаться с инженером тех. поддержки.\n"
	msg := tgbotapi.NewMessage(chatID, welcomeMessage)
	msg.ReplyMarkup = keyboards.GetMainKeyboard()
	bot.Send(msg)
}

func sendProduct(bot *tgbotapi.BotAPI, chatID int64) {
	previousState[chatID] = "production_instructions"

	chooseProductMessage := "Выберите продукт:"
	msg := tgbotapi.NewMessage(chatID, chooseProductMessage)
	msg.ReplyMarkup = keyboards.GetProductKeyboard()
	bot.Send(msg)
}

func sendDeploymentOptions(bot *tgbotapi.BotAPI, chatID int64) {
	previousState[chatID] = "deployment_options"

	deploymentMessage := "Выберите тип инсталляции:"
	msg := tgbotapi.NewMessage(chatID, deploymentMessage)
	msg.ReplyMarkup = keyboards.GetDeploymentOptionsKeyboard()
	bot.Send(msg)
}

func sendInstructions(bot *tgbotapi.BotAPI, chatID int64, product string) {
	previousState[chatID] = product
	chooseFunction := "Что подсказать? \n" +
		"- Cистемные требования \n" +
		"- Руководство по установке \n" +
		"- Руководство по администрированию \n"
	msg := tgbotapi.NewMessage(chatID, chooseFunction)
	msg.ReplyMarkup = keyboards.GetInstructionsKeyboard()
	bot.Send(msg)
}

func sendIsCertificates(bot *tgbotapi.BotAPI, chatID int64) {
	isCertificates := "Проверка сертификата сервера (server.crt): \n" +
		"openssl x509 -in server.crt -text -noout \n" +
		"Проверка цепочки сертификатов (ca.crt): \n" +
		"openssl x509 -in ca.crt -text -noout \n" +
		"Проверка приватного ключа (server.nopass.key): \n" +
		"openssl rsa -in server.nopass.key -check \n"
	msg := tgbotapi.NewMessage(chatID, isCertificates)
	bot.Send(msg)
}

func sendUnzippingISO(bot *tgbotapi.BotAPI, chatID int64) {
	unzippingISO := "Для разархивирования образа .iso используется инструмент \"bsdtar\": \n" +
		"apt-get install bsdtar \n" +
		"bsdtar -xvf путь_к_файлу.iso -C директория_для_извлечения \n"
	msg := tgbotapi.NewMessage(chatID, unzippingISO)
	bot.Send(msg)
}

func sendSupportEngineerContact(bot *tgbotapi.BotAPI, chatID int64) {
	errorMessage := "Направьте описание проблемы или ошибки инженеру \nТГ: @IgorMaksimov2000\nПочта: igor.maksimov@myoffice.team \n\n" +
		"Формат сообщения должен включать: \n" +
		"1. Описание ошибки/вопроса.\n" +
		"2. Выводы команд pip3 list и ansible --version. (Выполненные в корневой директории инсталляции, например, /root/install_pgs) \n" +
		"3. Конфигурационные файлы, которые были использованы при инсталляции. (Например, hosts.yml для PGS, hosts.yml/main.yml для CO) \n" +
		"4. Логи ошибок (Например, для развертывания СО сервера это будет файл deploy_co.log). \n\n" +
		"Спасибо! Инженер ответит вам в течение 10 минут. \n"
	msg := tgbotapi.NewMessage(chatID, errorMessage)
	bot.Send(msg)
}

func sendConfigFile(bot *tgbotapi.BotAPI, chatID int64, filePath, fileName string) {
	// Проверяем, существует ли файл
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		msg := tgbotapi.NewMessage(chatID, "Файл не найден.")
		if _, err := bot.Send(msg); err != nil {
			log.Println("Error sending message:", err)
		}
		return
	}

	// Открываем файл для отправки
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error opening file:", err)
		msg := tgbotapi.NewMessage(chatID, "Не удалось открыть файл.")
		if _, err := bot.Send(msg); err != nil {
			log.Println("Error sending message:", err)
		}
		return
	}
	defer file.Close()

	// Отправляем файл пользователю
	document := tgbotapi.NewDocument(chatID, tgbotapi.FileReader{
		Name:   fileName,
		Reader: file,
	})
	if _, err := bot.Send(document); err != nil {
		log.Println("Error sending document:", err)
	}
}
