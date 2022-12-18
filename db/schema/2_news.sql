CREATE TABLE news (
    id serial PRIMARY KEY,
    resturant text NOT NULL,
    date timestamp WITH time zone NOT NULL,
    language language NOT NULL DEFAULT 'none',
    name text NOT NULL,
    menu jsonb NOT NULL,
    fetched_at timestamp WITH time zone NOT NULL DEFAULT NOW()
);