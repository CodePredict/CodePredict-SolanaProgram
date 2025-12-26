# Solana Program для Polymarket

Прогнозный рынок на блокчейне Solana, реализованный с использованием чистой архитектуры на Go.

## Архитектура

Проект следует принципам чистой архитектуры (Clean Architecture) и разделен на следующие слои:

### Domain Layer (`internal/domain/`)
- **entities/**: Бизнес-сущности (Market, Position, Account)
- **repositories/**: Интерфейсы репозиториев
- **services/**: Интерфейсы бизнес-логики

### Application Layer (`internal/application/`)
- **usecases/**: Сценарии использования (CreateMarket, ResolveMarket, CreatePosition)

### Infrastructure Layer (`internal/infrastructure/`)
- **solana/**: Реализация работы с Solana (Program, AccountManager)
- **repositories/**: Реализация репозиториев для Solana
- **services/**: Реализация сервисов

### Presentation Layer (`internal/presentation/`)
- **instructions/**: Обработчики Solana инструкций

### Shared (`pkg/`)
- **errors/**: Общие типы ошибок
- **utils/**: Утилиты

## Структура проекта

```
.
├── cmd/
│   └── program/
│       └── main.go              # Точка входа программы
├── internal/
│   ├── domain/                  # Доменный слой
│   │   ├── entities/            # Сущности
│   │   ├── repositories/        # Интерфейсы репозиториев
│   │   └── services/            # Интерфейсы сервисов
│   ├── application/             # Слой приложения
│   │   └── usecases/            # Сценарии использования
│   ├── infrastructure/          # Слой инфраструктуры
│   │   ├── solana/              # Solana интеграция
│   │   ├── repositories/        # Реализация репозиториев
│   │   └── services/            # Реализация сервисов
│   └── presentation/            # Слой представления
│       └── instructions/        # Обработчики инструкций
├── pkg/                         # Общие пакеты
│   ├── errors/                  # Обработка ошибок
│   └── utils/                   # Утилиты
├── go.mod                       # Зависимости
├── Makefile                     # Команды сборки
└── README.MD                    # Документация
```

## Основные сущности

### Market (Рынок)
Представляет прогнозный рынок с возможностью создания, закрытия и разрешения.

### Position (Позиция)
Представляет позицию пользователя на рынке (YES или NO).

## Инструкции Solana

1. **CreateMarket**: Создание нового рынка
2. **ResolveMarket**: Разрешение рынка (Yes/No)
3. **CreatePosition**: Создание позиции на рынке
4. **CloseMarket**: Закрытие рынка

## Установка и запуск

### Требования
- Go 1.21 или выше
- Solana CLI (для деплоя)

### Установка зависимостей
```bash
make deps
```

### Сборка
```bash
make build
```

### Тестирование
```bash
make test
```

### Форматирование кода
```bash
make fmt
```

### Линтинг
```bash
make lint
```

## Разработка

### Добавление новой функциональности

1. **Доменный слой**: Определите сущности и интерфейсы в `internal/domain/`
2. **Слой приложения**: Создайте use case в `internal/application/usecases/`
3. **Инфраструктура**: Реализуйте репозитории в `internal/infrastructure/repositories/`
4. **Представление**: Добавьте обработчик инструкций в `internal/presentation/instructions/`

## Примечания

- Программа использует Program Derived Addresses (PDA) для хранения данных
- Сериализация данных выполняется в формате Borsh (стандарт для Solana)
- В production версии необходимо добавить:
  - Полную реализацию сериализации/десериализации
  - Индексацию для поиска рынков и позиций
  - Валидацию прав доступа
  - Обработку ошибок и логирование
  - Тесты

## Лицензия

MIT

