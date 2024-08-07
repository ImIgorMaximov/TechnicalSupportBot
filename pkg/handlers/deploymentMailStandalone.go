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

func sendPrivateKeyInsertPSN(bot *tgbotapi.BotAPI, chatID int64) {
	privateKeyInsertPSN := "Необходимо убедиться, что публичныЙ ключ на машине PSN находятся папке /root/.ssh/authorized_keys.\n" +
		"Если ключ отсутствует, то создайте с помощью команды: \n\n" +
		"ssh-keygen\n\n" +
		"Затем скопируйте публичный ключ из файла /root/.ssh/id_rsa.pub в папку /root/.ssh/authorized_keys:\n\n" +
		"ssh-copy-id -i /root/.ssh/id_rsa.pub root@<IP_адрес_или_домен_машины_Operator> \n"
	msg := tgbotapi.NewMessage(chatID, privateKeyInsertPSN)
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}

func sendDNSOptionsPSN(bot *tgbotapi.BotAPI, chatID int64) {
	dnsPGS := "Перед началом установки необходимо настроить DNS-сервер.\n" +
		"В случае использования переменной окружения (env) в конфигурационном файле hosts.yml записи будут иметь вид: \n\n" +
		"mailadmin-<env>.<default_domain> - Адрес веб-панели администрирования PSN \n" +
		"autoconfig-<env>.<default_domain> - Адрес сервиса автоконфигурирования для подключения клиентов MyOffice\n" +
		"mail-<env>.<default_domain> - Адрес главной страницы веб-интерфейса почты\n" +
		"cab-<env>.<default_domain> - Адрес для подключения к глобальной адресной книги\n" +
		"imap-<env>.<default_domain> - Адрес для подключения с сервису imap\n" +
		"smtp-<env>.<default_domain> - Адрес для подключения с сервису smtp\n" +
		"pbm-<env>.<default_domain> - Адрес для подключения с API сервису PBM\n\n" +
		"Если переменная окружения (env) не задана, записи примут вид:\n\n" +
		"mailadmin.<default_domain>\n" +
		"authoconfig.<default_domain>\n" +
		"mail.<default_domain>\n" +
		"cab.<default_domain>\n" +
		"imap.<default_domain>\n" +
		"smtp.<default_domain>\n" +
		"pbm.<default_domain>\n"
	msg := tgbotapi.NewMessage(chatID, dnsPGS)
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}
