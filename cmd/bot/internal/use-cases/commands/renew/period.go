package renew

import (
	"fmt"

	sharedDomain "github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/configs/currency"
	"github.com/quickpowered/thebeyond-master/internal/domain"
)

func (h handler) showPeriods(bot bin.Interface, p *sharedDomain.Payload, tariff domain.Tariff, tariffID int, msg i18n.RenewLocale) error {
	controlMsg := i18n.ControlMessages[p.Account.Language]
	periodTargets := h.TariffsRepo.GetPeriodTargets()
	keyboard := types.NewKeyboard()
	for _, days := range periodTargets {
		period, exists := h.TariffsRepo.GetPeriod(days)
		if !exists {
			continue
		}

		var label string
		totalPrice, extraDays := h.TariffsRepo.CalculateRenewal(p.Account, p.Account.Currency, tariff, tariffID, days)
		currencySymbol := currency.Currencies[p.Account.Currency].Emoji
		if p.Account.IsActive() && p.Account.Tariff != tariffID {
			if totalPrice == 0 {
				totalTotalDays := days + extraDays
				label = fmt.Sprintf("%s %d дн.", period.Emoji, totalTotalDays)
			} else {
				label = fmt.Sprintf("%s %s — %.2f%s", period.Emoji, h.getTariffPeriodName(msg, days), totalPrice, currencySymbol)
			}
		} else {
			label = fmt.Sprintf(
				"%s %s — %.2f%s", period.Emoji,
				h.getTariffPeriodName(msg, days+extraDays),
				totalPrice, currencySymbol,
			)
		}

		keyboard.NewRow(types.NewCallbackButton(label, fmt.Sprintf("renew %s %d", p.Args[1], days)))
	}

	keyboard.NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, CMD))

	text := "Выберите срок подписки:\n\n"
	var traffic string
	switch {
	case tariff.Traffic >= 5000:
		traffic = "Безлимит"
	case tariff.Traffic >= 1000:
		traffic = fmt.Sprintf("%d TB/мес.", tariff.Traffic/1000)
	default:
		traffic = fmt.Sprintf("%d GB/мес.", tariff.Traffic)
	}

	text += fmt.Sprintf("🚌 Трафик: %s", traffic)
	if tariff.Devices > 0 {
		text += fmt.Sprintf("\n🗺 Подключений: %d", tariff.Devices)
	}

	return bot.SendMessage(p.Message.Chat(), text, keyboard)
}
