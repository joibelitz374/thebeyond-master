-- name: CreatePayment :exec
INSERT INTO payment (account_id, amount, expires_at) VALUES ($1, $2, $3);

-- name: GetPayments :many
SELECT
    amount,
    created_at,
    expires_at
FROM
    payment
WHERE
    account_id = $1
    AND expires_at > NOW();

-- name: DeletePayment :exec
DELETE FROM payment
WHERE account_id = $1;
