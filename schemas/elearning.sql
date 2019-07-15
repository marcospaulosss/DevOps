CREATE EXTENSION IF NOT EXISTS "citext";


DROP TABLE IF EXISTS albums_shelves;
DROP TABLE IF EXISTS sections_tracks;
DROP TABLE IF EXISTS subjects_tracks;
DROP TABLE IF EXISTS sections;
DROP TABLE IF EXISTS tracks;
DROP TABLE IF EXISTS shelves;
DROP TABLE IF EXISTS albums;
DROP TABLE IF EXISTS subjects;


CREATE TABLE shelves (
    id           SERIAL PRIMARY KEY,
    title        CITEXT UNIQUE NOT NULL CHECK (char_length(title) >= 1),
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ
);
INSERT INTO shelves (id, title) VALUES
(1, 'Conhecimentos Navais p/ Oficial Temporário da Marinha (SMV)'),
(2, 'Estatuto e Ética da OAB'),
(3, 'Inglês para CBM-BA');
ALTER SEQUENCE shelves_id_seq RESTART WITH 1000;


CREATE TABLE albums (
    id           SERIAL PRIMARY KEY,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ,
    title        CITEXT NOT NULL CHECK (char_length(title) >= 1),
    description  VARCHAR(255) NOT NULL DEFAULT '',
    image        VARCHAR(255) NOT NULL DEFAULT '',
    is_published BOOLEAN DEFAULT false,
    published_at TIMESTAMPTZ DEFAULT NOW()
);
CREATE index albums_title_index ON albums (title);
INSERT INTO albums (id, title, image, is_published, published_at) VALUES 
(4, 'Forças armadas', '024DD405-6710-4193-2B7F-E62FCD82D3F2.jpeg', true, NOW()),
(5, 'Princípios', '05AFA8BA-038A-EC58-415C-0FFF633D1A44.jpeg', true, NOW()),
(6, 'Ética', '9d558032-d692-4f8b-9e7f-8dc9a7b5588c.jpeg', false, NOW()),
(7, 'Regência verbal e nominal', 'C81F8156-CDA6-5D60-2BA1-D62CAEE78B3E.jpeg', true, NOW());
ALTER SEQUENCE albums_id_seq RESTART WITH 1000;

CREATE TABLE tracks (
    id serial   PRIMARY KEY,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ,
    title       CITEXT NOT NULL CHECK (char_length(title) >= 1),
    description VARCHAR NOT NULL DEFAULT '',
    teachers    VARCHAR NOT NULL DEFAULT '',
    media       VARCHAR NOT NULL DEFAULT '',
    duration    INTEGER NOT NULL DEFAULT 0
);
INSERT INTO tracks (id, title, description, duration, teachers, media, created_at) VALUES
(8, 'Forças Armadas (FFAA)', 'Missão Constitucional; Hierarquia e disciplina; e Comandante Supremo das Forças Armadas.', 30, 'Alan Hirt', '56c28a87-d8dc-4c4f-a3d4-a4d1a8a52839.mp3', '2019-03-03'),
(9, 'Doutrina de Liderança da Marinha', '', 30, 'Luiz Felipe da Rocha', '23a49560-cffb-4629-8790-c497eed7fde2.mp3', '2019-03-04'),
(10, 'Infrações disciplinares e Órgãos da OAB', '', 30, 'Daniela Garrido', '067771e0-1f60-4e62-b760-4f512218436f.mp3', '2019-03-06'),
(11, 'Princípios; Inscrição na OAB; Direitos do Advogado', '', 30, 'Daniela Garrido', '1ee24ce3-227f-40ff-a9b9-612d8099ceda.mp3', '2019-03-06'),
(12, 'Processo Disciplinar; Infrações; Sanções; Responsabilização; Competência e Órgãos da OAB', '', 30, 'Daniela Garrido', '5b72651a-fa10-439f-bac0-ebc2325a1409.mp3', '2019-03-06'),
(13, 'Técnicas de Interpretação de Textos e Cognatos', '', 40, 'Ena Loiola', 'fd844dcc-5981-47b1-a445-5f55a312c49d.mp3', '2019-03-07'),
(14, 'Formação de Palavras, Substantivos, Artigos', '', 45, 'Ena Loiola', '7da96347-4251-4934-9d06-e36ee6b15807.mp3', '2019-03-07'),
(15, 'Conectivos', '', 42, 'Ena Loiola', 'c5c94c85-bd07-4aa9-9c73-9dde53f42251.mp3', '2019-03-07'),
(16, 'Verbos Auxiliares, Frasais e Modais', '', 54, 'Ena Loiola', '7ad13ee1-27f9-473f-a8c4-e5a43079f19e.mp3', '2019-03-07'),
(17, 'Tempos Verbais Parte 1', '', 121, 'Ena Loiola', 'd53ce904-608e-4dac-9f6a-00097aaa627c.mp3', '2019-03-07'),
(18, 'Tempos Verbais Parte 2', '', 62, 'Ena Loiola', 'b91137e1-0db3-434f-9244-630a1acea6a9.mp3', '2019-03-07'),
(19, 'Expressões Idiomáticas', '', 64, 'Ena Loiola', 'dbdabe78-4ab8-415f-a378-0877797217cb.mp3', '2019-03-07');
ALTER SEQUENCE tracks_id_seq RESTART WITH 1000;


CREATE TABLE sections (
    id SERIAL   PRIMARY KEY,
    album_id    INTEGER REFERENCES albums ON DELETE CASCADE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title       CITEXT NOT NULL CHECK (char_length(title) >= 1),
    description VARCHAR DEFAULT '',
    position    INTEGER NOT NULL DEFAULT 0,
    UNIQUE      (album_id, title)
);
INSERT INTO sections (id, album_id, title, position) VALUES
(101, 4, 'Capítulo 1', 1), (102, 4, 'Capítulo 2', 0),
(103, 5, 'Capítulo 1', 0),
(104, 6, 'Introdução', 0),
(105, 7, 'Parte 1', 2), (106, 7, 'Resumo', 3), (107, 7, 'Sobre', 0);


CREATE TABLE albums_shelves (
    shelf_id       INTEGER REFERENCES shelves ON DELETE CASCADE,
    album_id       INTEGER REFERENCES albums ON DELETE CASCADE,
    album_position INTEGER NOT NULL DEFAULT 0,
    UNIQUE         (album_id, shelf_id)
);
INSERT INTO albums_shelves (shelf_id, album_id, album_position) VALUES (1, 4, 0), (2, 6, 0), (2, 5, 1), (3, 7, 0); 


CREATE TABLE sections_tracks (
    section_id     INTEGER REFERENCES sections ON DELETE CASCADE,
    track_id       INTEGER REFERENCES tracks ON DELETE CASCADE,
    track_position INTEGER NOT NULL DEFAULT 0,
    UNIQUE         (track_id, section_id)
);
INSERT INTO sections_tracks (section_id, track_id, track_position) VALUES
(101, 8, 0), (102, 9, 1),
(103, 10, 0), (103, 11, 2),
(104, 12, 0), (104, 15, 1),
(105, 13, 1), (105, 14, 0), (105, 15, 2), (105, 16, 4), (105, 17, 3), (105, 18, 5), (105, 19, 6),
(106, 13, 0),
(107, 15, 0);


DROP TABLE IF EXISTS preferences;
CREATE TABLE preferences (
    id SERIAL  PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    type VARCHAR(100) NOT NULL UNIQUE,
    content jsonb NOT NULL
);
INSERT INTO preferences (type, content) VALUES ('home', '{"shelves":[2,1,3]}');

CREATE TABLE subjects (
    id           SERIAL PRIMARY KEY,
    title        CITEXT UNIQUE NOT NULL CHECK (char_length(title) >= 1),
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ,
    UNIQUE         (title)
);
INSERT INTO subjects (title) VALUES
('Português'),
('Matématica'),
('Inglês');

CREATE TABLE subjects_tracks (
    subject_id       INTEGER REFERENCES subjects ON DELETE CASCADE,
    track_id       INTEGER REFERENCES tracks ON DELETE CASCADE,
    UNIQUE         (track_id)
);
INSERT INTO subjects_tracks (subject_id, track_id) VALUES (1, 8), (1, 9), (2, 10), (2, 11);
