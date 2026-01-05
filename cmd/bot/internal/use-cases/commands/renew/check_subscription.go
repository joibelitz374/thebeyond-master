package renew

import (
	"context"
	"fmt"
	"time"

	telegramsdk "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/telegram"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/pkg/utils"
	"go.uber.org/zap"
)

func (h handler) checkSubscription(bot bin.Interface, p *domain.Payload) error {
	controlMsg := i18n.ControlMessages[p.Account.Language]
	freemiumMsg := i18n.TelegramChannelBonusMessages[p.Account.Language]

	tgBot, ok := bot.(*telegram.Adapter)
	if !ok {
		return fmt.Errorf("failed to cast bot to telegram adapter")
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()

	member, err := tgBot.GetAPI().(*telegramsdk.Bot).GetChatMember(ctx, &telegramsdk.GetChatMemberParams{
		ChatID: -1003309480333,
		UserID: int64(p.Message.Sender()),
	})
	if err != nil {
		return err
	}

	if member.Left != nil && member.Left.Status == models.ChatMemberTypeLeft {
		if err := tgBot.AnswerCallbackQuery(p.Message.CallbackQueryID(), freemiumMsg.NotSubscribedToChannel); err != nil {
			return err
		}

		return bot.SendMessage(p.Message.Chat(), "🆓 "+freemiumMsg.SubscriberBonus, types.NewKeyboard().
			NewRow(types.NewCallbackButton("🔄 "+freemiumMsg.Check, CMD+" check-subscription")).
			NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, CMD)))
	}

	if p.Account.FreemiumStatus == "available" {
		return bot.SendMessage(p.Message.Chat(), "❎ "+freemiumMsg.TariffAlreadyActive, types.NewKeyboard().
			NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, CMD)))
	}

	regions, err := h.SubscriptionService.GetRegions(p.Account.Region)
	if err != nil {
		h.Logger.Error("failed to get regions", zap.Error(err))
		return err
	}

	email := utils.NewEmail(p.Account.ID)

	if !p.Account.IsActive() {
		for _, region := range regions {
			h.Logger.Info("Preemptively removing client before add", zap.String("region", string(region)), zap.String("email", email))
			if remErr := h.XRayManagerService.RemoveClient(ctx, region, email); remErr != nil {
				h.Logger.Warn("Preemptive remove failed, but continuing", zap.Error(remErr), zap.String("region", string(region)), zap.String("email", email))
			}
		}

		select {
		case <-time.After(2 * time.Second):
		case <-ctx.Done():
			return ctx.Err()
		}

		for _, region := range regions {
			h.Logger.Info("Adding client after preemptive remove", zap.String("region", string(region)), zap.String("email", email))
			if addErr := h.XRayManagerService.AddClient(ctx, region, p.Account.KeyID, email); addErr != nil {
				h.Logger.Error("Failed to add client even after remove", zap.Error(addErr), zap.String("region", string(region)), zap.String("email", email))
				return addErr
			}
		}
	}

	if err := h.SubscriptionService.StartFreemium(ctx, p.Account.ID); err != nil {
		return err
	}

	return bot.SendMessage(p.Message.Chat(), freemiumMsg.ThanksForSubscription, types.NewKeyboard().
		NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, CMD)))
}
