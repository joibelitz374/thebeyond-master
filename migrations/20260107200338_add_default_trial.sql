ALTER TABLE account
ALTER COLUMN tariff SET DEFAULT 1,
ALTER COLUMN subscription_expires_at SET DEFAULT NOW() + INTERVAL '3 days',
ALTER COLUMN subscription_status SET DEFAULT 'available';
