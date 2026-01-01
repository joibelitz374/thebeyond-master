package commands

import (
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
)

const PRIVACY_CMD = "privacy"

type privacyHandler struct {
	deps.Dependencies
}

func NewPrivacyHandler(deps deps.Dependencies) privacyHandler {
	return privacyHandler{deps}
}

func (h privacyHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	msg := i18n.PrivacyPolicyMessages[p.Account.Language]
	controlMsg := i18n.ControlMessages[p.Account.Language]
	return bot.SendMessage(p.Message.Chat(), msg.Policy, types.NewKeyboard().
		NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, ABOUT_CMD)))
}
