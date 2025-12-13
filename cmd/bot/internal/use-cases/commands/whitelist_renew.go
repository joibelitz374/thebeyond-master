package commands

import (
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
)

const WHITELIST_RENEW_CMD = "whitelist_renew"

type whitelistRenewHandler struct {
	deps.Dependencies
}

func NewWhitelistRenewHandler(deps deps.Dependencies) *whitelistRenewHandler {
	return &whitelistRenewHandler{deps}
}

func (c *whitelistRenewHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	return bot.SendMessage(p.Message.Chat(), "OK!")
}
