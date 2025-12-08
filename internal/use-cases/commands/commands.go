package commands

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	tgBot "github.com/go-telegram/bot"
	"github.com/jackc/pgx/v5"
	"github.com/quickpowered/frilly/config/language"
	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/i18n"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/repositories/serverlocations"
	"github.com/quickpowered/frilly/internal/types/update"
	"github.com/quickpowered/frilly/internal/use-cases/commands/deps"
	"github.com/quickpowered/frilly/internal/use-cases/promo"
	"github.com/quickpowered/frilly/pkg/consts"
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

func NewUseCase(deps_ deps.Dependencies, promoUseCase promo.UseCase, serverLocationsRepo serverlocations.Repository) useCase {
	return useCase{deps_, promoUseCase, map[string]Command{
		ACCOUNT_CMD:    NewAccountHandler(deps_, serverLocationsRepo),
		HOWTOUSE_CMD:   NewHowToUseHandler(deps_),
		NEWKEY_CMD:     NewNewKeyHandler(deps_),
		RENEW_CMD:      NewRenewHandler(deps_),
		REGION_CMD:     NewRegionHandler(deps_),
		LANGUAGE_CMD:   NewLanguageHandler(deps_),
		CURRENCY_CMD:   NewCurrencyHandler(deps_),
		PROTOCOL_CMD:   NewProtocolHandler(deps_),
		LOCATION_CMD:   NewLocationCmd(deps_),
		REFUND_CMD:     NewRefundHandler(deps_),
		PROMO_CMD:      NewPromoCmd(deps_),
		TOP_PROMO_CMD:  NewTopHandler(deps_),
		REPOST_CMD:     NewRepostHandler(deps_),
		SET_PRICES_CMD: NewSetPricesHandler(deps_),
		WELCOME_CMD:    NewWelcomeHandler(deps_),
	}}
}

func (uc useCase) Run(bot bin.Interface, payload *domain.Payload) (err error) {
	platform := bot.GetPlatform()
	sender := payload.Message.Sender()
	text := payload.Message.Text()

	if len(text) == 0 {
		return nil
	}

	if text[0] == '/' {
		text = text[1:]
	}

	payload.Args = strings.Split(text, " ")
	cmd := strings.ToLower(payload.Args[0])

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	payload.Account, err = uc.AccountService.Get(ctx, platform, sender)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			if payload.Account.Language == "" {
				payload.Account.Language = "en"
			}

			promo, discount, err := uc.createAccount(bot, platform, sender, payload.Args, cmd)
			if err != nil {
				return err
			}

			targetPromoMsg := i18n.PromoMessages[language.Language(payload.Account.Language)]
			if promo.Name != "" {
				bot.SendMessage(payload.Message.Chat(), fmt.Sprintf("🎫 %s %s %s!\n%s: %d%%",
					targetPromoMsg.Promo, strings.ToUpper(promo.Name), targetPromoMsg.Activated,
					targetPromoMsg.Discount, discount))
			}

			payload.Args = []string{LANGUAGE_CMD}
			return uc.commands[LANGUAGE_CMD].Execute(bot, payload)
		} else {
			uc.Logger.Error("failed to get account", zap.Error(err))
			return err
		}
	}

	if payload.Account.ID == 0 {
		return fmt.Errorf("account not found")
	}

	if payload.Account.Language == "" {
		return uc.commands[LANGUAGE_CMD].Execute(bot, payload)
	}

	if payload.Account.Currency == "" {
		return uc.commands[CURRENCY_CMD].Execute(bot, payload)
	}

	if cmd, ok := uc.commands[cmd]; ok {
		uc.Logger.Info("new command ",
			zap.Int("platform", int(platform)),
			zap.Int("chat_id", payload.Message.Chat().GetID()),
			zap.Int("thread_id", payload.Message.Chat().GetThreadID()),
			zap.Int("message_id", payload.Message.ID()),
			zap.Int("sender", sender),
			zap.String("text", text))

		if err := cmd.Execute(bot, payload); err != nil {
			if err := bot.SendMessage(payload.Message.Chat(), "Internal error"); err != nil {
				uc.Logger.Error("failed to send message", zap.Error(err))
			}

			return err
		}
	}

	return nil
}

func (uc useCase) createAccount(
	bot bin.Interface,
	platform consts.Platform,
	senderID int,
	args []string,
	cmd string,
) (promo domain.Promo, discount int, err error) {
	if cmd == START_CMD && len(args) >= 2 {
		promo, discount, err = uc.promoUseCase.Get(strings.ToLower(args[1]))
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	var nullablePromoName *string
	if promo.Name != "" {
		nullablePromoName = &promo.Name
	}

	trialTime := time.Now().Add(time.Hour * 72)
	newAccountID, keyID, err := uc.AccountService.Create(ctx, platform, senderID, trialTime, nullablePromoName, discount)
	if err != nil {
		uc.Logger.Error("failed to create account", zap.Error(err))
		return promo, discount, err
	}

	if err := uc.XRayClient.AddUser(fmt.Sprintf("id%d@user", newAccountID), keyID); err != nil {
		uc.Logger.Error("failed to add user", zap.Error(err))
		return promo, discount, err
	}

	if err := bot.SendMessage(update.Chat{ID: 924536264}, fmt.Sprintf("Registered <a href=\"tg://user?id=%d\">new user</a>", senderID)); err != nil {
		return promo, discount, err
	}

	if nullablePromoName != nil {
		if err := uc.promoUseCase.RegisterReferral(
			bot.GetAPI().(*tgBot.Bot),
			senderID, newAccountID, promo,
			consts.PromoTargetClient,
		); err != nil {
			return promo, discount, err
		}
	}

	return promo, discount, nil
}
