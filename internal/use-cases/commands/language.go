package commands

import (
	"context"
	"strconv"
	"time"

	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/types"
	"github.com/quickpowered/frilly/internal/use-cases/commands/tools"
	"github.com/quickpowered/frilly/internal/use-cases/components"
	"github.com/quickpowered/frilly/pkg/consts"
	"github.com/quickpowered/frilly/pkg/values"
)

const LANGUAGE_CMD = "language"

var languagesOrder = [][]string{
	{"en", "de", "nl"},
	{"sv", "no", "da"},
	{"es", "fr"},
	{"pt", "it"},
	{"ru", "ua", "pl"},
	{"cs", "bg", "sr"},
	{"hr", "sk", "sl"},
	{"lt", "lv"},
	{"et", "fi"},
	{"el", "ro", "hu"},
	{"ar", "fa"},
	{"tr", "he"},
	{"zh", "ja", "ko"},
	{"vi", "th", "id"},
	{"ms", "tl"},
	{"hi", "ur", "bn"},
	{"ta", "te", "mr"},
}

var vkLanguagesPages = [][][]string{
	{
		{"en", "de"},
		{"nl", "da"},
		{"sv", "no"},
		{"fi", "et"},
	},
	{
		{"fr", "it"},
		{"es", "pt"},
		{"ro", "el"},
	},
	{
		{"ru", "ua"},
		{"pl", "cs"},
		{"sk", "hu"},
		{"lt", "lv"},
	},
	{
		{"sl", "hr"},
		{"sr", "bg"},
		{"tr", "he"},
		{"ar", "fa"},
	},
	{
		{"hi", "ur"},
		{"bn", "mr"},
		{"ta", "te"},
	},
	{
		{"zh", "ja"},
		{"ko", "vi"},
		{"th", "id"},
		{"ms", "tl"},
	},
}

type LanguageCmd struct {
	tools.Modules
	component components.Component
}

func NewLanguageCmd(modules tools.Modules) *LanguageCmd {
	return &LanguageCmd{modules, components.NewLanguageComponent()}
}

func (c *LanguageCmd) Execute(bot bin.Interface, payload *domain.Payload) error {
	var isStart bool
	if payload.Account.Language == "" {
		payload.Account.Language = "en"
		isStart = true
	}

	componentText := c.component.Text(payload.Account.Language)
	opts := []any{tools.ToForward(bot, payload)}

	var page int
	if len(payload.NodeRoute) >= 2 {
		arg, err := strconv.Atoi(payload.NodeRoute[1])
		if err != nil {
			payload.Account.Language = payload.NodeRoute[1]
			getter := func() (values.Value, bool) { return values.GetLanguage(payload.Account.Language) }
			setter := func() error {
				ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
				defer cancel()
				return c.AccountService.SetLanguage(ctx, payload.Account.ID, payload.Account.Language)
			}

			if err := setValue(bot, payload, componentText, getter, setter, opts...); err != nil {
				return err
			}

			if isStart {
				payload.NodeRoute = []string{CURRENCY_CMD}
				return NewCurrencyCmd(c.Modules).Execute(bot, payload)
			}
			return nil
		} else {
			page = arg
			if page < 1 || page >= len(vkLanguagesPages) {
				return bot.SendMessage(payload.Message.GetChat(), "Invalid page number", opts...)
			}
		}
	}

	getter := func(code string) string {
		language, ok := values.GetLanguage(code)
		if !ok {
			return "Undefined"
		}
		return language.Flag + " " + language.Name
	}

	buttonRows, err := getButtons(bot, page, "language", vkLanguagesPages, languagesOrder, getter)
	if err != nil {
		bot.SendMessage(payload.Message.GetChat(), "Invalid language code", opts...)
		return err
	}

	if bot.GetPlatform() != consts.PlatformTelegram {
		buttonRows = append(buttonRows, addNavigationButtons(bot, page)...)
	}

	opts = append(opts, &types.Keyboard{ButtonRows: buttonRows})
	return bot.SendMessage(payload.Message.GetChat(), componentText[0]+":", opts...)
}

func getButtons(bot bin.Interface,
	page int, cmd string,
	pages [][][]string, order [][]string,
	getter func(string) string,
) ([][]types.Button, error) {
	if bot.GetPlatform() != consts.PlatformTelegram {
		order = pages[page]
	}

	buttonRows, idx := make([][]types.Button, len(order)), 0
	for _, rows := range order {
		for _, code := range rows {
			buttonRows[idx] = append(buttonRows[idx], types.Button{
				Text: getter(code),
				Data: cmd + " " + code,
			})
		}
		idx++
	}

	return buttonRows, nil
}

func addNavigationButtons(bot bin.Interface, page int) [][]types.Button {
	buttons := make([][]types.Button, 1)
	if page > 0 {
		buttons[0] = append(buttons[0], types.Button{
			Text: "⬅️",
			Data: "language " + strconv.Itoa(page-1),
		})
	}
	if page < len(vkLanguagesPages)-1 {
		buttons[0] = append(buttons[0], types.Button{
			Text: "➡️",
			Data: "language " + strconv.Itoa(page+1),
		})
	}

	return buttons
}
