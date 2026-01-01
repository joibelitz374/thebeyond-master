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
	return bot.SendMessage(p.Message.Chat(), "Выберите тип сети:", types.NewKeyboard().
		NewRow(types.NewCallbackButton("🏡 Домашняя", "network home")).
		NewRow(types.NewCallbackButton("🚌 Мобильная", "network mobile")).
		NewRow(types.NewCallbackButton("◀️ Назад", SETTINGS_CMD)))
}
