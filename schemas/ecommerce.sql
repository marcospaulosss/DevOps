CREATE EXTENSION IF NOT EXISTS "citext";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS products;
CREATE TABLE products (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    started_selling_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    finished_selling_at TIMESTAMPTZ NOT NULL DEFAULT NOW() + INTERVAL '1 month',
    usage_expires_at TIMESTAMPTZ NOT NULL DEFAULT NOW() + INTERVAL '12 month',
    name CITEXT NOT NULL UNIQUE CHECK (char_length(name) > 2),
    description TEXT DEFAULT '',
    stock INTEGER DEFAULT 1,
    sku SERIAL NOT NULL,
    image VARCHAR NOT NULL DEFAULT '',
    is_published BOOLEAN DEFAULT false,
    company VARCHAR NOT NULL DEFAULT 'estrategia sa',
    product_type VARCHAR NOT NULL DEFAULT 'pacote',
    payments_types jsonb NOT NULL,
    items jsonb NOT NULL,
    history jsonb
);

CREATE INDEX idx_products_history ON products USING gin (history);
CREATE INDEX idx_products_items ON products USING gin (items);
CREATE INDEX idx_products_payments_types ON products USING gin (payments_types);
CREATE INDEX idx_products_product_type ON products (product_type);

INSERT INTO products (id, name, stock, product_type, payments_types, items) VALUES ('575e7695-241d-4b96-8061-9f8a3771f3cc', 'Português para Receita Federal', 1, 'pacote', '[{"id": "creditcard", "name": "Cartão de Crédito", "price": 19990, "installments": 12}]', '[{"name": "Curso de português para Receita"}]');
INSERT INTO products (id, name, stock, product_type, history, payments_types, items) VALUES ('a0d354bc-b8c3-4038-a18d-12307e80759b', 'Elearning Audio', 1, 'assinatura', '[{"created_at": "2019-06-05 20:50:12.962209+00", "name": "TRT audios", "payments_types": "[\"id\": \"creditcard\", \"name\": \"Cartão\", \"price\": 600, \"installments\": 10}]", "stock": 5}]'::jsonb, '[{"id": "creditcard", "name": "Cartão de Crédito", "price": 39999, "installments": 12}]'::jsonb, '[{"name":"assinatural full", "frequency": "month", "interval": 1}]');

DROP TABLE IF EXISTS courses;
CREATE TABLE courses (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    name CITEXT NOT NULL UNIQUE CHECK (char_length(name) > 5)
); 

INSERT INTO courses (id, name) VALUES ('732f4510-2a38-40aa-905c-fb3d91b24807', 'Curso de português para Receita');

DROP TABLE IF EXISTS subscriptions;
CREATE TABLE subscriptions (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    name CITEXT NOT NULL UNIQUE CHECK (char_length(name) > 1),
    frequency VARCHAR NOT NULL DEFAULT 'month',
    interval INTEGER NOT NULL DEFAULT 1,
    contents jsonb NOT NULL
); 

INSERT INTO subscriptions (id, name, contents) VALUES ('4e7763e4-8b93-4c84-ba93-fca60dace896', 'Full', '{"audios":["*"]}'::jsonb);

