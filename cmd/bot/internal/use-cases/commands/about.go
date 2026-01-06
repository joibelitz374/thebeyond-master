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

const ABOUT_CMD = "about"

type aboutHandler struct {
	deps deps.Dependencies
	path string
}

func NewAboutHandler(deps deps.Dependencies) aboutHandler {
	path := os.Getenv("ABOUT_IMAGE_PATH")
	if path == "" {
		path = "./assets/about.png"
	}
	return aboutHandler{deps, path}
}

func (h aboutHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	image, err := os.ReadFile(h.path)
	if err != nil {
		return err
	}

	msg := i18n.AboutMessages[p.Account.Language]

	return bot.SendMessage(p.Message.Chat(), "",
		types.NewAttachments().AddFile(bytes.NewReader(image)),
		types.NewKeyboard().
			NewRow(types.NewCallbackButton("📑 "+msg.TermsOfService, TOS_CMD)).
			NewRow(types.NewCallbackButton("🛡 "+msg.PrivacyPolicy, PRIVACY_CMD)).
			NewRow(types.NewCallbackButton("🛂 "+msg.RefundPolicy, REFUND_CMD)).
			// NewRow(types.NewCallbackButton("💳 "+msg.Billing, "billing")).
			// NewRow(types.NewCallbackButton("🆘 "+msg.FAQ, "faq")).
			NewRow(types.NewURLButton("☎️ "+msg.Support, "https://t.me/beyondsecurenews?direct")))
}
