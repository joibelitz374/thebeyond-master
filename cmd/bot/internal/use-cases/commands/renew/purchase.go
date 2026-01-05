package renew

import (
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/configs/currency"
	sharedDomain "github.com/quickpowered/thebeyond-master/internal/domain"
)

func (h handler) purchaseOffer(bot bin.Interface, p *domain.Payload, tariffID int, tariff sharedDomain.Tariff, days int) error {
	if _, exists := h.TariffsRepo.GetPeriod(days); !exists {
		return bot.SendMessage(p.Message.Chat(), "Period not found")
	}

	price, extraDays := h.TariffsRepo.CalculateRenewal(p.Account, currency.XTR, tariff, tariffID, days)
	if price == 0 {
		return h.upgradeTariff(bot, p, tariffID, tariff, days+extraDays)
	}

	return h.telegramInvoice(bot, p, tariffID, days+extraDays, price)
}
