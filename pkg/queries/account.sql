-- name: GetAccountIDByPlatform :one
SELECT fk_account_id
FROM platform_account
WHERE platform_id = $1
AND external_account_id = $2;

-- name: GetAccountByID :one
SELECT
    id,
    cluster_id,
    key_id,
    short_id,
    region,
    language,
    currency,
    protocol,
    subscription_expires_at,
    last_key_refresh_at,
    promo,
    discount
FROM account
WHERE id = $1;

-- name: GetAccountByKeyID :one
SELECT
    id,
    cluster_id,
    key_id,
    short_id,
    region,
    language,
    currency,
    protocol,
    subscription_expires_at,
    last_key_refresh_at,
    promo,
    discount
FROM account
WHERE key_id = $1;

-- name: GetPlatformUserID :one
SELECT external_account_id
FROM platform_account
WHERE fk_account_id = $1;

-- name: GetExpiredAccounts :many
SELECT id
FROM account
WHERE subscription_expires_at < NOW()
AND key_id <> '';

-- name: GetAllAccountsByPlatform :many
SELECT external_account_id
FROM platform_account
WHERE platform_id = $1;

-- name: FindUsersForServiceCheck :many
SELECT a.id, pa.external_account_id, a.language
FROM account a
JOIN platform_account pa ON pa.fk_account_id = a.id
WHERE a.created_at BETWEEN NOW() - INTERVAL '10 minutes'
  AND NOW() - INTERVAL '5 minutes'
AND a.service_check_sent = 0;

-- name: CreateAccount :one
INSERT INTO account(key_id, short_id, subscription_expires_at, region, promo, discount)
VALUES($1, $2, $3, $4, $5, $6) RETURNING id;

-- name: CreatePlatformAccount :exec
INSERT INTO platform_account(platform_id, external_account_id, fk_account_id)
VALUES ($1, $2, $3);

-- name: MarkServiceCheckSent :exec
UPDATE account
SET service_check_sent = 1
WHERE id = $1;