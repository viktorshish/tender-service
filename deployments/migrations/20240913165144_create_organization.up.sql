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

INSERT INTO organization (id, name, description, type, created_at, updated_at)
VALUES
    ('44444444-4444-4444-4444-444444444444', 'Avito LLC', 'Крупная технологическая компания', 'LLC', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('55555555-5555-5555-5555-555555555555', 'Tech Corp JSC', 'Компания в сфере информационных технологий', 'JSC', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('66666666-6666-6666-6666-666666666666', 'Innovations IE', 'Малый бизнес в сфере инноваций', 'IE', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);