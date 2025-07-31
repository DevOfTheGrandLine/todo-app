### README.md

# To-Do Менеджер на Go

Консольное приложение для управления задачами с поддержкой команд:
- Добавление, просмотр, завершение и удаление задач
- Импорт/экспорт данных в форматах JSON и CSV
- Фильтрация задач по статусу

## Требования
- Go версии 1.20 или выше

## Установка и сборка
1. Клонируйте репозиторий:

git clone https://github.com/yourusername/todo-app.git
cd todo-app


2. Соберите приложение:

go build -o todo ./cmd/todo


3. Запустите исполняемый файл:

./todo [команда] [флаги]


## Использование

### Основные команды

**Добавить задачу:**

./todo add --desc="Описание задачи"


**Вывести список задач:**

# Все задачи
./todo list

# Выполненные задачи
./todo list --filter=done

# Невыполненные задачи
./todo list --filter=pending


**Отметить задачу выполненной:**

./todo complete --id=1


**Удалить задачу:**

./todo delete --id=1


**Экспортировать задачи:**

# В JSON
./todo export --format=json --out=backup.json

# В CSV
./todo export --format=csv --out=backup.csv


**Импортировать задачи:**

./todo load --file=import.json
./todo load --file=import.csv


## Структура проекта

todo-app/
├── cmd/
│   └── todo/             # Точка входа в приложение
│       └── main.go
├── internal/
│   ├── todo/             # Логика работы с задачами
│   │   ├── manager.go
│   │   └── task.go
│   └── storage/          # Работа с файловым хранилищем
│       ├── csv_storage.go
│       └── json_storage.go
├── go.mod                # Файл модуля Go
└── README.md             # Документация


## Особенности
- Автоматическое сохранение данных в файл `tasks.json`
- Автогенерация ID задач
- Поддержка русского языка в командах
- Детальная обработка ошибок

## Лицензия
[MIT License](LICENSE)