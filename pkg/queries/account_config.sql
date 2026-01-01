-- name: RegenerateKey :exec
UPDATE account SET key_id = $1
WHERE id = $2;

-- name: SetRegion :exec
UPDATE account SET region = $1
WHERE id = $2;

-- name: SetLanguage :exec
UPDATE account SET language = $1
WHERE id = $2;

-- name: SetCurrency :exec
UPDATE account SET currency = $1
WHERE id = $2;

-- name: SetProtocol :exec
UPDATE account SET protocol = $1
WHERE id = $2;

-- name: SetLocation :exec
UPDATE account SET server_location = $1
WHERE id = $2;
