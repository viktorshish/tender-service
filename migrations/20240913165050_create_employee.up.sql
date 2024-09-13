CREATE TABLE employee
(
    id         UUID PRIMARY KEY,
    username   VARCHAR(50) UNIQUE NOT NULL,
    first_name VARCHAR(50),
    last_name  VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Добавляем пользователей в таблицу employee
INSERT INTO employee (id, username, first_name, last_name, created_at, updated_at)
VALUES
    ('11111111-1111-1111-1111-111111111111', 'user1', 'John', 'Doe', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('22222222-2222-2222-2222-222222222222', 'user2', 'Jane', 'Smith', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('33333333-3333-3333-3333-333333333333', 'user3', 'Alice', 'Johnson', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('44444444-4444-4444-4444-444444444444', 'user4', 'Bob', 'Brown', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);