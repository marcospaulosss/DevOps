CREATE EXTENSION IF NOT EXISTS "citext";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS accounts;
CREATE TABLE users (
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(15) NOT NULL,
    active BOOLEAN DEFAULT TRUE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ
);
CREATE index users_name_index ON users (name);
CREATE UNIQUE index users_email_index ON users (email);
CREATE UNIQUE index users_phone_index ON users (phone);

INSERT INTO users (name, email, phone) VALUES (
    'My first user',
    'usersfirst@estrategiaconcursos.com.br',
    '5511999999999'
),(
    'My second user',
    'userssecond@estrategiaconcursos.com.br',
    '4511999999999'
),(
    'My third user',
    'usersthird@estrategiaconcursos.com.br',
    '4511999999998'
),(
    'My four user',
    'usersfour@estrategiaconcursos.com.br',
    '5511999999666'
),(
    'My fivve user',
    'usersfive@estrategiaconcursos.com.br',
    '5511987604321'
),(
    'My six user',
    'userssix@estrategiaconcursos.com.br',
    '5511987604121'
),(
    'My seven user',
    'usersseven@estrategiaconcursos.com.br',
    '5511983604121'
);

CREATE TABLE accounts (
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) DEFAULT '',
    email_code VARCHAR(6) DEFAULT '',
    phone VARCHAR(15) DEFAULT '',
    phone_code VARCHAR(6) DEFAULT '',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

INSERT INTO accounts (email, email_code, phone, phone_code) VALUES
('code@estrategiaconcursos.com.br', 987654, '5511999999999', 123456),
('code2@estrategiaconcursos.com.br', 957654, '5511999999998', 123457);

INSERT INTO accounts (email, email_code) VALUES ('code3@estrategiaconcursos.com.br', 987633);
INSERT INTO accounts (email, email_code) VALUES ('usersfirst@estrategiaconcursos.com.br', 123456);
INSERT INTO accounts (email, email_code) VALUES ('userssecond@estrategiaconcursos.com.br', 123456);
INSERT INTO accounts (email, email_code) VALUES ('userssix@estrategiaconcursos.com.br', 123456);
INSERT INTO accounts (phone, phone_code) VALUES ('5511999999966', 123466);
INSERT INTO accounts (phone, phone_code) VALUES ('5511999999666', 123466);
INSERT INTO accounts (phone, phone_code) VALUES ('4511999999998', 123466);
INSERT INTO accounts (phone, phone_code) VALUES ('5511983604121', 123456);
