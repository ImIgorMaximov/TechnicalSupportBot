/*
Package handlers предоставляет функции для обработки различных команд и кнопок, используемых в техническом боте поддержки.

Функции пакета предназначены для взаимодействия с пользователями в Telegram, помогая им выполнять задачи, связанные с установкой и поддержкой продуктов.
Каждая функция предоставляет пользователям информацию о процессе установки, настройке окружения и управлении сертификатами, а также дает возможность напрямую связаться с инженерами поддержки для решения более сложных вопросов.

Основные функции пакета включают:
- Отправка приветственного сообщения и навигации по меню.
- Предоставление инструкций по установке и настройке продуктов.
- Консультации по проверке сертификатов и конфигурационным файлам.
- Поддержка пользователей при возникновении ошибок и нестандартных ситуаций.

Этот пакет разработан с учетом удобства использования, чтобы пользователи могли быстро получать нужную информацию, оставаться в курсе процессов развертывания и иметь легкий доступ к поддержке.

Автор: Максимов Игорь
Email: imigormaximov@gmail.com
*/

package handlers

import (
	"fmt"
	"log"
	"os"
	"strings"
	"technicalSupportBot/pkg/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// sendWelcomeMessage отправляет приветственное сообщение с основным меню.
func sendWelcomeMessage(bot *tgbotapi.BotAPI, chatID int64) {

	welcomeMessage := "Добро пожаловать в чат бот тех. поддержки МойОфис! :) " +
		"Выберите необходимую функцию:\n" +
		"1. Инструкции по продуктам.\n" +
		"2. Развертывание продуктов. \n" +
		"3. Расчет сайзинга продуктов. \n" +
		"4. Связаться с инженером тех. поддержки.\n"
	msg := tgbotapi.NewMessage(chatID, welcomeMessage)
	msg.ReplyMarkup = keyboards.GetMainKeyboard()
	bot.Send(msg)
}

// sendProduct отправляет сообщение с выбором продукта.
func sendProduct(bot *tgbotapi.BotAPI, chatID int64) {

	chooseProductMessage := "Выберите продукт:"
	msg := tgbotapi.NewMessage(chatID, chooseProductMessage)
	msg.ReplyMarkup = keyboards.GetProductKeyboard()
	bot.Send(msg)
}

// sendDeploymentOptions отправляет сообщение с выбором типа установки.
func sendDeploymentOptions(bot *tgbotapi.BotAPI, chatID int64) {

	deploymentMessage := "Выберите тип инсталляции:"
	msg := tgbotapi.NewMessage(chatID, deploymentMessage)
	msg.ReplyMarkup = keyboards.GetDeploymentOptionsKeyboard()
	bot.Send(msg)
}

// sendInstructions отправляет сообщение с подсказками по инструкциям.
func sendInstructions(bot *tgbotapi.BotAPI, chatID int64) {

	chooseFunction := "Что подсказать? \n" +
		"- Cистемные требования \n" +
		"- Руководство по установке \n" +
		"- Руководство по администрированию \n"
	msg := tgbotapi.NewMessage(chatID, chooseFunction)
	msg.ReplyMarkup = keyboards.GetInstructionsKeyboard()
	bot.Send(msg)
}

// sendIsCertificates отправляет инструкции по проверке сертификатов.
func sendIsCertificates(bot *tgbotapi.BotAPI, chatID int64) {
	isCertificates := "Проверка сертификата сервера (server.crt): \n" +
		"openssl x509 -in server.crt -text -noout \n" +
		"Проверка цепочки сертификатов (ca.crt): \n" +
		"openssl x509 -in ca.crt -text -noout \n" +
		"Проверка приватного ключа (server.nopass.key): \n" +
		"openssl rsa -in server.nopass.key -check \n"
	msg := tgbotapi.NewMessage(chatID, isCertificates)
	bot.Send(msg)
}

// sendRoleDescriptionsPrivateCloudCluster2k отправляет описание ролей для кластеров.
func sendRoleDescriptionsPrivateCloudCluster2k(bot *tgbotapi.BotAPI, chatID int64) {
	isCertificates := "Описание ролей: \n" +
		"Operator - сервер, с которого производится установка всех компонентов;\n\n" +
		"LB - сервер балансировки нагрузки для всех компонентов; \n\n" +
		"core - в составе: управление редактированием, коллаборации и документного API, балансировка нагрузки компонента CO, " +
		"управления импортом, экспортом и индексированием документов, подсистема сервиса файлового API, подсистема сервиса push-нотификаций; \n\n" +
		"infra - сервер, объединящий инфраструктурные роли сбор логов и мониторинга компонента CO  \n\n" +
		"mq - сервер очереди сообщений и подписок;\n\n" +
		"imc - сервер кэширования сессий и хранения промежуточных результатов в памяти;\n\n" +
		"etcd - Подсистема конфигурации продукта при помощи Etcd;\n\n" +
		"PGS-APP - cервер вычислений, обработки запросов, API, keycloak, внутренний балансировщик;\n\n" +
		"STORAGE - Сервер хранения данных (GlusterFS);\n\n" +
		"STORAGE-A - арбитр серверов хранения данных;\n\n" +
		"PGS-BE - сервер индексного поиска, хранения кэша, конфигурации, очередей (elasticsearch, etcd, arangodb_agent, redis, rabbitmq);\n\n" +
		"PGS-DB - cервер БД пользователей, метаданных, авторизации (PostgreSQL, ArangoDB);\n\n" +
		"PGS-LOG - cервер сбора логов компонента PGS (registry, syslog. роли мониторинга);\n\n"
	msg := tgbotapi.NewMessage(chatID, isCertificates)
	bot.Send(msg)
}

// sendUnzippingISO отправляет инструкции по разархивированию образа ISO.
func sendUnzippingISO(bot *tgbotapi.BotAPI, chatID int64) {
	unzippingISO := "Сопоставьте контрольную сумму с сайта и скаченного образа .iso для MD5: \n" +
		"md5sum /path/to/file.iso \n" +
		"Создайте директорию, например /mnt/iso\n" +
		"mkdir /mnt/iso\n" +
		"Смонтируйте образ : \n" +
		"mount -o loop /path/to/file.iso /mnt/iso \n\n" +
		"*Для разархивирования образа .iso также можно воспользоваться инструментом \"bsdtar\": \n" +
		"apt-get install bsdtar \n" +
		"bsdtar -xvf путь_к_файлу.iso -C директория_для_извлечения \n"
	msg := tgbotapi.NewMessage(chatID, unzippingISO)
	bot.Send(msg)
}

// sendSupportEngineerContact отправляет контактную информацию для связи с инженером поддержки.
func sendSupportEngineerContact(bot *tgbotapi.BotAPI, chatID int64) {
	errorMessage := "Направьте описание проблемы или вопроса инженеру \nТГ: @IgorMaksimov2000\nПочта: igor.maksimov@myoffice.team \n\n" +
		"Формат сообщения должен включать: \n" +
		"1. Описание ошибки/вопроса c приложением скриншотов.\n" +
		"2. Вывод команд pip3 list и ansible --version. (Выполненные в корневой директории инсталляции, например, /root/install_pgs) \n" +
		"3. Конфигурационные файлы, которые были использованы при инсталляции. (Например, hosts.yml для PGS, hosts.yml/main.yml для CO) \n" +
		"4. Логи ошибок (Например, при развертывания СО сервера это будет файл deploy_co.log). \n\n" +
		"Спасибо!\n"
	msg := tgbotapi.NewMessage(chatID, errorMessage)
	bot.Send(msg)
}

// sendStandaloneDownloadPackages отправляет инструкции для установки необходимых пакетов на отдельной машине.
func sendStandaloneDownloadPackages(bot *tgbotapi.BotAPI, chatID int64) {
	standaloneDownloadPackages := "Вся установка и настройка будет производиться на машине operator на примере системы Astra Linux Special Edition 1.7 «Орел» (базовый);\n" +
		"На ВМ c ролью operator обновите систему: \n" +
		"sudo su\n" +
		"apt update\n\n" +
		"Далее установим необходимые пакеты: \n" +
		"apt install -y python3-pip \n" +
		"python3 -m pip install ansible-core==2.11.12 \n" +
		"python3 -m pip install ansible==4.9.0 \n" +
		"python3 -m pip install jinja2==3.1.2 \n" +
		"python3 -m pip install yamllint \n\n" +
		"На этом все :) Двигаемся дальше..\n"
	msg := tgbotapi.NewMessage(chatID, standaloneDownloadPackages)
	msg.ReplyMarkup = keyboards.GetStandaloneNextStepKeyboard()
	bot.Send(msg)
}

// sendStandaloneDownloadPackages отправляет инструкции для установки необходимых пакетов на отдельной машине.
func sendIntegrationAD(bot *tgbotapi.BotAPI, chatID int64) {
	integrationAD := "После успешной установки PGS для настройки интеграции AD/aldPro необходимо произвести следующие действия: \n\n" +
		"1. Открыть доступ к компоненту Keycloak из внешней сети: \n\n" +
		"docker service update --publish-add published=8091,target=8080 pgs-keycloak_keycloak\n\n" +
		"2. Перезапустить сервисы pgs_aristoteles и pgs_euclid: \n\n" +
		"docker restart pgs_aristoteles pgs_euclid \n\n" +
		"3. Открыть веб-интерфейс Keycloak: \n" +
		"(адрес по умолчанию http://<ENV>.<DEFAULT_DOMAIN>:8091/auth) \n" +
		"4. Выбрать тенант (или realm), для которого нужна интеграция. \n" +
		"5. Нажать User Federation. \n" +
		"6. Из выпадающего меню выбрать провайдера LDAP (Add provider) с именем pgsldapnew. \n" +
		"7. Заполнить атрибуты. Пример заполнения указан в скриншоте. \n" +
		"8. Нажать Save и Synchronize all users. \n" +
		"9. Проверить отображение пользователей в админ панели. \n\n" +
		"** В aldPro вместо sAMAccountName используется person. \n" +
		"Для проверки корректности указанных фильтров выполните команду на сервере: \n\n" +
		"curl -u \"bind_user@domain:bind_password\" 'ldap://dc.domain.tld/OU=Пользователи,DC=domain,dc=tld??sub?(mail=user@domain.tld) \n"
	// Отправка сообщения с текстом
	msg := tgbotapi.NewMessage(chatID, integrationAD)
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Не удалось отправить сообщение: %v", err)
	}

	// Открытие файла с фото
	file, err := os.Open("/home/admin-msk/MyOfficeConfig/integrationAD.png")
	if err != nil {
		log.Printf("Не удалось открыть файл с фото: %v", err)
		return
	}
	defer file.Close()

	// Создание объекта для загрузки фото
	photo := tgbotapi.NewPhoto(chatID, tgbotapi.FileReader{
		Name:   "integrationAD.png", // имя файла, как будет показано в чате
		Reader: file,
	})

	// Отправка фото
	_, err = bot.Send(photo)
	if err != nil {
		log.Printf("Не удалось отправить фото: %v", err)
	}
}

// sendUnknownCommandMessage отправляет сообщение в случае ввода неизвестной команды.
func sendUnknownCommandMessage(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Неизвестная команда. Пожалуйста, выберите действие из меню.")
	bot.Send(msg)
}

// sendConfigFile отправляет конфигурационный файл пользователю с сервера.
func sendConfigFile(bot *tgbotapi.BotAPI, chatID int64, filePath, fileName string) {
	// Проверяем, существует ли файл
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		msg := tgbotapi.NewMessage(chatID, "Файл не найден.")
		if _, err := bot.Send(msg); err != nil {
			log.Println("Error sending message:", err)
		}
		return
	}

	// Открываем файл для отправки
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error opening file:", err)
		msg := tgbotapi.NewMessage(chatID, "Не удалось открыть файл.")
		if _, err := bot.Send(msg); err != nil {
			log.Println("Error sending message:", err)
		}
		return
	}
	defer file.Close()

	// Отправляем файл пользователю
	document := tgbotapi.NewDocument(chatID, tgbotapi.FileReader{
		Name:   fileName,
		Reader: file,
	})
	if _, err := bot.Send(document); err != nil {
		log.Println("Error sending document:", err)
	}
}

// formatSizingResults форматирует результаты сайзинга в текстовый вид для отправки пользователю.
func formatSizingResults(results map[string]map[string]string) string {
	var sb strings.Builder
	for component, data := range results {
		sb.WriteString(fmt.Sprintf("%s:\n", component))
		for key, value := range data {
			sb.WriteString(fmt.Sprintf("%s: %s\n", key, value))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
