package handlers

import (
	"errors"
	"log"
	"technicalSupportBot/pkg/deployment"
	"technicalSupportBot/pkg/instructions"
	"technicalSupportBot/pkg/keyboards"
	"technicalSupportBot/pkg/sizing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// State представляет состояние пользователя
type State struct {
	Previous string
	Current  string
	Product  string
	Action   string
	Type     string
}

// StateManager управляет состояниями пользователей
type StateManager struct {
	states map[int64]*State
}

// NewStateManager создает новый StateManager
func NewStateManager() *StateManager {
	return &StateManager{states: make(map[int64]*State)}
}

// GetState возвращает текущее состояние пользователя
func (sm *StateManager) GetState(chatID int64) *State {
	state, exists := sm.states[chatID]
	if !exists {
		state = &State{}
		sm.states[chatID] = state
	}
	return state
}

// SetState устанавливает новое состояние пользователя
func (sm *StateManager) SetState(chatID int64, previous, current string) {
	state := sm.GetState(chatID)
	state.Previous = previous
	state.Current = current
}

// HandleUpdate обрабатывает входящие сообщения от пользователей
func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update, data *Data, sm *StateManager) {
	chatID := update.Message.Chat.ID
	text := update.Message.Text

	log.Printf("Получено сообщение от chatID %d: %s", chatID, text)

	state := sm.GetState(chatID)

	userID := update.Message.From.ID

	switch text {
	case "/start", "В главное меню":
		sendWelcomeMessage(bot, chatID)
		sm.SetState(chatID, state.Current, "start")

	case "Инструкции по продуктам":
		state.Action = "instr"
		sendProduct(bot, chatID)
		sm.SetState(chatID, state.Current, "instr")

	case "Развертывание продуктов":
		state.Action = "deploy"
		sendProduct(bot, chatID)
		sm.SetState(chatID, state.Current, "deploy")

	case "Расчет сайзинга продуктов":
		state.Action = "sizing"
		sendProduct(bot, chatID)
		sm.SetState(chatID, state.Current, "sizing")

	case "Частное Облако":
		handlePrivateCloud(bot, chatID, sm)

	case "Squadus":
		handleSquadus(bot, chatID, sm)

	case "Mailion":
		handleMailion(bot, chatID, sm)

	case "Почта":
		handleMail(bot, chatID, sm)

	case "Системные требования":
		handleSystemRequirements(bot, chatID, sm)

	case "Руководство по установке":
		handleInstallationGuide(bot, chatID, sm)

	case "PGS":
		instructions.SendPGSInstallationGuide(bot, chatID)
		sm.SetState(chatID, state.Current, "pgs")

	case "CO":
		instructions.SendCOInstallationGuide(bot, chatID)
		sm.SetState(chatID, state.Current, "co")

	case "Руководство по администрированию":
		handleAdminGuide(bot, chatID, sm)

	case "Назад":
		HandleBackButton(bot, chatID, sm)

	case "Связаться с инженером тех. поддержки":
		sendSupportEngineerContact(bot, chatID)

	case "Standalone":
		handleStandalone(bot, chatID, sm, text)

	case "Количество пользователей":
		msg := tgbotapi.NewMessage(chatID, "Введите максимальное количество пользователей (например, 50):")
		bot.Send(msg)

		// Задаем ожидание
		data.Set(userID, "waiting_for", "awaitingMaxUserCountPrivateCloud")

	case "Количество активных пользователей":
		msg := tgbotapi.NewMessage(chatID, "Введите количество одновременно активных пользователей (например, 200):")
		bot.Send(msg)

		data.Set(userID, "waiting_for", "awaitingActiveUserCountPrivateCloud")

	case "Количество редактируемых документов":
		msg := tgbotapi.NewMessage(chatID, "Введите количество редактируемых документов (например, 10):")
		bot.Send(msg)

		data.Set(userID, "waiting_for", "awaitingDocumentCountPrivateCloud")

	case "Дисковую квоту в хранилище":
		msg := tgbotapi.NewMessage(chatID, "Введите дисковую квоту пользователей в хранилище (ГБ) (например, 2):")
		bot.Send(msg)

		data.Set(userID, "waiting_for", "awaitingStorageQuotaPrivateCloud")

	case "Cluster":
		handleCluster(bot, chatID, sm)

	case "Проверить корректность сертификатов и ключа":
		sendIsCertificates(bot, chatID)

	case "Описание ролей":
		sendRoleDescriptionsPrivateCloudCluster2k(bot, chatID)

	case "Пример конфига PGS - hosts.yml":
		sendConfigFile(bot, chatID, "/home/admin-msk/MyOfficeConfig/hostsPGS.yml", "hostsPGS.yml")

	case "Пример конфига PSN - hosts.yml":
		sendConfigFile(bot, chatID, "/home/admin-msk/MyOfficeConfig/hostsPSN.yml", "hostsPSN.yml")

	case "Пример конфига CO - main.yml":
		sendConfigFile(bot, chatID, "/home/admin-msk/MyOfficeConfig/mainCO.yml", "mainCO.yml")

	case "Пример конфига CO - hosts.yml":
		sendConfigFile(bot, chatID, "/home/admin-msk/MyOfficeConfig/hostsCO.yml", "hostsCO.yml")

	case "Пример конфига Squadus - hosts.yml":
		sendConfigFile(bot, chatID, "/home/admin-msk/MyOfficeConfig/hostsSquadus.yml", "hostsSquadus.yml")

	case "Пример конфига Squadus - main.yml":
		sendConfigFile(bot, chatID, "/home/admin-msk/MyOfficeConfig/hostsSquadus.yml", "mainSquadus.yml")

	case "Пример конфига Mailion - hosts.yml":
		sendConfigFile(bot, chatID, "/home/admin-msk/MyOfficeConfig/hostsMailion.yml", "hostsMailion.yml")

	case "Пример конфига Mailion - main.yml":
		sendConfigFile(bot, chatID, "/home/admin-msk/MyOfficeConfig/mainMailion.yml", "mainMailion.yml")

	case "Далее", "Установка CO", "Готово", "Запустить деплой":
		HandleNextStep(bot, chatID, sm)

	case "Распаковка ISO образа":
		sendUnzippingISO(bot, chatID)

	default:

		err := handleSettings(bot, sm, data, update)
		if err != nil {
			log.Println("handle settings:", err)

			sendWelcomeMessage(bot, chatID)
			sm.SetState(chatID, state.Current, "start")
			return
		}

	}
}

// handleStandalone обрабатывает запрос на Standalone
func handleStandalone(bot *tgbotapi.BotAPI, chatID int64, sm *StateManager, text string) {
	state := sm.GetState(chatID)
	log.Printf("handleStandalone: chatID %d, previousState %s, currentState %s", chatID, state.Previous, state.Current)

	// Обработка для перехода от состояния privateCloud
	if state.Product == "privateCloud" {
		if state.Action == "sizing" {
			sm.SetState(chatID, state.Current, "standalone")
			state.Type = "standalone"
			log.Printf("Текущее состояние: %s, Предыдущее состояние: %s, Действие: %s", state.Current, state.Previous, state.Action)

			sendStandaloneSettings(bot, chatID)
			// sizing.HandleUserInput(bot, chatID, &state.Current, nil)
			sm.SetState(chatID, state.Current, "settings")

			log.Printf("После вызова HandleUserInput. Текущее состояние: %s, Предыдущее состояние: %s.", state.Current, state.Previous)
		} else if state.Action == "deploy" {
			sm.SetState(chatID, state.Current, "standalone")
			state.Type = "standalone"
			log.Printf("Текущее состояние: %s, Предыдущее состояние: %s, Действие: %s", state.Current, state.Previous, state.Action)
			deployment.SendStandaloneRequirementsPrivateCloud(bot, chatID)
			sm.SetState(chatID, state.Current, "reqPrivateCloud")
			log.Printf("После вызова SendStandaloneRequirementsPrivateCloud. Текущее состояние: %s, Предыдущее состояние: %s.", state.Current, state.Previous)
		}
	} else if state.Product == "mail" {
		if state.Action == "sizing" {
			sm.SetState(chatID, state.Current, "standalone")
			state.Type = "standalone"
			log.Printf("Текущее состояние: %s, Предыдущее состояние: %s.", state.Current, state.Previous)
			sizing.HandleSizingMailStandalone(bot, chatID)
		} else if state.Action == "deploy" {
			sm.SetState(chatID, state.Current, "standalone")
			state.Type = "standalone"
			log.Printf("Текущее состояние: %s, Предыдущее состояние: %s. Отправка пакетов для самостоятельной загрузки.", state.Current, state.Previous)
			deployment.SendStandaloneRequirementsPSN(bot, chatID)
			sm.SetState(chatID, state.Current, "reqPsn")
			log.Printf("После вызова SendStandaloneRequirementsPSN. Текущее состояние: %s, Предыдущее состояние: %s.", state.Current, state.Previous)
		}
	} else if state.Product == "squadus" {
		if state.Action == "deploy" {
			sm.SetState(chatID, state.Current, "standalone")
			state.Type = "standalone"
			log.Printf("Текущее состояние: %s, Предыдущее состояние: %s. Отправка пакетов для самостоятельной загрузки.", state.Current, state.Previous)
			deployment.SendStandaloneRequirementsSquadus(bot, chatID)
			sm.SetState(chatID, state.Current, "reqSquadus")
			log.Printf("После вызова SendStandaloneRequirementsSquadus. Текущее состояние: %s, Предыдущее состояние: %s.", state.Current, state.Previous)
		}
	} else if state.Product == "mailion" {
		if state.Action == "deploy" {
			sm.SetState(chatID, state.Current, "standalone")
			state.Type = "standalone"
			log.Printf("Текущее состояние: %s, Предыдущее состояние: %s. Отправка пакетов для самостоятельной загрузки.", state.Current, state.Previous)
			deployment.SendStandaloneRequirementsMailion(bot, chatID)
			sm.SetState(chatID, state.Current, "reqMailion")
			log.Printf("После вызова SendStandaloneRequirementsMailion. Текущее состояние: %s, Предыдущее состояние: %s.", state.Current, state.Previous)
		}
	}
}

// handleCluster обрабатывает запрос на Cluster
func handleCluster(bot *tgbotapi.BotAPI, chatID int64, sm *StateManager) {
	state := sm.GetState(chatID)
	log.Printf("handleCluster: chatID %d, previousState %s, currentState %s", chatID, state.Previous, state.Current)

	msg := tgbotapi.NewMessage(chatID, "Извините, раздел находится в разработке😢")
	bot.Send(msg)
}

// handlePrivateCloud обрабатывает запрос на Частное Облако
func handlePrivateCloud(bot *tgbotapi.BotAPI, chatID int64, sm *StateManager) {
	state := sm.GetState(chatID)
	state.Product = "privateCloud"
	log.Printf("handlePrivateCloud: chatID %d, previousState %s, currentState %s, productState %s", chatID, state.Previous, state.Current, state.Product)

	if state.Action == "instr" {
		sendInstructions(bot, chatID)
		sm.SetState(chatID, state.Current, "privateCloud")
		log.Printf("Переключение состояния на privateCloud после инструкции: chatID %d, previousState %s, currentState %s", chatID, state.Previous, state.Current)
	} else if state.Action == "deploy" || state.Current == "sizing" {
		sendDeploymentOptions(bot, chatID)
		sm.SetState(chatID, state.Current, "privateCloud")
		log.Printf("Переключение состояния на privateCloud после выбора развертывания или сайзинга: chatID %d, previousState %s, currentState %s", chatID, state.Previous, state.Current)
	}
}

// handleMail обрабатывает запрос на Почту
func handleMail(bot *tgbotapi.BotAPI, chatID int64, sm *StateManager) {
	state := sm.GetState(chatID)
	state.Product = "mail"
	log.Printf("handleMail: chatID %d, previousState %s, currentState %s, productState %s", chatID, state.Previous, state.Current, state.Product)

	if state.Action == "instr" {
		sendInstructions(bot, chatID)
		sm.SetState(chatID, state.Current, "mail")
		log.Printf("Переключение состояния на mail после инструкции: chatID %d, previousState %s, currentState %s", chatID, state.Previous, state.Current)
	} else if state.Action == "deploy" || state.Current == "sizing" {
		sendDeploymentOptions(bot, chatID)
		sm.SetState(chatID, state.Current, "mail")
		log.Printf("Переключение состояния на mail после выбора развертывания или сайзинга: chatID %d, previousState %s, currentState %s", chatID, state.Previous, state.Current)
	}
}

// handleMailion обрабатывает запрос на Mailion
func handleMailion(bot *tgbotapi.BotAPI, chatID int64, sm *StateManager) {
	state := sm.GetState(chatID)
	state.Product = "mailion"
	log.Printf("handleMailion: chatID %d, previousState %s, currentState %s, productState %s", chatID, state.Previous, state.Current, state.Product)
	if state.Current == "instr" {
		sendInstructions(bot, chatID)
		sm.SetState(chatID, state.Current, "mailion")
		log.Printf("Переключение состояния на mailion после инструкции: chatID %d, previousState %s, currentState %s", chatID, state.Previous, state.Current)
	} else if state.Action == "deploy" || state.Current == "sizing" {
		sendDeploymentOptions(bot, chatID)
		sm.SetState(chatID, state.Current, "mailion")
		log.Printf("Переключение состояния на mailion после выбора развертывания или сайзинга: chatID %d, previousState %s, currentState %s", chatID, state.Previous, state.Current)
	}
}

// handleSquadus обрабатывает запрос на Squadus
func handleSquadus(bot *tgbotapi.BotAPI, chatID int64, sm *StateManager) {
	state := sm.GetState(chatID)
	state.Product = "squadus"
	log.Printf("handleSquadus: chatID %d, previousState %s, currentState %s, productState %s", chatID, state.Previous, state.Current, state.Product)
	if state.Current == "instr" {
		sendInstructions(bot, chatID)
		sm.SetState(chatID, state.Current, "squadus")
		log.Printf("Переключение состояния на squadus после инструкции: chatID %d, previousState %s, currentState %s", chatID, state.Previous, state.Current)
	} else if state.Action == "deploy" || state.Current == "sizing" {
		sendDeploymentOptions(bot, chatID)
		sm.SetState(chatID, state.Current, "squadus")
		log.Printf("Переключение состояния на squadus после выбора развертывания или сайзинга: chatID %d, previousState %s, currentState %s", chatID, state.Previous, state.Current)
	}
}

func handlePrivateKeyInsert(bot *tgbotapi.BotAPI, chatID int64, sm *StateManager) {
	state := sm.GetState(chatID)

	if state.Product == "privateCloud" {
		deployment.SendPrivateKeyInsertPrivateCloud(bot, chatID)
		sm.SetState(chatID, state.Current, "privateKeyInsertPrivateCloud")
	} else if state.Product == "mail" {
		deployment.SendPrivateKeyInsertPSN(bot, chatID)
		sm.SetState(chatID, state.Current, "privateKeyInsertPSN")
	} else if state.Product == "squadus" {
		deployment.SendPrivateKeyInsertSquadus(bot, chatID)
		sm.SetState(chatID, state.Current, "privateKeyInsertSquadus")
	} else if state.Product == "mailion" {
		deployment.SendPrivateKeyInsertMailion(bot, chatID)
		sm.SetState(chatID, state.Current, "privateKeyInsertMailion")
	}
}

// handleSystemRequirements обрабатывает запрос на системные требования
func handleSystemRequirements(bot *tgbotapi.BotAPI, chatID int64, sm *StateManager) {
	state := sm.GetState(chatID)
	log.Printf("handleSystemRequirements: chatID %d, previousState %s, currentState %s", chatID, state.Previous, state.Current)

	if state.Current == "privateCloud" {
		instructions.SendSystemRequirementsPivateCloud(bot, chatID)
		state.Current = "requirementsPrivateCloud"
	} else if state.Current == "squadus" {
		instructions.SendSystemRequirementsSquadus(bot, chatID)
		state.Current = "requirementsSquadus"
	} else if state.Current == "mailion" {
		instructions.SendSystemRequirementsMailion(bot, chatID)
		state.Current = "requirementsMailion"
	} else if state.Current == "mail" {
		instructions.SendSystemRequirementsMail(bot, chatID)
		state.Current = "requirementsMail"
	}
}

// handleInstallationGuide обрабатывает запрос на руководство по установке
func handleInstallationGuide(bot *tgbotapi.BotAPI, chatID int64, sm *StateManager) {
	state := sm.GetState(chatID)
	log.Printf("handleInstallationGuide: chatID %d, previousState %s, currentState %s", chatID, state.Previous, state.Current)

	if state.Current == "privateCloud" {
		instructions.SendInstallationGuideOptionsPrivateCloud(bot, chatID)
		state.Current = "installationGuidePrivateCloud"
	} else if state.Current == "squadus" {
		instructions.SendInstallationGuideSquadus(bot, chatID)
		state.Current = "installationGuideSquadus"
	} else if state.Current == "mailion" {
		instructions.SendInstallationGuideMailion(bot, chatID)
		state.Current = "installationGuideMailion"
	} else if state.Current == "mail" {
		instructions.SendInstallationGuideMail(bot, chatID)
		state.Current = "installationGuideMail"
	}
}

// handleAdminGuide обрабатывает запрос на руководство по администрированию
func handleAdminGuide(bot *tgbotapi.BotAPI, chatID int64, sm *StateManager) {
	state := sm.GetState(chatID)
	log.Printf("handleAdminGuide: chatID %d, previousState %s, currentState %s", chatID, state.Previous, state.Current)

	if state.Current == "privateCloud" {
		instructions.SendAdminGuidePrivateCloud(bot, chatID)
		state.Current = "adminGuidePrivateCloud"
	} else if state.Current == "squadus" {
		instructions.SendAdminGuideSquadus(bot, chatID)
		state.Current = "adminGuideSquadus"
	} else if state.Current == "mailion" {
		instructions.SendAdminGuideMailion(bot, chatID)
		state.Current = "adminGuideMailion"
	} else if state.Current == "mail" {
		instructions.SendAdminGuideMail(bot, chatID)
		state.Current = "adminGuideMail"
	}
}

func handleSettings(bot *tgbotapi.BotAPI, sm *StateManager, data *Data,
	update tgbotapi.Update) error {

	state := sm.GetState(update.Message.Chat.ID)
	if state.Type != "standalone" {
		return errors.New("")
	}

	userID := update.Message.From.ID

	// получение ожидаемых параметров
	val, ok := data.Get(userID, "waiting_for")
	if ok {
		if val == "awaitingMaxUserCountPrivateCloud" {
			state.Current = val
			sm.SetState(update.Message.Chat.ID, state.Current, val)
			// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Максимальное количество пользователей сохранено! Теперь введите количество активных пользователей.")
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Принятые параметры:\n"+
				"✔️ Количество пользователей \n"+
				"- Количество одновременно активных пользователей \n"+
				"- Количество редактируемых документов \n"+
				"- Дисковую квоту пользователей в хранилище")

			data.Set(userID, "MaxUserCountPrivateCloud", update.Message.Text)
			data.Set(userID, "waiting_for", "")

			msg.ReplyMarkup = keyboards.GetStandalonePrivateCloudKeyboard()
			bot.Send(msg)

		} else if val == "awaitingActiveUserCountPrivateCloud" {
			sm.SetState(update.Message.Chat.ID, state.Current, val)
			// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Максимальное количество пользователей сохранено! Теперь введите количество активных пользователей.")
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Принятые параметры:\n"+
				"✔️ Количество пользователей \n"+
				"✔️ Количество одновременно активных пользователей \n"+
				"- Количество редактируемых документов \n"+
				"- Дисковую квоту пользователей в хранилище")

			data.Set(userID, "ActiveUserCountPrivateCloud", update.Message.Text)
			data.Set(userID, "waiting_for", "")
			msg.ReplyMarkup = keyboards.GetStandalonePrivateCloudKeyboard()
			bot.Send(msg)
		} else if val == "awaitingDocumentCountPrivateCloud" {
			sm.SetState(update.Message.Chat.ID, state.Current, val)
			// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Kоличество активных пользователей сохранено! Теперь введите количество редактируемых документов")
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Принятые параметры:\n"+
				"✔️ Количество пользователей \n"+
				"✔️ Количество одновременно активных пользователей \n"+
				"✔️ Количество редактируемых документов \n"+
				"- Дисковую квоту пользователей в хранилище")

			data.Set(userID, "DocumentCountPrivateCloud", update.Message.Text)
			data.Set(userID, "waiting_for", "")

			msg.ReplyMarkup = keyboards.GetStandalonePrivateCloudKeyboard()
			bot.Send(msg)
		} else if val == "awaitingStorageQuotaPrivateCloud" {
			sm.SetState(update.Message.Chat.ID, state.Current, val)
			// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Дисковая квота в хранилище сохранено! Теперь введите количество редактируемых документов")
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Принятые параметры:\n"+
				"✔️ Количество пользователей \n"+
				"✔️ Количество одновременно активных пользователей \n"+
				"✔️ Количество редактируемых документов \n"+
				"✔️ Дисковую квоту пользователей в хранилище")

			data.Set(userID, "StorageQuotaPrivateCloud", update.Message.Text)
			data.Set(userID, "waiting_for", "")

			// msg.ReplyMarkup = keyboards.GetStandalonePrivateCloudKeyboard()
			bot.Send(msg)

			state := sm.GetState(update.Message.Chat.ID)
			log.Println(data.UserData[userID])
			sizing.HandleUserInput(bot, update.Message.Chat.ID, &state.Current, data.UserData[userID])
		}
	} else {
		return errors.New("not exist")
	}

	return nil
}
