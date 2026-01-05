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
	"github.com/quickpowered/thebeyond-master/pkg/utils"
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
	msg := i18n.NewKeyMessages[p.Account.Language]
	controlMsg := i18n.ControlMessages[p.Account.Language]

	if len(p.Args) >= 2 && p.Args[1] == "confirm" {
		return h.confirm(bot, p, msg, controlMsg)
	}

	return bot.SendMessage(p.Message.Chat(), msg.ConfirmPrompt, types.NewKeyboard().
		NewRow(
			types.NewCallbackButton("◀️ "+controlMsg.Back, MENU_CMD),
			types.NewCallbackButton("🎾 Confirm", NEWKEY_CMD+" confirm"),
		))
}

func (h newKeyHandler) confirm(bot bin.Interface, p *domain.Payload, msg i18n.NewKeyLocale, controlMsg i18n.ControlLocale) error {
	if !p.Account.IsActive() {
		return bot.SendMessage(p.Message.Chat(), "You don't have a subscription", types.NewKeyboard().
			NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, MENU_CMD)))
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	keyID, err := h.AccountService.RegenerateKey(ctx, p.Account.ID)
	if err != nil {
		return err
	}

	regions, err := h.SubscriptionService.GetRegions(p.Account.Region)
	if err != nil {
		return err
	}

	for _, region := range regions {
		if err := h.XRayManagerService.RemoveClient(ctx, region, utils.NewEmail(p.Account.ID)); err != nil {
			log.Error("failed to remove client", zap.Error(err))
		}

		if err := h.XRayManagerService.AddClient(ctx, region, keyID, utils.NewEmail(p.Account.ID)); err != nil {
			return err
		}
	}

	return bot.SendMessage(p.Message.Chat(), msg.SuccessMessage, types.NewKeyboard().
		NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, MENU_CMD)))
}
