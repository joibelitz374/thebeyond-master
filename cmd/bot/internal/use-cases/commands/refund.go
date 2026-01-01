package commands

import (
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
)

const REFUND_CMD = "refund"

type refundHandler struct {
	deps.Dependencies
}

func NewRefundHandler(deps deps.Dependencies) refundHandler {
	return refundHandler{deps}
}

func (h refundHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	msg := i18n.RefundPolicyMessages[p.Account.Language]
	controlMsg := i18n.ControlMessages[p.Account.Language]
	return bot.SendMessage(p.Message.Chat(), msg.Policy, types.NewKeyboard().
		NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, ABOUT_CMD)))
}
