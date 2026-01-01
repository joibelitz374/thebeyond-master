package commands

import (
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
)

const TOS_CMD = "tos"

type tosHandler struct {
	deps.Dependencies
}

func NewTosHandler(deps deps.Dependencies) tosHandler {
	return tosHandler{deps}
}

func (h tosHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	msg := i18n.TosMessages[p.Account.Language]
	controlMsg := i18n.ControlMessages[p.Account.Language]
	return bot.SendMessage(p.Message.Chat(), msg.Terms, types.NewKeyboard().
		NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, ABOUT_CMD)))
}
