package deps

import (
	"github.com/quickpowered/frilly/internal/repositories/serverlocations"
	"github.com/quickpowered/frilly/internal/repositories/subscriptions"
	"github.com/quickpowered/frilly/internal/repositories/xray"
	"github.com/quickpowered/frilly/internal/services/account"
	"github.com/quickpowered/frilly/internal/services/payment"
	"github.com/quickpowered/frilly/internal/services/promo"
	"go.uber.org/zap"
)

type Dependencies struct {
	AccountService      account.Interface
	PaymentService      payment.Interface
	PromoService        promo.Interface
	SubscriptionsRepo   subscriptions.Repository
	ServerLocationsRepo serverlocations.Repository
	XRayClient          xray.Interface
	Logger              *zap.Logger
}

func NewDependencies(
	accountService account.Interface,
	paymentService payment.Interface,
	promoService promo.Interface,
	subscriptionsRepo subscriptions.Repository,
	serverLocationsRepo serverlocations.Repository,
	xrayClient xray.Interface,
	logger *zap.Logger,
) Dependencies {
	return Dependencies{
		accountService, paymentService, promoService,
		subscriptionsRepo, serverLocationsRepo,
		xrayClient, logger,
	}
}
