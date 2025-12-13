package commands

import (
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
)

const TOP_PROMO_CMD = "toppromo"

type topHandler struct {
	deps.Dependencies
}

func NewTopHandler(deps deps.Dependencies) *topHandler {
	return &topHandler{deps}
}

func (c *topHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	return nil
}
