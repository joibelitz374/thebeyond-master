package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/pkg/consts"
)

type AccountInterface interface {
	GetAllByPlatform(ctx context.Context, platform consts.Platform) (accounts []int, err error)
	Get(ctx context.Context, platform consts.Platform, platformAccountID int) (domain.Account, error)
	GetByAccountID(ctx context.Context, accountID int) (account domain.Account, err error)
	GetPlatformUserID(ctx context.Context, accountID int) (externalAccountID int, err error)
	GetExpiredUsers(ctx context.Context) ([]int, error)
	Create(ctx context.Context, platform consts.Platform, platformAccountID int, keyID, shortID string, expiresAt time.Time, promo *string, discount int) (int, error)
	AddSubscriptionExpiresAt(ctx context.Context, accountID int, duration time.Duration) error
	RemoveSubscriptionExpiresAt(ctx context.Context, accountID int, duration time.Duration) error
	CancelSubscriptions(ctx context.Context, accountID int) error
	SetDiscount(ctx context.Context, accountID int, discount int) error
	RegenerateKey(ctx context.Context, accountID int, keyID string) error
	SetRegion(ctx context.Context, accountID int, region string) error
	SetLanguage(ctx context.Context, accountID int, language string) error
	SetCurrency(ctx context.Context, accountID int, currency string) error
	SetProtocol(ctx context.Context, accountID int, protocol string) error
	SetLocation(ctx context.Context, accountID int, location string) error
}

type accountDB struct {
	pool *pgxpool.Pool
}

func NewAccount(pool *pgxpool.Pool) AccountInterface {
	return accountDB{pool}
}

func (db accountDB) GetAllByPlatform(ctx context.Context, platform consts.Platform) (accounts []int, err error) {
	rows, err := db.pool.Query(ctx, GET_ALL_ACCOUNTS_SQL, platform)
	if err != nil {
		return accounts, err
	}

	defer rows.Close()

	for rows.Next() {
		var accountID int
		if err := rows.Scan(&accountID); err != nil {
			return accounts, err
		}
		accounts = append(accounts, accountID)
	}

	return accounts, nil
}

func (db accountDB) Get(ctx context.Context, platform consts.Platform, platformAccountID int) (account domain.Account, err error) {
	txn, err := db.pool.Begin(ctx)
	if err != nil {
		return account, err
	}

	defer txn.Rollback(ctx)

	var accountID int
	if err := txn.QueryRow(ctx, GET_ACCOUNT_BY_PLATFORM_SQL, platform, platformAccountID).Scan(&accountID); err != nil {
		return account, err
	}

	if err := txn.QueryRow(ctx, GET_ACCOUNT_BY_ID_SQL, accountID).Scan(
		&account.ID,
		&account.KeyID,
		&account.ShortID,
		&account.SubscriptionExpiresAt,
		&account.Region,
		&account.Language,
		&account.Currency,
		&account.Protocol,
		&account.ServerLocation,
		&account.LastKeyRefreshAt,
		&account.Promo,
		&account.Discount,
	); err != nil {
		return account, err
	}

	return account, txn.Commit(ctx)
}

func (db accountDB) GetByAccountID(ctx context.Context, accountID int) (account domain.Account, err error) {
	if err := db.pool.QueryRow(ctx, GET_ACCOUNT_BY_ID_SQL, accountID).Scan(
		&account.ID,
		&account.KeyID,
		&account.ShortID,
		&account.SubscriptionExpiresAt,
		&account.Region,
		&account.Language,
		&account.Currency,
		&account.Protocol,
		&account.ServerLocation,
		&account.LastKeyRefreshAt,
		&account.Promo,
		&account.Discount,
	); err != nil {
		return account, err
	}
	return account, nil
}

func (db accountDB) GetPlatformUserID(ctx context.Context, accountID int) (externalAccountID int, err error) {
	return externalAccountID, db.pool.QueryRow(ctx, GET_PLATFORM_USER_ID_SQL, accountID).
		Scan(&externalAccountID)
}

func (db accountDB) GetExpiredUsers(ctx context.Context) (accounts []int, err error) {
	rows, err := db.pool.Query(ctx, GET_EXPIRED_ACCOUNTS_SQL)
	if err != nil {
		return accounts, err
	}

	defer rows.Close()

	for rows.Next() {
		var accountID int
		if err := rows.Scan(&accountID); err != nil {
			return accounts, err
		}
		accounts = append(accounts, accountID)
	}

	return accounts, nil
}

func (db accountDB) Create(ctx context.Context, platform consts.Platform, platformAccountID int, keyID, shortID string, expiresAt time.Time, promo *string, discount int) (int, error) {
	txn, err := db.pool.Begin(ctx)
	if err != nil {
		return 0, err
	}

	defer txn.Rollback(ctx)

	var newAccountID int
	if err := txn.QueryRow(ctx, INSERT_ACCOUNT_SQL, keyID, shortID, expiresAt, promo, discount).Scan(&newAccountID); err != nil {
		return 0, err
	}

	if _, err := txn.Exec(ctx, INSERT_PLATFORM_ACCOUNT_SQL, platform, platformAccountID, newAccountID); err != nil {
		return 0, err
	}

	return newAccountID, txn.Commit(ctx)
}

func (db accountDB) AddSubscriptionExpiresAt(ctx context.Context, accountID int, duration time.Duration) error {
	_, err := db.pool.Exec(ctx, ADD_SUBSCRIPTION_EXPIRES_AT_SQL, duration.Seconds(), accountID)
	return err
}

func (db accountDB) RemoveSubscriptionExpiresAt(ctx context.Context, accountID int, duration time.Duration) error {
	_, err := db.pool.Exec(ctx, REMOVE_SUBSCRIPTION_EXPIRES_AT_SQL, duration.Seconds(), accountID)
	return err
}

func (db accountDB) CancelSubscriptions(ctx context.Context, accountID int) error {
	_, err := db.pool.Exec(ctx, CANCEL_SUBSCRIPTIONS_SQL, accountID)
	return err
}

func (db accountDB) SetDiscount(ctx context.Context, accountID int, discount int) error {
	_, err := db.pool.Exec(ctx, SET_DISCOUNT_SQL, discount, accountID)
	return err
}

func (db accountDB) RegenerateKey(ctx context.Context, accountID int, keyID string) error {
	_, err := db.pool.Exec(ctx, SET_KEYID_SQL, keyID, accountID)
	return err
}

func (db accountDB) SetRegion(ctx context.Context, accountID int, region string) error {
	_, err := db.pool.Exec(ctx, SET_REGION_SQL, region, accountID)
	return err
}

func (db accountDB) SetLanguage(ctx context.Context, accountID int, language string) error {
	_, err := db.pool.Exec(ctx, SET_LANGUAGE_SQL, language, accountID)
	return err
}

func (db accountDB) SetCurrency(ctx context.Context, accountID int, currency string) error {
	_, err := db.pool.Exec(ctx, SET_CURRENCY_SQL, currency, accountID)
	return err
}

func (db accountDB) SetProtocol(ctx context.Context, accountID int, protocol string) error {
	_, err := db.pool.Exec(ctx, SET_PROTOCOL_SQL, protocol, accountID)
	return err
}

func (db accountDB) SetLocation(ctx context.Context, accountID int, location string) error {
	_, err := db.pool.Exec(ctx, SET_LOCATION_SQL, location, accountID)
	return err
}
