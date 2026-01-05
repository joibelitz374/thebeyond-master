package commands

import (
	"bytes"
	"os"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/renew"
)

const MENU_CMD = "start"

type menuHandler struct {
	deps             deps.Dependencies
	tariffImagePaths map[int]string
}

func NewMenuHandler(deps deps.Dependencies) menuHandler {
	return menuHandler{deps, map[int]string{
		-1: "/shared/assets/expired.png",
		0:  "/shared/assets/free.png",
		1:  "/shared/assets/personal.png",
		2:  "/shared/assets/home.png",
		3:  "/shared/assets/unlimited.png",
	}}
}

func (h menuHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	msg := i18n.AccountMessages[p.Account.Language]
	var tariffImagePath string
	if p.Account.IsActive() {
		tariffImagePath = h.tariffImagePaths[p.Account.Tariff]
	} else {
		tariffImagePath = h.tariffImagePaths[-1]
	}

	image, err := os.ReadFile(tariffImagePath)
	if err != nil {
		return err
	}

	return bot.SendMessage(p.Message.Chat(), "",
		types.NewAttachments().AddFile(bytes.NewReader(image)),
		types.NewKeyboard().
			NewRow(types.NewCallbackButton("🛜 "+msg.Connect, CONNECT_CMD)).
			NewRow(
				types.NewCallbackButton("🛍 "+msg.Tariffs, renew.CMD),
				types.NewCallbackButton("🔄 "+msg.NewKey, NEWKEY_CMD),
			).
			NewRow(
				types.NewURLButton("📢 "+msg.News, "https://t.me/beyondsecurenews"),
				types.NewURLButton("💬 "+msg.Reviews, "https://t.me/thebeyondreviews"),
			).
			NewRow(types.NewCallbackButton("🌈 "+msg.FREEDAYS+" 🔥", PROMO_CMD)).
			NewRow(types.NewCallbackButton("💊 "+msg.Settings, SETTINGS_CMD)))
}
