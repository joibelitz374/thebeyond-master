-- name: AddSubscriptionExpiresAt :exec
UPDATE account
SET subscription_expires_at = COALESCE(subscription_expires_at, NOW()) + make_interval(secs => $1)
WHERE id = $2;

-- name: RemoveSubscriptionExpiresAt :exec
UPDATE account
SET subscription_expires_at = COALESCE(subscription_expires_at, NOW()) - make_interval(secs => $1)
WHERE id = $2;

-- name: CancelSubscriptions :exec
UPDATE account
SET subscription_expires_at = NULL
WHERE id = $1;
