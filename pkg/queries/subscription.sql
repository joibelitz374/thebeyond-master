-- name: GetNeedRefreshTraffic :many
SELECT id FROM account WHERE last_traffic_refresh_at < NOW() - INTERVAL '30 days';

-- name: GetAccountsToDisable :many
SELECT id FROM (
    SELECT
        id,
        subscription_expires_at,
        is_disabled,
        used_uplink + used_downlink AS used_bytes,
        ((CASE
            WHEN subscription_status = 'available'
                THEN CASE tariff
                    WHEN 1 THEN 150 WHEN 2 THEN 400 WHEN 3 THEN 5120
                    ELSE 0
                END
            ELSE 0
        END)
        + CASE WHEN freemium_status = 'available' THEN 10 ELSE 0 END)
        * 1024::BIGINT * 1024 * 1024 AS traffic_limit_bytes
    FROM account
) AS a
WHERE NOT is_disabled AND ((
    subscription_expires_at IS NOT NULL
    AND subscription_expires_at < now()
) OR (traffic_limit_bytes > 0 AND used_bytes >= traffic_limit_bytes));

-- name: GetAccountsToEnable :many
SELECT id, key_id FROM (
    SELECT
        id,
        key_id,
        subscription_expires_at,
        is_disabled,
        used_uplink + used_downlink AS used_bytes,
        (
            (CASE
                WHEN subscription_status = 'available'
                    THEN CASE tariff
                        WHEN 1 THEN 150 WHEN 2 THEN 400 WHEN 4 THEN 5120
                        ELSE 0
                    END
                ELSE 0
            END)
            + CASE WHEN freemium_status = 'available' THEN 10 ELSE 0 END
        ) * 1024::BIGINT * 1024 * 1024 AS traffic_limit_bytes
    FROM account
) AS a
WHERE is_disabled
  AND subscription_expires_at IS NOT NULL
  AND subscription_expires_at >= now()
  AND (traffic_limit_bytes = 0 OR used_bytes < traffic_limit_bytes);

-- name: ResetLastTrafficRefreshAt :exec
UPDATE account SET last_traffic_refresh_at = NOW(),
    used_uplink = 0, used_downlink = 0
WHERE id = $1;

-- name: GetFremiumAccounts :many
SELECT a.id, pa.external_account_id
FROM platform_account pa
JOIN account a ON a.id = pa.fk_account_id
WHERE pa.platform_id = 0
  AND a.freemium_status = 'available';

-- name: StartFreemium :exec
UPDATE account SET freemium_status = 'available' WHERE id = $1;

-- name: RemoveFreemium :exec
UPDATE account SET freemium_status = 'unavailable' WHERE id = $1;

-- name: SetSubscriptionExpiresAt :exec
UPDATE account SET subscription_expires_at = $1 WHERE id = $2;

-- name: AddSubscriptionExpiresAt :exec
UPDATE account
SET
    subscription_expires_at
    = coalesce(subscription_expires_at, now()) + make_interval(secs => $1)
WHERE id = $2;

-- name: RemoveSubscriptionExpiresAt :exec
UPDATE account
SET
    subscription_expires_at
    = coalesce(subscription_expires_at, now()) - make_interval(secs => $1)
WHERE id = $2;

-- name: CancelSubscriptions :exec
UPDATE account
SET
    tariff = 0,
    subscription_status = 'unavailable',
    subscription_expires_at = NULL,
    is_disabled = true
WHERE id = $1;
