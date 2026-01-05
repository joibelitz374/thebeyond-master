package manageSubscriptions

import (
	"context"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager"
	"github.com/quickpowered/thebeyond-master/internal/services/account"
	"github.com/quickpowered/thebeyond-master/internal/services/subscription"
	"go.uber.org/zap"
)

type UseCase struct {
	accountService      account.Interface
	subscriptionService subscription.Interface
	xraymanagerRepo     xraymanager.Repository
	logger              *zap.Logger
}

func NewUseCase(accountService account.Interface, subscriptionService subscription.Interface, xraymanagerRepo xraymanager.Repository, logger *zap.Logger) *UseCase {
	return &UseCase{accountService, subscriptionService, xraymanagerRepo, logger}
}

func (uc *UseCase) ResetTraffic(ctx context.Context, bot bin.Interface) error {
	ids, err := uc.subscriptionService.GetNeedRefreshTraffic(ctx)
	if err != nil {
		return err
	}

	uc.logger.Debug("reset subscriptions traffic found", zap.Int("count", len(ids)))

	for _, accountID := range ids {
		if err := uc.subscriptionService.ResetLastTrafficRefreshAt(ctx, accountID); err != nil {
			uc.logger.Error("reset subscriptions traffic error", zap.Int("account_id", accountID), zap.Error(err))
		}
	}

	return nil
}
