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

// SendStandaloneSquadusConfigure отправляет инструкции по настройке файла hosts.yml & main.yml для Squadus.
func SendSquadusConfigure(bot *tgbotapi.BotAPI, chatID int64) {
	configure := "Заполним файл hosts.yml в директории /root/install_squadus/:\n" +
		"vim /root/install_squadus/hosts.yml\n\n" +
		"В секцию hosts добавьте доменное имя вашего Squadus-сервера: \n" +
		"hosts:\n" +
		"\t\tsquadus.myoffice-app.ru: \n" +
		"Операцию необходимо проделать со всеми сервисами: squadus_apps, squadus_converter, squadus_db, squadus_ha, squadus_infra...\n\n" +
		"Приступаем заполнять конфиг main.yml в директории /root/install_squadus/group_vars/squadus_setup :\n" +
		"vim /root/install_squadus/group_vars/squadus_setup/main.yml\n\n" +
		"Заполните переменные окружения: \n" +
		"squadus_domain: \"myoffice-app.ru\" \n" +
		"При использовании domain_env co_domain_module примет ввид: \n" +
		"domain_module: \"{service}-env.{domain}\" \n\n" +
		"Сгенерируйте пароль для argon_mongodb_password: \"N9pXuaR6CNImDqFz4Pve\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для boron_mongodb_password: \"yooMiechai4hilohbie9\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для caddy2_web_auth_users: admin: \"Ooyesheis6eeyei9ooK5\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для calcium_mongodb_password: \"che7ohnee3eed1Chahf6\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для cobalt_mongodb_password: \"eefeo1IeR7yoolei3kei\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для copper_mongodb_password: \"kei6ohrie7aiHah9ooTi\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для docker_registry_password: \"oEeUL8kvCkkuxUia8ZaL\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для grafana_admin_password: \"ie7Ca4zequaezishooc6\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для iridium_mongodb_password: \"Shaja1goo6Ik0Eitaita\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для jibri_recorder_password: \"Ohyoog1iuxa6ual8fiiD\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для jibri_auth_password: \"Ohyoog1iuxa6ual8fiiD\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для jicofo_auth_password: \"oPa9uuwiezoo2iheiHoh\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для jicofo_component_secret: \"ahDapef9eifaiQu8laov\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для jitsi_jwt_secret: \"Ooquae1ephohsoa9UiHa\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для jitsi_transcriber_auth_password: \"Eengeimoo7as2aebeeSh\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для jvb_auth_password: \"eewof8aekiewei8Uoloh\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для jvb_component_secret: \"Iequ3doom7keeghai6Ah\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для keepalived_vrrp_instances:squadus_ha:password: \"zuX3ohm3\" (pwgen 8 1)\n\n" +
		"Сгенерируйте пароль для keepalived_vrrp_instances:squadus_meet_ha:password: \"aNg7ohc4\" (pwgen 8 1)\n\n" +
		"Сгенерируйте пароль для krypton_mongodb_password: \"iwee1phaeCohf7Heephi\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для minio_access_key: \"ooP9leiB7iridioK4Aeg\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для minio_secret_key: \"Uo1lebai3quodae2meep\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для minio_backup_access_key: \"voo3cho6LeeThutieNof\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для mongodb_root_password: \"ioc4peiPho4aiYoobie6\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для mongodb_secured_key: \"ixai0eenoonge0ho4Nae\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для mongodb_backup_mongodb_pbm_user_password: \"tiu2Eol4NohPh5ohtah7\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для mongodb_exporter_password: \"taey7Mie6oodei0aPhaeD\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для nats_authorization_password: \"quae8ei5Bie6ohj3mohr\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для nats_cluster_authorization_password: \"theewie5ohSheeshaequ\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для nginx_web_auth_users:admin:password: \"Ooyesheis6eeyei9ooK5\" (pwgen 20 1)\n\n" +
		"Сгенерируйте и вставьте приватный ключ для opendkim_domain_keys: \"im.myoffice-app.ru\" : opendkim-genkey -s im. -d im.myoffice-app.ru && cat im.myoffice-app.ru.private:\n\n" +
		"Сгенерируйте пароль для postgresql_password: \"znrNTzBWKeulm1xk76geqJBbTxH0lVH3v\" (pwgen 33 1 -s)\n\n" +
		"Сгенерируйте пароль для postgresql_exporter_password: \"RCmEYO4GveKc0WljJD2qOAIErIiSiPvRp\" (pwgen 33 1 -s)\n\n" +
		"Сгенерируйте пароль для postgresql_backup_password: \"LYgnwgpRpWSPD0vDbtdCbVq6z0byrjd9j\" (pwgen 33 1 -s)\n\n" +
		"Сгенерируйте пароль для scandium_password: \"uxei3ohkoo7NieMigoo8\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для radon_mongodb_password: \"oNohcoo4zaejaudieK9X\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для redis_password: \"pu2Fai2Bie6aiphiezie\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для sodium_mongodb_password: \"ahnai8Iew0ahx7aa9Xie\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для squadus_mongodb_password: \"ohqu0Zeevie9giezu9og\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для squadus_federation_admin_password: \"WrHnlIzpECyd4dNBaAhmIbwTkAbfoZAEq\" (pwgen 33 1 -s)\n\n" +
		"Сгенерируйте пароль для squadus_federation_as_token: \"f1VKHcd1k7FHefpLkIvKkzQw3eQeVS6i5\" (pwgen 33 1 -s)\n\n" +
		"Сгенерируйте пароль для squadus_federation_hs_token: \"DEYIjhSaBoRuVwC23tUhq4Fn5KIurSyf5\" (pwgen 33 1 -s)\n\n" +
		"Сгенерируйте пароль для squadus_jiratrigger_jira_password: \"5F5H6Fh3XSMcbZLaSbFvxbfXlkDbPCuU2\" (pwgen 33 1 -s)\n\n" +
		"Сгенерируйте пароль для squadus_jwt_secret: \"Vagi2uk2CheCahbohphe\" (pwgen 20 1)\n\n" +
		"Заполните smtp-адрес squadus_smtp_from_email: \"noreply@myoffice-app.ru\"\n\n" +
		"Заполните squash_servus_token: \"kohLeighe4AengeeNohn\"\n\n" +
		"Сгенерируйте пароль для squash_replication_token: \"wnaJyC0Ib1hoUBo9LY7k8DWONgb8ysj7Q\" (pwgen 33 1 -s)\n\n" +
		"Сгенерируйте пароль для synapse_macaroon_secret_key: \"0tiRDVnXaGLhU3gtg2CdfVogtBF1HGobi\" (pwgen 33 1 -s)\n\n" +
		"Сгенерируйте пароль для synapse_postgresql_password: \"1Vvc1EiSGHYeVL4m2eOzoyiEXQGo2mpTE\" (pwgen 33 1 -s)\n\n" +
		"Сгенерируйте пароль для synapse_registration_shared_secret: \"IIayYERMNKfoa7wZbn5X7xIADYJ1T2Ozy\" (pwgen 33 1 -s)\n\n" +
		"Сгенерируйте пароль для scandium_mongodb_password: \"le5Suul0lukookieQuai\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для tennessine_mongodb_password: \"ohvawah6ohzishiiNg3p\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для turnserver_cli_password: \"RcBaN5mKibWe25h6NWiN\" (pwgen 20 1 -s)\n\n" +
		"Сгенерируйте пароль для turnserver_secret: \"oogieyahneiBienei8ey\" (pwgen 20 1)\n\n" +
		"*В примерах используется редактор vim \n" +
		"*При необходимости выберите пример конфига, нажав соответствующую кнопку. \n"

	msg := tgbotapi.NewMessage(chatID, configure)
	msg.ReplyMarkup = keyboards.GetCOStandaloneConfigKeyboard()
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
