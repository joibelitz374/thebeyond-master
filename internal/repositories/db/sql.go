package db

const GET_ALL_ACCOUNTS_SQL = `
	SELECT external_account_id
	FROM platform_account
	WHERE platform_id = $1;`

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
		last_key_refresh_at,
		promo,
		discount
	FROM account
	WHERE id = $1;`

const GET_PLATFORM_USER_ID_SQL = `
	SELECT external_account_id
	FROM platform_account
	WHERE fk_account_id = $1;`

const GET_EXPIRED_ACCOUNTS_SQL = `
	SELECT id
	FROM account
	WHERE subscription_expires_at < NOW()
	AND key_id <> '';`

const INSERT_ACCOUNT_SQL = `
	INSERT INTO account(key_id, short_id, subscription_expires_at, region, promo, discount)
	VALUES($1, $2, $3, 'ru', $4, $5) RETURNING id;`

const INSERT_PLATFORM_ACCOUNT_SQL = `
	INSERT INTO platform_account(platform_id, external_account_id, fk_account_id)
	VALUES ($1, $2, $3);`

const ACTIVATION_SUBSCRIPTION_SQL = `
	UPDATE account
	SET key_id = $1, subscription_expires_at = $2
	WHERE id = $3;`

const SET_KEYID_SQL = `
	UPDATE account
	SET key_id = $1
	WHERE id = $2;`

const ADD_SUBSCRIPTION_EXPIRES_AT_SQL = `
	UPDATE account
	SET subscription_expires_at = COALESCE(subscription_expires_at, NOW()) + make_interval(secs => $1)
	WHERE id = $2;`

const REMOVE_SUBSCRIPTION_EXPIRES_AT_SQL = `
	UPDATE account
	SET subscription_expires_at = COALESCE(subscription_expires_at, NOW()) - make_interval(secs => $1)
	WHERE id = $2;`

const CANCEL_SUBSCRIPTIONS_SQL = `
	UPDATE account
	SET subscription_expires_at = NULL
	WHERE id = $1;`

const SET_DISCOUNT_SQL = `
	UPDATE account
	SET discount = $1
	WHERE id = $2;`

const SET_REGION_SQL = `
	UPDATE account
	SET region = $1
	WHERE id = $2;`

const SET_LANGUAGE_SQL = `
	UPDATE account
	SET language = $1
	WHERE id = $2;`

const SET_CURRENCY_SQL = `
	UPDATE account
	SET currency = $1
	WHERE id = $2;`

const SET_PROTOCOL_SQL = `
	UPDATE account
	SET protocol = $1
	WHERE id = $2;`

const SET_LOCATION_SQL = `
	UPDATE account
	SET location = $1
	WHERE id = $2;`

const INSERT_PAYMENT_SQL = `
	INSERT INTO payment (account_id, amount, expires_at)
	VALUES($1, $2, $3);`

const GET_PAYMENTS_SQL = `
	SELECT amount, created_at, expires_at
	FROM payment
	WHERE account_id = $1
	AND expires_at > NOW();`

const DELETE_PAYMENT_SQL = `
	DELETE FROM payment
	WHERE account_id = $1;`

const GET_PROMOS_SQL = `
	SELECT name, creator, level, clients, buyers
	FROM promo
	WHERE creator = $1;`

const GET_PROMO_SQL = `
	SELECT name, creator, level, clients, buyers
	FROM promo
	WHERE name = $1;`

const INSERT_PROMO_SQL = `
	INSERT INTO promo(name, creator)
	VALUES($1, $2);`

const GET_PROMO_CLIENTS_SQL = `
	SELECT COUNT(*)
	FROM promo_client
	WHERE promo = $1;`

const UPGRADE_LEVEL_SQL = `
	UPDATE promo
	SET level = level + 1,
	clients = 0,
	buyers = 0
	WHERE name = $1;
`

const INCREASE_PROMO_CLIENTS_SQL = `
	UPDATE promo
	SET clients = clients + 1
	WHERE name = $1;`

const INCREASE_PROMO_BUYERS_SQL = `
	UPDATE promo
	SET buyers = buyers + 1
	WHERE name = $1;`

const GET_PROMO_BUYERS_SQL = `
	SELECT COUNT(*)
	FROM promo_buyer
	WHERE promo = $1;`

const CHECK_PROMO_BUYER_SQL = `
	SELECT fk_buyer_id
	FROM promo_buyer
	WHERE promo = $1
	AND fk_buyer_id = $2;`

const INSERT_PROMO_BUYER_SQL = `
	INSERT INTO promo_buyer(promo, fk_buyer_id)
	VALUES($1, $2);`
