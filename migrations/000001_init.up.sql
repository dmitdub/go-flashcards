CREATE SCHEMA IF NOT EXISTS flashcards;

CREATE TABLE flashcards.users (
    id       SERIAL               PRIMARY KEY,
    version  BIGINT      NOT NULL DEFAULT 1,
    nickname VARCHAR(30) NOT NULL UNIQUE CHECK (char_length(nickname) BETWEEN 3 AND 30),
    phone    VARCHAR(15)          UNIQUE CHECK (
        phone IS NULL OR (
            phone ~ '^\+[0-9]+$'
            AND char_length(phone) BETWEEN 10 AND 15
        )
    )
);

CREATE TABLE flashcards.decks (
    id          SERIAL                PRIMARY KEY,
    version     BIGINT       NOT NULL DEFAULT 1,
    user_id     INTEGER      NOT NULL REFERENCES flashcards.users(id) ON DELETE CASCADE,
    title       VARCHAR(100) NOT NULL CHECK (char_length(title) BETWEEN 1 AND 100),
    description VARCHAR(500)          CHECK (
        description IS NULL OR char_length(description) BETWEEN 1 AND 500
    )
);

CREATE TABLE flashcards.cards (
    id               SERIAL                PRIMARY KEY,
    version          BIGINT       NOT NULL DEFAULT 1,
    deck_id          INTEGER      NOT NULL REFERENCES flashcards.decks(id) ON DELETE CASCADE,
    word             VARCHAR(200) NOT NULL CHECK (char_length(word) BETWEEN 1 AND 200),
    translation      VARCHAR(500) NOT NULL CHECK (char_length(translation) BETWEEN 1 AND 500),
    level            INTEGER      NOT NULL DEFAULT 0 CHECK (level BETWEEN 0 AND 5),
    next_review_at   TIMESTAMPTZ  NOT NULL,
    last_reviewed_at TIMESTAMPTZ,
    
    CHECK (
        (level = 0 AND last_reviewed_at IS NULL)
        OR
        (level BETWEEN 1 AND 5 AND last_reviewed_at IS NOT NULL)
    )
);