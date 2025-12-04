package tools

import (
	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/repositories/xray"
	"github.com/quickpowered/frilly/internal/services/account"
	"github.com/quickpowered/frilly/internal/services/payment"
	"go.uber.org/zap"
)

type Command interface {
	Execute(bot bin.Interface, payload *domain.Payload) error
}

type Modules struct {
	PaymentService payment.Interface
	AccountService account.Interface
	XRayClient     xray.Interface
	Logger         *zap.Logger
}

func NewModules(paymentService payment.Interface, accountService account.Interface, xrayClient xray.Interface, logger *zap.Logger) Modules {
	return Modules{paymentService, accountService, xrayClient, logger}
}
