package renew

import (
	"fmt"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/configs/currency"
	"github.com/quickpowered/thebeyond-master/configs/language"
)

func (h handler) showTariffs(bot bin.Interface, p *domain.Payload) error {
	selectTariffsMsg := i18n.SelectTariffsMessages[p.Account.Language]
	return bot.SendMessage(p.Message.Chat(), selectTariffsMsg.SelectTariff+":", h.listOfTariffs(p.Account.Language, p.Account.Currency, selectTariffsMsg))
}

func (h handler) listOfTariffs(language language.Language, acurrency currency.Currency, selectTariffsMsg i18n.SelectTariffsLocale) *types.Keyboard {
	tariffsMsg := i18n.TariffsMessages[language]
	controlMsg := i18n.ControlMessages[language]
	keyboard := types.NewKeyboard()

	for _, id := range h.TariffsRepo.GetTargets() {
		tariff, ok := h.TariffsRepo.Get(id)
		if !ok {
			continue
		}

		var label string
		if id == 0 {
			label += tariff.Emoji + " " + selectTariffsMsg.Free
		} else {
			currencySymbol := currency.Currencies[acurrency].Emoji
			label += fmt.Sprintf("%s %s — %.2f%s",
				tariff.Emoji, tariffsMsg[id],
				tariff.Prices[acurrency], currencySymbol)
			if tariff.Discount != "" {
				label += " (" + tariff.Discount + ")"
			}
		}

		keyboard.NewRow(types.NewCallbackButton(label, fmt.Sprintf("renew %d", id)))
	}

	keyboard.NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, "start"))
	return keyboard
}
