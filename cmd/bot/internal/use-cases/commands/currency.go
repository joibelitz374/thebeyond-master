package commands

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/telegram"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	"github.com/quickpowered/thebeyond-master/configs"
	"github.com/quickpowered/thebeyond-master/configs/currency"
	"github.com/quickpowered/thebeyond-master/configs/language"
)

const CURRENCY_CMD = "currency"

type currencyHandler struct {
	deps.Dependencies
	menuHandler Command
}

func NewCurrencyHandler(deps deps.Dependencies) currencyHandler {
	menuHandler := NewMenuHandler(deps)
	return currencyHandler{deps, menuHandler}
}

func (h currencyHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	isNewcomer := p.Account.IsNewcomer()
	language := language.Language(p.Account.Language)
	msg := i18n.CurrencyMessages[language]
	controlMsg := i18n.ControlMessages[language]
	opts := []any{deps.ToForward(bot, p)}
	var page int

	if len(p.Args) >= 2 {
		trialMsg := i18n.TrialMessages[language]
		return h.changeCurrency(bot, p, msg, trialMsg, isNewcomer)
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

	if !isNewcomer {
		keyboard.ButtonRows = append(keyboard.ButtonRows, []types.Button{{
			Text: "◀️ " + controlMsg.Back,
			Data: SETTINGS_CMD,
		}})
	}

	opts = append(opts, keyboard)
	return bot.SendMessage(p.Message.Chat(), msg.ChooseCurrency+":", opts...)
}

func (h currencyHandler) changeCurrency(
	bot bin.Interface,
	p *domain.Payload,
	msg i18n.CurrencyLocale,
	trialMsg i18n.TrialLocale,
	isNewcomer bool,
) error {
	id, err := strconv.Atoi(p.Args[1])
	if err != nil {
		p.Account.Currency = p.Args[1]

		getter := func() (configs.ItemInfo, bool) {
			return currency.Get(currency.Currency(p.Account.Currency))
		}
		setter := func() error {
			ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
			defer cancel()
			return h.AccountService.SetCurrency(ctx, p.Account.ID, p.Account.Currency)
		}

		if err := setValue(bot, p, msg.CurrencyChangedTo, getter, setter); err != nil {
			return err
		}

		if isNewcomer {
			if err := bot.SendMessage(p.Message.Chat(), trialMsg.Message); err != nil {
				return err
			}

			p.Args = []string{MENU_CMD}
			return h.menuHandler.Execute(bot, p)
		}
		p.Args = []string{CURRENCY_CMD}
		return h.Execute(bot, p)
	}

	page := id - 1
	if page < 0 || page >= len(currency.LimitedTargets) {
		tgBot, ok := bot.(*telegram.Adapter)
		if !ok {
			return errors.New("failed to cast bot to telegram adapter")
		}

		if err := tgBot.AnswerCallbackQuery(p.Message.CallbackQueryID(), "Invalid page number"); err != nil {
			return err
		}

		p.Args = []string{CURRENCY_CMD}
		return h.Execute(bot, p)
	}

	return nil
}
