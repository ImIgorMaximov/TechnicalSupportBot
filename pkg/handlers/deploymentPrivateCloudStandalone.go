package handlers

import (
	"technicalSupportBot/pkg/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
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
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}

func sendPrivateKeyInsert(bot *tgbotapi.BotAPI, chatID int64) {
	privateKeyInsert := "Необходимо убедиться, что публичные ключи машин PGS и CO находятся на машине Operator в папке /root/.ssh/authorized_keys.\n" +
		"Если ключи отсутствуют, создайте пары ключей на машинах PGS и CO с помощью команды: \n\n" +
		"ssh-keygen\n\n" +
		"Затем скопируйте публичные ключи из файлов /root/.ssh/id_rsa.pub на машину Operator в папку /root/.ssh/authorized_keys:\n\n" +
		"ssh-copy-id -i /root/.ssh/id_rsa.pub root@<IP_адрес_или_домен_машины_Operator> \n"
	msg := tgbotapi.NewMessage(chatID, privateKeyInsert)
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}

func sendDNSOptions(bot *tgbotapi.BotAPI, chatID int64) {
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

func sendStandaloneDownloadDistribution(bot *tgbotapi.BotAPI, chatID int64) {
	downloadPackages := "Первая установка будет произведена на машину PGS.\n" +
		"После установки необходимых пакетов на машине operator подготовьте архив, который выдается инженером @IgorMaksimov или Аккаунт Менеджером.\n" +
		"Далее создайте директорию с помощью команды: \n" +
		"mkdir install_MyOffice_PGS\n\n" +
		"Распакуйте данный архив командой:\n" +
		"tar xf MyOffice_PGS_version.tgz -C install_MyOffice_PGS \n" +
		"*vesion - введите соответствующую версию продукта \n\n" +
		"После этого перейдите в каталог install_MyOffice_PGS: \n" +
		"cd install_MyOffice_PGS\n"
	msg := tgbotapi.NewMessage(chatID, downloadPackages)
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}

func sendCertificatesAndKeys(bot *tgbotapi.BotAPI, chatID int64) {
	certificatesAndKeys := "Для работы веб-интерфейса PGS необходима установка SSL-сертификатов.\n" +
		"Рекомендуется использовать сертификаты, полученные от публичных центров сертификации.\n" +
		"Сертификаты необходимо разместить в каталоге, соответствующему доменному имени PGS.\n\n" +
		"Напримере домена myoffice-app.ru : \n" +
		"cd /root/install_MyOffice_PGS/certificates\n" +
		"mkdir myoffice-app.ru\n\n" +
		"Вставьте серитификаты в директорию, соответствующую вашему доменному имени.\n\n Список необходимых сертификатов: \n" +
		"server.crt - содержит SSL-сертификат для *.<default_domain> и все промежуточные сертификаты, кроме корневого доверенного. \n" +
		"server.nopass.key - Приватный ключ сертификата, не требующий кодовой фразы. \n" +
		"ca.crt - файл сертификата удостоверяющего центра.\n\n" +
		"Проверить наличия сертификатов и ключа:\n" +
		"ls -la /root/install_MyOffice_PGS/certificates/myoffice-app.ru\n\n" +
		"Далее начинаем заполнять конфигурационные файлы!:)\n"
	msg := tgbotapi.NewMessage(chatID, certificatesAndKeys)
	msg.ReplyMarkup = keyboards.GetIsCertificatesKeyboard()
	bot.Send(msg)
}

func sendStandalonePGSConfigure(bot *tgbotapi.BotAPI, chatID int64) {
	pgsConfigure := "Необходимо скопировать шаблон файла inventory в корневой каталог дистрибутива и заполнить секции hosts и vars.\n\n" +
		"Операция копирования выполняется с помощью команды:\n" +
		"cp /root/install_MyOffice_PGS/inventory/hosts-sa.yaml hosts.yml \n" +
		"Далее заполним файл hosts.yml : \n" +
		"vim /root/install_MyOffice_PGS/hosts.yml\n\n" +
		"В секцию hosts добавьте доменное имя вашего PGS-сервера: \n" +
		"hosts:\n" +
		"\tpgs.myoffice-app.ru: \n" +
		"Аналогично проделать с другими сервисами: search, redis, storage, nginx, etcd...\n" +
		"Далее в секцию vars необходимо заполнить следующие переменные:\n" +
		"DEFAULT_DOMAIN: \"myoffice-app.ru\"\n" +
		"ENV: \"\" - *если используется переменная окружения\n" +
		"Сгенерируйте и внесите пароли для сервисов (команда: pwgen 13 7) : \n" +
		"KEYCLOAK_PASSWORD: \"81mToSPFJ8ezr8\"\n" +
		"KEYCLOAK_REALM_PASSWORD: \"MVh2PiA2S5cPk\"\n" +
		"KEYCLOAK_POSTGRES_PASSWORD: \"7Afd3G12P5VyUg\"\n" +
		"ARANGODB_PASSWORD: \"55ab8qk7ES4P4LX\"\n" +
		"RABBITMQ_PASSWORD: \"BdyYgDwLLY8M5U9\"\n" +
		"REDIS_PASSWORD: \"S73uo3iH3qFRdnf\"\n" +
		"GRAFANA_ADMIN_PASSWORD: \"oPpKvc6We3mES6\"\n\n" +
		"В секции co заполнить \"FS App encryption settings\" : \n" +
		"FS_APP_ENCRYPTION_SALT: \"2DD4E59B582AF71F\"\n" +
		"AUTH_ENCRYPTION_SALT: \"2DD4E59B582AF71F\"\n" +
		"APP_ADMIN_PASSWORD: \"6dbYv6qVJrqiVB\"\n\n" +
		"*В примерах используется редактор vim \n" +
		"*При необходимости выберите пример конфига, нажав соответствующую кнопку. \n"
	msg := tgbotapi.NewMessage(chatID, pgsConfigure)
	msg.ReplyMarkup = keyboards.GetPGSStandaloneConfig()
	bot.Send(msg)
}

func sendPGSDeploy(bot *tgbotapi.BotAPI, chatID int64) {
	pgsDeploy := "Для запуска установки PGS необходимо перейти в каталог /root/install_MyOffice_PGS/ и выполнить следующую команду:\n" +
		"./deploy.sh hosts.yml\n\n" +
		"Ожидаем результат! При возниковении ошибок при инсталляции обращайтесь к инженеру!\n"
	msg := tgbotapi.NewMessage(chatID, pgsDeploy)
	msg.ReplyMarkup = keyboards.GetCOInstallation()
	bot.Send(msg)
}

func sendCOInstallation(bot *tgbotapi.BotAPI, chatID int64) {
	coInstallation := "Переходим к установке и настройке CO (Сервер совместного редактирования)\n" +
		"На машину operator перенести дистрибутив CO, который выдается инженером @IgorMaksimov или Аккаунт Менеджером. \n\n" +
		"Данный дистрибутив (.iso) включает: \n" +
		"co_ansible_bin_version.run - файл с  подсистемой управления конфигурациями\n" +
		"co_infra_version.run - файл с  с хранилищем Docker-контейнеров\n\n" +
		"Далее выполните запуск скрипта с хранилищем Docker-контейнеров:\n" +
		"bash co_infra_version.run \n\n" +
		"После завершения установки необходимо убедиться, что список содержит сообщения [ OK ] или [CHANGE]\n\n" +
		"Далее выполните запуск скрипта co_ansible_bin_version.run :\n" +
		"bash co_ansible_bin_version.run \n\n" +
		"После завершения установки необходимо убедиться, что список содержит сообщения [ OK ] или [CHANGE]\n\n" +
		"Перейдите в каталог /install_co/ :\n" +
		"cd /root/install_co\n\n" +
		"Скопируйте файл /root/install_co/contrib/co/ansible.cfg в /root/install_co/ :\n" +
		"cp /root/install_co/contrib/co/ansible.cfg ansible.cfg\n\n" +
		"Скопируйте файл /root/install_co/contrib/co/standalone/hosts.yml в /root/install_co/:\n" +
		"cp /root/install_co/contrib/co/standalone/hosts.yml hosts.yml\n\n" +
		"Создайте каталог co_setup/ в директории /root/install_co/group_vars/:\n" +
		"mkdir /root/install_co/group_vars/co_setup/\n\n" +
		"Скопируйте в созданную директорию co_setup каталог с переменными для заполнения :\n" +
		"cp -r /root/install_co/contrib/co/standalone/group_vars/co_setup/* /root/install_co/group_vars/co_setup/\n\n" +
		"Далее начинаем заполнять конфигурационные файлы..\n"

	msg := tgbotapi.NewMessage(chatID, coInstallation)
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}

func sendCOConfigure(bot *tgbotapi.BotAPI, chatID int64) {
	coConfigure := "Заполним файл hosts.yml в директории /root/install_co/:\n" +
		"vim /root/install_co/hosts.yml\n\n" +
		"В секцию hosts добавьте доменное имя вашего CO-сервера: \n" +
		"hosts:\n" +
		"\tco.myoffice-app.ru: \n" +
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
		"Сравните значения переменных из конфигурационного файла hosts.yml PGS, они должны сопвадать:\n" +
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

	msg := tgbotapi.NewMessage(chatID, coConfigure)
	msg.ReplyMarkup = keyboards.GetCOStandaloneConfigKeyboard()
	bot.Send(msg)
}
