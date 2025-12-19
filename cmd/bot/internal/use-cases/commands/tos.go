package commands

import (
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	"github.com/quickpowered/thebeyond-master/configs/language"
)

const TOS_CMD = "tos"

type tosHandler struct {
	deps.Dependencies
}

func NewTosHandler(deps deps.Dependencies) tosHandler {
	return tosHandler{deps}
}

func (h tosHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	language := language.Language(p.Account.Language)
	msg := i18n.TosMessages[language]
	controlMsg := i18n.ControlMessages[language]
	opts := []any{deps.ToForward(bot, p), &types.Keyboard{
		ButtonRows: [][]types.Button{{{Text: "◀️ " + controlMsg.Back, Data: ABOUT_CMD}}},
	}}
	return bot.SendMessage(p.Message.Chat(), msg.Terms, opts...)
}
