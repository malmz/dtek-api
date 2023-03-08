-- name: CreateNews :copyfrom
INSERT INTO news (
    title_sv,
    body_sv,
    title_en,
    body_en,
    committee,
    author,
    created_at,
    updated_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
);