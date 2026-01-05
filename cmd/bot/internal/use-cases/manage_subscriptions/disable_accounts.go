package manageSubscriptions

import (
	"context"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types/update"
	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager/dto"
	"github.com/quickpowered/thebeyond-master/pkg/utils"
	"go.uber.org/zap"
)

func (uc *UseCase) DisableAccounts(ctx context.Context, bot bin.Interface) error {
	ids, err := uc.subscriptionService.GetAccountsToDisable(ctx)
	if err != nil {
		return err
	}

	uc.logger.Debug("disable subscriptions found", zap.Int("count", len(ids)))

	for _, id := range ids {
		email := utils.NewEmail(id)
		if err := uc.xraymanagerRepo.RemoveClient(ctx, dto.RegionRussia, email); err != nil {
			uc.logger.Error("remove user", zap.Int("account_id", id), zap.Error(err))
		}

		if err := uc.subscriptionService.Cancel(ctx, id); err != nil {
			uc.logger.Error("cancel subscription", zap.Int("account_id", id), zap.Error(err))
			continue
		}

		externalAccountID, err := uc.accountService.GetPlatformUserID(ctx, id)
		if err != nil {
			uc.logger.Error("failed to get expire subsription external account", zap.Int("account_id", id), zap.Error(err))
			continue
		}

		keyboard := types.NewKeyboard()
		keyboard.NewRow(types.NewCallbackButton("🛍 Продлить", "renew"))
		if err := bot.SendMessage(update.Chat{ID: externalAccountID}, "🕒 Ваша подписка истекла, либо исчерпан лимит трафика!", keyboard); err != nil {
			uc.logger.Error("failed to send expire subsription message", zap.Int("account_id", id), zap.Error(err))
			continue
		}

		uc.logger.Info("disable subscription",
			zap.Int("account_id", id),
			zap.String("email", email))
	}

	return nil
}
