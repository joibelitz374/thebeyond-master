CREATE TABLE IF NOT EXISTS account (
    id SERIAL NOT NULL PRIMARY KEY,
    key_id TEXT NOT NULL,
    short_id TEXT NOT NULL,
    subscription_expires_at TIMESTAMPTZ DEFAULT NOW() + INTERVAL '1 day',
    region TEXT NOT NULL,
    language TEXT NOT NULL DEFAULT '',
    currency TEXT NOT NULL DEFAULT '',
    protocol TEXT NOT NULL DEFAULT 'vxr',
    server_location TEXT NOT NULL DEFAULT 'de',
    last_key_refresh_at TIMESTAMPTZ DEFAULT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS platform_account (
    platform_id SMALLINT NOT NULL,
    external_account_id BIGINT NOT NULL,
    fk_account_id INTEGER NOT NULL REFERENCES account ON DELETE CASCADE,
    PRIMARY KEY (platform_id, external_account_id)
);

CREATE TABLE IF NOT EXISTS payment (
    account_id INTEGER NOT NULL REFERENCES account ON DELETE CASCADE,
    amount INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMPTZ
);
