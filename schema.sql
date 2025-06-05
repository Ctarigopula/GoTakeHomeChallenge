CREATE TABLE clients
(
    id       SERIAL PRIMARY KEY,
    name     TEXT,
    metadata JSONB
);

CREATE TABLE client_users
(
    id         SERIAL PRIMARY KEY,
    client_id INTEGER,
    first_name TEXT,
    last_name TEXT,
    metadata  JSONB
);

CREATE TABLE client_messages
(
    id         SERIAL PRIMARY KEY,
    created    TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL,
    deleted_at_order INTEGER,
    user_ids   INTEGER[],
    metadata   JSONB
);
