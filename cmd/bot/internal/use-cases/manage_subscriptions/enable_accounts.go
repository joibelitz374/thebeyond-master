package manageSubscriptions

import (
	"context"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager/dto"
	"github.com/quickpowered/thebeyond-master/pkg/utils"
	"go.uber.org/zap"
)

func (uc *UseCase) EnableAccounts(ctx context.Context, bot bin.Interface) error {
	ids, err := uc.subscriptionService.GetAccountsToEnable(ctx)
	if err != nil {
		return err
	}

	uc.logger.Debug("enable subscriptions found", zap.Int("count", len(ids)))

	for _, account := range ids {
		email := utils.NewEmail(account.ID)
		if err := uc.xraymanagerRepo.AddClient(ctx, dto.RegionRussia, account.KeyID, email); err != nil {
			uc.logger.Error("remove user", zap.Error(err),
				zap.Int("account_id", account.ID),
				zap.String("key_id", account.KeyID))
		}

		uc.logger.Info("enabled subscription",
			zap.Int("account_id", account.ID),
			zap.String("key_id", account.KeyID),
			zap.String("email", email))
	}

	return nil
}
