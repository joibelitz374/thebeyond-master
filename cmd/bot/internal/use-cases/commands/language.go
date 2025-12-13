package commands

import (
	"context"
	"strconv"
	"time"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	"github.com/quickpowered/thebeyond-master/configs"
	"github.com/quickpowered/thebeyond-master/configs/language"
)

const LANGUAGE_CMD = "language"

type languageHandler struct {
	deps.Dependencies
	currencyHandler Command
}

func NewLanguageHandler(deps deps.Dependencies) languageHandler {
	currencyHandler := NewCurrencyHandler(deps)
	return languageHandler{deps, currencyHandler}
}

func (h languageHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	isNewcomer := p.Account.IsNewcomer()
	if isNewcomer {
		p.Account.Language = "en"
	}

	msg := i18n.LanguageMessages[language.Language(p.Account.Language)]
	opts := []any{deps.ToForward(bot, p)}

	var page int
	if len(p.Args) >= 2 {
		return h.changeLanguage(bot, p, msg, isNewcomer, opts...)
	}

	getter := func(code string) string {
		language, ok := language.Get(language.Language(code))
		if !ok {
			return "Undefined"
		}
		return language.Flag + " " + language.Name
	}

	keyboard, err := buildKeyboardList(bot.GetPlatform(), page, LANGUAGE_CMD, language.LimitedTargets, language.Targets, getter)
	if err != nil {
		return err
	}

	opts = append(opts, keyboard)
	return bot.SendMessage(p.Message.Chat(), msg.ChooseLanguage+":", opts...)
}

func (h languageHandler) changeLanguage(
	bot bin.Interface,
	p *domain.Payload,
	msg i18n.LanguageLocale,
	isNewcomer bool,
	opts ...any,
) error {
	id, err := strconv.Atoi(p.Args[1])
	if err != nil {
		p.Account.Language = p.Args[1]

		getter := func() (configs.ItemInfo, bool) { return language.Get(language.Language(p.Account.Language)) }
		setter := func() error {
			ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
			defer cancel()
			return h.AccountService.SetLanguage(ctx, p.Account.ID, p.Account.Language)
		}

		if err := setValue(bot, p, msg.LanguageChangedTo, getter, setter, opts...); err != nil {
			return err
		}

		if isNewcomer {
			p.Args = []string{CURRENCY_CMD}
			return h.currencyHandler.Execute(bot, p)
		}
		return nil
	}

	page := id
	if page < 1 || page >= len(language.LimitedTargets) {
		return bot.SendMessage(p.Message.Chat(), "Invalid page number", opts...)
	}

	return nil
}
