package commands

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/interfaces"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/middlewares"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/renew"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/promo"
)

const START_CMD = "start"

type UseCase interface {
	Run(bot bin.Interface, payload *domain.Payload) error
}

type useCase struct {
	deps.Dependencies
	promoUseCase promo.UseCase
	commands     map[string]interfaces.Command
}

func NewUseCase(deps_ deps.Dependencies, promoUseCase promo.UseCase) useCase {
	wrap := func(cmd interfaces.Command, name string) interfaces.Command {
		return middlewares.WithLogging(middlewares.WithDeletion(cmd, deps_.Logger), deps_.Logger, name)
	}

	commands := map[string]interfaces.Command{
		MENU_CMD:     wrap(NewMenuHandler(deps_), MENU_CMD),
		ABOUT_CMD:    wrap(NewAboutHandler(deps_), ABOUT_CMD),
		SETTINGS_CMD: wrap(NewSettingsHandler(deps_), SETTINGS_CMD),
		TOS_CMD:      wrap(NewTosHandler(deps_), TOS_CMD),
		PRIVACY_CMD:  wrap(NewPrivacyHandler(deps_), PRIVACY_CMD),
		REFUND_CMD:   wrap(NewRefundHandler(deps_), REFUND_CMD),
		CONNECT_CMD:  wrap(NewConnectHandler(deps_), CONNECT_CMD),
		PROMO_CMD:    wrap(NewPromoCmd(deps_), PROMO_CMD),
		renew.CMD:    wrap(renew.NewHandler(deps_), renew.CMD),
		NEWKEY_CMD:   wrap(NewNewKeyHandler(deps_), NEWKEY_CMD),
		REGION_CMD:   wrap(NewRegionHandler(deps_), REGION_CMD),
		// NETWORK_CMD:  NewNetworkHandler(deps_),
		LANGUAGE_CMD: wrap(NewLanguageHandler(deps_), LANGUAGE_CMD),
		CURRENCY_CMD: wrap(NewCurrencyHandler(deps_), CURRENCY_CMD),
		// PULL_CMD:       NewPullHandler(deps_),
		PROTOCOL_CMD: wrap(NewProtocolHandler(deps_), PROTOCOL_CMD),
		SUPPORT_CMD:  wrap(NewSupportHandler(deps_), SUPPORT_CMD),
		REPOST_CMD:   wrap(NewRepostHandler(deps_), REPOST_CMD),
	}

	return useCase{deps_, promoUseCase, commands}
}

func (uc useCase) Run(bot bin.Interface, p *domain.Payload) (err error) {
	text := strings.TrimSpace(p.Message.Text())
	if text == "" {
		return nil
	}

	if text[0] == '/' {
		text = text[1:]
	}

	p.Args = strings.Fields(text)
	if len(p.Args) == 0 {
		return nil
	}

	cmdName := strings.ToLower(p.Args[0])

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	p.Account, err = uc.AccountService.Get(ctx, bot.GetPlatform(), p.Message.Sender())
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return uc.createAccount(bot, bot.GetPlatform(), p)
		} else {
			return fmt.Errorf("failed to get account: %w", err)
		}
	}

	if p.Account.ID == 0 {
		return fmt.Errorf("account not found")
	}

	switch {
	case p.Account.Language == "":
		cmdName = LANGUAGE_CMD
	case p.Account.Currency == "":
		cmdName = CURRENCY_CMD
	}

	cmd, ok := uc.commands[cmdName]
	if !ok {
		return bot.SendMessage(p.Message.Chat(), "Command not found", types.NewKeyboard().
			NewRow(types.NewCallbackButton("🔄 Menu", MENU_CMD)))
	}

	if err := cmd.Execute(bot, p); err != nil {
		return bot.SendMessage(p.Message.Chat(), "Internal error", types.NewKeyboard().
			NewRow(types.NewCallbackButton("🔄 Menu", MENU_CMD)))
	}

	return nil
}
