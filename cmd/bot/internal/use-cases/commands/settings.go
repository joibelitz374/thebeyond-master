package commands

import (
	"bytes"
	"os"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	"github.com/quickpowered/thebeyond-master/configs/language"
)

const SETTINGS_CMD = "settings"

type settingsHandler struct {
	deps deps.Dependencies
	path string
}

func NewSettingsHandler(deps deps.Dependencies) settingsHandler {
	path := os.Getenv("CONFIG_IMAGE_PATH")
	if path == "" {
		path = "./assets/config.png"
	}

	return settingsHandler{deps, path}
}

func (h settingsHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	image, err := os.ReadFile(h.path)
	if err != nil {
		return err
	}

	language := language.Language(p.Account.Language)
	msg := i18n.SettingsMessages[language]
	controlMsg := i18n.ControlMessages[language]
	opts := []any{
		deps.ToForward(bot, p),
		h.buildKeyboard(msg, controlMsg),
		types.NewAttachments().AddFile(bytes.NewReader(image)),
	}
	return bot.SendMessage(p.Message.Chat(), "", opts...)
}

func (h settingsHandler) buildKeyboard(msg i18n.SettingsLocale, controlMsg i18n.ControlLocale) *types.Keyboard {
	return &types.Keyboard{
		ButtonRows: [][]types.Button{
			// {{Text: "⭐️ " + msg.Custom, Data: "custom"}},
			{
				// {Text: "💟 " + msg.Region, Data: "region"},
				// {Text: "🗺 " + msg.Locations, Data: "location"},
			},
			{
				// {Text: "🧠 " + msg.Smart, Data: "mode smart"},
				// {Text: "🏡 " + msg.Network, Data: "network"},
			},
			{
				{Text: "💬 " + msg.Language, Data: "language"},
				{Text: "🤑 " + msg.Currency, Data: "currency"},
			},
			{{Text: "◀️ " + controlMsg.Back, Data: MENU_CMD}},
		},
	}
}
