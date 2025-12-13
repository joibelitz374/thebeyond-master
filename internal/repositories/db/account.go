package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/quickpowered/thebeyond-master/internal/domain"
	"github.com/quickpowered/thebeyond-master/pkg/consts"
	"github.com/quickpowered/thebeyond-master/pkg/sqlc"
)

type AccountInterface interface {
	Get(ctx context.Context, platform consts.Platform, platformAccountID int) (domain.Account, error)
	GetByAccountID(ctx context.Context, accountID int) (domain.Account, error)
	GetByKeyID(ctx context.Context, keyID string) (domain.Account, error)
	GetPlatformUserID(ctx context.Context, accountID int) (int, error)
	GetExpiredUsers(ctx context.Context) ([]int, error)
	GetAllByPlatform(ctx context.Context, platform consts.Platform) ([]int, error)
	Create(ctx context.Context, platform consts.Platform, platformAccountID int, keyID, shortID string, expiresAt time.Time, promo *string, discount int) (int, error)
	AddSubscriptionExpiresAt(ctx context.Context, accountID int, duration time.Duration) error
	RemoveSubscriptionExpiresAt(ctx context.Context, accountID int, duration time.Duration) error
	CancelSubscriptions(ctx context.Context, accountID int) error
	SetDiscount(ctx context.Context, accountID int, discount int) error
	ResetDiscount(ctx context.Context, accountID int) error
	RegenerateKey(ctx context.Context, accountID int, keyID string) error
	SetRegion(ctx context.Context, accountID int, region string) error
	SetLanguage(ctx context.Context, accountID int, language string) error
	SetCurrency(ctx context.Context, accountID int, currency string) error
	SetProtocol(ctx context.Context, accountID int, protocol string) error
	SetLocation(ctx context.Context, accountID int, location string) error
}

type accountDB struct {
	pool    *pgxpool.Pool
	queries *sqlc.Queries
}

func NewAccount(pool *pgxpool.Pool) AccountInterface {
	return accountDB{pool, sqlc.New(pool)}
}

func (db accountDB) Get(ctx context.Context, platform consts.Platform, platformAccountID int) (domain.Account, error) {
	txn, err := db.pool.Begin(ctx)
	if err != nil {
		return domain.Account{}, err
	}

	defer txn.Rollback(ctx)

	qtxn := db.queries.WithTx(txn)

	accountID, err := qtxn.GetAccountIDByPlatform(ctx, sqlc.GetAccountIDByPlatformParams{
		PlatformID:        int16(platform),
		ExternalAccountID: int64(platformAccountID),
	})
	if err != nil {
		return domain.Account{}, err
	}

	account, err := qtxn.GetAccountByID(ctx, int32(accountID))
	if err != nil {
		return domain.Account{}, nil
	}

	var promo *string
	if account.Promo.Valid {
		promo = &account.Promo.String
	}

	return domain.Account{
		ID:                    int(account.ID),
		ClusterID:             int(account.ClusterID.Int16),
		KeyID:                 account.KeyID,
		ShortID:               account.ShortID,
		Region:                account.Region,
		Language:              account.Language,
		Currency:              account.Currency,
		Protocol:              account.Protocol,
		SubscriptionExpiresAt: &account.SubscriptionExpiresAt.Time,
		LastKeyRefreshAt:      &account.LastKeyRefreshAt.Time,
		Promo:                 promo,
		Discount:              int(account.Discount.Int32),
	}, txn.Commit(ctx)
}

func (db accountDB) GetByAccountID(ctx context.Context, accountID int) (domain.Account, error) {
	account, err := db.queries.GetAccountByID(ctx, int32(accountID))

	var promo *string
	if account.Promo.Valid {
		promo = &account.Promo.String
	}

	return domain.Account{
		ID:                    int(account.ID),
		ClusterID:             int(account.ClusterID.Int16),
		KeyID:                 account.KeyID,
		ShortID:               account.ShortID,
		Region:                account.Region,
		Language:              account.Language,
		Currency:              account.Currency,
		Protocol:              account.Protocol,
		SubscriptionExpiresAt: &account.SubscriptionExpiresAt.Time,
		LastKeyRefreshAt:      &account.LastKeyRefreshAt.Time,
		Promo:                 promo,
		Discount:              int(account.Discount.Int32),
	}, err
}

func (db accountDB) GetByKeyID(ctx context.Context, keyID string) (domain.Account, error) {
	account, err := db.queries.GetAccountByKeyID(ctx, keyID)

	var promo *string
	if account.Promo.Valid {
		promo = &account.Promo.String
	}

	return domain.Account{
		ID:                    int(account.ID),
		ClusterID:             int(account.ClusterID.Int16),
		KeyID:                 account.KeyID,
		ShortID:               account.ShortID,
		Region:                account.Region,
		Language:              account.Language,
		Currency:              account.Currency,
		Protocol:              account.Protocol,
		SubscriptionExpiresAt: &account.SubscriptionExpiresAt.Time,
		LastKeyRefreshAt:      &account.LastKeyRefreshAt.Time,
		Promo:                 promo,
		Discount:              int(account.Discount.Int32),
	}, err
}

func (db accountDB) GetPlatformUserID(ctx context.Context, accountID int) (int, error) {
	externalAccountID, err := db.queries.GetPlatformUserID(ctx, int32(accountID))
	return int(externalAccountID), err
}

func (db accountDB) GetExpiredUsers(ctx context.Context) ([]int, error) {
	rows, err := db.queries.GetExpiredAccounts(ctx)
	if err != nil {
		return nil, err
	}

	accounts := make([]int, len(rows))
	for i, row := range rows {
		accounts[i] = int(row)
	}

	return accounts, nil
}

func (db accountDB) GetAllByPlatform(ctx context.Context, platform consts.Platform) (accounts []int, err error) {
	rows, err := db.queries.GetAllAccountsByPlatform(ctx, int16(platform))
	if err != nil {
		return accounts, err
	}

	accounts = make([]int, len(rows))
	for i, row := range rows {
		accounts[i] = int(row)
	}

	return accounts, nil
}

func (db accountDB) Create(ctx context.Context, platform consts.Platform, platformAccountID int, keyID, shortID string, expiresAt time.Time, promo *string, discount int) (int, error) {
	txn, err := db.pool.Begin(ctx)
	if err != nil {
		return 0, err
	}

	defer txn.Rollback(ctx)

	createAccountParams := sqlc.CreateAccountParams{
		KeyID:                 keyID,
		ShortID:               shortID,
		SubscriptionExpiresAt: pgtype.Timestamptz{Valid: true, Time: expiresAt},
		Region:                "ru",
		Discount:              pgtype.Int4{Valid: true, Int32: int32(discount)},
	}
	if promo != nil {
		createAccountParams.Promo = pgtype.Text{Valid: true, String: *promo}
	}

	newAccountID, err := db.queries.CreateAccount(ctx, createAccountParams)
	if err != nil {
		return 0, err
	}

	if err := db.queries.CreatePlatformAccount(ctx, sqlc.CreatePlatformAccountParams{
		PlatformID:        int16(platform),
		ExternalAccountID: int64(platformAccountID),
		FkAccountID:       int32(newAccountID),
	}); err != nil {
		return 0, err
	}

	return int(newAccountID), txn.Commit(ctx)
}

func (db accountDB) AddSubscriptionExpiresAt(ctx context.Context, accountID int, duration time.Duration) error {
	return db.queries.AddSubscriptionExpiresAt(ctx, sqlc.AddSubscriptionExpiresAtParams{
		ID:   int32(accountID),
		Secs: duration.Seconds(),
	})
}

func (db accountDB) RemoveSubscriptionExpiresAt(ctx context.Context, accountID int, duration time.Duration) error {
	return db.queries.RemoveSubscriptionExpiresAt(ctx, sqlc.RemoveSubscriptionExpiresAtParams{
		ID:   int32(accountID),
		Secs: duration.Seconds(),
	})
}

func (db accountDB) CancelSubscriptions(ctx context.Context, accountID int) error {
	return db.queries.CancelSubscriptions(ctx, int32(accountID))
}

func (db accountDB) SetDiscount(ctx context.Context, accountID int, discount int) error {
	return db.queries.SetDiscount(ctx, sqlc.SetDiscountParams{
		ID:       int32(accountID),
		Discount: pgtype.Int4{Valid: true, Int32: int32(discount)},
	})
}

func (db accountDB) ResetDiscount(ctx context.Context, accountID int) error {
	return db.queries.ResetDiscount(ctx, int32(accountID))
}

func (db accountDB) RegenerateKey(ctx context.Context, accountID int, keyID string) error {
	return db.queries.RegenerateKey(ctx, sqlc.RegenerateKeyParams{
		ID:    int32(accountID),
		KeyID: keyID,
	})
}

func (db accountDB) SetRegion(ctx context.Context, accountID int, region string) error {
	return db.queries.SetRegion(ctx, sqlc.SetRegionParams{
		ID:     int32(accountID),
		Region: region,
	})
}

func (db accountDB) SetLanguage(ctx context.Context, accountID int, language string) error {
	return db.queries.SetLanguage(ctx, sqlc.SetLanguageParams{
		ID:       int32(accountID),
		Language: language,
	})
}

func (db accountDB) SetCurrency(ctx context.Context, accountID int, currency string) error {
	return db.queries.SetCurrency(ctx, sqlc.SetCurrencyParams{
		ID:       int32(accountID),
		Currency: currency,
	})
}

func (db accountDB) SetProtocol(ctx context.Context, accountID int, protocol string) error {
	return db.queries.SetProtocol(ctx, sqlc.SetProtocolParams{
		ID:       int32(accountID),
		Protocol: protocol,
	})
}

func (db accountDB) SetLocation(ctx context.Context, accountID int, location string) error {
	return db.queries.SetLocation(ctx, sqlc.SetLocationParams{
		ID:             int32(accountID),
		ServerLocation: location,
	})
}
