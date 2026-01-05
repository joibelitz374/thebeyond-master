package renew

import (
	"context"
	"fmt"
	"time"

	sharedDomain "github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/internal/domain"
	"go.uber.org/zap"
)

func (h handler) upgradeTariff(bot bin.Interface, p *sharedDomain.Payload, tariffID int, tariff domain.Tariff, days int) error {
	tariffNamesMsg := i18n.TariffNamesMessages[p.Account.Language]
	tariffChangeMsg := i18n.TariffChangeMessages[p.Account.Language]
	controlMsg := i18n.ControlMessages[p.Account.Language]

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	if err := h.SubscriptionService.SetTariff(ctx, p.Account.ID, tariffID); err != nil {
		h.Logger.Error("failed to set tariff", zap.Error(err))
		return bot.SendMessage(p.Message.Chat(), "Failed to change plan")
	}

	duration := time.Duration(days) * 24 * time.Hour
	if err := h.SubscriptionService.SetExpiresAt(ctx, p.Account.ID, duration); err != nil {
		h.Logger.Error("failed to add duration", zap.Error(err))
		return bot.SendMessage(p.Message.Chat(), "Failed to extend subscription")
	}

	text := fmt.Sprintf("🆙 "+tariffChangeMsg.TariffChangedTo+" %s %s (+%d "+tariffChangeMsg.DaysShort+")", tariff.Emoji, tariffNamesMsg[tariffID], days)
	return bot.SendMessage(p.Message.Chat(), text, types.NewKeyboard().
		NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, CMD)))
}
