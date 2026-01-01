package commands

import (
	"bytes"
	"os"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
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
	msg := i18n.SettingsMessages[p.Account.Language]
	controlMsg := i18n.ControlMessages[p.Account.Language]

	image, err := os.ReadFile(h.path)
	if err != nil {
		return err
	}

	return bot.SendMessage(p.Message.Chat(), "",
		types.NewAttachments().AddFile(bytes.NewReader(image)),
		types.NewKeyboard().
			// NewRow(types.NewCallbackButton("⭐️ "+msg.Custom, "custom")).
			// NewRow(
			// 	types.NewCallbackButton("💟 "+msg.Region, "region"),
			// 	types.NewCallbackButton("🗺 "+msg.Locations, "location"),
			// ).
			// NewRow(
			// 	types.NewCallbackButton("🧠 "+msg.SMART, "mode smart"),
			// 	types.NewCallbackButton("🏡 "+msg.Network, "network"),
			// ).
			NewRow(
				types.NewCallbackButton("💬 "+msg.Language, LANGUAGE_CMD),
				types.NewCallbackButton("🤑 "+msg.Currency, CURRENCY_CMD),
			).
			NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, MENU_CMD)))
}
