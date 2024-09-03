/*
Package deployment предоставляет функции для процесса развертывания Почтового сервера PSN,

	Эти функции предназначены для использования в боте, который помогает пользователям
	пошагово выполнять установку PSN, предоставляя необходимые инструкции.
	Функции используют библиотеку tgbotapi для взаимодействия с Telegram API.

	Автор: Максимов Игорь
	Email: imigormaximov@gmail.com
*/
package deployment

import (
	"technicalSupportBot/pkg/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// SendStandaloneRequirementsPSN отправляет сообщение с требованиями для установки PSN в режиме Standalone.
// Сообщение включает аппаратные и системные требования, а также контактные данные для получения
// дополнительной помощи по сайзингу от инженера.
//
// Параметры:
// - bot: Экземпляр бота Telegram.
// - chatID: Идентификатор чата Telegram, в который будет отправлено сообщение.
func SendStandaloneRequirementsPSN(bot *tgbotapi.BotAPI, chatID int64) {
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
		"- РЕД ОС 7.3.2 Муром (версия ФСТЭК);\n" +
		"- CentOS 7.9;\n" +
		"Нажмите далее для продолжения. :)\n"

	msg := tgbotapi.NewMessage(chatID, requirements)
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}

// SendPrivateKeyInsertPSN отправляет сообщение с инструкцией по установке публичного ключа на машину PSN.
// Сообщение включает команды для генерации ключа и добавления его в файл authorized_keys на PSN.
func SendPrivateKeyInsertPSN(bot *tgbotapi.BotAPI, chatID int64) {
	privateKeyInsert := "Необходимо убедиться, что публичныЙ ключ на машине PSN находятся папке /root/.ssh/authorized_keys.\n" +
		"Если ключ отсутствует, то создайте с помощью команды: \n\n" +
		"ssh-keygen\n\n" +
		"Затем скопируйте публичный ключ из файла /root/.ssh/id_rsa.pub в папку /root/.ssh/authorized_keys:\n\n" +
		"ssh-copy-id -i /root/.ssh/id_rsa.pub root@<IP_адрес_или_домен_машины_Operator> \n"
	msg := tgbotapi.NewMessage(chatID, privateKeyInsert)
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}

// SendDNSOptionsPSN отправляет сообщение с рекомендациями по настройке DNS-сервера перед установкой PSN.
// Сообщение описывает, как настраивать DNS-записи для различных сервисов PSN в зависимости от использования переменных окружения.
func SendDNSOptionsPSN(bot *tgbotapi.BotAPI, chatID int64) {
	dns := "Перед началом установки необходимо настроить DNS-сервер.\n" +
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
	msg := tgbotapi.NewMessage(chatID, dns)
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}

// SendStandaloneDownloadDistributionPSN отправляет сообщение с инструкцией по загрузке и распаковке дистрибутива PSN.
// Сообщение описывает, как подготовить директорию и распаковать архив с дистрибутивом PSN.
func SendStandaloneDownloadDistributionPSN(bot *tgbotapi.BotAPI, chatID int64) {
	standaloneDownloadDistribution := "После установки необходимых пакетов на машине PSN или operator подготовьте архив, который выдается инженером или Аккаунт Менеджером.\n" +
		"Далее создайте директорию с помощью команды: \n\n" +
		"mkdir  install_psn\n\n" +
		"Распакуйте данный архив командой:\n\n" +
		"tar xvzf MyOffice_PSN_SRV-XXX.tgz -C install_psn \n" +
		"*vesion - введите соответствующую версию продукта \n\n" +
		"После этого перейдите в каталог install_psn: \n\n" +
		"cd install_psn\n"
	msg := tgbotapi.NewMessage(chatID, standaloneDownloadDistribution)
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}

// SendCertificatesAndKeysPSN отправляет сообщение с инструкциями по установке SSL-сертификатов для веб-интерфейса PSN.
// Сообщение описывает, какие сертификаты и ключи необходимы и где их разместить.
func SendCertificatesAndKeysPSN(bot *tgbotapi.BotAPI, chatID int64) {
	certificatesAndKeys := "Для работы веб-интерфейса PSN необходима установка SSL-сертификатов.\n" +
		"Рекомендуется использовать сертификаты, полученные от публичных центров сертификации.\n" +
		"Сертификаты необходимо разместить в каталоге, соответствующему доменному имени PSN.\n\n" +
		"Напримере домена myoffice-app.ru : \n\n" +
		"cd /root/install-psn/certificates\n" +
		"mkdir myoffice-app.ru\n\n" +
		"Вставьте серитификаты в директорию, соответствующую вашему доменному имени.\n\n Список необходимых сертификатов: \n\n" +
		"server.crt - содержит SSL-сертификат для *.<default_domain> и все промежуточные сертификаты, кроме корневого доверенного. \n" +
		"server.nopass.key - Приватный ключ сертификата, не требующий кодовой фразы. \n" +
		"ca.crt - файл сертификата удостоверяющего центра.\n\n" +
		"Проверить наличия сертификатов и ключа:\n\n" +
		"ls -la /root/install-psn/certificates/myoffice-app.ru\n\n" +
		"Далее начинаем заполнять конфигурационные файлы!:)\n"
	msg := tgbotapi.NewMessage(chatID, certificatesAndKeys)
	msg.ReplyMarkup = keyboards.GetIsCertificatesKeyboard()
	bot.Send(msg)
}

// SendStandalonePSNConfigure отправляет сообщение с инструкцией по заполнению конфигурационного файла hosts.yml.
// В этой функции подробно описываются шаги, необходимые для настройки различных сервисов PSN.
func SendStandalonePSNConfigure(bot *tgbotapi.BotAPI, chatID int64) {
	configure := "Заполним конфигурационный файл hosts.yml : \n\n" +
		"vim /root/install_psn/inventory/hosts.yml\n\n" +
		"В секцию hosts добавьте доменное имя вашего PGS-сервера: \n" +
		"hosts:\n" +
		"\t\tpsn.myoffice-app.ru: \n" +
		"Аналогично проделать с другими сервисами: etcd, redis, postgres, ldap...\n\n" +
		"Далее в секцию vars необходимо заполнить следующие переменные:\n\n" +
		"external_domain: \"myoffice-app.ru\"\n\n" +
		"domain_module: \"{service}-{domain}\" - при данном значении, когда между {service} и {domain} стоит дефис, то домен psn.myoffice-app.ru будет разрешаться admin-psn.myoffice-app.ru\n\n" +
		"cert_path: \"certificates/myoffice-app.ru/\" - укажите директорию размещения сертификатов \n\n" +
		"Сгенерируйте и внесите пароли для переменных (команда: pwgen 13 10) : \n\n" +
		"postgres_superuser: \"1uUkvcjs6FLggc\"\n\n" +
		"postgres_replica_user: \"fpRZwbRN5hoQPqM\"\n\n" +
		"postgres_db_user: \"o6CPe2UDJffp4rw\"\n\n" +
		"redis_user: \"sQKwsWBikHr81t\"\n\n" +
		"rabbitmq_user: \"GGMQSzG4TMm5R37\"\n\n" +
		"ds389_manager_user: \"1Je4tr1Srp2PXF1\"\n\n" +
		"ds389_replicator_user: \"imnaBFw7u17zCt4\"\n\n" +
		"dovecot_adm_user: \"kmgiyJ2TH8Mry9Z\"\n\n" +
		"psnapi_adm_user: \"T8tSKZZK6eVGjU5\"\n\n" +
		"etcd_browser_user: \"4svRPsCbzuaDxs\"\n\n" +
		"db_secret_key: \"Qwerty1234567890\"\n\n" +
		"internal_secret_key: \"Qwerty1234567890\"\n\n" +
		"auth_jwt_key: \"Qwerty1234567890\"\n\n" +
		"*В примерах используется редактор vim \n" +
		"*При необходимости выберите пример конфига, нажав соответствующую кнопку. \n"
	msg := tgbotapi.NewMessage(chatID, configure)
	msg.ReplyMarkup = keyboards.GetPSNStandaloneConfig()
	bot.Send(msg)
}

// SendPSNDeploy отправляет сообщение с инструкцией по запуску процесса установки PSN.
// Сообщение включает команды, необходимые для выполнения скрипта развертывания, и рекомендации по обращению к инженеру в случае ошибок.
func SendPSNDeploy(bot *tgbotapi.BotAPI, chatID int64) {
	deploy := "Для запуска установки PSN необходимо перейти в каталог /root/install-psn/ и выполнить следующую команду:\n\n" +
		"./deploy.sh inventory/hosts.yml\n\n" +
		"Ожидаем результат! При возниковении ошибок или вопросов свяжитесь с инженером!\n"
	msg := tgbotapi.NewMessage(chatID, deploy)
	msg.ReplyMarkup = keyboards.GetFinishKeyboard()
	bot.Send(msg)
}
