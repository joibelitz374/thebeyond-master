package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/quickpowered/thebeyond-master/internal/domain"
	"github.com/quickpowered/thebeyond-master/pkg/sqlc"
)

type PaymentInterface interface {
	Create(ctx context.Context, accountID, amount int, expiresAt time.Time) error
	Get(ctx context.Context, accountID int, expiresAt time.Time) ([]domain.Payment, error)
	Delete(ctx context.Context, accountID int) error
}

type paymentDB struct {
	pool    *pgxpool.Pool
	queries *sqlc.Queries
}

func NewPayment(pool *pgxpool.Pool) PaymentInterface {
	return paymentDB{pool, sqlc.New(pool)}
}

func (db paymentDB) Create(ctx context.Context, accountID, amount int, expiresAt time.Time) error {
	return db.queries.CreatePayment(ctx, sqlc.CreatePaymentParams{
		AccountID: int32(accountID),
		Amount:    int32(amount),
		ExpiresAt: pgtype.Timestamptz{Valid: true, Time: expiresAt},
	})
}

func (db paymentDB) Get(ctx context.Context, accountID int, expiresAt time.Time) ([]domain.Payment, error) {
	rows, err := db.queries.GetPayments(ctx, int32(accountID))
	if err != nil {
		return nil, err
	}

	payments := make([]domain.Payment, len(rows))
	for i, payment := range rows {
		payments[i] = domain.Payment{
			Amount:    int(payment.Amount),
			CreatedAt: payment.CreatedAt.Time,
			ExpiresAt: payment.ExpiresAt.Time,
		}
	}

	return payments, err
}

func (db paymentDB) Delete(ctx context.Context, accountID int) error {
	return db.queries.DeletePayment(ctx, int32(accountID))
}
