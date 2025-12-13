package commands

import (
	"fmt"
	"strconv"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
)

const SET_PRICES_CMD = "setprices"

type setPricesHandler struct {
	deps.Dependencies
}

func NewSetPricesHandler(deps deps.Dependencies) *setPricesHandler {
	return &setPricesHandler{deps}
}

func (c *setPricesHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	if p.Message.Sender() != 924536264 {
		return bot.SendMessage(p.Message.Chat(), "You are not allowed to use this command")
	}

	if len(p.Args) < 3 {
		return bot.SendMessage(p.Message.Chat(), "Usage: /setprices [+/-] [multiplier]")
	}

	multiplier, err := strconv.ParseFloat(p.Args[2], 64)
	if err != nil {
		return bot.SendMessage(p.Message.Chat(), "Invalid multiplier")
	}

	switch p.Args[1] {
	case "+":
		if err := c.SubscriptionsRepo.IncreasePrices(multiplier); err != nil {
			return err
		}
	case "-":
		if err := c.SubscriptionsRepo.DecreasePrices(multiplier); err != nil {
			return err
		}
	default:
		return bot.SendMessage(p.Message.Chat(), "Invalid operator")
	}

	return bot.SendMessage(p.Message.Chat(), fmt.Sprintf("Successfully %s prices by %.2f", p.Args[1], multiplier))
}
