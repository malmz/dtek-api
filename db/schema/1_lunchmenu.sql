-- +goose Up
CREATE TYPE language AS ENUM ('none', 'sv', 'en');

CREATE TABLE lunch_menus (
    id serial PRIMARY KEY,
    resturant text NOT NULL,
    date date NOT NULL,
    language language NOT NULL DEFAULT 'none',
    name text NOT NULL,
    menu jsonb NOT NULL,
    fetched_at timestamp with time zone NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE lunch_menus;

DROP TYPE language;