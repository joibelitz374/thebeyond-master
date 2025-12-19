package commands

import (
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
)

const NETWORK_CMD = "network"

type networkHandler struct {
	deps.Dependencies
}

func NewNetworkHandler(deps deps.Dependencies) networkHandler {
	return networkHandler{deps}
}

func (h networkHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	opts := []any{deps.ToForward(bot, p)}
	opts = append(opts, &types.Keyboard{
		ButtonRows: [][]types.Button{
			{{
				Text: "🏡 Домашняя",
				Data: "network home",
			}},
			{{
				Text: "🚌 Мобильная",
				Data: "network mobile",
			}},
			{{Text: "◀️ Назад", Data: SETTINGS_CMD}},
		},
	})
	return bot.SendMessage(p.Message.Chat(), "Выберите тип сети:", opts...)
}
