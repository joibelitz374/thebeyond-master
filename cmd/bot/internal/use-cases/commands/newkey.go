package commands

import (
	"context"
	"time"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	"github.com/quickpowered/thebeyond-master/configs/language"
	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager/dto"
	"github.com/quickpowered/thebeyond-master/pkg/email"
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

			keyID, err := h.AccountService.RegenerateKey(ctx, p.Account.ID)
			if err != nil {
				return err
			}

			if err := h.XRayManagerRepo.RemoveClient(ctx, dto.ClusterID(p.Account.ClusterID), dto.NodeTypeBlacklist, email.NewAccount(p.Account.ID)); err != nil {
				return err
			}

			if err := h.XRayManagerRepo.AddClient(ctx, dto.ClusterID(p.Account.ClusterID), dto.NodeTypeBlacklist, keyID, email.NewAccount(p.Account.ID)); err != nil {
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
