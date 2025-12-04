package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/quickpowered/frilly/internal/domain"
)

type PaymentInterface interface {
	Create(ctx context.Context, accountID, amount int, expiresAt time.Time) error
	Get(ctx context.Context, accountID int, expiresAt time.Time) (payments []domain.Payment, err error)
	Delete(ctx context.Context, accountID int) error
}

type paymentDB struct {
	pool *pgxpool.Pool
}

func NewPayment(pool *pgxpool.Pool) PaymentInterface {
	return paymentDB{pool}
}

func (db paymentDB) Create(ctx context.Context, accountID, amount int, expiresAt time.Time) error {
	_, err := db.pool.Exec(ctx, INSERT_PAYMENT_SQL, accountID, amount, expiresAt)
	return err
}

func (db paymentDB) Get(ctx context.Context, accountID int, expiresAt time.Time) (payments []domain.Payment, err error) {
	rows, err := db.pool.Query(ctx, GET_PAYMENTS_SQL, accountID)
	if err != nil {
		return nil, err
	}

	payments = make([]domain.Payment, 0)
	for rows.Next() {
		payment := domain.Payment{}
		if err = rows.Scan(&payment.Amount, &payment.CreatedAt, &payment.ExpiresAt); err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	return payments, err
}

func (db paymentDB) Delete(ctx context.Context, accountID int) error {
	_, err := db.pool.Exec(ctx, DELETE_PAYMENT_SQL, accountID)
	return err
}
