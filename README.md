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

Создайте в корне проекта файл *.env* и поместите в него данные подключения к серверу и БД:

```
SERVER_ADDRESS — адрес и порт, который будет слушать HTTP сервер при запуске. Пример: 0.0.0.0:8080.
POSTGRES_USERNAME — имя пользователя для подключения к PostgreSQL.
POSTGRES_PASSWORD — пароль для подключения к PostgreSQL.
POSTGRES_HOST — хост для подключения к PostgreSQL (например, localhost).
POSTGRES_PORT — порт для подключения к PostgreSQL (например, 5432).
POSTGRES_DATABASE — имя базы данных PostgreSQL, которую будет использовать приложение.
```

## Запуск сервера

```bash
go run main.go
```
