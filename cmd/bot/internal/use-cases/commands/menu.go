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
	"github.com/quickpowered/thebeyond-master/pkg/utils"
)

const MENU_CMD = "start"

type menuHandler struct {
	deps deps.Dependencies
	path string
}

func NewMenuHandler(deps deps.Dependencies) menuHandler {
	path := os.Getenv("MENU_IMAGE_PATH")
	if path == "" {
		path = "./assets/menu.png"
	}

	return menuHandler{deps, path}
}

func (h menuHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	image, err := os.ReadFile(h.path)
	if err != nil {
		return err
	}

	msg := i18n.AccountMessages[language.Language(p.Account.Language)]
	opts := []any{
		deps.ToForward(bot, p),
		h.buildKeyboard(msg),
		types.NewAttachments().AddFile(bytes.NewReader(image)),
	}

	remaining := utils.HumanizeDuration(utils.TimeRemaining(*p.Account.SubscriptionExpiresAt))
	if remaining == "0 minutes" {
		remaining = "waiting prolongation"
	}

	return bot.SendMessage(p.Message.Chat(), "🕒 "+msg.Remaining+": "+remaining, opts...)
}

func (h menuHandler) buildKeyboard(msg i18n.AccountLocale) *types.Keyboard {
	return &types.Keyboard{
		ButtonRows: [][]types.Button{
			{{Text: "🛜 " + msg.Connect, Data: "connect"}},
			{{Text: "🛍 " + msg.Renew, Data: "renew"}, {Text: "🔄 " + msg.NewKey, Data: "newkey"}},
			{{Text: "🌈 " + msg.FREEDAYS + " 🔥", Data: "promo"}},
			{{Text: "ℹ️ " + msg.AboutServiceAndTerms, Data: "about"}},
			{{Text: "💊 " + msg.Settings, Data: "settings"}},
		},
	}
}
