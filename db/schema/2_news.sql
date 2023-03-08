-- +goose Up
CREATE TABLE news (
    id serial PRIMARY KEY,
    title_sv text NOT NULL,
    body_sv text NOT NULL,
    title_en text NOT NULL,
    body_en text NOT NULL,
    committee text,
    author text,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
);

-- +goose Down
DROP TABLE news;