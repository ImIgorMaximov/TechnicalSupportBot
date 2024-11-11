# TechnicalSupportBot

Telegram-бот технической поддержки МойОфис, предназначенный для предоставления пользователям доступа к инструкциям, расчету сайзинга и помощи по установке продуктов, таких как Частное Облако, Mailion, Squadus, и Почта.

Бот позволяет пользователям:

    Получать инструкции по установке, администрированию и системным требованиям для продуктов.
    Помогать пошагово установить соответствующие продукты в типе Standalone инсталляции.
    Рассчитывать и отправлять файл с расчетом сайзинга для соответствующих продуктов.

# Структура проекта

Проект имеет следующую структуру:

TechnicalSupportBot/
│
├── cmd/
│   └── main.go
│
├── docs/
│   └── README.md
│
├── pkg/
│   ├── deployment/
│   │   ├── deploymentMailionStandalone.go
│   │   ├── deploymentMailStandalone.go
│   │   ├── deploymentPrivateCloudStandalone.go
│   │   └── deploymentSquadusStandalone.go
│   │
│   ├── handlers/
│   │   ├── back.go
│   │   ├── handlers.go
│   │   ├── mainMessage.go
│   │   └── next.go
│   │
│   ├── instructions/
│   │   ├── mailionInstructions.go
│   │   ├── mailInstructions.go
│   │   ├── privateCloudInstructions.go
│   │   └── squadusInstructions.go
│   │
│   ├── keyboards/
│   │   └── keyboards.go
│   │
│   ├── sizing/
│   │   ├── sizingMailion.go
│   │   ├── sizingPrivateCloudStandalone.go
│   │   ├── sizingPSNStandalone.go
│   │   └── sizingSquadus.go
│
├── go.mod
└── go.sum

# Описание пакетов
cmd/main.go

Главный файл проекта, в котором инициализируется и запускается Telegram-бот, обрабатывающий входящие запросы от пользователей.
pkg/deployment/

Этот пакет содержит функции, связанные с развертыванием различных продуктов в типе Standalone инсталляции. Он включает в себя следующие файлы:

    deploymentMailionStandalone.go — развертывание для Mailion.
    deploymentMailStandalone.go — развертывание для Почты.
    deploymentPrivateCloudStandalone.go — развертывание для Частного Облака.
    deploymentSquadusStandalone.go — развертывание для Squadus.

pkg/handlers/

Пакет для обработки команд и сообщений от пользователей. Включает следующие файлы:

    back.go — обработка команды "Назад".
    handlers.go — основной файл для обработки различных команд и состояний.
    mainMessage.go — отправка главных сообщений и меню пользователям.
    next.go — переход между состояниями и шагами в процессе установки или расчета.

pkg/instructions/

Функции для отправки ссылок на различные инструкции и руководства через Telegram-бота. Эти файлы предоставляют инструкции для каждого из продуктов:

    mailionInstructions.go — инструкция для Mailion.
    mailInstructions.go — инструкция для Почты.
    privateCloudInstructions.go — инструкция для Частного Облака.
    squadusInstructions.go — инструкция для Squadus.

pkg/keyboards/

Этот пакет управляет клавишами и клавиатурами, используемыми в Telegram-боте для взаимодействия с пользователем. Включает файл:

    keyboards.go — создание и настройка клавиатур для бота.

pkg/sizing/

Пакет для расчета сайзинга в зависимости от введенных пользователем данных. Включает следующие файлы:

    sizingMailion.go — расчет сайзинга для Mailion.
    sizingPrivateCloudStandalone.go — расчет сайзинга для Частного Облака.
    sizingPSNStandalone.go — расчет сайзинга для PSN.
    sizingSquadus.go — расчет сайзинга для Squadus.

# Установка

Для запуска бота вам потребуется Go (версии 1.18 и выше).

Клонируйте репозиторий:

    git clone https://github.com/ImIgorMaximov/TechnicalSupportBot

Перейдите в директорию проекта:

    cd TechnicalSupportBot

Установите зависимости:

    go mod tidy

Установите переменную окружения с токеном бота:

    export TELEGRAM_BOT_TOKEN="your-telegram-bot-token"

Запустите бота:

    go run cmd/main.go

# Использование

После запуска бота, пользователи смогут:

    1. Получать инструкции по установке и администрированию продуктов через Telegram.
    2. Процесс установки будет пошаговым, с пояснениями и подсказками на каждом этапе.
    3. Бот также предоставит возможность рассчитать сайзинг для продуктов на основе введенных данных и отправит соответствующий PDF файл.
    4. Клавиатуры в Telegram будут помогать пользователю навигировать между этапами и выбирать нужные опции.
