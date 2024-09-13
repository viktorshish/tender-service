CREATE TYPE organization_type AS ENUM (
    'IE',
    'LLC',
    'JSC'
    );

CREATE TABLE organization
(
    id          UUID PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    description TEXT,
    type        organization_type,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Добавляем организации в таблицу organization
INSERT INTO organization (id, name, description, type, created_at, updated_at)
VALUES
    ('55555555-5555-5555-5555-555555555555', 'Avito LLC', 'Технологическая компания', 'LLC', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('66666666-6666-6666-6666-666666666666', 'Tech Corp', 'Компания-разработчик ПО', 'JSC', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);