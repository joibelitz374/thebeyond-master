-- name: GetAccountIDByPlatform :one
SELECT fk_account_id
FROM platform_account
WHERE
    platform_id = $1
    AND external_account_id = $2;

-- name: GetAccountByID :one
SELECT
    id,
    cluster_id,
    key_id,
    short_id,
    tariff,
    discount,
    promo,
    freemium_status,
    subscription_status,
    subscription_expires_at,
    devices,
    used_uplink,
    used_downlink,
    protocol,
    region,
    language,
    currency,
    last_key_refresh_at,
    last_traffic_refresh_at
FROM account
WHERE id = $1;

-- name: GetAccountByKeyID :one
SELECT
    id,
    cluster_id,
    key_id,
    short_id,
    tariff,
    discount,
    promo,
    freemium_status,
    subscription_status,
    subscription_expires_at,
    devices,
    used_uplink,
    used_downlink,
    protocol,
    region,
    language,
    currency,
    last_key_refresh_at,
    last_traffic_refresh_at
FROM account
WHERE key_id = $1;

-- name: GetPlatformUserID :one
SELECT external_account_id
FROM platform_account
WHERE fk_account_id = $1;

-- name: GetAllAccountsByPlatform :many
SELECT external_account_id
FROM platform_account
WHERE platform_id = $1;

-- name: FindUsersForServiceCheck :many
SELECT
    a.id,
    pa.external_account_id,
    a.language
FROM account AS a
INNER JOIN platform_account AS pa ON a.id = pa.fk_account_id
WHERE
    a.created_at <= NOW() - INTERVAL '5 minutes'
    AND a.service_check_sent = 0;

-- name: CreateAccount :one
INSERT INTO account (key_id, short_id, region, promo, discount)
VALUES ($1, $2, $3, $4, $5) RETURNING id;

-- name: CreatePlatformAccount :exec
INSERT INTO platform_account (platform_id, external_account_id, fk_account_id)
VALUES ($1, $2, $3);

-- name: MarkServiceCheckSent :exec
UPDATE account SET service_check_sent = 1
WHERE id = $1;

-- name: SetTariff :exec
UPDATE account SET tariff = $1
WHERE id = $2;
