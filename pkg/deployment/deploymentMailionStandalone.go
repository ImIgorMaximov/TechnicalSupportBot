/*
Package deployment предоставляет функции для процесса развертывания Standalone Mailion.

Эти функции предназначены для использования в боте, который помогает пользователям
пошагово выполнять установку Standalone Mailion, предоставляя необходимые инструкции.
Функции используют библиотеку tgbotapi для взаимодействия с Telegram API.

Автор: Максимов Игорь
Email: imigormaximov@gmail.com
*/

package deployment

import (
	"technicalSupportBot/pkg/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// SendStandaloneRequirementsSquadus отправляет пользователю информацию о системных и аппаратных требованиях
// для установки Standalone Squadus с указанным сайзингом.
func SendStandaloneRequirementsMailion(bot *tgbotapi.BotAPI, chatID int64) {
	requirements := "Аппаратные и системные требования для установки Standalone Mailion c сайзингом:\n\n" +
		"Максимальное кол-во пользователей - 50; \n" +
		"*Данный сайзинг является примером, для более детального расчета обратитесь к инженеру @IgorMaksimov2000\n\n" +
		"Аппаратные требования: \n" +
		"Виртуальная машина с ролью - operator (Для управления процессом установки)\n\n" +
		"Operator: 12 (CPU, vCPU); 32 GB (RAM), 50 (HDD), 50 GB (SSD)\n" +
		"Cистемные требования (OS): \n" +
		"- Astra Linux Special Edition 1.7 «Орел» (базовый);\n" +
		"- РЕД ОС 7.3 Муром (версия ФСТЭК);\n" +
		"- CentOS 7.7;\n" +
		"- Ubuntu 20.04\n" +
		"- CentOS 7.7+\n" +
		"Нажмите далее для продолжения. :)\n"

	msg := tgbotapi.NewMessage(chatID, requirements)
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}

// SendPrivateKeyInsertMailion отправляет сообщение с инструкцией по установке публичного ключа на ВМ.
// Сообщение включает команды для генерации ключа и добавления его в файл authorized_keys на ВМ.
func SendPrivateKeyInsertMailion(bot *tgbotapi.BotAPI, chatID int64) {
	privateKeyInsert := "Необходимо убедиться, что публичныЙ ключ на ВМ находятся папке /root/.ssh/authorized_keys.\n" +
		"Если ключ отсутствует, то создайте с помощью команды: \n\n" +
		"ssh-keygen\n\n" +
		"Затем скопируйте публичный ключ из файла /root/.ssh/id_rsa.pub в папку /root/.ssh/authorized_keys:\n\n" +
		"ssh-copy-id -i /root/.ssh/id_rsa.pub root@<IP_адрес_или_домен_машины_Operator> \n"
	msg := tgbotapi.NewMessage(chatID, privateKeyInsert)
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}

// SendDNSOptionsMailion отправляет сообщение с рекомендациями по настройке DNS-сервера перед установкой Mailion.
// Сообщение описывает, как настраивать DNS-записи для различных сервисов Squadus в зависимости от использования переменных окружения.
func SendDNSOptionsMailion(bot *tgbotapi.BotAPI, chatID int64) {
	dns := "Перед началом установки необходимо настроить DNS-сервер.\n" +
		"В случае использования переменной окружения (env) в конфигурационном файле main.yml записи будут иметь вид: \n\n" +
		"api-<env>.<default_domain> - CNAME, значение @\n" +
		"auth-<env>.<default_domain> - CNAME, значение @\n" +
		"autoconfig-<env>.<default_domain> - CNAME, значение @\n" +
		"avatars-<env>.<default_domain> - CNAME, значение @\n" +
		"caldav-<env>.<default_domain> - CNAME, значение @\n" +
		"carddav-<env>.<default_domain> - CNAME, значение @\n" +
		"grpc-<env>.<default_domain> - CNAME, значение @\n" +
		"imap-<env>.<default_domain> - CNAME, значение @\n" +
		"mail-<env>.<default_domain> - CNAME, значение @\n" +
		"preview-<env>.<default_domain> - CNAME, значение @\n" +
		"relay-<env>.<default_domain> - A, значение <ucs_mail_relay_vip>, *адрес самого сервера этой группы\n" +
		"resources-<env>.<default_domain> - CNAME, значение @ \n" +
		"secured-<env>.<default_domain> - CNAME, значение @\n" +
		"smtp-<env>.<default_domain> - A, значение <ucs_mail_vip>\n" +
		"db-<env>.<default_domain> - CNAME, значение @\n\n" +
		"Если переменная окружения (env) не задана, записи примут вид:\n\n" +
		"api.<default_domain> \n" +
		"auth.<default_domain> \n" +
		"autoconfig.<default_domain> \n" +
		"avatars.<default_domain> \n" +
		"caldav.<default_domain> \n" +
		"carddav.<default_domain> \n" +
		"grpc.<default_domain> \n" +
		"imap.<default_domain> \n" +
		"mail.<default_domain> \n" +
		"preview.<default_domain> \n" +
		"relay.<default_domain> \n" +
		"resources.<default_domain> \n" +
		"secured.<default_domain> \n" +
		"smtp.<default_domain> \n" +
		"db.<default_domain> \n\n"
	msg := tgbotapi.NewMessage(chatID, dns)
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}

// SendMailionInstallation отправляет сообщение с инструкцией по загрузке и распаковке дистрибутива Mailion.
func SendStandaloneDownloadDistributionMailion(bot *tgbotapi.BotAPI, chatID int64) {
	installation := "Переходим к установке и настройке Mailion сервера.\n\n" +
		"На машину operator перенести дистрибутив Mailion, который выдается инженером или Аккаунт Менеджером. \n\n" +
		"Данный дистрибутив (.iso) включает: \n\n" +
		"mailion_ansible_bin_version.run - файл с  подсистемой управления конфигурациями\n" +
		"mailion_infra_version.run - файл с  с хранилищем Docker-контейнеров\n\n" +
		"Далее выполните запуск скрипта с хранилищем Docker-контейнеров:\n\n" +
		"bash mailion_infra_version.run \n\n" +
		"После завершения установки необходимо убедиться, что список содержит сообщения [ OK ] или [CHANGE]\n\n" +
		"Далее выполните запуск скрипта mailion_ansible_bin_version.run :\n\n" +
		"bash mailion_ansible_bin_version.run \n\n" +
		"После завершения установки необходимо убедиться, что список содержит сообщения [ OK ] или [CHANGE]\n\n" +
		"Перейдите в каталог /install_mailion/ :\n\n" +
		"cd /root/install_mailion\n\n" +
		"Скопируйте файл /root/install_mailion/contrib/mailion/ansible.cfg в /root/install_mailion/ :\n\n" +
		"cp /root/install_mailion/contrib/mailion/ansible.cfg ansible.cfg\n\n" +
		"Скопируйте файл /root/install_mailion/contrib/mailion/standalone/hosts.yml в /root/install_mailion/:\n\n" +
		"cp /root/install_mailion/contrib/mailion/standalone/hosts.yml hosts.yml\n\n" +
		"Создайте в каталоге /root/install_mailion/group_vars/ директорию, например, mailion_setup: \n\n" +
		"mkdir /root/install_mailion/group_vars/mailion_setup \n\n" +
		"Перенесите заготовку файлов параметров group_vars с помощью команды: :\n\n" +
		"cp -r /root/install_mailion/contrib/mailion/standalone/group_vars/ucs_setup/* /root/install_mailion/group_vars/mailion_setup\n\n"

	msg := tgbotapi.NewMessage(chatID, installation)
	msg.ReplyMarkup = keyboards.GetUnzippingISOKeyboard()
	bot.Send(msg)
}

// SendCertificatesAndKeysMailion отправляет инструкции по установке SSL-сертификатов для Mailion.
func SendCertificatesAndKeysMailion(bot *tgbotapi.BotAPI, chatID int64) {
	certificatesAndKeys := "Для работы Mailion необходима установка SSL-сертификатов.\n" +
		"Рекомендуется использовать сертификаты, полученные от публичных центров сертификации.\n" +
		"Сертификаты необходимо разместить в каталоге mailion.\n\n" +
		"cd /root/install_mailion/certificates\n" +
		"Вставьте серитификаты в директорию certificates.\n\n Список необходимых сертификатов: \n\n" +
		"server.crt - сертификат внешнего домена. \n" +
		"server.nopass.key - ключ внешнего домена. \n" +
		"ca.pem - цепочка сертификатов промежуточных центров сертификации.\n\n" +
		"Проверить наличия сертификатов и ключа:\n" +
		"ls -la /root/install_mailion/certificates/\n\n"
	msg := tgbotapi.NewMessage(chatID, certificatesAndKeys)
	msg.ReplyMarkup = keyboards.GetIsCertificatesKeyboard()
	bot.Send(msg)
}

// SendStandaloneMailionConfigure отправляет инструкции по настройке файла hosts.yml & main.yml для Mailion.
func SendStandaloneMailionConfigure(bot *tgbotapi.BotAPI, chatID int64) {
	configure := "Заполним файл hosts.yml в директории /root/install_mailion/:\n" +
		"vim /root/install_mailion/hosts.yml\n\n" +
		"В секцию hosts добавьте доменное имя вашего Squadus-сервера: \n" +
		"hosts:\n" +
		"\t\tmailion.myoffice-app.ru: \n" +
		"Операцию необходимо проделать со всеми сервисами.\n\n" +
		"Приступаем заполнять конфиг main.yml в директории /root/install_mailion/group_vars/mailion_setup :\n" +
		"vim /root/install_mailion/group_vars/mailion_setup/main.yml\n\n" +
		"Заполните домен: \n" +
		"mailion_external_domain: \"myoffice-app.ru\" \n" +
		"При использовании переменной окружения переменная domain_module примет ввид: \n" +
		"mailion_external_domain: \"{service}-env.{domain}\" \n\n" +
		"Добавьте список доменов, которые будет обслуживаться Mailion : \n" +
		"mailion_supported_domains: \n" +
		"	- myoffice-app.ru \n\n" +
		"RSPAMD конфигурация : \n" +
		"rspamd_dkim_hosts: : \n" +
		"	myoffice-app.ru : \n" +
		"		dkim_key: \"{{ lookup('file', '/root/install_mailion/files/dkim.key') }}\" : \n\n" +
		"Cгенерируйте пароли для служб или оставьте по умолчанию. Примеры паролей можно посмотреть в конфигурационном файле main.yml.\n\n" +
		"*В примерах используется редактор vim \n" +
		"*При необходимости выберите пример конфига, нажав соответствующую кнопку. \n"

	msg := tgbotapi.NewMessage(chatID, configure)
	msg.ReplyMarkup = keyboards.GetMailionStandaloneConfigKeyboard()
	bot.Send(msg)
}

// SendMailionDeploy отправляет команду для развертывания Mailion.
func SendMailionDeploy(bot *tgbotapi.BotAPI, chatID int64) {
	deploy := "Для запуска установки Mailion необходимо перейти в каталог /root/install_mailion/ и выполнить следующую команду:\n\n" +
		"ansible-playbook playbooks/main.yml --diff\n\n" +
		"Ожидаем результат! При возниковении ошибок или вопросов обращайтесь к инженеру!\n"
	msg := tgbotapi.NewMessage(chatID, deploy)
	msg.ReplyMarkup = keyboards.GetFinishKeyboard()
	bot.Send(msg)
}
