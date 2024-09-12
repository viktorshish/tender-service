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

- [Gin](https://gin-gonic.com/docs/) - фреймворк для создания HTTP API.

- [GORM](https://gorm.io/docs/) фреймворк для работы с ORM.

- [Godotenv](https://github.com/joho/godotenv) для работы с переменными окружения из .env файла.


```bash
go get gorm.io/gorm gorm.io/driver/postgres github.com/gin-gonic/gin github.com/joho/godotenv
go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
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

### Сущности

Для работы с моделями необходимо установить расширение uuid-ossp:

```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

Пользователь (User):

```sql
CREATE TABLE employee (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(50) UNIQUE NOT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

Организация (Organization):

```sql
CREATE TYPE organization_type AS ENUM (
    'IE',
    'LLC',
    'JSC'
);

CREATE TABLE organization (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    type organization_type,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE organization_responsible (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,
    user_id UUID REFERENCES employee(id) ON DELETE CASCADE
);
```

Тендеры (Tender):

```sql
CREATE TYPE status_type AS ENUM (
    'CREATED',
    'PUBLISHED',
    'CLOSED', 
    'CANCELED'
);

CREATE TABLE tenders (
     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
     name VARCHAR(255) NOT NULL,
     description TEXT,
     status status_type NOT NULL,
     version INT DEFAULT 1,
     organization_id UUID NOT NULL,
     responsible_id UUID NOT NULL,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     FOREIGN KEY (organization_id) REFERENCES organization(id) ON DELETE CASCADE,
     FOREIGN KEY (responsible_id) REFERENCES organization_responsible(id) ON DELETE CASCADE
);
```
## Запуск сервера

```bash
go run main.go
```
