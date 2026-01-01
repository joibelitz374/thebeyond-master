package renew

import (
	"strconv"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/configs/currency"
)

func (h handler) purchaseOffer(bot bin.Interface, p *domain.Payload, msg i18n.RenewLocale) error {
	tariffID, err := strconv.Atoi(p.Args[1])
	if err != nil {
		return bot.SendMessage(p.Message.Chat(), "Invalid tariff ID")
	}

	tariff, ok := h.TariffsRepo.Get(tariffID)
	if !ok {
		return bot.SendMessage(p.Message.Chat(), "No tariff found")
	}

	days, err := strconv.Atoi(p.Args[2])
	if err != nil {
		return bot.SendMessage(p.Message.Chat(), "Invalid days value")
	}

	price, extraDays := h.TariffsRepo.CalculateRenewal(p.Account, currency.XTR, tariff, tariffID, days)
	if price == 0 {
		return h.upgradeTariff(bot, p, tariffID, tariff, days+extraDays)
	}

	if _, exists := h.TariffsRepo.GetPeriod(days); !exists {
		return bot.SendMessage(p.Message.Chat(), "Period not found")
	}

	if err := h.telegramInvoice(bot, p, tariffID, days+extraDays, price, msg); err != nil {
		return err
	}

	return bot.SendMessage(p.Message.Chat(), "Return to main menu:", types.NewKeyboard().
		NewRow(types.NewCallbackButton("🔄 Menu", "start")))
}
