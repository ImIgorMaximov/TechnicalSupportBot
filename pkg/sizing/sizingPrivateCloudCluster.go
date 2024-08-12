package sizing

import (
	"fmt"
	"log"

	"technicalSupportBot/pkg/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xuri/excelize/v2"
)

var previousStateCluster = make(map[int64]string)
var clusterInputValues = make(map[int64][]string)

func HandleClusterMoreThan2k(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Введите максимальное количество аккаунтов (Например, 1000) :")
	bot.Send(msg)
	previousStateCluster[chatID] = "awaitingMaxAccountCount"
}

func HandleClusterMoreThan2kInput(bot *tgbotapi.BotAPI, chatID int64, userInput string) {

	log.Printf("Пользователь %d начал ввод для диапазона <2k", chatID)
	log.Printf("Текущее состояние: %s", previousStateCluster[chatID])

	switch previousStateCluster[chatID] {
	case "awaitingMaxAccountCount":
		clusterInputValues[chatID] = []string{userInput}
		msg := tgbotapi.NewMessage(chatID, "Введите количество одновременно активных пользователей:")
		bot.Send(msg)
		previousStateCluster[chatID] = "awaitingActiveUserCount"
		log.Printf("Ожидание ввода максимального количества аккаунтов для пользователя %d", chatID)

	case "awaitingActiveUserCount":
		clusterInputValues[chatID] = append(clusterInputValues[chatID], userInput)
		msg := tgbotapi.NewMessage(chatID, "Введите количество документов, редактируемых одновременно:")
		bot.Send(msg)
		previousStateCluster[chatID] = "awaitingDocumentCount"

	case "awaitingDocumentCount":
		clusterInputValues[chatID] = append(clusterInputValues[chatID], userInput)
		msg := tgbotapi.NewMessage(chatID, "Использовать S3 хранилище для инсталляций с общим размером хранилища более 100ТБ? (да/нет):")
		bot.Send(msg)
		previousStateCluster[chatID] = "awaitingS3Storage"

	case "awaitingS3Storage":
		clusterInputValues[chatID] = append(clusterInputValues[chatID], userInput)
		msg := tgbotapi.NewMessage(chatID, "Введите дисковую квоту пользователя (ГБ):")
		bot.Send(msg)
		previousStateCluster[chatID] = "awaitingUserDiskQuota"

	case "awaitingUserDiskQuota":
		clusterInputValues[chatID] = append(clusterInputValues[chatID], userInput)
		msg := tgbotapi.NewMessage(chatID, "Введите дисковую квоту для общих папок (ГБ):")
		bot.Send(msg)
		previousStateCluster[chatID] = "awaitingSharedFolderQuota"

	case "awaitingSharedFolderQuota":
		clusterInputValues[chatID] = append(clusterInputValues[chatID], userInput)
		log.Printf("Данные для расчета: %v", clusterInputValues[chatID])
		calculateAndSendClusterSizing(bot, chatID)

	default:
		log.Printf("Ошибка: неизвестное состояние %s", previousStateCluster[chatID])
	}
}

func calculateAndSendClusterSizing(bot *tgbotapi.BotAPI, chatID int64) {
	filePath := "/home/admin-msk/Documents/sizingPrivateCloudCluster.xlsx"
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Println("Ошибка открытия файла:", err)
		msg := tgbotapi.NewMessage(chatID, "Произошла ошибка при открытии файла.")
		bot.Send(msg)
		return
	}
	defer f.Close()

	// Заполнение ячеек данными
	err = f.SetCellValue("Cluster<2k", "D4", clusterInputValues[chatID][0])
	err = f.SetCellValue("Cluster<2k", "F6", clusterInputValues[chatID][1])
	err = f.SetCellValue("Cluster<2k", "F7", clusterInputValues[chatID][2])
	err = f.SetCellValue("Cluster<2k", "D8", clusterInputValues[chatID][3])
	err = f.SetCellValue("Cluster<2k", "F11", clusterInputValues[chatID][4])
	err = f.SetCellValue("Cluster<2k", "D12", clusterInputValues[chatID][5])

	if err != nil {
		log.Println("Ошибка записи в файл:", err)
		msg := tgbotapi.NewMessage(chatID, "Произошла ошибка при записи в файл.")
		bot.Send(msg)
		return
	}

	if err := f.Save(); err != nil {
		log.Println("Ошибка сохранения файла:", err)
		msg := tgbotapi.NewMessage(chatID, "Произошла ошибка при сохранении файла.")
		bot.Send(msg)
		return
	}

	// Извлечение результатов
	operatorVM, _ := f.GetCellValue("Cluster<2k", "C21")
	operatorCPU, _ := f.GetCellValue("Cluster<2k", "D21")
	operatorRAM, _ := f.GetCellValue("Cluster<2k", "E21")
	operatorSSD, _ := f.GetCellValue("Cluster<2k", "F21")
	operatorHDD, _ := f.GetCellValue("Cluster<2k", "G21")

	lbVM, _ := f.GetCellValue("Cluster<2k", "C22")
	lbCPU, _ := f.GetCellValue("Cluster<2k", "D22")
	lbRAM, _ := f.GetCellValue("Cluster<2k", "E22")
	lbSSD, _ := f.GetCellValue("Cluster<2k", "F22")
	lbHDD, _ := f.GetCellValue("Cluster<2k", "G22")

	coCoreVM, _ := f.GetCellValue("Cluster<2k", "C23")
	coCoreCPU, _ := f.GetCellValue("Cluster<2k", "D23")
	coCoreRAM, _ := f.GetCellValue("Cluster<2k", "E23")
	coCoreSSD, _ := f.GetCellValue("Cluster<2k", "F23")
	coCoreHDD, _ := f.GetCellValue("Cluster<2k", "G23")

	coInfraVM, _ := f.GetCellValue("Cluster<2k", "C24")
	coInfraCPU, _ := f.GetCellValue("Cluster<2k", "D24")
	coInfraRAM, _ := f.GetCellValue("Cluster<2k", "E24")
	coInfraSSD, _ := f.GetCellValue("Cluster<2k", "F24")
	coInfraHDD, _ := f.GetCellValue("Cluster<2k", "G24")

	coMqImcEtcdVM, _ := f.GetCellValue("Cluster<2k", "C25")
	coMqImcEtcdCPU, _ := f.GetCellValue("Cluster<2k", "D25")
	coMqImcEtcdRAM, _ := f.GetCellValue("Cluster<2k", "E25")
	coMqImcEtcdSSD, _ := f.GetCellValue("Cluster<2k", "F25")
	coMqImcEtcdHDD, _ := f.GetCellValue("Cluster<2k", "G25")

	pgsAppVM, _ := f.GetCellValue("Cluster<2k", "C26")
	pgsAppCPU, _ := f.GetCellValue("Cluster<2k", "D26")
	pgsAppRAM, _ := f.GetCellValue("Cluster<2k", "E26")
	pgsAppSSD, _ := f.GetCellValue("Cluster<2k", "F26")
	pgsAppHDD, _ := f.GetCellValue("Cluster<2k", "G26")

	pgsStorageVM, _ := f.GetCellValue("Cluster<2k", "C27")
	pgsStorageCPU, _ := f.GetCellValue("Cluster<2k", "D27")
	pgsStorageRAM, _ := f.GetCellValue("Cluster<2k", "E27")
	pgsStorageSSD, _ := f.GetCellValue("Cluster<2k", "F27")
	pgsStorageHDD, _ := f.GetCellValue("Cluster<2k", "G27")

	pgsStorageAVM, _ := f.GetCellValue("Cluster<2k", "C28")
	pgsStorageACPU, _ := f.GetCellValue("Cluster<2k", "D28")
	pgsStorageARAM, _ := f.GetCellValue("Cluster<2k", "E28")
	pgsStorageASSD, _ := f.GetCellValue("Cluster<2k", "F28")
	pgsStorageAHDD, _ := f.GetCellValue("Cluster<2k", "G28")

	pgsBEVM, _ := f.GetCellValue("Cluster<2k", "C29")
	pgsBECPU, _ := f.GetCellValue("Cluster<2k", "D29")
	pgsBERAM, _ := f.GetCellValue("Cluster<2k", "E29")
	pgsBESSD, _ := f.GetCellValue("Cluster<2k", "F29")
	pgsBEHDD, _ := f.GetCellValue("Cluster<2k", "G29")

	pgsDBVM, _ := f.GetCellValue("Cluster<2k", "C30")
	pgsDBCPU, _ := f.GetCellValue("Cluster<2k", "D30")
	pgsDBRAM, _ := f.GetCellValue("Cluster<2k", "E30")
	pgsDBSSD, _ := f.GetCellValue("Cluster<2k", "F30")
	pgsDBHDD, _ := f.GetCellValue("Cluster<2k", "G30")

	pgsLOGVM, _ := f.GetCellValue("Cluster<2k", "C31")
	pgsLOGCPU, _ := f.GetCellValue("Cluster<2k", "D31")
	pgsLOGRAM, _ := f.GetCellValue("Cluster<2k", "E31")
	pgsLOGSSD, _ := f.GetCellValue("Cluster<2k", "F31")
	pgsLOGHDD, _ := f.GetCellValue("Cluster<2k", "G31")

	// Отправка результата пользователю
	resultMsg := fmt.Sprintf(
		"Результаты расчета сайзинга для продукта Частное Облако Cluster:\n\n"+
			"Operator: кол-во ВМ - %s, CPU - %s, RAM - %s ГБ, SSD - %s ГБ, HDD - %s ГБ;\n\n"+
			"LB: кол-во ВМ - %s, CPU - %s, RAM - %s ГБ, SSD - %s ГБ, HDD - %s ГБ;\n\n"+
			"CO core: кол-во ВМ - %s, CPU - %s, RAM - %s ГБ, SSD - %s ГБ, HDD - %s ГБ;\n\n"+
			"CO infra: кол-во ВМ - %s, CPU - %s, RAM - %s ГБ, SSD - %s ГБ, HDD - %s ГБ;\n\n"+
			"CO mq-imc-etcd: кол-во ВМ - %s, CPU - %s, RAM - %s ГБ, SSD - %s ГБ, HDD - %s ГБ;\n\n"+
			"PGS-APP: кол-во ВМ - %s, CPU - %s, RAM - %s ГБ, SSD - %s ГБ, HDD - %s ГБ;\n\n"+
			"PGS-STORAGE: кол-во ВМ - %s, CPU - %s, RAM - %s ГБ, SSD - %s ГБ, HDD - %s ГБ;\n\n"+
			"PGS-STORAGE-A: кол-во ВМ - %s, CPU - %s, RAM - %s ГБ, SSD - %s ГБ, HDD - %s ГБ;\n\n"+
			"PGS-BE: кол-во ВМ - %s, CPU - %s, RAM - %s ГБ, SSD - %s ГБ, HDD - %s ГБ;\n\n"+
			"PGS-DB: кол-во ВМ - %s, CPU - %s, RAM - %s ГБ, SSD - %s ГБ, HDD - %s ГБ;\n\n"+
			"PGS-LOG: кол-во ВМ - %s, CPU - %s, RAM - %s ГБ, SSD - %s ГБ, HDD - %s ГБ;\n\n",
		operatorVM, operatorCPU, operatorRAM, operatorSSD, operatorHDD,
		lbVM, lbCPU, lbRAM, lbSSD, lbHDD,
		coCoreVM, coCoreCPU, coCoreRAM, coCoreSSD, coCoreHDD,
		coInfraVM, coInfraCPU, coInfraRAM, coInfraSSD, coInfraHDD,
		coMqImcEtcdVM, coMqImcEtcdCPU, coMqImcEtcdRAM, coMqImcEtcdSSD, coMqImcEtcdHDD,
		pgsAppVM, pgsAppCPU, pgsAppRAM, pgsAppSSD, pgsAppHDD,
		pgsStorageVM, pgsStorageCPU, pgsStorageRAM, pgsStorageSSD, pgsStorageHDD,
		pgsStorageAVM, pgsStorageACPU, pgsStorageARAM, pgsStorageASSD, pgsStorageAHDD,
		pgsBEVM, pgsBECPU, pgsBERAM, pgsBESSD, pgsBEHDD,
		pgsDBVM, pgsDBCPU, pgsDBRAM, pgsDBSSD, pgsDBHDD,
		pgsLOGVM, pgsLOGCPU, pgsLOGRAM, pgsLOGSSD, pgsLOGHDD,
	)

	msg := tgbotapi.NewMessage(chatID, resultMsg)
	bot.Send(msg)

	keyboard := keyboards.GetMainMenuWithPrivateCloudCluster2kRolesKeyboard()
	msgWithKeyboard := tgbotapi.NewMessage(chatID, "Выберите следующую опцию:")
	msgWithKeyboard.ReplyMarkup = keyboard
	if _, err := bot.Send(msgWithKeyboard); err != nil {
		log.Printf("Ошибка отправки клавиатуры: %v", err)
	}
}
