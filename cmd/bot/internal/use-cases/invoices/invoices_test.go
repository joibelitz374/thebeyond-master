package invoices_test

// import (
// 	"errors"
// 	"fmt"
// 	"testing"
// 	"time"

// 	"github.com/go-telegram/bot/models"
// 	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/invoices"
// 	ccurrency "github.com/quickpowered/thebeyond-master/configs/currency"
// 	"github.com/quickpowered/thebeyond-master/internal/domain"
// 	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager/dto"
// 	"github.com/quickpowered/thebeyond-master/pkg/consts"
// 	"github.com/quickpowered/thebeyond-master/pkg/logger"
// 	"github.com/quickpowered/thebeyond-master/pkg/mocks"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func createPreCheckoutUpdate(currency, payload string, userID int64) *models.Update {
// 	return &models.Update{
// 		PreCheckoutQuery: &models.PreCheckoutQuery{
// 			ID:             "test-query-id",
// 			Currency:       currency,
// 			InvoicePayload: payload,
// 			From:           &models.User{ID: userID},
// 		},
// 	}
// }

// var errTestComplete = errors.New("test complete - stopping before bot call")

// func TestHandlePreCheckoutQuery_InvalidCurrency(t *testing.T) {
// 	accountMock := new(mocks.MockAccountService)
// 	paymentMock := new(mocks.MockPaymentService)
// 	promoMock := new(mocks.MockPromoUseCase)
// 	tariffsMock := new(mocks.MockSubscriptionsRepo)
// 	xrayMock := new(mocks.MockXrayManagerRepo)
// 	uc := invoices.NewUseCase(accountMock, paymentMock, promoMock, tariffsMock, xrayMock, logger.New("debug", "development"))

// 	update := createPreCheckoutUpdate("USD", "d:30", 12345)
// 	err := uc.HandlePreCheckoutQuery(nil, update)

// 	assert.Error(t, err)
// 	assert.Equal(t, "invalid currency", err.Error())
// }

// func TestHandlePreCheckoutQuery_InvalidPayload(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		payload string
// 	}{
// 		{"empty payload", ""},
// 		{"single value", "30"},
// 		{"wrong prefix", "x:30"},
// 		{"missing days value", "d:"},
// 		{"too many parts", "d:30;t:2;extra"},
// 		{"wrong second prefix", "d:30;x:2"},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			accountMock := new(mocks.MockAccountService)
// 			paymentMock := new(mocks.MockPaymentService)
// 			promoMock := new(mocks.MockPromoUseCase)
// 			subsMock := new(mocks.MockSubscriptionsRepo)
// 			xrayMock := new(mocks.MockXrayManagerRepo)

// 			subsMock.On("GetPeriod", mock.Anything).Return(domain.Period{}, false)

// 			uc := invoices.NewUseCase(accountMock, paymentMock, promoMock, subsMock, xrayMock, logger.New("debug", "development"))

// 			update := createPreCheckoutUpdate("XTR", tt.payload, 12345)
// 			err := uc.HandlePreCheckoutQuery(nil, update)

// 			assert.Error(t, err, "Expected error for payload '%s', got nil", tt.payload)
// 		})
// 	}
// }

// func TestHandlePreCheckoutQuery_InvalidDays(t *testing.T) {
// 	accountMock := new(mocks.MockAccountService)
// 	paymentMock := new(mocks.MockPaymentService)
// 	promoMock := new(mocks.MockPromoUseCase)
// 	subsMock := new(mocks.MockSubscriptionsRepo)
// 	xrayMock := new(mocks.MockXrayManagerRepo)

// 	subsMock.On("GetPeriod", mock.Anything).Return(domain.Period{}, false)
// 	uc := invoices.NewUseCase(accountMock, paymentMock, promoMock, subsMock, xrayMock, logger.New("debug", "development"))

// 	update := createPreCheckoutUpdate("XTR", "d:abc", 12345)
// 	err := uc.HandlePreCheckoutQuery(nil, update)

// 	assert.Error(t, err, "Expected error for non-numeric days, got nil")
// }

// func TestHandlePreCheckoutQuery_InvalidTariffID(t *testing.T) {
// 	accountMock := new(mocks.MockAccountService)
// 	paymentMock := new(mocks.MockPaymentService)
// 	promoMock := new(mocks.MockPromoUseCase)
// 	subsMock := new(mocks.MockSubscriptionsRepo)
// 	xrayMock := new(mocks.MockXrayManagerRepo)

// 	subsMock.On("GetPeriod", mock.Anything).Return(domain.Period{}, false)
// 	uc := invoices.NewUseCase(accountMock, paymentMock, promoMock, subsMock, xrayMock, logger.New("debug", "development"))

// 	update := createPreCheckoutUpdate("XTR", "d:30;t:abc", 12345)
// 	err := uc.HandlePreCheckoutQuery(nil, update)

// 	assert.Error(t, err, "Expected error for non-numeric tariff ID, got nil")
// }

// func TestHandlePreCheckoutQuery_SubscriptionNotFound(t *testing.T) {
// 	accountMock := new(mocks.MockAccountService)
// 	paymentMock := new(mocks.MockPaymentService)
// 	promoMock := new(mocks.MockPromoUseCase)
// 	subsMock := new(mocks.MockSubscriptionsRepo)
// 	xrayMock := new(mocks.MockXrayManagerRepo)

// 	// Mock: subscription not found
// 	subsMock.On("GetPeriod", 45).Return(domain.Period{}, false)

// 	uc := invoices.NewUseCase(accountMock, paymentMock, promoMock, subsMock, xrayMock, logger.New("debug", "development"))

// 	update := createPreCheckoutUpdate("XTR", "d:45", 12345)
// 	err := uc.HandlePreCheckoutQuery(nil, update)

// 	assert.Error(t, err)
// 	assert.Equal(t, "invalid subscription", err.Error())
// 	subsMock.AssertExpectations(t)
// }

// func TestHandlePreCheckoutQuery_InvalidTariffPeriod(t *testing.T) {
// 	accountMock := new(mocks.MockAccountService)
// 	paymentMock := new(mocks.MockPaymentService)
// 	promoMock := new(mocks.MockPromoUseCase)
// 	subsMock := new(mocks.MockSubscriptionsRepo)
// 	xrayMock := new(mocks.MockXrayManagerRepo)

// 	subsMock.On("GetPeriod", 45).Return(domain.Period{}, false)
// 	uc := invoices.NewUseCase(accountMock, paymentMock, promoMock, subsMock, xrayMock, logger.New("debug", "development"))

// 	// 45 days is not a valid period (only 30, 90, 180, 365)
// 	update := createPreCheckoutUpdate("XTR", "d:45;t:2", 12345)
// 	err := uc.HandlePreCheckoutQuery(nil, update)

// 	assert.Error(t, err, "Expected error for invalid period, got nil")
// }

// func TestHandlePreCheckoutQuery_InvalidTariffPlan(t *testing.T) {
// 	accountMock := new(mocks.MockAccountService)
// 	paymentMock := new(mocks.MockPaymentService)
// 	promoMock := new(mocks.MockPromoUseCase)
// 	subsMock := new(mocks.MockSubscriptionsRepo)
// 	xrayMock := new(mocks.MockXrayManagerRepo)

// 	accountMock.On("Get", mock.Anything, consts.PlatformTelegram, 12345).
// 		Return(domain.Account{ID: 1, KeyID: "test-key"}, nil)

// 	subsMock.On("GetPeriod", 30).Return(domain.Period{Emoji: "🚀"}, true)
// 	subsMock.On("Get", 99).Return(domain.Tariff{}, false)

// 	uc := invoices.NewUseCase(accountMock, paymentMock, promoMock, subsMock, xrayMock, logger.New("debug", "development"))

// 	// Tariff 99 doesn't exist
// 	update := createPreCheckoutUpdate("XTR", "d:30;t:99", 12345)
// 	err := uc.HandlePreCheckoutQuery(nil, update)

// 	assert.Error(t, err)
// 	assert.Equal(t, "invalid tariff", err.Error())
// }

// func TestHandlePreCheckoutQuery_AccountServiceError(t *testing.T) {
// 	expectedErr := errors.New("database connection failed")

// 	accountMock := new(mocks.MockAccountService)
// 	paymentMock := new(mocks.MockPaymentService)
// 	promoMock := new(mocks.MockPromoUseCase)
// 	subsMock := new(mocks.MockSubscriptionsRepo)
// 	xrayMock := new(mocks.MockXrayManagerRepo)

// 	accountMock.On("Get", mock.Anything, consts.PlatformTelegram, 12345).
// 		Return(domain.Account{}, expectedErr)
// 	subsMock.On("GetPeriod", 30).Return(domain.Period{Emoji: "🚀"}, true)
// 	subsMock.On("Get", 1).Return(domain.Tariff{Prices: map[ccurrency.Currency]float64{ccurrency.XTR: 99.0}}, true)
// 	uc := invoices.NewUseCase(accountMock, paymentMock, promoMock, subsMock, xrayMock, logger.New("debug", "development"))

// 	update := createPreCheckoutUpdate("XTR", "d:30", 12345)
// 	err := uc.HandlePreCheckoutQuery(nil, update)

// 	assert.Equal(t, expectedErr, err)
// }

// func TestHandlePreCheckoutQuery_PaymentServiceError(t *testing.T) {
// 	expectedErr := errors.New("payment creation failed")
// 	expires := time.Now().Add(24 * time.Hour)

// 	accountMock := new(mocks.MockAccountService)
// 	paymentMock := new(mocks.MockPaymentService)
// 	promoMock := new(mocks.MockPromoUseCase)
// 	subsMock := new(mocks.MockSubscriptionsRepo)
// 	xrayMock := new(mocks.MockXrayManagerRepo)

// 	accountMock.On("Get", mock.Anything, consts.PlatformTelegram, 12345).
// 		Return(domain.Account{ID: 1, KeyID: "test-key", SubscriptionExpiresAt: &expires}, nil)
// 	accountMock.On("SetTariff", mock.Anything, 1, 1).Return(nil)

// 	subsMock.On("GetPeriod", 30).Return(domain.Period{Emoji: "🚀"}, true)
// 	subsMock.On("Get", 1).Return(domain.Tariff{Prices: map[ccurrency.Currency]float64{ccurrency.XTR: 99.0}}, true)
// 	subsMock.On("CalculateRenewal", mock.Anything, ccurrency.XTR, mock.Anything, 1, 30).
// 		Return(99.0, 0, "")

// 	paymentMock.On("Create", mock.Anything, 1, 99, mock.Anything).Return(expectedErr)

// 	uc := invoices.NewUseCase(accountMock, paymentMock, promoMock, subsMock, xrayMock, logger.New("debug", "development"))

// 	update := createPreCheckoutUpdate("XTR", "d:30", 12345)
// 	err := uc.HandlePreCheckoutQuery(nil, update)

// 	assert.Equal(t, expectedErr, err)
// }

// func TestHandlePreCheckoutQuery_XrayManagerAddClientError(t *testing.T) {
// 	expectedErr := errors.New("xray add client failed")

// 	accountMock := new(mocks.MockAccountService)
// 	paymentMock := new(mocks.MockPaymentService)
// 	promoMock := new(mocks.MockPromoUseCase)
// 	subsMock := new(mocks.MockSubscriptionsRepo)
// 	xrayMock := new(mocks.MockXrayManagerRepo)

// 	accountMock.On("Get", mock.Anything, consts.PlatformTelegram, 12345).
// 		Return(domain.Account{ID: 1, KeyID: "test-key", SubscriptionExpiresAt: nil}, nil)
// 	accountMock.On("SetTariff", mock.Anything, 1, 1).Return(nil)

// 	subsMock.On("GetPeriod", 30).Return(domain.Period{Emoji: "🚀"}, true)
// 	subsMock.On("Get", 1).Return(domain.Tariff{Prices: map[ccurrency.Currency]float64{ccurrency.XTR: 99.0}}, true)
// 	subsMock.On("CalculateRenewal", mock.Anything, ccurrency.XTR, mock.Anything, 1, 30).
// 		Return(99.0, 0, "")

// 	paymentMock.On("Create", mock.Anything, 1, 99, mock.Anything).Return(nil)
// 	xrayMock.On("AddClient", mock.Anything, dto.RegionRussia, "test-key", mock.Anything).Return(expectedErr)

// 	uc := invoices.NewUseCase(accountMock, paymentMock, promoMock, subsMock, xrayMock, logger.New("debug", "development"))

// 	update := createPreCheckoutUpdate("XTR", "d:30", 12345)
// 	err := uc.HandlePreCheckoutQuery(nil, update)

// 	assert.Equal(t, expectedErr, err)
// }

// func TestHandlePreCheckoutQuery_NewUserFlow(t *testing.T) {
// 	accountMock := new(mocks.MockAccountService)
// 	paymentMock := new(mocks.MockPaymentService)
// 	promoMock := new(mocks.MockPromoUseCase)
// 	subsMock := new(mocks.MockSubscriptionsRepo)
// 	xrayMock := new(mocks.MockXrayManagerRepo)

// 	// New user - no subscription yet
// 	accountMock.On("Get", mock.Anything, consts.PlatformTelegram, 12345).
// 		Return(domain.Account{ID: 42, KeyID: "user-key-123", SubscriptionExpiresAt: nil}, nil)
// 	accountMock.On("SetTariff", mock.Anything, 42, 1).Return(nil)

// 	subsMock.On("GetPeriod", 30).Return(domain.Period{Emoji: "🚀"}, true)
// 	subsMock.On("Get", 1).Return(domain.Tariff{Prices: map[ccurrency.Currency]float64{ccurrency.XTR: 99.0}}, true)
// 	subsMock.On("CalculateRenewal", mock.Anything, ccurrency.XTR, mock.Anything, 1, 30).
// 		Return(99.0, 0, "")

// 	paymentMock.On("Create", mock.Anything, 42, 99, mock.Anything).Return(nil)
// 	xrayMock.On("AddClient", mock.Anything, dto.RegionRussia, "user-key-123", mock.Anything).Return(nil)
// 	accountMock.On("AddSubscriptionExpiresAt", mock.Anything, 42, 30*24*time.Hour).Return(nil)
// 	accountMock.On("ResetDiscount", mock.Anything, 42).Return(errTestComplete)

// 	uc := invoices.NewUseCase(accountMock, paymentMock, promoMock, subsMock, xrayMock, logger.New("debug", "development"))

// 	update := createPreCheckoutUpdate("XTR", "d:30", 12345)
// 	err := uc.HandlePreCheckoutQuery(nil, update)

// 	assert.Equal(t, errTestComplete, err)
// 	accountMock.AssertExpectations(t)
// 	paymentMock.AssertExpectations(t)
// 	xrayMock.AssertExpectations(t)
// }

// func TestHandlePreCheckoutQuery_ExistingUserFlow(t *testing.T) {
// 	expires := time.Now().Add(10 * 24 * time.Hour)

// 	accountMock := new(mocks.MockAccountService)
// 	paymentMock := new(mocks.MockPaymentService)
// 	promoMock := new(mocks.MockPromoUseCase)
// 	subsMock := new(mocks.MockSubscriptionsRepo)
// 	xrayMock := new(mocks.MockXrayManagerRepo)

// 	accountMock.On("Get", mock.Anything, consts.PlatformTelegram, 12345).
// 		Return(domain.Account{ID: 42, KeyID: "user-key-123", SubscriptionExpiresAt: &expires}, nil)
// 	accountMock.On("SetTariff", mock.Anything, 42, 1).Return(nil)

// 	subsMock.On("GetPeriod", 30).Return(domain.Period{Emoji: "🚀"}, true)
// 	subsMock.On("Get", 1).Return(domain.Tariff{Prices: map[ccurrency.Currency]float64{ccurrency.XTR: 99.0}}, true)
// 	subsMock.On("CalculateRenewal", mock.Anything, ccurrency.XTR, mock.Anything, 1, 30).
// 		Return(99.0, 0, "")

// 	paymentMock.On("Create", mock.Anything, 42, 99, mock.Anything).Return(nil)
// 	// AddClient should NOT be called for existing user
// 	accountMock.On("AddSubscriptionExpiresAt", mock.Anything, 42, 30*24*time.Hour).Return(nil)
// 	accountMock.On("ResetDiscount", mock.Anything, 42).Return(errTestComplete)

// 	uc := invoices.NewUseCase(accountMock, paymentMock, promoMock, subsMock, xrayMock, logger.New("debug", "development"))

// 	update := createPreCheckoutUpdate("XTR", "d:30", 12345)
// 	err := uc.HandlePreCheckoutQuery(nil, update)

// 	assert.Equal(t, errTestComplete, err)
// 	xrayMock.AssertNotCalled(t, "AddClient", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
// }

// func TestHandlePreCheckoutQuery_DiscountApplied(t *testing.T) {
// 	expires := time.Now().Add(24 * time.Hour)

// 	accountMock := new(mocks.MockAccountService)
// 	paymentMock := new(mocks.MockPaymentService)
// 	promoMock := new(mocks.MockPromoUseCase)
// 	subsMock := new(mocks.MockSubscriptionsRepo)
// 	xrayMock := new(mocks.MockXrayManagerRepo)

// 	accountMock.On("Get", mock.Anything, consts.PlatformTelegram, 12345).
// 		Return(domain.Account{
// 			ID:                    42,
// 			KeyID:                 "user-key-123",
// 			SubscriptionExpiresAt: &expires,
// 			Discount:              20, // 20% discount
// 		}, nil)
// 	accountMock.On("SetTariff", mock.Anything, 42, 1).Return(nil)

// 	subsMock.On("GetPeriod", 30).Return(domain.Period{Emoji: "🚀"}, true)
// 	subsMock.On("Get", 1).Return(domain.Tariff{Prices: map[ccurrency.Currency]float64{ccurrency.XTR: 99.0}}, true)
// 	subsMock.On("CalculateRenewal", mock.Anything, ccurrency.XTR, mock.Anything, 1, 30).
// 		Return(99.0, 0, "")

// 	// Price should be 99 - 20% = 79.2 -> 79
// 	paymentMock.On("Create", mock.Anything, 42, 79, mock.Anything).Return(nil)
// 	accountMock.On("AddSubscriptionExpiresAt", mock.Anything, 42, 30*24*time.Hour).Return(nil)
// 	accountMock.On("ResetDiscount", mock.Anything, 42).Return(errTestComplete)

// 	uc := invoices.NewUseCase(accountMock, paymentMock, promoMock, subsMock, xrayMock, logger.New("debug", "development"))

// 	update := createPreCheckoutUpdate("XTR", "d:30", 12345)
// 	_ = uc.HandlePreCheckoutQuery(nil, update)
// 	paymentMock.AssertExpectations(t)
// }

// func TestHandlePreCheckoutQuery_TariffBasedPayload(t *testing.T) {
// 	accountMock := new(mocks.MockAccountService)
// 	paymentMock := new(mocks.MockPaymentService)
// 	promoMock := new(mocks.MockPromoUseCase)
// 	subsMock := new(mocks.MockSubscriptionsRepo)
// 	xrayMock := new(mocks.MockXrayManagerRepo)

// 	accountMock.On("Get", mock.Anything, consts.PlatformTelegram, 12345).
// 		Return(domain.Account{ID: 42, KeyID: "user-key-123", SubscriptionExpiresAt: nil, Tariff: 0}, nil)
// 	accountMock.On("SetTariff", mock.Anything, 42, 2).Return(nil)

// 	subsMock.On("GetPeriod", 30).Return(domain.Period{Emoji: "🚀"}, true)
// 	subsMock.On("Get", 2).Return(domain.Tariff{Prices: map[ccurrency.Currency]float64{ccurrency.XTR: 179.0}}, true)
// 	subsMock.On("CalculateRenewal", mock.Anything, ccurrency.XTR, mock.Anything, 2, 30).
// 		Return(179.0, 0, "")

// 	paymentMock.On("Create", mock.Anything, 42, 179, mock.Anything).Return(nil)
// 	xrayMock.On("AddClient", mock.Anything, dto.RegionRussia, "user-key-123", mock.Anything).Return(nil)
// 	accountMock.On("AddSubscriptionExpiresAt", mock.Anything, 42, 30*24*time.Hour).Return(nil)
// 	accountMock.On("ResetDiscount", mock.Anything, 42).Return(errTestComplete)

// 	uc := invoices.NewUseCase(accountMock, paymentMock, promoMock, subsMock, xrayMock, logger.New("debug", "development"))

// 	// d:30;t:2 = 30 days, tariff ID 2
// 	update := createPreCheckoutUpdate("XTR", "d:30;t:2", 12345)
// 	_ = uc.HandlePreCheckoutQuery(nil, update)
// 	accountMock.AssertCalled(t, "SetTariff", mock.Anything, 42, 2)
// }

// func TestHandlePreCheckoutQuery_TariffUpgrade(t *testing.T) {
// 	expires := time.Now().Add(15 * 24 * time.Hour)

// 	accountMock := new(mocks.MockAccountService)
// 	paymentMock := new(mocks.MockPaymentService)
// 	promoMock := new(mocks.MockPromoUseCase)
// 	subsMock := new(mocks.MockSubscriptionsRepo)
// 	xrayMock := new(mocks.MockXrayManagerRepo)

// 	accountMock.On("Get", mock.Anything, consts.PlatformTelegram, 12345).
// 		Return(domain.Account{
// 			ID:                    42,
// 			KeyID:                 "user-key-123",
// 			SubscriptionExpiresAt: &expires,
// 			Tariff:                1, // Эконом (99р)
// 		}, nil)
// 	accountMock.On("SetTariff", mock.Anything, 42, 2).Return(nil)

// 	subsMock.On("GetPeriod", 30).Return(domain.Period{Emoji: "🚀"}, true)
// 	subsMock.On("Get", 2).Return(domain.Tariff{Prices: map[ccurrency.Currency]float64{ccurrency.XTR: 179.0}}, true)
// 	subsMock.On("CalculateRenewal", mock.Anything, ccurrency.XTR, mock.MatchedBy(func(t domain.Tariff) bool {
// 		return t.Prices[ccurrency.XTR] == 179.0
// 	}), 2, 30).Return(150.0, 0, "test upgrade")

// 	// Price should be 150
// 	paymentMock.On("Create", mock.Anything, 42, 150, mock.Anything).Return(nil)

// 	accountMock.On("AddSubscriptionExpiresAt", mock.Anything, 42, 30*24*time.Hour).Return(nil)
// 	accountMock.On("ResetDiscount", mock.Anything, 42).Return(errTestComplete)

// 	uc := invoices.NewUseCase(accountMock, paymentMock, promoMock, subsMock, xrayMock, logger.New("debug", "development"))

// 	// Upgrading to tariff 2 (Личный) for 30 days
// 	update := createPreCheckoutUpdate("XTR", "d:30;t:2", 12345)
// 	_ = uc.HandlePreCheckoutQuery(nil, update)
// 	paymentMock.AssertExpectations(t)
// }

// func TestHandlePreCheckoutQuery_TariffDowngrade(t *testing.T) {
// 	expires := time.Now().Add(20 * 24 * time.Hour)

// 	accountMock := new(mocks.MockAccountService)
// 	paymentMock := new(mocks.MockPaymentService)
// 	promoMock := new(mocks.MockPromoUseCase)
// 	subsMock := new(mocks.MockSubscriptionsRepo)
// 	xrayMock := new(mocks.MockXrayManagerRepo)

// 	accountMock.On("Get", mock.Anything, consts.PlatformTelegram, 12345).
// 		Return(domain.Account{
// 			ID:                    42,
// 			KeyID:                 "user-key-123",
// 			SubscriptionExpiresAt: &expires,
// 			Tariff:                3, // Домашний (299р)
// 		}, nil)
// 	accountMock.On("SetTariff", mock.Anything, 42, 1).Return(nil)

// 	subsMock.On("GetPeriod", 30).Return(domain.Period{Emoji: "🚀"}, true)
// 	subsMock.On("Get", 1).Return(domain.Tariff{Prices: map[ccurrency.Currency]float64{ccurrency.XTR: 99.0}}, true)
// 	subsMock.On("CalculateRenewal", mock.Anything, ccurrency.XTR, mock.Anything, 1, 30).
// 		Return(0.0, 5, "surplus")

// 	// Price should be 0 for downgrade with credit surplus
// 	paymentMock.On("Create", mock.Anything, 42, 0, mock.Anything).Return(nil)
// 	accountMock.On("AddSubscriptionExpiresAt", mock.Anything, 42, 35*24*time.Hour).Return(nil)
// 	accountMock.On("ResetDiscount", mock.Anything, 42).Return(errTestComplete)

// 	uc := invoices.NewUseCase(accountMock, paymentMock, promoMock, subsMock, xrayMock, logger.New("debug", "development"))

// 	// Downgrading to tariff 1 (Эконом) for 30 days
// 	update := createPreCheckoutUpdate("XTR", "d:30;t:1", 12345)
// 	_ = uc.HandlePreCheckoutQuery(nil, update)
// 	paymentMock.AssertCalled(t, "Create", mock.Anything, 42, 0, mock.Anything)
// }

// func TestHandlePreCheckoutQuery_AllPeriods(t *testing.T) {
// 	periods := []int{30, 90, 180, 365}

// 	for _, days := range periods {
// 		t.Run(fmt.Sprintf("period_%d_days", days), func(t *testing.T) {
// 			accountMock := new(mocks.MockAccountService)
// 			paymentMock := new(mocks.MockPaymentService)
// 			promoMock := new(mocks.MockPromoUseCase)
// 			subsMock := new(mocks.MockSubscriptionsRepo)
// 			xrayMock := new(mocks.MockXrayManagerRepo)

// 			accountMock.On("Get", mock.Anything, consts.PlatformTelegram, 12345).
// 				Return(domain.Account{ID: 1, KeyID: "test-key", SubscriptionExpiresAt: nil}, nil)
// 			accountMock.On("SetTariff", mock.Anything, 1, 2).Return(nil)

// 			subsMock.On("GetPeriod", days).Return(domain.Period{Emoji: "🚀"}, true)
// 			subsMock.On("Get", 2).Return(domain.Tariff{Prices: map[ccurrency.Currency]float64{ccurrency.XTR: 179.0}}, true)
// 			subsMock.On("CalculateRenewal", mock.Anything, ccurrency.XTR, mock.Anything, 2, days).
// 				Return(179.0, 1, "") // Mock 1 extra day for variety

// 			paymentMock.On("Create", mock.Anything, 1, 179, mock.Anything).Return(nil)
// 			xrayMock.On("AddClient", mock.Anything, dto.RegionRussia, "test-key", mock.Anything).Return(nil)
// 			accountMock.On("AddSubscriptionExpiresAt", mock.Anything, 1, time.Duration(days+1)*24*time.Hour).Return(nil)
// 			accountMock.On("ResetDiscount", mock.Anything, 1).Return(errTestComplete)

// 			uc := invoices.NewUseCase(accountMock, paymentMock, promoMock, subsMock, xrayMock, logger.New("debug", "development"))

// 			payload := fmt.Sprintf("d:%d;t:2", days)
// 			update := createPreCheckoutUpdate("XTR", payload, 12345)
// 			_ = uc.HandlePreCheckoutQuery(nil, update)
// 			accountMock.AssertCalled(t, "AddSubscriptionExpiresAt", mock.Anything, 1, time.Duration(days+1)*24*time.Hour)
// 		})
// 	}
// }

// func TestHandlePreCheckoutQuery_AddSubscriptionExpiresAtError(t *testing.T) {
// 	expectedErr := errors.New("subscription update failed")
// 	expires := time.Now().Add(24 * time.Hour)

// 	accountMock := new(mocks.MockAccountService)
// 	paymentMock := new(mocks.MockPaymentService)
// 	promoMock := new(mocks.MockPromoUseCase)
// 	subsMock := new(mocks.MockSubscriptionsRepo)
// 	xrayMock := new(mocks.MockXrayManagerRepo)

// 	accountMock.On("Get", mock.Anything, consts.PlatformTelegram, 12345).
// 		Return(domain.Account{ID: 42, KeyID: "user-key-123", SubscriptionExpiresAt: &expires}, nil)
// 	accountMock.On("SetTariff", mock.Anything, 42, 1).Return(nil)

// 	subsMock.On("GetPeriod", 30).Return(domain.Period{Emoji: "🚀"}, true)
// 	subsMock.On("Get", 1).Return(domain.Tariff{Prices: map[ccurrency.Currency]float64{ccurrency.XTR: 99.0}}, true)
// 	subsMock.On("CalculateRenewal", mock.Anything, ccurrency.XTR, mock.Anything, 1, 30).
// 		Return(99.0, 0, "")

// 	paymentMock.On("Create", mock.Anything, 42, 99, mock.Anything).Return(nil)
// 	accountMock.On("AddSubscriptionExpiresAt", mock.Anything, 42, 30*24*time.Hour).Return(expectedErr)

// 	uc := invoices.NewUseCase(accountMock, paymentMock, promoMock, subsMock, xrayMock, logger.New("debug", "development"))

// 	update := createPreCheckoutUpdate("XTR", "d:30", 12345)
// 	err := uc.HandlePreCheckoutQuery(nil, update)

// 	assert.Equal(t, expectedErr, err)
// }

// func TestHandlePreCheckoutQuery_SetTariffError(t *testing.T) {
// 	expectedErr := errors.New("tariff update failed")

// 	accountMock := new(mocks.MockAccountService)
// 	paymentMock := new(mocks.MockPaymentService)
// 	promoMock := new(mocks.MockPromoUseCase)
// 	subsMock := new(mocks.MockSubscriptionsRepo)
// 	xrayMock := new(mocks.MockXrayManagerRepo)

// 	accountMock.On("Get", mock.Anything, consts.PlatformTelegram, 12345).
// 		Return(domain.Account{ID: 1, KeyID: "test-key"}, nil)

// 	subsMock.On("GetPeriod", 30).Return(domain.Period{Emoji: "🚀"}, true)
// 	subsMock.On("Get", 2).Return(domain.Tariff{Prices: map[ccurrency.Currency]float64{ccurrency.XTR: 179.0}}, true)

// 	accountMock.On("SetTariff", mock.Anything, 1, 2).Return(expectedErr)

// 	uc := invoices.NewUseCase(accountMock, paymentMock, promoMock, subsMock, xrayMock, logger.New("debug", "development"))

// 	update := createPreCheckoutUpdate("XTR", "d:30;t:2", 12345)
// 	err := uc.HandlePreCheckoutQuery(nil, update)

// 	assert.Equal(t, expectedErr, err)
// }
