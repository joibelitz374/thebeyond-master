package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/quickpowered/frilly/config/language"
	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/i18n"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/types"
	"github.com/quickpowered/frilly/internal/use-cases/commands/deps"
)

const NEWKEY_CMD = "newkey"

type newKeyHandler struct {
	deps.Dependencies
}

func NewNewKeyHandler(deps deps.Dependencies) newKeyHandler {
	return newKeyHandler{deps}
}

func (h newKeyHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	msg := i18n.NewKeyMessages[language.Language(p.Account.Language)]
	opts := []any{deps.ToForward(bot, p)}

	if len(p.Args) >= 2 {
		if p.Args[1] == "confirm" {
			if !p.Account.IsActive() {
				return bot.SendMessage(p.Message.Chat(), "You don't have a subscription", opts...)
			}

			ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
			defer cancel()

			if err := h.XRayClient.RemoveUser(fmt.Sprintf("id%d@user", p.Account.ID)); err != nil {
				return err
			}

			keyID, err := h.AccountService.RegenerateKey(ctx, p.Account.ID)
			if err != nil {
				return err
			}

			if err := h.XRayClient.AddUser(fmt.Sprintf("id%d@user", p.Account.ID), keyID); err != nil {
				return err
			}

			return bot.SendMessage(p.Message.Chat(), msg.SuccessMessage, opts...)
		}
	}

	opts = append(opts, &types.Keyboard{
		ButtonRows: [][]types.Button{
			{{Text: "🎾 Confirm", Data: "newkey confirm"}},
		},
	})

	return bot.SendMessage(p.Message.Chat(), msg.ConfirmPrompt, opts...)
}
