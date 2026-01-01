-- name: GetPromosByCreator :many
SELECT
    name,
    creator,
    level,
    clients
FROM promo
WHERE creator = $1;

-- name: GetPromo :one
SELECT
    name,
    creator,
    level,
    clients
FROM promo
WHERE name = $1;

-- name: CreatePromo :exec
INSERT INTO promo (name, creator) VALUES ($1, $2);

-- name: UpgradePromoLevel :exec
UPDATE promo SET level = level + 1, clients = 0
WHERE name = $1;

-- name: IncreasePromoClients :exec
UPDATE promo SET clients = clients + 1
WHERE name = $1;
