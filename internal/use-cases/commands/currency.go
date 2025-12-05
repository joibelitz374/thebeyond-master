package commands

import (
	"context"
	"strconv"
	"time"

	"github.com/quickpowered/frilly/config"
	"github.com/quickpowered/frilly/config/currency"
	"github.com/quickpowered/frilly/config/language"
	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/i18n"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/use-cases/commands/deps"
)

const CURRENCY_CMD = "currency"

type currencyHandler struct {
	deps.Dependencies
	welcomeHandler Command
}

func NewCurrencyHandler(deps deps.Dependencies) currencyHandler {
	welcomeHandler := NewWelcomeHandler(deps)
	return currencyHandler{deps, welcomeHandler}
}

func (h currencyHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	msg := i18n.CurrencyMessages[language.Language(p.Account.Language)]
	opts := []any{deps.ToForward(bot, p)}

	isNewcomer := p.Account.IsNewcomer()

	var page int
	if len(p.Args) >= 2 {
		return h.changeCurrency(bot, p, msg, isNewcomer, opts...)
	}

	getter := func(code string) string {
		currency, ok := currency.Get(currency.Currency(code))
		if !ok {
			return "N/A"
		}
		return currency.Flag + " " + currency.Name
	}

	keyboard, err := buildKeyboardList(bot.GetPlatform(), page, CURRENCY_CMD, currency.LimitedTargets, currency.Targets, getter)
	if err != nil {
		return err
	}

	opts = append(opts, keyboard)
	return bot.SendMessage(p.Message.Chat(), msg.ChooseCurrency+":", opts...)
}

func (h currencyHandler) changeCurrency(
	bot bin.Interface,
	p *domain.Payload,
	msg i18n.CurrencyLocale,
	isNewcomer bool,
	opts ...any,
) error {
	id, err := strconv.Atoi(p.Args[1])
	if err != nil {
		p.Account.Currency = p.Args[1]

		getter := func() (config.ItemInfo, bool) {
			return currency.Get(currency.Currency(p.Account.Currency))
		}
		setter := func() error {
			ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
			defer cancel()
			return h.AccountService.SetCurrency(ctx, p.Account.ID, p.Account.Currency)
		}

		if err := setValue(bot, p, msg.CurrencyChangedTo, getter, setter, opts...); err != nil {
			return err
		}

		if isNewcomer {
			p.Args = []string{WELCOME_CMD}
			return h.welcomeHandler.Execute(bot, p)
		}

		return nil
	}

	page := id - 1
	if page < 0 || page >= len(currency.LimitedTargets) {
		return bot.SendMessage(p.Message.Chat(), "Invalid page number", opts...)
	}

	return nil
}
