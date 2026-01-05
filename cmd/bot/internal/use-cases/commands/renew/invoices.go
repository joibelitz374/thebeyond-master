package renew

import (
	"errors"
	"fmt"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/telegram"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/pkg/consts"
)

func (h handler) telegramInvoice(bot bin.Interface, p *domain.Payload, tariffID, days int, price float64) error {
	periodNamesMsg := i18n.PeriodNamesMessages[p.Account.Language]
	controlMsg := i18n.ControlMessages[p.Account.Language]
	invoiceMsg := i18n.InvoiceMessages[p.Account.Language]

	if bot.GetPlatform() != consts.PlatformTelegram {
		return bot.SendMessage(p.Message.Chat(), "This command is only available in Telegram", types.NewKeyboard().
			NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, fmt.Sprintf("%s t:%d;d:-1", CMD, tariffID))))
	}

	tgBot, ok := bot.(*telegram.Adapter)
	if !ok {
		return errors.New("unable to obtain Telegram bot")
	}

	periodName := h.getTariffPeriodName(periodNamesMsg, days)
	payload := fmt.Sprintf("d:%d;t:%d", days, tariffID)
	return tgBot.SendInvoice(p.Message.Chat(), periodName, invoiceMsg.Invoice, payload, "XTR", int(price), types.NewKeyboard().
		NewRow(types.NewPayButton(fmt.Sprintf("Pay %d ⭐", int(price)))).
		NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, fmt.Sprintf("%s t:%d;d:-1", CMD, tariffID))))
}
