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

func (h handler) showPeriods(bot bin.Interface, p *sharedDomain.Payload, tariffID int, tariff domain.Tariff) error {
	periodsMsg := i18n.PeriodsMessages[p.Account.Language]
	periodNamesMsg := i18n.PeriodNamesMessages[p.Account.Language]
	accountMsg := i18n.AccountMessages[p.Account.Language]
	controlMsg := i18n.ControlMessages[p.Account.Language]
	periodTargets := h.TariffsRepo.GetPeriodTargets()

	keyboard := types.NewKeyboard()
	for _, row := range periodTargets {
		buttonRows := []types.Button{}
		for _, days := range row {
			period, exists := h.TariffsRepo.GetPeriod(days)
			if !exists {
				continue
			}

			totalPrice, extraDays := h.TariffsRepo.CalculateRenewal(p.Account, p.Account.Currency, tariff, tariffID, days)
			currencySymbol := currency.Currencies[p.Account.Currency].Emoji
			buttonRows = append(buttonRows, types.NewCallbackButton(fmt.Sprintf(
				"%s %s — %.2f%s", period.Emoji,
				h.getTariffPeriodName(periodNamesMsg, days+extraDays),
				totalPrice, currencySymbol,
			), fmt.Sprintf("%s t:%d;d:%d", CMD, tariffID, days)))
		}
		keyboard.NewRow(buttonRows...)
	}

	keyboard.NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, fmt.Sprintf("%s t:%d", CMD, tariffID)))

	text := periodsMsg.SelectPeriod + ":\n"
	var traffic string
	switch {
	case tariff.Traffic >= 5000:
		traffic = periodsMsg.Unlimited
	case tariff.Traffic >= 1000:
		traffic = fmt.Sprintf("%d TB/%s", tariff.Traffic/1000, accountMsg.MonthShort)
	default:
		traffic = fmt.Sprintf("%d GB/%s", tariff.Traffic, accountMsg.MonthShort)
	}

	text += fmt.Sprintf("🚌 <b>"+periodsMsg.Traffic+"</b>: %s", traffic)
	return bot.SendMessage(p.Message.Chat(), text, keyboard)
}

func (h handler) getTariffPeriodName(msg i18n.PeriodNamesLocale, days int) (title string) {
	switch days {
	case 7:
		title = msg.Week
	case 14:
		title = msg.TwoWeeks
	case 30:
		title = msg.Month
	case 90:
		title = msg.Season
	case 180:
		title = msg.HalfYear
	case 365:
		title = msg.Year
	case 730:
		title = msg.TwoYears
	default:
		title = fmt.Sprintf("%d дн.", days)
	}
	return title
}
