-- name: SetDiscount :exec
UPDATE account SET discount = $1
WHERE id = $2;

-- name: ResetDiscount :exec
UPDATE account SET discount = 0
WHERE id = $1;
