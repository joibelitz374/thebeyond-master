package account

import (
	"context"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/repositories/db"
	"github.com/quickpowered/frilly/pkg/consts"
	"github.com/quickpowered/frilly/pkg/values"
)

type Interface interface {
	Get(ctx context.Context, platform consts.Platform, platformAccountID int) (domain.Account, error)
	GetExpiredUsers(ctx context.Context) ([]domain.Account, error)
	Create(ctx context.Context, platform consts.Platform, platformAccountID int, expiresAt time.Time) (int, string, error)
	SetSubscriptionExpiresAt(ctx context.Context, accountID int, expiresAt *time.Time) error
	RegenerateKey(ctx context.Context, accountID int) (string, error)
	SetRegion(ctx context.Context, accountID int, region string) error
	SetLanguage(ctx context.Context, accountID int, language string) error
	SetCurrency(ctx context.Context, accountID int, currency string) error
	SetProtocol(ctx context.Context, accountID int, protocol string) error
	SetLocation(ctx context.Context, accountID int, location string) error
}

type service struct {
	repo db.AccountInterface
}

func NewService(repo db.AccountInterface) Interface {
	return service{repo}
}

func (s service) Get(ctx context.Context, platform consts.Platform, platformAccountID int) (domain.Account, error) {
	return s.repo.Get(ctx, platform, platformAccountID)
}

func (s service) GetExpiredUsers(ctx context.Context) ([]domain.Account, error) {
	return s.repo.GetExpiredUsers(ctx)
}

func (s service) Create(ctx context.Context, platform consts.Platform, platformAccountID int, expiresAt time.Time) (int, string, error) {
	keyID, err := gonanoid.Generate(consts.NANOID_ALPHABET, consts.NANOID_LENGTH)
	if err != nil {
		return 0, "", err
	}

	shortID := values.RandomShortId()

	accountID, err := s.repo.Create(ctx, platform, platformAccountID, keyID, shortID, expiresAt)
	return accountID, keyID, err
}

func (s service) SetSubscriptionExpiresAt(ctx context.Context, accountID int, expiresAt *time.Time) error {
	return s.repo.SetSubscriptionExpiresAt(ctx, accountID, expiresAt)
}

func (s service) RegenerateKey(ctx context.Context, accountID int) (string, error) {
	keyID, err := gonanoid.Generate(consts.NANOID_ALPHABET, consts.NANOID_LENGTH)
	if err != nil {
		return "", err
	}

	return keyID, s.repo.RegenerateKey(ctx, accountID, keyID)
}

func (s service) SetRegion(ctx context.Context, accountID int, region string) error {
	return s.repo.SetRegion(ctx, accountID, region)
}

func (s service) SetLanguage(ctx context.Context, accountID int, language string) error {
	return s.repo.SetLanguage(ctx, accountID, language)
}

func (s service) SetCurrency(ctx context.Context, accountID int, currency string) error {
	return s.repo.SetCurrency(ctx, accountID, currency)
}

func (s service) SetProtocol(ctx context.Context, accountID int, protocol string) error {
	return s.repo.SetProtocol(ctx, accountID, protocol)
}

func (s service) SetLocation(ctx context.Context, accountID int, location string) error {
	return s.repo.SetLocation(ctx, accountID, location)
}
