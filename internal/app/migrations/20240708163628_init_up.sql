-- +goose Up
CREATE TABLE users (
    id serial NOT NULL PRIMARY KEY,
    username text NOT NULL UNIQUE,
    password text NOT NULL,
    registered_at timestamp with time zone DEFAULT now() NOT NULL
);

CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    release_year INT NOT NULL,
    genre VARCHAR(100),
    director VARCHAR(255),
    writer VARCHAR(255),
    actors TEXT,
    imdb_rating DECIMAL(3, 1),
    duration INT, -- in minutes
    language VARCHAR(100),
    country VARCHAR(100),
    budget BIGINT,
    box_office BIGINT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE carts (
    user_id  integer NOT NULL primary key,
    book_ids integer[] not null,
    created_at 		timestamp with time zone 	DEFAULT now() NOT NULL,
    updated_at 		timestamp with time zone,

    FOREIGN KEY (user_id) REFERENCES users(id)
);
-- +goose Down
DROP TABLE users;
DROP TABLE movies;