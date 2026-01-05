package renew

import (
	"fmt"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/configs/currency"
	sharedDomain "github.com/quickpowered/thebeyond-master/internal/domain"
	"github.com/quickpowered/thebeyond-master/pkg/utils"
)

func (h handler) showTariffs(bot bin.Interface, p *domain.Payload, tariffID int, tariff sharedDomain.Tariff) error {
	accountsMsg := i18n.AccountMessages[p.Account.Language]
	tariffsMsg := i18n.TariffsMessages[p.Account.Language]
	tariffNamesMsg := i18n.TariffNamesMessages[p.Account.Language]
	freemiumMsg := i18n.TelegramChannelBonusMessages[p.Account.Language]
	controlMsg := i18n.ControlMessages[p.Account.Language]

	days := 30
	currencySymbol := currency.Currencies[p.Account.Currency].Emoji
	keyboard := types.NewKeyboard()

	totalPrice, extraDays := h.TariffsRepo.CalculateRenewal(p.Account, p.Account.Currency, tariff, tariffID, days)
	if tariffID == 0 {
		keyboard.NewRow(types.NewCallbackButton("🔄 "+freemiumMsg.Check, fmt.Sprintf("%s check-subscription", CMD)))
	} else {
		if totalPrice == 0 && p.Account.Tariff != tariffID {
			totalDays := days + extraDays
			keyboard.NewRow(types.NewCallbackButton(fmt.Sprintf("🆙 %d %s", totalDays, tariffsMsg.DaysShort), fmt.Sprintf("%s t:%d;d:%d", CMD, tariffID, totalDays)))
		} else {
			keyboard.NewRow(types.NewCallbackButton("🛍 "+tariffsMsg.Buy, fmt.Sprintf("%s t:%d;d:-1", CMD, tariffID)))
		}
	}

	keyboard.NewRow(h.tariffsPagination(tariffID)...)
	keyboard.NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, "start"))

	text := "🗒 <b>" + tariffsMsg.Tariff + "</b>: " + tariffNamesMsg[tariffID] + "\n"
	if tariffID == 0 {
		text += "❇️ <b>" + tariffsMsg.ConditionPrefix + "</b>: " + tariffsMsg.SubscribePre + " <b><a href=\"https://t.me/beyondsecurenews\">" + tariffsMsg.NewsChannelLabel + "</a></b>\n"
	} else if totalPrice != 0 {
		text += fmt.Sprintf("💶 <b>"+tariffsMsg.Price+"</b>: %.2f%s\n", totalPrice, currencySymbol)
	}

	var maxTraffic string
	if tariffID != 3 {
		maxTraffic = fmt.Sprintf("%d GB/%s", tariff.Traffic, accountsMsg.MonthShort)
		text += "🚌 <b>" + tariffsMsg.Traffic + "</b>: " + maxTraffic + "\n"
	}

	return bot.SendMessage(p.Message.Chat(), text, keyboard)
}

func (h handler) tariffsPagination(tariffID int) []types.Button {
	rows := []types.Button{}
	if tariffID > 0 {
		rows = append(rows, types.NewCallbackButton("◀️", fmt.Sprintf("%s t:%d", CMD, tariffID-1)))
	}

	rows = append(rows, types.NewCallbackButton(utils.ToKeycap(tariffID+1), fmt.Sprintf("%s t:%d", CMD, tariffID)))

	if tariffID < 3 {
		rows = append(rows, types.NewCallbackButton("▶️", fmt.Sprintf("%s t:%d", CMD, tariffID+1)))
	}

	return rows
}
