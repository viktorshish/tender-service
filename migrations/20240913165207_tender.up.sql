CREATE TYPE status_type AS ENUM (
    'CREATED',
    'PUBLISHED',
    'CLOSED',
    'CANCELED'
    );

CREATE TABLE tender
(
    id              UUID PRIMARY KEY,
    name            VARCHAR(255) NOT NULL,
    description     TEXT,
    service_type    VARCHAR(255) NOT NULL,
    status          status_type  NOT NULL DEFAULT 'CREATED',
    version         INT          NOT NULL DEFAULT 1,
    organization_id UUID         NOT NULL REFERENCES organization (id) ON DELETE CASCADE,
    responsible_id  UUID         NOT NULL REFERENCES employee (id) ON DELETE SET NULL,
    creator_id      UUID         NOT NULL REFERENCES employee (id) ON DELETE CASCADE,
    created_at      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);
