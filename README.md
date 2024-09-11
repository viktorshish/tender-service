tender-service

# tender-service

Сервис проведения тендеров

## Создание базы данных
Должна быть установлена PostgreSQL

```bash
psql postgres
$ CREATE DATABASE tender_service;
$ \l
$ \q
```

## Настройка проекта
Инициализируйте модуль

```bash
go mod init tender-service
```

Добавь необходимые зависимости:

•	GORM для работы с ORM.

•	Gin для создания HTTP API.

•	Godotenv для работы с переменными окружения из .env файла.

```bash
go get gorm.io/gorm gorm.io/driver/postgres github.com/gin-gonic/gin github.com/joho/godotenv
```

## Запуск сервера

```bash
go run main.go
```
