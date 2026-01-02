package commands

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types/update"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/renew"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/promo"
	"go.uber.org/zap"
)

const START_CMD = "start"

type UseCase interface {
	Run(bot bin.Interface, payload *domain.Payload) error
}

type Command interface {
	Execute(bot bin.Interface, payload *domain.Payload) error
}

type useCase struct {
	deps.Dependencies
	promoUseCase promo.UseCase
	commands     map[string]Command
}

func NewUseCase(deps_ deps.Dependencies, promoUseCase promo.UseCase) useCase {
	commands := map[string]Command{
		MENU_CMD:     NewMenuHandler(deps_),
		ABOUT_CMD:    NewAboutHandler(deps_),
		SETTINGS_CMD: NewSettingsHandler(deps_),
		TOS_CMD:      NewTosHandler(deps_),
		PRIVACY_CMD:  NewPrivacyHandler(deps_),
		REFUND_CMD:   NewRefundHandler(deps_),
		CONNECT_CMD:  NewConnectHandler(deps_),
		PROMO_CMD:    NewPromoCmd(deps_),
		renew.CMD:    renew.NewHandler(deps_),
		NEWKEY_CMD:   NewNewKeyHandler(deps_),
		REGION_CMD:   NewRegionHandler(deps_),
		// NETWORK_CMD:  NewNetworkHandler(deps_),
		LANGUAGE_CMD: NewLanguageHandler(deps_),
		CURRENCY_CMD: NewCurrencyHandler(deps_),
		// PULL_CMD:       NewPullHandler(deps_),
		PROTOCOL_CMD: NewProtocolHandler(deps_),
		SUPPORT_CMD:  NewSupportHandler(deps_),
		REPOST_CMD:   NewRepostHandler(deps_),
	}

	return useCase{deps_, promoUseCase, commands}
}

func (uc useCase) Run(bot bin.Interface, p *domain.Payload) (err error) {
	text := p.Message.Text()
	if len(text) == 0 {
		return nil
	}

	if len(text) == 0 {
		return nil
	}

	if text[0] == '/' {
		text = text[1:]
	}

	p.Args = strings.Split(text, " ")
	p.Args[0] = strings.ToLower(p.Args[0])

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
		p.Args[0] = LANGUAGE_CMD
	case p.Account.Currency == "":
		p.Args[0] = CURRENCY_CMD
	}

	cmd, ok := uc.commands[p.Args[0]]
	if !ok {
		buttonRows := [][]types.Button{{{Text: "🔄 Menu", Data: MENU_CMD}}}
		return bot.SendMessage(p.Message.Chat(), "Command not found", &types.Keyboard{ButtonRows: buttonRows})
	}

	defer func() {
		uc.deleteMessage(bot, p.Message.Chat(), p.Message.ID())
	}()

	if err := cmd.Execute(bot, p); err != nil {
		buttonRows := [][]types.Button{{{Text: "🔄 Menu", Data: MENU_CMD}}}
		uc.Logger.Error("Internal command error", zap.Error(err))
		return bot.SendMessage(p.Message.Chat(), "Internal error", &types.Keyboard{ButtonRows: buttonRows})
	}

	uc.Logger.Info("command executed",
		zap.Int("sender", p.Message.Sender()),
		zap.Int("message_id", p.Message.ID()),
		zap.String("command", p.Args[0]),
		zap.String("args", strings.Join(p.Args[1:], " ")))

	return nil
}

func (uc useCase) deleteMessage(bot bin.Interface, chat update.ChatInterface, messageID int) error {
	if err := bot.DeleteMessage(chat, messageID); err != nil {
		uc.Logger.Error("failed to delete message", zap.Error(err))
	}
	return nil
}
