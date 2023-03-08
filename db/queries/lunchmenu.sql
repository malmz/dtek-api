-- name: CreateLunchMenus :copyfrom
INSERT INTO lunch_menus (
        resturant,
        date,
        language,
        name,
        menu
    )
VALUES ($1, $2, $3, $4, $5);

-- name: GetLunchByDate :one
SELECT *
FROM lunch_menus
WHERE resturant = $1
    AND date = $2
    AND language = $3
    OR language = 'none';

-- name: GetLunchByDateRange :many
SELECT *
FROM lunch_menus
WHERE resturant = $1
    AND date >= $2
    AND date <= $3
    AND language = $4
    OR language = 'none';