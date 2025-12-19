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

	language := language.Language(p.Account.Language)
	msg := i18n.AboutMessages[language]
	opts := []any{
		deps.ToForward(bot, p),
		h.buildKeyboard(msg, language),
		types.NewAttachments().AddFile(bytes.NewReader(image)),
	}

	return bot.SendMessage(p.Message.Chat(), "", opts...)
}

func (h aboutHandler) buildKeyboard(msg i18n.AboutLocale, language language.Language) *types.Keyboard {
	controlMsg := i18n.ControlMessages[language]
	return &types.Keyboard{
		ButtonRows: [][]types.Button{
			{{Text: "📑 " + msg.TermsOfService, Data: "tos"}},
			{{Text: "🛡 " + msg.PrivacyPolicy, Data: "privacy"}},
			{{Text: "🛂 " + msg.RefundPolicy, Data: "refund"}},
			// {{Text: "💳 Оплата и биллинг", Data: "billing"}},
			// {{Text: "🆘 FAQ / Справка", Data: "faq"}},
			{{Text: "☎️ " + msg.Support, URL: "https://t.me/beyondsecurenews?direct"}},
			{{Text: "◀️ " + controlMsg.Back, Data: MENU_CMD}},
		},
	}
}
