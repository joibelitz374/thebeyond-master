package subscription

import (
	"context"
	"errors"
	"time"

	"github.com/quickpowered/thebeyond-master/internal/domain"
	"github.com/quickpowered/thebeyond-master/internal/repositories/db"
	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager/dto"
)

type Interface interface {
	GetNeedRefreshTraffic(ctx context.Context) ([]int, error)
	GetAccountsToDisable(ctx context.Context) ([]int, error)
	GetAccountsToEnable(ctx context.Context) ([]domain.Account, error)
	EnableAccount(ctx context.Context, accountID int) error
	GetRegions(region string) ([]dto.Region, error)
	ResetLastTrafficRefreshAt(ctx context.Context, accountID int) error
	GetFremiumAccounts(ctx context.Context) ([]domain.ExternalAccount, error)
	StartFreemium(ctx context.Context, accountID int) error
	RemoveFreemium(ctx context.Context, accountID int) error
	SetTariff(ctx context.Context, accountID int, tariff int) error
	SetExpiresAt(ctx context.Context, accountID int, duration time.Duration) error
	AddExpiresAt(ctx context.Context, accountID int, duration time.Duration) error
	RemoveExpiresAt(ctx context.Context, accountID int, duration time.Duration) error
	Cancel(ctx context.Context, accountID int) error
	SetDiscount(ctx context.Context, accountID int, discount int) error
	ResetDiscount(ctx context.Context, accountID int) error
}

type service struct {
	repo db.SubscriptionInterface
}

func NewService(repo db.SubscriptionInterface) Interface {
	return service{repo}
}

func (s service) GetNeedRefreshTraffic(ctx context.Context) ([]int, error) {
	return s.repo.GetNeedRefreshTraffic(ctx)
}

func (s service) GetAccountsToDisable(ctx context.Context) ([]int, error) {
	return s.repo.GetAccountsToDisable(ctx)
}

func (s service) GetAccountsToEnable(ctx context.Context) ([]domain.Account, error) {
	return s.repo.GetAccountsToEnable(ctx)
}

func (s service) GetRegions(region string) ([]dto.Region, error) {
	regions := []dto.Region{}
	switch region {
	case "ru":
		regions = append(regions, dto.RegionRussia)
	default:
		return nil, errors.New("region not found")
	}
	return regions, nil
}

func (s service) EnableAccount(ctx context.Context, accountID int) error {
	return s.repo.EnableAccount(ctx, accountID)
}

func (s service) ResetLastTrafficRefreshAt(ctx context.Context, accountID int) error {
	return s.repo.ResetLastTrafficRefreshAt(ctx, accountID)
}

func (s service) GetFremiumAccounts(ctx context.Context) ([]domain.ExternalAccount, error) {
	return s.repo.GetFremiumAccounts(ctx)
}

func (s service) StartFreemium(ctx context.Context, accountID int) error {
	return s.repo.StartFreemium(ctx, accountID)
}

func (s service) RemoveFreemium(ctx context.Context, accountID int) error {
	return s.repo.RemoveFreemium(ctx, accountID)
}

func (s service) SetTariff(ctx context.Context, accountID int, tariff int) error {
	return s.repo.SetTariff(ctx, accountID, tariff)
}

func (s service) SetExpiresAt(ctx context.Context, accountID int, duration time.Duration) error {
	return s.repo.SetExpiresAt(ctx, accountID, duration)
}

func (s service) AddExpiresAt(ctx context.Context, accountID int, duration time.Duration) error {
	return s.repo.AddExpiresAt(ctx, accountID, duration)
}

func (s service) RemoveExpiresAt(ctx context.Context, accountID int, duration time.Duration) error {
	return s.repo.RemoveExpiresAt(ctx, accountID, duration)
}

func (s service) Cancel(ctx context.Context, accountID int) error {
	return s.repo.Cancel(ctx, accountID)
}

func (s service) SetDiscount(ctx context.Context, accountID int, discount int) error {
	return s.repo.SetDiscount(ctx, accountID, discount)
}

func (s service) ResetDiscount(ctx context.Context, accountID int) error {
	return s.repo.ResetDiscount(ctx, accountID)
}
