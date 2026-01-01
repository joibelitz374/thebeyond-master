package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/quickpowered/thebeyond-master/internal/domain"
	"github.com/quickpowered/thebeyond-master/pkg/sqlc"
)

type SubscriptionInterface interface {
	GetNeedRefreshTraffic(ctx context.Context) ([]int, error)
	GetAccountsToDisable(ctx context.Context) ([]int, error)
	GetAccountsToEnable(ctx context.Context) ([]domain.Account, error)
	ResetLastTrafficRefreshAt(ctx context.Context, accountID int) error
	GetFremiumAccounts(ctx context.Context) ([]domain.ExternalAccount, error)
	StartFreemium(ctx context.Context, accountID int) error
	RemoveFreemium(ctx context.Context, accountID int) error
	SetExpiresAt(ctx context.Context, accountID int, duration time.Duration) error
	SetTariff(ctx context.Context, accountID int, tariff int) error
	AddExpiresAt(ctx context.Context, accountID int, duration time.Duration) error
	RemoveExpiresAt(ctx context.Context, accountID int, duration time.Duration) error
	Cancel(ctx context.Context, accountID int) error
	SetDiscount(ctx context.Context, accountID int, discount int) error
	ResetDiscount(ctx context.Context, accountID int) error
}

type subscriptionDB struct {
	pool    *pgxpool.Pool
	queries *sqlc.Queries
}

func NewSubscription(pool *pgxpool.Pool) SubscriptionInterface {
	return subscriptionDB{pool, sqlc.New(pool)}
}

func (db subscriptionDB) GetNeedRefreshTraffic(ctx context.Context) ([]int, error) {
	rows, err := db.queries.GetNeedRefreshTraffic(ctx)
	if err != nil {
		return nil, err
	}

	accountIDs := make([]int, len(rows))
	for i, row := range rows {
		accountIDs[i] = int(row)
	}

	return accountIDs, nil
}

func (db subscriptionDB) GetAccountsToDisable(ctx context.Context) ([]int, error) {
	rows, err := db.queries.GetAccountsToDisable(ctx)
	if err != nil {
		return nil, err
	}

	accountIDs := make([]int, len(rows))
	for i, row := range rows {
		accountIDs[i] = int(row)
	}

	return accountIDs, nil
}

func (db subscriptionDB) GetAccountsToEnable(ctx context.Context) ([]domain.Account, error) {
	rows, err := db.queries.GetAccountsToEnable(ctx)
	if err != nil {
		return nil, err
	}

	accounts := make([]domain.Account, len(rows))
	for i, row := range rows {
		accounts[i] = domain.Account{
			ID:    int(row.ID),
			KeyID: row.KeyID,
		}
	}

	return accounts, nil
}

func (db subscriptionDB) ResetLastTrafficRefreshAt(ctx context.Context, accountID int) error {
	return db.queries.ResetLastTrafficRefreshAt(ctx, int32(accountID))
}

func (db subscriptionDB) GetFremiumAccounts(ctx context.Context) ([]domain.ExternalAccount, error) {
	rows, err := db.queries.GetFremiumAccounts(ctx)
	if err != nil {
		return nil, err
	}

	accounts := make([]domain.ExternalAccount, len(rows))
	for i, account := range rows {
		accounts[i] = domain.ExternalAccount{
			ID:                int(account.ID),
			ExternalAccountID: int(account.ExternalAccountID),
		}
	}

	return accounts, nil
}

func (db subscriptionDB) StartFreemium(ctx context.Context, accountID int) error {
	return db.queries.StartFreemium(ctx, int32(accountID))
}

func (db subscriptionDB) RemoveFreemium(ctx context.Context, accountID int) error {
	return db.queries.RemoveFreemium(ctx, int32(accountID))
}

func (db subscriptionDB) SetTariff(ctx context.Context, accountID int, tariff int) error {
	return db.queries.SetTariff(ctx, sqlc.SetTariffParams{
		ID:     int32(accountID),
		Tariff: int16(tariff),
	})
}

func (db subscriptionDB) SetExpiresAt(ctx context.Context, accountID int, duration time.Duration) error {
	expiresAt := time.Now().UTC().Add(duration)
	return db.queries.SetSubscriptionExpiresAt(ctx, sqlc.SetSubscriptionExpiresAtParams{
		ID:                    int32(accountID),
		SubscriptionExpiresAt: pgtype.Timestamptz{Time: expiresAt, Valid: true},
	})
}

func (db subscriptionDB) AddExpiresAt(ctx context.Context, accountID int, duration time.Duration) error {
	return db.queries.AddSubscriptionExpiresAt(ctx, sqlc.AddSubscriptionExpiresAtParams{
		ID:   int32(accountID),
		Secs: duration.Seconds(),
	})
}

func (db subscriptionDB) RemoveExpiresAt(ctx context.Context, accountID int, duration time.Duration) error {
	return db.queries.RemoveSubscriptionExpiresAt(ctx, sqlc.RemoveSubscriptionExpiresAtParams{
		ID:   int32(accountID),
		Secs: duration.Seconds(),
	})
}

func (db subscriptionDB) Cancel(ctx context.Context, accountID int) error {
	return db.queries.CancelSubscriptions(ctx, int32(accountID))
}

func (db subscriptionDB) SetDiscount(ctx context.Context, accountID int, discount int) error {
	return db.queries.SetDiscount(ctx, sqlc.SetDiscountParams{
		ID:       int32(accountID),
		Discount: pgtype.Int4{Valid: true, Int32: int32(discount)},
	})
}

func (db subscriptionDB) ResetDiscount(ctx context.Context, accountID int) error {
	return db.queries.ResetDiscount(ctx, int32(accountID))
}
