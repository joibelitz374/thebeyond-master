package mocks

import (
	"github.com/quickpowered/thebeyond-master/configs/currency"
	"github.com/quickpowered/thebeyond-master/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockSubscriptionsRepo struct {
	mock.Mock
}

func (m *MockSubscriptionsRepo) Get(id int) (domain.Tariff, bool) {
	args := m.Called(id)
	return args.Get(0).(domain.Tariff), args.Bool(1)
}

func (m *MockSubscriptionsRepo) GetPeriod(days int) (domain.Period, bool) {
	args := m.Called(days)
	return args.Get(0).(domain.Period), args.Bool(1)
}

func (m *MockSubscriptionsRepo) GetTargets() []int {
	args := m.Called()
	return args.Get(0).([]int)
}

func (m *MockSubscriptionsRepo) GetPeriodTargets() []int {
	args := m.Called()
	return args.Get(0).([]int)
}

func (m *MockSubscriptionsRepo) CalculateRenewal(account domain.Account, currency currency.Currency, targetTarrif domain.Tariff, targetTarrifID, days int) (finalPrice float64, extraDays int) {
	args := m.Called(account, currency, targetTarrif, targetTarrifID, days)
	return args.Get(0).(float64), args.Int(1)
}
