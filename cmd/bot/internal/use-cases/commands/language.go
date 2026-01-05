package commands

import (
	"context"
	"time"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/interfaces"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	"github.com/quickpowered/thebeyond-master/configs"
	"github.com/quickpowered/thebeyond-master/configs/language"
)

const LANGUAGE_CMD = "language"

type languageHandler struct {
	deps.Dependencies
	currencyHandler interfaces.Command
}

func NewLanguageHandler(deps deps.Dependencies) languageHandler {
	currencyHandler := NewCurrencyHandler(deps)
	return languageHandler{deps, currencyHandler}
}

func (h languageHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	isNewcomer := p.Account.IsNewcomer()
	if isNewcomer {
		p.Account.Language = language.English
	}

	msg := i18n.LanguageMessages[p.Account.Language]
	controlMsg := i18n.ControlMessages[p.Account.Language]
	if len(p.Args) >= 2 {
		return h.changeLanguage(bot, p, msg, isNewcomer)
	}

	buttons, err := getButtons(LANGUAGE_CMD, language.Targets, func(code string) string {
		language, ok := language.Get(language.Language(code))
		if !ok {
			return "Undefined"
		}
		return language.Flag + " " + language.Name
	})
	if err != nil {
		return err
	}

	if !isNewcomer {
		buttons = append(buttons, []types.Button{types.NewCallbackButton("◀️ "+controlMsg.Back, SETTINGS_CMD)})
	}

	return bot.SendMessage(p.Message.Chat(), msg.ChooseLanguage+":", &types.Keyboard{ButtonRows: buttons})
}

func (h languageHandler) changeLanguage(bot bin.Interface, p *domain.Payload, msg i18n.LanguageLocale, isNewcomer bool) error {
	p.Account.Language = language.Language(p.Args[1])

	getter := func() (configs.ItemInfo, bool) {
		return language.Get(language.Language(p.Account.Language))
	}

	setter := func() error {
		ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
		defer cancel()
		return h.AccountService.SetLanguage(ctx, p.Account.ID, string(p.Account.Language))
	}

	if err := setValue(bot, p, msg.LanguageChangedTo, getter, setter); err != nil {
		return err
	}

	if isNewcomer {
		p.Args = []string{CURRENCY_CMD}
		return h.currencyHandler.Execute(bot, p)
	}

	p.Args = []string{LANGUAGE_CMD}
	return h.Execute(bot, p)
}
