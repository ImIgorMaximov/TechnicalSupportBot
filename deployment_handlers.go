package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func sendDeploymentOptions(bot *tgbotapi.BotAPI, chatID int64) {
	previousState[chatID] = "deployment"

	chooseProductMessage := "Выберите продукт для развертывания:"
	msg := tgbotapi.NewMessage(chatID, chooseProductMessage)
	msg.ReplyMarkup = getProductDeploymentKeyboard()
	bot.Send(msg)
}

func sendProductDeploymentInstructions(bot *tgbotapi.BotAPI, chatID int64, product string) {
	chooseInstallationTypeMessage := "Выберите тип установки для " + product + ":"
	msg := tgbotapi.NewMessage(chatID, chooseInstallationTypeMessage)
	msg.ReplyMarkup = getInstallationTypeKeyboard()
	bot.Send(msg)
}

func sendStandaloneRequirements(bot *tgbotapi.BotAPI, chatID int64, product string) {
	requirements := "Аппаратные и системные требования для установки Standalone Частное Облако c сайзингом:\n" +
		"Максимальное кол-во пользователей - 50; \n" +
		"Количество одновременно активных пользователей - 10; \n" +
		"Количество документов, редактируемых одновременно - 10; \n" +
		"Дисковая квота пользователя в хранилище, Гб - 1; \n" +
		"*Данный сайзинг является примером, для более детального расчета обратитесь к инженеру @IgorMaksimov2000\n\n" +
		"Аппаратные требования: \n" +
		"3 Виртуальные машины с ролями - operator (Для управления процессом установки), PGS (Система хранения данных), CO (Система редактирования и совместной работы)\n" +
		"Operator: 1 (CPU, vCPU); 4 GB (RAM), 50 GB (SSD)\n" +
		"PGS: 8 (CPU, vCPU); 20 GB (RAM), 150 GB (SSD)\n" +
		"CO: 8 (CPU, vCPU); 20 GB (RAM), 100 GB (SSD)\n" +
		"Cистемные требования (OS): \n" +
		"- Astra Linux Special Edition 1.7 «Орел» (базовый);\n" +
		"- РЕД ОС 7.3 Муром (версия ФСТЭК);\n" +
		"- CentOS 7.9;\n" +
		"- Ubuntu 22.04\n" +
		"Нажмите далее для продолжения. :)\n"

	msg := tgbotapi.NewMessage(chatID, requirements)
	msg.ReplyMarkup = getStandaloneNextStepKeyboard()
	bot.Send(msg)
}

func sendClusterDevelopmentMessage(bot *tgbotapi.BotAPI, chatID int64) {
	developmentMessage := "Данный раздел находится в разработке.. Приносим извинения😥"
	msg := tgbotapi.NewMessage(chatID, developmentMessage)
	msg.ReplyMarkup = getBackKeyboard()
	bot.Send(msg)
}

func sendStandaloneDownloadPackages(bot *tgbotapi.BotAPI, chatID int64) {
	downloadPackages := "Отлично! Тачки подготовлены! Двигаемся дальше..\n" +
		"PS. Вся установка и настройка будет производиться на машине operator на примере системы Astra Linux Special Edition 1.7 «Орел» (базовый);\n" +
		"На ВМ c ролью operator обновите систему: \n" +
		"sudo su\n" +
		"apt update\n" +
		"Далее установим необходимые пакеты: \n" +
		"apt install -y python3-pip \n" +
		"python3 -m pip install ansible-core==2.11.12 \n" +
		"python3 -m pip install ansible==4.9.0 \n" +
		"python3 -m pip install jinja2==3.1.2 \n" +
		"python3 -m pip install yamllint \n" +
		"На этом все :) Двигаемся дальше..\n"
	msg := tgbotapi.NewMessage(chatID, downloadPackages)
	msg.ReplyMarkup = getStandaloneDownloadPackages()
	bot.Send(msg)
}
