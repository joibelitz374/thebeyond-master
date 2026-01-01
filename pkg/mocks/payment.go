package mocks

import (
	"context"
	"time"

	"github.com/quickpowered/thebeyond-master/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockPaymentService struct {
	mock.Mock
}

func (m *MockPaymentService) Create(ctx context.Context, accountID, amount int, expiresAt time.Time) error {
	args := m.Called(ctx, accountID, amount, expiresAt)
	return args.Error(0)
}

func (m *MockPaymentService) Get(ctx context.Context, accountID int, expiresAt time.Time) ([]domain.Payment, error) {
	args := m.Called(ctx, accountID, expiresAt)
	return args.Get(0).([]domain.Payment), args.Error(1)
}

func (m *MockPaymentService) Delete(ctx context.Context, accountID int) error {
	args := m.Called(ctx, accountID)
	return args.Error(0)
}
