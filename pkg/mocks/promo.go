package mocks

import (
	"github.com/go-telegram/bot"
	"github.com/quickpowered/thebeyond-master/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockPromoUseCase struct {
	mock.Mock
}

func (m *MockPromoUseCase) Get(promoName string) (domain.Promo, int, error) {
	args := m.Called(promoName)
	return args.Get(0).(domain.Promo), args.Int(1), args.Error(2)
}

func (m *MockPromoUseCase) RegisterReferral(botAPI *bot.Bot, senderID, accountID int, promo domain.Promo) error {
	args := m.Called(botAPI, senderID, accountID, promo)
	return args.Error(0)
}
