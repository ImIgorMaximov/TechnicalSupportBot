/*
Package deployment предоставляет функции для процесса развертывания Standalone Squadus.

Эти функции предназначены для использования в боте, который помогает пользователям
пошагово выполнять установку Standalone Squadus, предоставляя необходимые инструкции.
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
func SendStandaloneRequirementsSquadus(bot *tgbotapi.BotAPI, chatID int64) {
	requirements := "Аппаратные и системные требования для установки Standalone Squadus c сайзингом:\n\n" +
		"Максимальное кол-во пользователей - 50; \n" +
		"*Данный сайзинг является примером, для более детального расчета обратитесь к инженеру @IgorMaksimov2000\n\n" +
		"Аппаратные требования: \n" +
		"Виртуальная машина с ролью - operator (Для управления процессом установки)\n\n" +
		"Operator: 18 (CPU, vCPU); 24 GB (RAM), 100 GB (SSD)\n" +
		"Cистемные требования (OS): \n" +
		"- Astra Linux Special Edition 1.7 «Орел» (базовый);\n" +
		"- РЕД ОС 7.3 Муром (версия ФСТЭК);\n" +
		"- CentOS 7.7;\n" +
		"- Ubuntu 20.04\n" +
		"Нажмите далее для продолжения. :)\n"

	msg := tgbotapi.NewMessage(chatID, requirements)
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}

// SendPrivateKeyInsertSquadus отправляет сообщение с инструкцией по установке публичного ключа на ВМ.
// Сообщение включает команды для генерации ключа и добавления его в файл authorized_keys на ВМ.
func SendPrivateKeyInsertSquadus(bot *tgbotapi.BotAPI, chatID int64) {
	privateKeyInsert := "Необходимо убедиться, что публичныЙ ключ на ВМ находятся папке /root/.ssh/authorized_keys.\n" +
		"Если ключ отсутствует, то создайте с помощью команды: \n\n" +
		"ssh-keygen\n\n" +
		"Затем скопируйте публичный ключ из файла /root/.ssh/id_rsa.pub в папку /root/.ssh/authorized_keys:\n\n" +
		"ssh-copy-id -i /root/.ssh/id_rsa.pub root@<IP_адрес_или_домен_машины_Operator> \n"
	msg := tgbotapi.NewMessage(chatID, privateKeyInsert)
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}

// SendDNSOptionsSquadus отправляет сообщение с рекомендациями по настройке DNS-сервера перед установкой Squadus.
// Сообщение описывает, как настраивать DNS-записи для различных сервисов Squadus в зависимости от использования переменных окружения.
func SendDNSOptionsSquadus(bot *tgbotapi.BotAPI, chatID int64) {
	dns := "Перед началом установки необходимо настроить DNS-сервер.\n" +
		"В случае использования переменной окружения (env) в конфигурационном файле hosts.yml записи будут иметь вид: \n\n" +
		"im-<env>.<default_domain> \n" +
		"go-<env>.<default_domain> \n" +
		"meet-<env>.<default_domain> \n" +
		"scc-<env>.<default_domain> \n" +
		"preview-<env>.<default_domain> \n" +
		"turn-<env>.<default_domain> \n" +
		"editor-<env>.<default_domain> \n\n" +
		"Если переменная окружения (env) не задана, записи примут вид:\n\n" +
		"im.<default_domain>\n" +
		"go.<default_domain>\n" +
		"meet.<default_domain>\n" +
		"scc.<default_domain>\n" +
		"preview.<default_domain>\n" +
		"turn.<default_domain>\n" +
		"editor.<default_domain>\n"
	msg := tgbotapi.NewMessage(chatID, dns)
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}

// SendSquadusInstallation отправляет сообщение с инструкцией по загрузке и распаковке дистрибутива Squadus.
func SendStandaloneDownloadDistributionSquadus(bot *tgbotapi.BotAPI, chatID int64) {
	installation := "Переходим к установке и настройке Squadus сервера.\n\n" +
		"На машину operator перенести дистрибутив Squadus, который выдается инженером или Аккаунт Менеджером. \n\n" +
		"Данный дистрибутив (.iso) включает: \n\n" +
		"squadus_ansible_bin_version.run - файл с  подсистемой управления конфигурациями\n" +
		"squadus_infra_version.run - файл с  с хранилищем Docker-контейнеров\n\n" +
		"Далее выполните запуск скрипта с хранилищем Docker-контейнеров:\n\n" +
		"bash squadus_infra_version.run \n\n" +
		"После завершения установки необходимо убедиться, что список содержит сообщения [ OK ] или [CHANGE]\n\n" +
		"Далее выполните запуск скрипта squadus_ansible_bin_version.run :\n\n" +
		"bash squadus_ansible_bin_version.run \n\n" +
		"После завершения установки необходимо убедиться, что список содержит сообщения [ OK ] или [CHANGE]\n\n" +
		"Перейдите в каталог /install_squadus/ :\n\n" +
		"cd /root/install_squadus\n\n" +
		"Скопируйте файл /root/install_squadus/contrib/squadus/ansible.cfg в /root/install_squadus/ :\n\n" +
		"cp /root/install_squadus/contrib/squadus/ansible.cfg ansible.cfg\n\n" +
		"Скопируйте файл /root/install_squadus/contrib/squadus/standalone_hosts.yml в /root/install_squadus/:\n\n" +
		"cp /root/install_squadus/contrib/squadus/standalone_hosts.yml hosts.yml\n\n" +
		"Перенесите заготовку файлов параметров group_vars с помощью команды: :\n\n" +
		"cp -r /root/install_squadus/contrib/squadus/group_vars/squadus_setup /root/install_squadus/group_vars/\n\n"

	msg := tgbotapi.NewMessage(chatID, installation)
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}

// SendCertificatesAndKeysSquadus отправляет инструкции по установке SSL-сертификатов для Squadus.
func SendCertificatesAndKeysSquadus(bot *tgbotapi.BotAPI, chatID int64) {
	certificatesAndKeys := "Для работы Squadus необходима установка SSL-сертификатов.\n" +
		"Рекомендуется использовать сертификаты, полученные от публичных центров сертификации.\n" +
		"Сертификаты необходимо разместить в каталоге certificates.\n\n" +
		"cd /root/install_squadus/certificates\n" +
		"Вставьте серитификаты в директорию certificates.\n\n Список необходимых сертификатов: \n\n" +
		"server.crt - сертификат внешнего домена. \n" +
		"server.nopass.key - ключ внешнего домена. \n" +
		"ca.сrt - цепочка сертификатов промежуточных центров сертификации.\n\n" +
		"Проверить наличия сертификатов и ключа:\n" +
		"ls -la /root/install_squadus/certificates/\n\n"
	msg := tgbotapi.NewMessage(chatID, certificatesAndKeys)
	msg.ReplyMarkup = keyboards.GetIsCertificatesKeyboard()
	bot.Send(msg)
}

// SendSquadusDeploy отправляет команду для развертывания Squadus.
func SendSquadusDeploy(bot *tgbotapi.BotAPI, chatID int64) {
	deploy := "Для запуска установки Squadus необходимо перейти в каталог /root/install_squadus/ и выполнить следующую команду:\n\n" +
		"ansible-playbook playbooks/main.yml --diff\n\n" +
		"Ожидаем результат! При возниковении ошибок или вопросов обращайтесь к инженеру!\n"
	msg := tgbotapi.NewMessage(chatID, deploy)
	msg.ReplyMarkup = keyboards.GetFinishKeyboard()
	bot.Send(msg)
}
