package commands

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3/log"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	"github.com/quickpowered/thebeyond-master/configs/language"
	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager/dto"
	"github.com/quickpowered/thebeyond-master/pkg/email"
	"go.uber.org/zap"
)

const NEWKEY_CMD = "newkey"

type newKeyHandler struct {
	deps.Dependencies
}

func NewNewKeyHandler(deps deps.Dependencies) newKeyHandler {
	return newKeyHandler{deps}
}

func (h newKeyHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	language := language.Language(p.Account.Language)
	msg := i18n.NewKeyMessages[language]
	controlMsg := i18n.ControlMessages[language]
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

			for _, region := range []dto.Region{dto.RegionRussia} {
				if err := h.XRayManagerRepo.RemoveClient(ctx, region, email.NewAccount(p.Account.ID)); err != nil {
					log.Error("failed to remove client", zap.Error(err))
				}

				if err := h.XRayManagerRepo.AddClient(ctx, region, keyID, email.NewAccount(p.Account.ID)); err != nil {
					log.Error("failed to add client", zap.Error(err))
				}
			}

			opts = append(opts, &types.Keyboard{ButtonRows: [][]types.Button{{
				{Text: "◀️ " + controlMsg.Back, Data: MENU_CMD},
			}}})
			return bot.SendMessage(p.Message.Chat(), msg.SuccessMessage, opts...)
		}
	}

	opts = append(opts, &types.Keyboard{
		ButtonRows: [][]types.Button{
			{{Text: "🎾 Confirm", Data: "newkey confirm"}},
			{{Text: "◀️ " + controlMsg.Back, Data: MENU_CMD}},
		},
	})

	return bot.SendMessage(p.Message.Chat(), msg.ConfirmPrompt, opts...)
}
