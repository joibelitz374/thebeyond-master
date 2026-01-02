package commands

import (
	"context"
	"time"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/renew"
	"github.com/quickpowered/thebeyond-master/configs"
	"github.com/quickpowered/thebeyond-master/configs/currency"
)

const CURRENCY_CMD = "currency"

type currencyHandler struct {
	deps.Dependencies
	renewHandler Command
}

func NewCurrencyHandler(deps deps.Dependencies) currencyHandler {
	renewHandler := renew.NewHandler(deps)
	return currencyHandler{deps, renewHandler}
}

func (h currencyHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	isNewcomer := p.Account.IsNewcomer()
	msg := i18n.CurrencyMessages[p.Account.Language]
	controlMsg := i18n.ControlMessages[p.Account.Language]
	if len(p.Args) >= 2 {
		return h.changeCurrency(bot, p, msg, isNewcomer)
	}

	buttons, err := getButtons(CURRENCY_CMD, currency.Targets, func(code string) string {
		currency, ok := currency.Get(currency.Currency(code))
		if !ok {
			return "N/A"
		}
		return currency.Flag + " " + currency.Name
	})
	if err != nil {
		return err
	}

	if !isNewcomer {
		buttons = append(buttons, []types.Button{types.NewCallbackButton("◀️ "+controlMsg.Back, SETTINGS_CMD)})
	}

	return bot.SendMessage(p.Message.Chat(), msg.ChooseCurrency+":", &types.Keyboard{ButtonRows: buttons})
}

func (h currencyHandler) changeCurrency(bot bin.Interface, p *domain.Payload, msg i18n.CurrencyLocale, isNewcomer bool) error {
	p.Account.Currency = currency.Currency(p.Args[1])

	getter := func() (configs.ItemInfo, bool) {
		return currency.Get(p.Account.Currency)
	}

	setter := func() error {
		ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
		defer cancel()
		return h.AccountService.SetCurrency(ctx, p.Account.ID, string(p.Account.Currency))
	}

	if err := setValue(bot, p, msg.CurrencyChangedTo, getter, setter); err != nil {
		return err
	}

	if isNewcomer {
		p.Args = []string{renew.CMD, "check-subscription"}
		return h.renewHandler.Execute(bot, p)
	}

	p.Args = []string{CURRENCY_CMD}
	return h.Execute(bot, p)
}
