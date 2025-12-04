package payment

import (
	"context"
	"time"

	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/repositories/db"
)

type Interface interface {
	Create(ctx context.Context, accountID, amount int, expiresAt time.Time) error
	Get(ctx context.Context, accountID int, expiresAt time.Time) (payments []domain.Payment, err error)
	Delete(ctx context.Context, accountID int) error
}

type service struct {
	repo db.PaymentInterface
}

func NewService(repo db.PaymentInterface) Interface {
	return &service{repo}
}

func (s *service) Create(ctx context.Context, accountID, amount int, expiresAt time.Time) error {
	return s.repo.Create(ctx, accountID, amount, expiresAt)
}

func (s *service) Get(ctx context.Context, accountID int, expiresAt time.Time) (payments []domain.Payment, err error) {
	return s.repo.Get(ctx, accountID, expiresAt)
}

func (s *service) Delete(ctx context.Context, accountID int) error {
	return s.repo.Delete(ctx, accountID)
}
