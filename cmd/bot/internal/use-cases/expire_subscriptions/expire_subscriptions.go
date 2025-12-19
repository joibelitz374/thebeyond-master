package expireSubscriptions

import (
	"context"

	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager"
	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager/dto"
	"github.com/quickpowered/thebeyond-master/internal/services/account"
	"github.com/quickpowered/thebeyond-master/pkg/email"
	"go.uber.org/zap"
)

type UseCase struct {
	accountService  account.Interface
	xraymanagerRepo xraymanager.Repository
	logger          *zap.Logger
}

func NewUseCase(accountService account.Interface, xraymanagerRepo xraymanager.Repository, logger *zap.Logger) *UseCase {
	return &UseCase{accountService, xraymanagerRepo, logger}
}

func (uc *UseCase) Run(ctx context.Context) error {
	ids, err := uc.accountService.GetExpiredUsers(ctx)
	if err != nil {
		return err
	}

	uc.logger.Debug("expired subscriptions found",
		zap.Int("count", len(ids)))

	for _, id := range ids {
		email := email.NewAccount(id)
		uc.logger.Debug("expired subscription",
			zap.Int("account_id", id),
			zap.String("email", email))

		for _, region := range []dto.Region{dto.RegionRussia} {
			if err := uc.xraymanagerRepo.RemoveClient(ctx, region, email); err != nil {
				uc.logger.Error("remove user", zap.Error(err), zap.Int("account_id", id))
			}
		}

		if err := uc.accountService.CancelSubscriptions(ctx, id); err != nil {
			uc.logger.Error("cancel subscription", zap.Error(err), zap.Int("account_id", id))
		}
	}
	return nil
}
