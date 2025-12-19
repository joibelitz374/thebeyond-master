package account

import (
	"context"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/quickpowered/thebeyond-master/internal/domain"
	"github.com/quickpowered/thebeyond-master/internal/repositories/db"
	"github.com/quickpowered/thebeyond-master/pkg/consts"
	"github.com/quickpowered/thebeyond-master/pkg/values"
)

type Interface interface {
	Get(ctx context.Context, platform consts.Platform, platformAccountID int) (domain.Account, error)
	GetAllByPlatform(ctx context.Context, platform consts.Platform) (accounts []int, err error)
	GetByAccountID(ctx context.Context, accountID int) (account domain.Account, err error)
	GetByKeyID(ctx context.Context, keyID string) (domain.Account, error)
	GetPlatformUserID(ctx context.Context, accountID int) (externalAccountID int, err error)
	GetExpiredUsers(ctx context.Context) ([]int, error)
	FindUsersForServiceCheck(ctx context.Context) (accounts []domain.ServiceCheck, err error)
	Create(ctx context.Context, platform consts.Platform, platformAccountID int, expiresAt time.Time, promo *string, discount int) (int, string, error)
	MarkServiceCheckSent(ctx context.Context, accountID int) error
	AddSubscriptionExpiresAt(ctx context.Context, accountID int, duration time.Duration) error
	RemoveSubscriptionExpiresAt(ctx context.Context, accountID int, duration time.Duration) error
	CancelSubscriptions(ctx context.Context, accountID int) error
	SetDiscount(ctx context.Context, accountID int, discount int) error
	ResetDiscount(ctx context.Context, accountID int) error
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

func (s service) GetByAccountID(ctx context.Context, accountID int) (account domain.Account, err error) {
	return s.repo.GetByAccountID(ctx, accountID)
}

func (s service) GetByKeyID(ctx context.Context, keyID string) (account domain.Account, err error) {
	return s.repo.GetByKeyID(ctx, keyID)
}

func (s service) GetPlatformUserID(ctx context.Context, accountID int) (externalAccountID int, err error) {
	return s.repo.GetPlatformUserID(ctx, accountID)
}

func (s service) GetExpiredUsers(ctx context.Context) ([]int, error) {
	return s.repo.GetExpiredUsers(ctx)
}

func (s service) GetAllByPlatform(ctx context.Context, platform consts.Platform) (accounts []int, err error) {
	return s.repo.GetAllByPlatform(ctx, platform)
}

func (s service) FindUsersForServiceCheck(ctx context.Context) (accounts []domain.ServiceCheck, err error) {
	return s.repo.FindUsersForServiceCheck(ctx)
}

func (s service) Create(ctx context.Context, platform consts.Platform, platformAccountID int, expiresAt time.Time, promo *string, discount int) (int, string, error) {
	keyID, err := gonanoid.Generate(consts.NANOID_ALPHABET, consts.NANOID_LENGTH)
	if err != nil {
		return 0, "", err
	}

	shortID := values.RandomShortId()

	accountID, err := s.repo.Create(ctx, platform, platformAccountID, keyID, shortID, expiresAt, promo, discount)
	return accountID, keyID, err
}

func (s service) MarkServiceCheckSent(ctx context.Context, accountID int) error {
	return s.repo.MarkServiceCheckSent(ctx, accountID)
}

func (s service) AddSubscriptionExpiresAt(ctx context.Context, accountID int, duration time.Duration) error {
	return s.repo.AddSubscriptionExpiresAt(ctx, accountID, duration)
}

func (s service) RemoveSubscriptionExpiresAt(ctx context.Context, accountID int, duration time.Duration) error {
	return s.repo.RemoveSubscriptionExpiresAt(ctx, accountID, duration)
}

func (s service) CancelSubscriptions(ctx context.Context, accountID int) error {
	return s.repo.CancelSubscriptions(ctx, accountID)
}

func (s service) SetDiscount(ctx context.Context, accountID int, discount int) error {
	return s.repo.SetDiscount(ctx, accountID, discount)
}

func (s service) ResetDiscount(ctx context.Context, accountID int) error {
	return s.repo.ResetDiscount(ctx, accountID)
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
