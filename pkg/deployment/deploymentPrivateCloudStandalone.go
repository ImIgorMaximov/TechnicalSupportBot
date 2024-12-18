/*
Package deployment предоставляет функции для процесса развертывания продуктов: Mailion, Частное Облако, Почта, Squadus.

Файл deploymentPrivateCloudStandalone.go предназначен для пошаговой установки продукта Частное Облако по типу инсталляции Standalone.
Функции используют библиотеку tgbotapi для взаимодействия с Telegram API.

Автор: Максимов Игорь
Email: imigormaximov@gmail.com
*/

package deployment

import (
	"technicalSupportBot/pkg/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// SendStandaloneRequirementsPrivateCloud отправляет пользователю информацию о системных и аппаратных требованиях
// для установки Standalone Private Cloud с указанным сайзингом.
func SendStandaloneRequirementsPrivateCloud(bot *tgbotapi.BotAPI, chatID int64) {
	requirements := "Аппаратные и системные требования для установки Standalone Частное Облако c сайзингом:\n\n" +
		"Максимальное кол-во пользователей - 50; \n" +
		"Количество одновременно активных пользователей - 10; \n" +
		"Количество документов, редактируемых одновременно - 10; \n" +
		"Дисковая квота пользователя в хранилище, Гб - 1; \n" +
		"*Данный сайзинг является примером, для более детального расчета обратитесь к инженеру @IgorMaksimov2000\n\n" +
		"Аппаратные требования: \n" +
		"3 Виртуальные машины с ролями - operator (Для управления процессом установки), PGS (Система хранения данных), CO (Система редактирования и совместной работы)\n\n" +
		"Operator: 1 (CPU, vCPU); 4 GB (RAM), 50 GB (SSD)\n" +
		"PGS: 8 (CPU, vCPU); 20 GB (RAM), 150 GB (SSD)\n" +
		"CO: 8 (CPU, vCPU); 20 GB (RAM), 100 GB (SSD)\n\n" +
		"Cистемные требования (OS): \n" +
		"- Astra Linux Special Edition 1.7 «Орел» (базовый);\n" +
		"- РЕД ОС 7.3 Муром (версия ФСТЭК);\n" +
		"- Ubuntu 22.04\n" +
		"Нажмите далее для продолжения. :)\n"

	msg := tgbotapi.NewMessage(chatID, requirements)
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboardWithIntegrationAD()
	bot.Send(msg)
}

// SendStandaloneDownloadPackages отправляет пользователю инструкции по установке необходимых пакетов на машину operator.
func SendStandaloneDownloadPackages(bot *tgbotapi.BotAPI, chatID int64) {
	downloadPackages := "Отлично! Двигаемся дальше..\n" +
		"PS. Вся установка и настройка будет производиться на машине operator на примере системы Astra Linux Special Edition 1.7.5 «Орел» (базовый);\n" +
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
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}

// SendPrivateKeyInsert отправляет инструкции по добавлению публичных ключей машин PGS и CO на машину Operator.
func SendPrivateKeyInsertPrivateCloud(bot *tgbotapi.BotAPI, chatID int64) {
	privateKeyInsert := "Необходимо убедиться, что публичные ключи машин PGS и CO находятся на машине Operator в папке /root/.ssh/authorized_keys.\n" +
		"Если ключи отсутствуют, создайте пары ключей на машинах PGS и CO с помощью команды: \n\n" +
		"ssh-keygen\n\n" +
		"Затем скопируйте публичные ключи из файлов /root/.ssh/id_rsa.pub на машину Operator в папку /root/.ssh/authorized_keys:\n\n" +
		"ssh-copy-id -i /root/.ssh/id_rsa.pub root@<IP_адрес_или_домен_машины_Operator> \n"
	msg := tgbotapi.NewMessage(chatID, privateKeyInsert)
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}

// SendDNSOptionsPGS отправляет инструкции по настройке DNS-сервера для PGS.
func SendDNSOptionsPGS(bot *tgbotapi.BotAPI, chatID int64) {
	dns := "Перед началом установки необходимо настроить DNS-сервер, указав адрес сервера установки Nginx.\n" +
		"В случае использования переменной окружения (env) в конфигурационном файле hosts.yml записи будут иметь вид: \n\n" +
		"admin-<env>.<default_domain> - Адрес веб-панели администрирования PGS \n" +
		"pgs-<env>.<default_domain> - Адрес точки входа для API\n\n" +
		"Если переменная окружения (env) не задана, записи примут вид:\n\n" +
		"admin.<default_domain>\n" +
		"pgs.<default_domain>\n\n"
	msg := tgbotapi.NewMessage(chatID, dns)
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}

// SendStandaloneDownloadDistributionPrivateCloud отправляет инструкции по подготовке архива с дистрибутивом для машины PGS.
func SendStandaloneDownloadDistributionPrivateCloud(bot *tgbotapi.BotAPI, chatID int64) {
	downloadDistr := "Первая установка будет произведена на машину PGS.\n" +
		"После установки необходимых пакетов на машине operator подготовьте архив, который выдается инженером или Аккаунт Менеджером.\n" +
		"Далее создайте директорию с помощью команды: \n\n" +
		"mkdir install_MyOffice_PGS\n\n" +
		"Распакуйте данный архив командой:\n\n" +
		"tar xf MyOffice_PGS_version.tgz -C install_MyOffice_PGS \n" +
		"*vesion - введите соответствующую версию продукта \n\n" +
		"После этого перейдите в каталог install_MyOffice_PGS: \n\n" +
		"cd install_MyOffice_PGS\n"
	msg := tgbotapi.NewMessage(chatID, downloadDistr)
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}

// SendCertificatesAndKeysPGS отправляет инструкции по установке SSL-сертификатов для PGS.
func SendCertificatesAndKeysPGS(bot *tgbotapi.BotAPI, chatID int64) {
	certificatesAndKeys := "Для работы веб-интерфейса PGS необходима установка SSL-сертификатов.\n" +
		"Рекомендуется использовать сертификаты, полученные от публичных центров сертификации.\n" +
		"Сертификаты необходимо разместить в каталоге, соответствующему доменному имени PGS.\n\n" +
		"Напримере домена myoffice-app.ru : \n\n" +
		"cd /root/install_MyOffice_PGS/certificates\n" +
		"mkdir myoffice-app.ru\n\n" +
		"Вставьте серитификаты в директорию, соответствующую вашему доменному имени.\n\n Список необходимых сертификатов: \n\n" +
		"server.crt - содержит SSL-сертификат для *.<default_domain> и все промежуточные сертификаты, кроме корневого доверенного. \n" +
		"server.nopass.key - Приватный ключ сертификата, не требующий кодовой фразы. \n" +
		"ca.crt - файл сертификата удостоверяющего центра.\n\n" +
		"Проверить наличия сертификатов и ключа:\n\n" +
		"ls -la /root/install_MyOffice_PGS/certificates/myoffice-app.ru\n\n" +
		"Далее начинаем заполнять конфигурационные файлы!:)\n"
	msg := tgbotapi.NewMessage(chatID, certificatesAndKeys)
	msg.ReplyMarkup = keyboards.GetIsCertificatesKeyboard()
	bot.Send(msg)
}

// SendStandalonePGSConfigure отправляет инструкции по настройке файла hosts.yml для PGS.
func SendStandalonePGSConfigure(bot *tgbotapi.BotAPI, chatID int64) {
	configure := "Необходимо скопировать шаблон файла inventory в корневой каталог дистрибутива и заполнить секции hosts и vars.\n\n" +
		"Операция копирования выполняется с помощью команды:\n\n" +
		"cp /root/install_MyOffice_PGS/inventory/hosts-sa.yaml hosts.yml \n\n" +
		"Далее заполним файл hosts.yml : \n\n" +
		"vim /root/install_MyOffice_PGS/hosts.yml\n\n" +
		"В секцию hosts добавьте доменное имя вашего PGS-сервера: \n" +
		"hosts:\n" +
		"\t\tpgs.myoffice-app.ru: \n" +
		"Аналогично проделать с другими сервисами: search, redis, storage, nginx, etcd...\n\n" +
		"Далее в секцию vars необходимо заполнить следующие переменные:\n\n" +
		"DEFAULT_DOMAIN: \"myoffice-app.ru\"\n\n" +
		"ENV: \"\" - *если используется переменная окружения\n\n" +
		"Сгенерируйте и внесите пароли для сервисов (команда: pwgen 13 7) : \n\n" +
		"KEYCLOAK_PASSWORD: \"81mToSPFJ8ezr8\"\n\n" +
		"KEYCLOAK_REALM_PASSWORD: \"MVh2PiA2S5cPk\"\n\n" +
		"KEYCLOAK_POSTGRES_PASSWORD: \"7Afd3G12P5VyUg\"\n\n" +
		"ARANGODB_PASSWORD: \"55ab8qk7ES4P4LX\"\n\n" +
		"RABBITMQ_PASSWORD: \"BdyYgDwLLY8M5U9\"\n\n" +
		"REDIS_PASSWORD: \"S73uo3iH3qFRdnf\"\n\n" +
		"GRAFANA_ADMIN_PASSWORD: \"oPpKvc6We3mES6\"\n\n" +
		"В секции co заполнить \"FS App encryption settings\" : \n\n" +
		"FS_APP_ENCRYPTION_SALT: \"2DD4E59B582AF71F\"\n\n" +
		"AUTH_ENCRYPTION_SALT: \"2DD4E59B582AF71F\"\n\n" +
		"APP_ADMIN_PASSWORD: \"6dbYv6qVJrqiVB\"\n\n" +
		"*В примерах используется редактор vim \n" +
		"*При необходимости выберите пример конфига, нажав соответствующую кнопку. \n"
	msg := tgbotapi.NewMessage(chatID, configure)
	msg.ReplyMarkup = keyboards.GetPGSStandaloneConfig()
	bot.Send(msg)
}

// SendPGSDeploy отправляет инструкции по развертыванию PGS.
func SendPGSDeploy(bot *tgbotapi.BotAPI, chatID int64) {
	deploy := "Для запуска установки PGS необходимо перейти в каталог /root/install_MyOffice_PGS/ и выполнить следующую команду:\n\n" +
		"./deploy.sh hosts.yml\n\n" +
		"Ожидаем результат! При возниковении ошибок или вопросов свяжитесь с инженером!\n"
	msg := tgbotapi.NewMessage(chatID, deploy)
	msg.ReplyMarkup = keyboards.GetCOInstallation()
	bot.Send(msg)
}

// SendDNSOptionsCO отправляет инструкции по настройке DNS-сервера для CO.
func SendDNSOptionsCO(bot *tgbotapi.BotAPI, chatID int64) {
	dns := "Перед началом установки необходимо настроить DNS-сервер.\n" +
		"В случае использования переменной окружения (env) в конфигурационном файле main.yml записи будут иметь вид: \n\n" +
		"auth-<domain_env>.<domain_name> \n" +
		"cdn-<domain_env>.<domain_name> \n" +
		"coapi-<domain_env>.<domain_name> \n" +
		"docs-<domain_env>.<domain_name> \n" +
		"files-<domain_env>.<domain_name> \n" +
		"links-<domain_env>.<domain_name> \n" +
		"_https._tcp-<domain_env>.<domain_name> \n\n" +
		"Если переменная окружения (env) не задана, записи примут вид:\n\n" +
		"auth.<domain_name>\n" +
		"cdn.<domain_name> \n" +
		"coapi.<domain_name> \n" +
		"docs.<domain_name> \n" +
		"files.<domain_name> \n" +
		"links.<domain_name> \n" +
		"_https._tcp.<domain_name> \n"
	msg := tgbotapi.NewMessage(chatID, dns)
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}

// SendCOInstallation отправляет сообщение с инструкцией по загрузке и распаковке дистрибутива CO.
func SendCOInstallation(bot *tgbotapi.BotAPI, chatID int64) {
	installation := "Переходим к установке и настройке CO (Сервер совместного редактирования).\n\n" +
		"На машину operator перенести дистрибутив CO, который выдается инженером или Аккаунт Менеджером. \n\n" +
		"Данный дистрибутив (.iso) включает: \n\n" +
		"co_ansible_bin_version.run - файл с  подсистемой управления конфигурациями\n" +
		"co_infra_version.run - файл с  с хранилищем Docker-контейнеров\n\n" +
		"Далее выполните запуск скрипта с хранилищем Docker-контейнеров:\n\n" +
		"bash co_infra_version.run \n\n" +
		"После завершения установки необходимо убедиться, что список содержит сообщения [ OK ] или [CHANGE]\n\n" +
		"Далее выполните запуск скрипта co_ansible_bin_version.run :\n\n" +
		"bash co_ansible_bin_version.run \n\n" +
		"После завершения установки необходимо убедиться, что список содержит сообщения [ OK ] или [CHANGE]\n\n" +
		"Перейдите в каталог /install_co/ :\n\n" +
		"cd /root/install_co\n\n" +
		"Скопируйте файл /root/install_co/contrib/co/ansible.cfg в /root/install_co/ :\n\n" +
		"cp /root/install_co/contrib/co/ansible.cfg ansible.cfg\n\n" +
		"Скопируйте файл /root/install_co/contrib/co/standalone/hosts.yml в /root/install_co/:\n\n" +
		"cp /root/install_co/contrib/co/standalone/hosts.yml hosts.yml\n\n" +
		"Создайте каталог co_setup/ в директории /root/install_co/group_vars/:\n\n" +
		"mkdir /root/install_co/group_vars/co_setup/\n\n" +
		"Скопируйте в созданную директорию co_setup каталог с переменными для заполнения :\n\n" +
		"cp -r /root/install_co/contrib/co/standalone/group_vars/co_setup/* /root/install_co/group_vars/co_setup/\n\n"

	msg := tgbotapi.NewMessage(chatID, installation)
	msg.ReplyMarkup = keyboards.GetUnzippingISOKeyboard()
	bot.Send(msg)
}

// SendCertificatesAndKeysCO отправляет инструкции по установке SSL-сертификатов для CO.
func SendCertificatesAndKeysCO(bot *tgbotapi.BotAPI, chatID int64) {
	certificatesAndKeys := "Для работы CO необходима установка SSL-сертификатов.\n" +
		"Рекомендуется использовать сертификаты, полученные от публичных центров сертификации.\n" +
		"Сертификаты необходимо разместить в каталоге certificates.\n\n" +
		"cd /root/install_co/certificates\n" +
		"Вставьте серитификаты в директорию certificates.\n\n Список необходимых сертификатов: \n\n" +
		"server.crt - сертификат внешнего домена. \n" +
		"server.nopass.key - ключ внешнего домена. \n" +
		"ca.сrt - цепочка сертификатов промежуточных центров сертификации.\n\n" +
		"Проверить наличия сертификатов и ключа:\n" +
		"ls -la /root/install_co/certificates/\n\n"
	msg := tgbotapi.NewMessage(chatID, certificatesAndKeys)
	msg.ReplyMarkup = keyboards.GetIsCertificatesKeyboard()
	bot.Send(msg)
}

// SendStandaloneCOConfigure отправляет инструкции по настройке файла hosts.yml & main.yml для CO.
func SendCOConfigure(bot *tgbotapi.BotAPI, chatID int64) {
	configure := "Заполним файл hosts.yml в директории /root/install_co/:\n" +
		"vim /root/install_co/hosts.yml\n\n" +
		"В секцию hosts добавьте доменное имя вашего CO-сервера: \n" +
		"hosts:\n" +
		"\t\tco.myoffice-app.ru: \n" +
		"Операцию необходимо проделать со всеми сервисами: co_chatbot, co_etcd, co_mq, co_cvm, co_cu...\n\n" +
		"Приступаем заполнять конфиг main.yml в директории /root/install_co/group_vars/co_setup :\n" +
		"vim /root/install_co/group_vars/co_setup/main.yml\n\n" +
		"Заполните переменные окружения: \n" +
		"domain_name: \"myoffice-app.ru\" \n" +
		"При использовании domain_env co_domain_module примет ввид: \n" +
		"co_domain_module: \"{service}-{domain}\" \n\n" +
		"В docker_daemon_parameters: ->  insecure-registries: [\"operator.myoffice-app.ru:5000\"] внесите домен оператора\n\n" +
		"Аналогично в docker_image_registry: \"operator.myoffice-app.ru:5000\"\n\n" +
		"Сгенерируйте пароль для etcd_browser_password: \"ail8Et8uiph5iegahqui\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для elasticsearch_admin_password: \"quai4Aigohchoo4uu4uThaeQuaigh4Vu\" (pwgen 32 1)\n\n" +
		"Сгенерируйте пароль для kibana_elasticsearch_password:  \"aeyeicee3jo8Be1Kiegieph4shahjeiw\" (pwgen 32 1)\n\n" +
		"Сгенерируйте пароль для redis_password: \"zohh4thie9IjaGhue5le\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для openresty_api_password: \"gpQfhLNLdvp82Y\" (pwgen 20 1)\n\n" +
		"Сгенерируйте пароль для openresty_mail_oauth2_client_secret: \"Nae9ea7ohgieVa8A\" (pwgen 16 1)\n\n" +
		"Измените значения переменных fs_api_url, fs_app_url, fs_card_url на домен PGS машины: \n" +
		"fs_api_url: \"https://pgs.myoffice-app.ru/pgsapi\" \n" +
		"fs_app_url: \"https://pgs.myoffice-app.ru/pgsapi\" \n" +
		"fs_card_url: \"https://pgs.myoffice-app.ru/pgsapi\" \n\n" +
		"Сравните значения переменных из конфигурационного файла hosts.yml PGS, они должны сопвадать:\n\n" +
		"auth_encryption_key = AUTH_ENCRYPTION_KEY : \"D1A693EB309C968A6EBC41787703DAE3B9C69405E5AE0FE6BF9CE2FF36CB8343\" \n\n" +
		"auth_encryption_iv = AUTH_ENCRYPTION_IV : \"7E3F053970AD7DE1A4394E10AE0F4022\" \n\n" +
		"auth_encryption_salt = AUTH_ENCRYPTION_SALT : \"2DD4E59B582AF71F\" \n\n" +
		"fs_app_encryption_key = FS_APP_ENCRYPTION_KEY : \"D1A693EB309C968A6EBC41787703DAE3B9C69405E5AE0FE6BF9CE2FF36CB8343\" \n\n" +
		"fs_app_encryption_iv = FS_APP_ENCRYPTION_IV : \"7E3F053970AD7DE1A4394E10AE0F4022\" \n\n" +
		"fs_app_encryption_salt = FS_APP_ENCRYPTION_SALT : \"2DD4E59B582AF71F\" \n\n" +
		"fs_token_salt_ext = FS_TOKEN_SALT_EXT : \"ae1iQuioQu6pooWaleez9ve1ye2ohCah2ohcoMai3xeeS5ooGhee9ohcaifare2eighohG0AiphahJ\" \n\n" +
		"Сгенерируйте пароль для APP_ADMIN_PASSWORD: \"6dbYv6qVJrqiVB\" (pwgen 10 1)\n\n" +
		"*В примерах используется редактор vim \n" +
		"*При необходимости выберите пример конфига, нажав соответствующую кнопку. \n"

	msg := tgbotapi.NewMessage(chatID, configure)
	msg.ReplyMarkup = keyboards.GetCOStandaloneConfigKeyboard()
	bot.Send(msg)
}

// SendCODeploy отправляет команду для развертывания CO.
func SendCODeploy(bot *tgbotapi.BotAPI, chatID int64) {
	deploy := "Для запуска установки CO необходимо перейти в каталог /root/install_co/ и выполнить следующую команду:\n\n" +
		"ansible-playbook playbooks/main.yml --diff\n\n" +
		"Ожидаем результат! При возниковении ошибок или вопросов обращайтесь к инженеру!\n"
	msg := tgbotapi.NewMessage(chatID, deploy)
	msg.ReplyMarkup = keyboards.GetFinishKeyboardWithIntegrationAD()
	bot.Send(msg)
}
