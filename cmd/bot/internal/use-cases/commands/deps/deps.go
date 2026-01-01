package deps

import (
	"github.com/quickpowered/thebeyond-master/internal/repositories/tariffs"
	"github.com/quickpowered/thebeyond-master/internal/services/account"
	"github.com/quickpowered/thebeyond-master/internal/services/payment"
	"github.com/quickpowered/thebeyond-master/internal/services/promo"
	"github.com/quickpowered/thebeyond-master/internal/services/subscription"
	"github.com/quickpowered/thebeyond-master/internal/services/xraymanager"
	"go.uber.org/zap"
)

type Dependencies struct {
	AccountService      account.Interface
	SubscriptionService subscription.Interface
	PaymentService      payment.Interface
	PromoService        promo.Interface
	TariffsRepo         tariffs.Repository
	XRayManagerService  xraymanager.Interface
	Logger              *zap.Logger
}

func NewDependencies(
	accountService account.Interface,
	subscriptionService subscription.Interface,
	paymentService payment.Interface,
	promoService promo.Interface,
	tariffsRepo tariffs.Repository,
	xraymanagerService xraymanager.Interface,
	logger *zap.Logger,
) Dependencies {
	return Dependencies{
		accountService, subscriptionService,
		paymentService, promoService,
		tariffsRepo, xraymanagerService,
		logger,
	}
}
