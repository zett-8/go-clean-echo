-- +goose Up
-- +goose StatementBegin
CREATE table authors (
    id SERIAL,
    name text,
    country text,
    CONSTRAINT author_pk PRIMARY KEY(id)
);

CREATE TABLE books (
    id SERIAL,
    name text,
    author_id int,
    CONSTRAINT book_pk PRIMARY KEY(id),
    CONSTRAINT book_to_author FOREIGN KEY (author_id) REFERENCES authors(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE books;
DROP TABLE authors;
-- +goose StatementEnd
