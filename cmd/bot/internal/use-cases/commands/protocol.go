package commands

import (
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
)

const PROTOCOL_CMD = "protocol"

type protocolHandler struct {
	deps.Dependencies
}

func NewProtocolHandler(deps deps.Dependencies) protocolHandler {
	return protocolHandler{deps}
}

func (h protocolHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	text := "Выберите протокол:\n\n" +
		"Amnezia WG (AWG) — усовершенствованная версия WireGuard. Это — быстрый и безопасный VPN. Ограничивается нами количеством устройств;\n\n" +
		"XRay — усовершенствованная версия V2Ray с протоколами для регионов с наивысшей цензурой. Это — более медленный, но трудно обнаруживаемый регуляторами прокси. Ограничивается нами потребляемым трафиком;"
	return bot.SendMessage(p.Message.Chat(), text, types.NewKeyboard().
		NewRow(types.NewCallbackButton("🎮 Amnezia WG", "protocol awg")).
		NewRow(types.NewCallbackButton("🍿 XRay", "protocol vxr")))
}
