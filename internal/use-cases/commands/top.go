package commands

import (
	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/use-cases/commands/deps"
)

const TOP_PROMO_CMD = "toppromo"

type topHandler struct {
	deps.Dependencies
}

func NewTopHandler(deps deps.Dependencies) *topHandler {
	return &topHandler{deps}
}

func (c *topHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	// opts := []any{tools.ToForward(bot, p)}
	// return bot.SendMessage(p.Message.GetChat(), "Top N:", opts...)
	return nil
}
