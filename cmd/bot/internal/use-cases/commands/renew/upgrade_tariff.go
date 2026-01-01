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

	tariffsMsg := i18n.TariffsMessages[p.Account.Language]
	return bot.SendMessage(p.Message.Chat(), fmt.Sprintf("🆙 Тариф изменён на %s %s (+%d дн.)",
		tariff.Emoji, tariffsMsg[tariffID], days),
		types.NewKeyboard().
			NewRow(types.NewCallbackButton("🔄 Menu", "start")))
}
