package deps

import (
	"github.com/quickpowered/thebeyond-master/internal/repositories/subscriptions"
	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager"
	"github.com/quickpowered/thebeyond-master/internal/services/account"
	"github.com/quickpowered/thebeyond-master/internal/services/payment"
	"github.com/quickpowered/thebeyond-master/internal/services/promo"
	"go.uber.org/zap"
)

type Dependencies struct {
	AccountService    account.Interface
	PaymentService    payment.Interface
	PromoService      promo.Interface
	SubscriptionsRepo subscriptions.Repository
	XRayManagerRepo   xraymanager.Repository
	Logger            *zap.Logger
}

func NewDependencies(
	accountService account.Interface,
	paymentService payment.Interface,
	promoService promo.Interface,
	subscriptionsRepo subscriptions.Repository,
	xraymanagerRepo xraymanager.Repository,
	logger *zap.Logger,
) Dependencies {
	return Dependencies{
		accountService, paymentService, promoService,
		subscriptionsRepo, xraymanagerRepo, logger,
	}
}
