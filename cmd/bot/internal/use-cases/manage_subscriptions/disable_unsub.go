package manageSubscriptions

import (
	"context"
	"fmt"
	"time"

	telegramsdk "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/telegram"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types/update"
	"go.uber.org/zap"
)

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
