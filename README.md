# tender-service

Сервис проведения тендеров

## Installation

```bash
# create .env and fill your data
$ cp .env.dist .env

# run app 
$ make start

# migrate
$ make migrate-up

# enjoy!
```

## Tests

1.	*GET /api/ping* - 
Проверка доступности сервера. Возвращает “pong”.

2.	*POST /api/tenders/new* - 
Создание нового тендера.

3.	*PATCH /api/tenders/:id/publish* - 
Публикация тендера (доступно только ответственным за организацию).

4.	*PATCH /api/tenders/:id/cancel* - 
Отмена тендера (доступно только автору и ответственным за организацию).

5.	*GET /api/tenders* - 
Получение списка тендеров с возможностью фильтрации по типу услуг (query параметр serviceType).


