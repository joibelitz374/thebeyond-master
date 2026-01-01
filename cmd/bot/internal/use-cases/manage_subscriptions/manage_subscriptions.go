package manageSubscriptions

import (
	"context"
	"fmt"
	"time"

	telegramsdk "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/telegram"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types/update"
	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager"
	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager/dto"
	"github.com/quickpowered/thebeyond-master/internal/services/account"
	"github.com/quickpowered/thebeyond-master/internal/services/subscription"
	"github.com/quickpowered/thebeyond-master/pkg/utils"
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

func (uc *UseCase) DisableUnsub(ctx context.Context, bot bin.Interface) error {
	subscribers, err := uc.subscriptionService.GetFremiumAccounts(ctx)
	if err != nil {
		return err
	}

	uc.logger.Debug("disable unsub found", zap.Int("count", len(subscribers)))

	tgBot, ok := bot.(*telegram.Adapter)
	if !ok {
		return fmt.Errorf("failed to cast bot to telegram adapter")
	}

	for _, subscriber := range subscribers {
		ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
		defer cancel()

		member, err := tgBot.GetAPI().(*telegramsdk.Bot).GetChatMember(ctx, &telegramsdk.GetChatMemberParams{
			ChatID: -1003309480333,
			UserID: int64(subscriber.ExternalAccountID),
		})
		if err != nil {
			return err
		}

		if member.Left != nil && member.Left.Status == models.ChatMemberTypeLeft {
			if err := uc.subscriptionService.RemoveFreemium(ctx, subscriber.ID); err != nil {
				return err
			}

			return bot.SendMessage(update.Chat{ID: subscriber.ExternalAccountID}, "🏐 Вы отписались от нашего Telegram канала! (-10 GB)")
		}
	}

	return nil
}

func (uc *UseCase) ResetTraffic(ctx context.Context, bot bin.Interface) error {
	ids, err := uc.subscriptionService.GetNeedRefreshTraffic(ctx)
	if err != nil {
		return err
	}

	uc.logger.Debug("reset subscriptions traffic found", zap.Int("count", len(ids)))

	for _, accountID := range ids {
		if err := uc.subscriptionService.ResetLastTrafficRefreshAt(ctx, accountID); err != nil {
			uc.logger.Error("reset subscriptions traffic error", zap.Error(err),
				zap.Int("account_id", accountID))
		}
	}

	return nil
}

func (uc *UseCase) DisableAccounts(ctx context.Context, bot bin.Interface) error {
	ids, err := uc.subscriptionService.GetAccountsToDisable(ctx)
	if err != nil {
		return err
	}

	uc.logger.Debug("disable subscriptions found", zap.Int("count", len(ids)))

	for _, id := range ids {
		email := utils.NewEmail(id)
		if err := uc.xraymanagerRepo.RemoveClient(ctx, dto.RegionRussia, email); err != nil {
			uc.logger.Error("remove user", zap.Error(err), zap.Int("account_id", id))
		}

		if err := uc.subscriptionService.Cancel(ctx, id); err != nil {
			uc.logger.Error("cancel subscription", zap.Error(err), zap.Int("account_id", id))
			continue
		}

		externalAccountID, err := uc.accountService.GetPlatformUserID(ctx, id)
		if err != nil {
			uc.logger.Error("failed to get expire subsription external account", zap.Error(err))
			continue
		}

		keyboard := types.NewKeyboard()
		keyboard.NewRow(types.NewCallbackButton("🛍 Продлить", "renew"))
		if err := bot.SendMessage(update.Chat{ID: externalAccountID}, "🕒 Ваша подписка истекла, либо исчерпан лимит трафика!", keyboard); err != nil {
			uc.logger.Error("failed to send expire subsription message", zap.Error(err))
			continue
		}

		uc.logger.Debug("disable subscription",
			zap.Int("account_id", id),
			zap.String("email", email))
	}

	return nil
}

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

		uc.logger.Debug("enabled subscription",
			zap.Int("account_id", account.ID),
			zap.String("key_id", account.KeyID),
			zap.String("email", email))
	}

	return nil
}
