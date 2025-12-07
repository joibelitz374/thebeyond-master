package promo

import (
	"context"

	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/repositories/db"
)

type Interface interface {
	Create(ctx context.Context, name string, creator int) error
	Get(ctx context.Context, name string) (promo domain.Promo, err error)
	GetByCreator(ctx context.Context, creator int) (promos []domain.Promo, err error)
	UpgradeLevel(ctx context.Context, name string) error
	CheckBuyer(ctx context.Context, name string, buyerID int) error
	AddBuyer(ctx context.Context, name string, buyerID int) error
	IncreaseClients(ctx context.Context, name string) error
	IncreaseBuyers(ctx context.Context, name string) error
}

type service struct {
	repo db.PromoInterface
}

func NewService(repo db.PromoInterface) Interface {
	return &service{repo}
}

func (s *service) Create(ctx context.Context, name string, creator int) error {
	return s.repo.Create(ctx, name, creator)
}

func (s *service) Get(ctx context.Context, name string) (promo domain.Promo, err error) {
	return s.repo.Get(ctx, name)
}

func (s *service) GetByCreator(ctx context.Context, creator int) (promos []domain.Promo, err error) {
	return s.repo.GetByCreator(ctx, creator)
}

func (s *service) UpgradeLevel(ctx context.Context, name string) error {
	return s.repo.UpgradeLevel(ctx, name)
}

func (s *service) CheckBuyer(ctx context.Context, name string, buyerID int) error {
	return s.repo.CheckBuyer(ctx, name, buyerID)
}

func (s *service) AddBuyer(ctx context.Context, name string, buyerID int) error {
	return s.repo.AddBuyer(ctx, name, buyerID)
}

func (s *service) IncreaseClients(ctx context.Context, name string) error {
	return s.repo.IncreaseClients(ctx, name)
}

func (s *service) IncreaseBuyers(ctx context.Context, name string) error {
	return s.repo.IncreaseBuyers(ctx, name)
}
