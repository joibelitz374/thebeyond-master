package db

const GET_ACCOUNT_BY_PLATFORM_SQL = `
	SELECT fk_account_id
	FROM platform_account
	WHERE platform_id = $1
	AND external_account_id = $2;`

const GET_ACCOUNT_BY_ID_SQL = `
	SELECT
		id,
		key_id,
		short_id,
		subscription_expires_at,
		region,
		language,
		currency,
		protocol,
		server_location,
		last_key_refresh_at
	FROM account
	WHERE id = $1;`

const GET_EXPIRED_USERS_SQL = "SELECT id FROM account WHERE subscription_expires_at < NOW() AND key_id <> '';"

const INSERT_ACCOUNT_SQL = "INSERT INTO account(key_id, short_id, subscription_expires_at, region) VALUES($1, $2, $3, 'ru') RETURNING id;"

const INSERT_PLATFORM_ACCOUNT_SQL = `INSERT INTO platform_account(platform_id, external_account_id, fk_account_id) VALUES ($1, $2, $3);`

const ACTIVATION_SUBSCRIPTION_SQL = "UPDATE account SET key_id = $1, subscription_expires_at = $2 WHERE id = $3;"

const SET_KEYID_SQL = "UPDATE account SET key_id = $1 WHERE id = $2;"

const SET_SUBSCRIPTION_EXPIRES_AT_SQL = "UPDATE account SET subscription_expires_at = $1 WHERE id = $2;"

const SET_REGION_SQL = "UPDATE account SET region = $1 WHERE id = $2;"

const SET_LANGUAGE_SQL = "UPDATE account SET language = $1 WHERE id = $2;"

const SET_CURRENCY_SQL = "UPDATE account SET currency = $1 WHERE id = $2;"

const SET_PROTOCOL_SQL = "UPDATE account SET protocol = $1 WHERE id = $2;"

const SET_LOCATION_SQL = "UPDATE account SET location = $1 WHERE id = $2;"

const INSERT_PAYMENT_SQL = "INSERT INTO payment (account_id, amount, expires_at) VALUES($1, $2, $3);"

const GET_PAYMENTS_SQL = "SELECT amount, created_at, expires_at FROM payment WHERE account_id = $1 AND expires_at > NOW();"

const DELETE_PAYMENT_SQL = "DELETE FROM payment WHERE account_id = $1;"
