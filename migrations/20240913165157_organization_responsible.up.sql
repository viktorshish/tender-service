CREATE TABLE organization_responsible
(
    id              UUID PRIMARY KEY,
    organization_id UUID REFERENCES organization (id) ON DELETE CASCADE,
    user_id         UUID REFERENCES employee (id) ON DELETE CASCADE
);

-- Добавляем записи о том, что user1 и user2 являются ответственными за свои организации
INSERT INTO organization_responsible (id, organization_id, user_id)
VALUES
    ('77777777-7777-7777-7777-777777777777', '55555555-5555-5555-5555-555555555555', '11111111-1111-1111-1111-111111111111'), -- user1 ответственен за Avito LLC
    ('88888888-8888-8888-8888-888888888888', '66666666-6666-6666-6666-666666666666', '22222222-2222-2222-2222-222222222222'); -- user2 ответственен за Tech Corp
