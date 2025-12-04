package commands

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/repositories/web"
	"github.com/quickpowered/frilly/internal/use-cases/commands/tools"
	"go.uber.org/zap"
)

type Commands interface {
	Run(bot bin.Interface, payload *domain.Payload) error
}

type UseCase struct {
	tools.Modules
	commands map[string]tools.Command
}

func NewUseCase(modules tools.Modules, exchangeRates web.ExchangeRatesInterface) *UseCase {
	return &UseCase{modules, map[string]tools.Command{
		ACCOUNT_CMD:  NewAccountCmd(modules),
		NEWKEY_CMD:   NewNewKeyCmd(modules),
		RENEW_CMD:    NewRenewCmd(modules, exchangeRates),
		REGION_CMD:   NewRegionCmd(modules),
		LANGUAGE_CMD: NewLanguageCmd(modules),
		CURRENCY_CMD: NewCurrencyCmd(modules),
		PROTOCOL_CMD: NewProtocolCmd(modules),
		LOCATION_CMD: NewLocationCmd(modules),
		REFUND_CMD:   NewRefundCmd(modules),
	}}
}

func (c *UseCase) Run(bot bin.Interface, payload *domain.Payload) (err error) {
	platform := bot.GetPlatform()
	sender := payload.Message.GetSender()
	text := payload.Message.GetText()

	c.Logger.Info("new message",
		zap.Int("platform", int(platform)),
		zap.Int("chat_id", payload.Message.GetChat().GetID()),
		zap.Int("thread_id", payload.Message.GetChat().GetThreadID()),
		zap.Int("message_id", payload.Message.GetID()),
		zap.Int("sender", sender),
		zap.String("text", text))

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	payload.Account, err = c.AccountService.Get(ctx, platform, sender)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
			defer cancel()

			newAccountID, keyID, err := c.AccountService.Create(ctx, platform, sender, time.Now().Add(time.Hour*24))
			if err != nil {
				c.Logger.Error("failed to create account", zap.Error(err))
				return err
			}

			payload.Account.ID = newAccountID
			payload.Account.KeyID = keyID

			if err := c.XRayClient.AddUser(fmt.Sprintf("id%d@user", payload.Account.ID), payload.Account.KeyID); err != nil {
				c.Logger.Error("failed to add user", zap.Error(err))
				return err
			}

			return c.commands[LANGUAGE_CMD].Execute(bot, payload)
		} else {
			c.Logger.Error("failed to get account", zap.Error(err))
			return err
		}
	}

	if payload.Account.ID == 0 {
		return fmt.Errorf("account not found")
	}

	if len(text) == 0 {
		return nil
	}

	if text[0] == '/' {
		text = text[1:]
	}

	payload.NodeRoute = strings.Split(text, " ")

	if payload.Account.Language == "" {
		return c.commands[LANGUAGE_CMD].Execute(bot, payload)
	}

	if payload.Account.Currency == "" {
		return c.commands[CURRENCY_CMD].Execute(bot, payload)
	}

	// if payload.Account.Region == "" {
	// 	return c.commands[REGION_CMD].Execute(bot, payload)
	// }

	if cmd, ok := c.commands[strings.ToLower(payload.NodeRoute[0])]; ok {
		if err := cmd.Execute(bot, payload); err != nil {
			if err := bot.SendMessage(payload.Message.GetChat(), "Internal error"); err != nil {
				c.Logger.Error("failed to send message", zap.Error(err))
			}

			return err
		}
	}

	return nil
}
