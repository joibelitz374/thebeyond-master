package commands

import (
	"time"

	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/types"
	"github.com/quickpowered/frilly/internal/use-cases/commands/tools"
	"github.com/quickpowered/frilly/internal/use-cases/components"
	"github.com/quickpowered/frilly/pkg/consts"
	"github.com/quickpowered/frilly/pkg/utils"
)

const ACCOUNT_CMD = "account"

type AccountCmd struct {
	tools.Modules
	components components.Component
}

func NewAccountCmd(modules tools.Modules) *AccountCmd {
	return &AccountCmd{modules, components.NewAccountComponent()}
}

func (c *AccountCmd) Execute(bot bin.Interface, payload *domain.Payload) error {
	componentText := c.components.Text(payload.Account.Language)
	text := componentText[0] + ":\n"
	opts := []any{tools.ToForward(bot, payload), &types.Keyboard{
		ButtonRows: [][]types.Button{
			{{Text: "🎾 Renew", Data: "renew"}},
			{{Text: "🔄 New Key", Data: "newkey"}},
			{{Text: "⛔️ Refund", Data: "refund"}},
		},
	}}

	if payload.Account.SubscriptionExpiresAt != nil && payload.Account.SubscriptionExpiresAt.After(time.Now()) {
		serverLocation := locations[payload.Account.ServerLocation]
		key := utils.GenerateVLESSURI(payload.Account.KeyID, serverLocation.IP, "web.max.ru", payload.Account.ShortID, serverLocation.Flag+" "+serverLocation.Name)

		switch bot.GetPlatform() {
		case consts.PlatformTelegram:
			key = "<blockquote><code>" + key + "</code></blockquote>"
		}

		humanizeDuration := utils.HumanizeDuration(utils.TimeRemaining(*payload.Account.SubscriptionExpiresAt))
		if humanizeDuration == "0 minutes" {
			humanizeDuration = "waiting prolongation"
		}

		text += "\n🛜 " + componentText[1] + ":\n" + key +
			"\n🕒 " + componentText[2] + ": " + humanizeDuration
	} else {
		text += "\n🗒 " + componentText[3]
	}

	return bot.SendMessage(payload.Message.GetChat(), text, opts...)
}
