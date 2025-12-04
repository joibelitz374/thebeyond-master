package commands

import (
	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/types"
	"github.com/quickpowered/frilly/internal/use-cases/commands/tools"
)

const PROTOCOL_CMD = "protocol"

type ProtocolCmd struct {
	tools.Modules
}

func NewProtocolCmd(modules tools.Modules) *ProtocolCmd {
	return &ProtocolCmd{modules}
}

func (c *ProtocolCmd) Execute(bot bin.Interface, payload *domain.Payload) error {
	opts := []any{tools.ToForward(bot, payload), types.DisableMentions, &types.Keyboard{
		ButtonRows: [][]types.Button{
			{{Text: "🍿 VLESS+XHTTP+Reality", Data: "protocol vxr"}},
			{{Text: "🎮 Hysteria 2", Data: "protocol hys2"}},
		},
	}}

	return bot.SendMessage(payload.Message.GetChat(), "Choose protocol:", opts...)
}
