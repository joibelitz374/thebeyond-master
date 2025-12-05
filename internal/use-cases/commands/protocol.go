package commands

import (
	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/types"
	"github.com/quickpowered/frilly/internal/use-cases/commands/deps"
)

const PROTOCOL_CMD = "protocol"

type protocolHandler struct {
	deps.Dependencies
}

func NewProtocolHandler(deps deps.Dependencies) protocolHandler {
	return protocolHandler{deps}
}

func (h protocolHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	opts := []any{deps.ToForward(bot, p), types.DisableMentions, &types.Keyboard{
		ButtonRows: [][]types.Button{
			{{Text: "🍿 VLESS+XHTTP+Reality", Data: "protocol vxr"}},
			{{Text: "🎮 Hysteria 2", Data: "protocol hys2"}},
		},
	}}

	return bot.SendMessage(p.Message.Chat(), "Choose protocol:", opts...)
}
