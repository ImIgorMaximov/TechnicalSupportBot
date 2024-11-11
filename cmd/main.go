/*
Package main

Пакет main реализует основной функционал запуска Telegram-бота для технической поддержки,
который обрабатывает входящие обновления от пользователей и управляет состояниями их сессий.

Основные функции:
- Создание и настройка экземпляра бота Telegram с использованием токена из переменной окружения.
- Настройка и управление состояниями пользователей с использованием StateManager, предоставляемого пакетом `handlers`.
- Асинхронная обработка сообщений и callback-запросов от пользователей с помощью горутин для оптимизации производительности.

Каждое обновление обрабатывается в отдельной горутине, что позволяет боту одновременно обрабатывать несколько сообщений и команд.

Автор: Максимов Игорь
Email: imigormaximov@gmail.com
*/

package main

import (
	"log"
	"os"
	"sync"
	"technicalSupportBot/pkg/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Получаем токен бота из переменной окружения
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("Переменная окружения 'TELEGRAM_BOT_TOKEN' не объявлена. Добавьте токен бота в переменные окружения.")
	}

	// Создаём экземпляр бота
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Ошибка при создании бота: %v", err)
	}

	// Включаем логирование (для отладки можно включить режим Debug)
	bot.Debug = true // Установите false для продакшена

	// Создаём StateManager для управления состояниями пользователей
	sm := handlers.NewStateManager()

	// Задаём параметры получения обновлений (время ожидания 60 секунд)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Получаем канал обновлений от Telegram
	updates := bot.GetUpdatesChan(u)

	// Используем WaitGroup для синхронизации горутин
	var wg = &sync.WaitGroup{}

	// Основной цикл обработки обновлений
	for update := range updates {
		// Игнорируем обновления без сообщений и callback-запросов
		if update.Message == nil && update.CallbackQuery == nil {
			continue
		}

		// Увеличиваем счётчик горутин в WaitGroup
		wg.Add(1)

		// Обрабатываем обновления в новой горутине
		go func(update tgbotapi.Update) {
			defer wg.Done() // Уменьшаем счётчик после завершения работы горутины

			// Обрабатываем обновление (HandleUpdate ничего не возвращает, поэтому не ожидаем значений)
			handlers.HandleUpdate(bot, update, sm)
		}(update)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
}
