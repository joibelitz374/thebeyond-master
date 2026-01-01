package mocks

import (
	"context"
	"time"

	"github.com/quickpowered/thebeyond-master/internal/domain"
	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager/dto"
	"github.com/quickpowered/thebeyond-master/pkg/consts"
	"github.com/stretchr/testify/mock"
)

type MockAccountService struct {
	mock.Mock
}

func (m *MockAccountService) Get(ctx context.Context, platform consts.Platform, platformAccountID int) (domain.Account, error) {
	args := m.Called(ctx, platform, platformAccountID)
	return args.Get(0).(domain.Account), args.Error(1)
}

func (m *MockAccountService) GetAllByPlatform(ctx context.Context, platform consts.Platform) ([]int, error) {
	args := m.Called(ctx, platform)
	return args.Get(0).([]int), args.Error(1)
}

func (m *MockAccountService) GetByAccountID(ctx context.Context, accountID int) (domain.Account, error) {
	args := m.Called(ctx, accountID)
	return args.Get(0).(domain.Account), args.Error(1)
}

func (m *MockAccountService) GetByKeyID(ctx context.Context, keyID string) (domain.Account, error) {
	args := m.Called(ctx, keyID)
	return args.Get(0).(domain.Account), args.Error(1)
}

func (m *MockAccountService) GetPlatformUserID(ctx context.Context, accountID int) (int, error) {
	args := m.Called(ctx, accountID)
	return args.Int(0), args.Error(1)
}

func (m *MockAccountService) GetExpiredUsers(ctx context.Context) ([]int, error) {
	args := m.Called(ctx)
	return args.Get(0).([]int), args.Error(1)
}

func (m *MockAccountService) GetRegions(region string) ([]dto.Region, error) {
	args := m.Called(region)
	return args.Get(0).([]dto.Region), args.Error(1)
}

func (m *MockAccountService) FindUsersForServiceCheck(ctx context.Context) ([]domain.ServiceCheck, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.ServiceCheck), args.Error(1)
}

func (m *MockAccountService) Create(ctx context.Context, platform consts.Platform, platformAccountID int, promo *string, discount int) (int, string, error) {
	args := m.Called(ctx, platform, platformAccountID, promo, discount)
	return args.Int(0), args.String(1), args.Error(2)
}

func (m *MockAccountService) MarkServiceCheckSent(ctx context.Context, accountID int) error {
	args := m.Called(ctx, accountID)
	return args.Error(0)
}

func (m *MockAccountService) AddSubscriptionExpiresAt(ctx context.Context, accountID int, duration time.Duration) error {
	args := m.Called(ctx, accountID, duration)
	return args.Error(0)
}

func (m *MockAccountService) RemoveSubscriptionExpiresAt(ctx context.Context, accountID int, duration time.Duration) error {
	args := m.Called(ctx, accountID, duration)
	return args.Error(0)
}

func (m *MockAccountService) CancelSubscriptions(ctx context.Context, accountID int) error {
	args := m.Called(ctx, accountID)
	return args.Error(0)
}

func (m *MockAccountService) SetDiscount(ctx context.Context, accountID int, discount int) error {
	args := m.Called(ctx, accountID, discount)
	return args.Error(0)
}

func (m *MockAccountService) ResetDiscount(ctx context.Context, accountID int) error {
	args := m.Called(ctx, accountID)
	return args.Error(0)
}

func (m *MockAccountService) RegenerateKey(ctx context.Context, accountID int) (string, error) {
	args := m.Called(ctx, accountID)
	return args.String(0), args.Error(1)
}

func (m *MockAccountService) SetRegion(ctx context.Context, accountID int, region string) error {
	args := m.Called(ctx, accountID, region)
	return args.Error(0)
}

func (m *MockAccountService) SetLanguage(ctx context.Context, accountID int, language string) error {
	args := m.Called(ctx, accountID, language)
	return args.Error(0)
}

func (m *MockAccountService) SetCurrency(ctx context.Context, accountID int, currency string) error {
	args := m.Called(ctx, accountID, currency)
	return args.Error(0)
}

func (m *MockAccountService) SetProtocol(ctx context.Context, accountID int, protocol string) error {
	args := m.Called(ctx, accountID, protocol)
	return args.Error(0)
}

func (m *MockAccountService) SetLocation(ctx context.Context, accountID int, location string) error {
	args := m.Called(ctx, accountID, location)
	return args.Error(0)
}

func (m *MockAccountService) SetTariff(ctx context.Context, accountID int, tariff int) error {
	args := m.Called(ctx, accountID, tariff)
	return args.Error(0)
}

func (m *MockAccountService) AddTraffic(ctx context.Context, accountID int, uplink, downlink int64) error {
	args := m.Called(ctx, accountID, uplink, downlink)
	return args.Error(0)
}

func (m *MockAccountService) StartFreemiumPeriod(ctx context.Context, accountID int) error {
	args := m.Called(ctx, accountID)
	return args.Error(0)
}
