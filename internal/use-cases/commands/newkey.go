package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/types"
	"github.com/quickpowered/frilly/internal/use-cases/commands/tools"
	"github.com/quickpowered/frilly/internal/use-cases/components"
)

const NEWKEY_CMD = "newkey"

type NewKeyCmd struct {
	tools.Modules
	component components.Component
}

func NewNewKeyCmd(modules tools.Modules) *NewKeyCmd {
	return &NewKeyCmd{modules, components.NewNewKeyComponent()}
}

func (c *NewKeyCmd) Execute(bot bin.Interface, payload *domain.Payload) error {
	componentText := c.component.Text(payload.Account.Language)
	opts := []any{tools.ToForward(bot, payload)}

	if len(payload.NodeRoute) >= 2 {
		if payload.NodeRoute[1] == "confirm" {
			if payload.Account.SubscriptionExpiresAt != nil && payload.Account.SubscriptionExpiresAt.Before(time.Now()) {
				return bot.SendMessage(payload.Message.GetChat(), "You don't have a subscription", opts...)
			}

			ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
			defer cancel()

			if err := c.XRayClient.RemoveUser(fmt.Sprintf("id%d@user", payload.Account.ID)); err != nil {
				return err
			}

			keyID, err := c.AccountService.RegenerateKey(ctx, payload.Account.ID)
			if err != nil {
				return err
			}

			fmt.Println("keyID", keyID)
			if err := c.XRayClient.AddUser(fmt.Sprintf("id%d@user", payload.Account.ID), keyID); err != nil {
				return err
			}

			return bot.SendMessage(payload.Message.GetChat(), componentText[1], opts...)
		}
	}

	opts = append(opts, &types.Keyboard{
		ButtonRows: [][]types.Button{
			{{Text: "🎾 Confirm", Data: "newkey confirm"}},
		},
	})

	return bot.SendMessage(payload.Message.GetChat(), componentText[0], opts...)
}
