package handlers

import (
	"technicalSupportBot/pkg/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func sendStandaloneRequirementsPSN(bot *tgbotapi.BotAPI, chatID int64) {
	requirements := "Аппаратные и системные требования для установки Standalone Почта c сайзингом:\n\n" +
		"Максимальное кол-во пользователей - 50; \n" +
		"Одновременно работающих пользователей, доля - 0,4; \n" +
		"Дисковая квота пользователя в почте, Гб - 1; \n" +
		"Писем в сутки на пользователя - 40; \n" +
		"Коэффициент спама, доля от предыдущей строки - 0,3; \n" +
		"Объем хранения логов PSN, Гб - 20; \n" +
		"*Данный сайзинг является примером, для более детального расчета обратитесь к инженеру @IgorMaksimov2000\n\n" +
		"Аппаратные требования: \n" +
		"1 Виртуальная машина с ролью - operator или PSN (Для управления процессом установки)\n\n" +
		"PSN: 4 (CPU, vCPU); 4 GB (RAM), 135 GB (HDD)\n\n" +
		"Cистемные требования (OS): \n" +
		"- Astra Linux Special Edition 1.7 «Орел» (базовый);\n" +
		"- РЕД ОС 7.3 Муром (версия ФСТЭК);\n" +
		"- CentOS 7.9;\n" +
		"- Ubuntu 22.04\n" +
		"Нажмите далее для продолжения. :)\n"

	msg := tgbotapi.NewMessage(chatID, requirements)
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}
