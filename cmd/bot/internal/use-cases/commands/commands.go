package commands

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	telegram "github.com/go-telegram/bot"
	"github.com/jackc/pgx/v5"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types/update"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/promo"
	"github.com/quickpowered/thebeyond-master/configs/language"
	sharedDomain "github.com/quickpowered/thebeyond-master/internal/domain"
	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager/dto"
	"github.com/quickpowered/thebeyond-master/pkg/consts"
	"github.com/quickpowered/thebeyond-master/pkg/email"
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
	return useCase{deps_, promoUseCase, map[string]Command{
		ACCOUNT_CMD: NewAccountHandler(deps_),
		CONNECT_CMD: NewConnectHandler(deps_),
		NEWKEY_CMD:  NewNewKeyHandler(deps_),
		RENEW_CMD:   NewRenewHandler(deps_),
		// WHITELIST_RENEW_CMD: NewWhitelistRenewHandler(deps_),
		// PULL_CMD:       NewPullHandler(deps_),
		// DEVICES_CMD:    NewDevicesHandler(deps_),
		REGION_CMD:     NewRegionHandler(deps_),
		LANGUAGE_CMD:   NewLanguageHandler(deps_),
		CURRENCY_CMD:   NewCurrencyHandler(deps_),
		PROTOCOL_CMD:   NewProtocolHandler(deps_),
		REFUND_CMD:     NewRefundHandler(deps_),
		PROMO_CMD:      NewPromoCmd(deps_),
		REPOST_CMD:     NewRepostHandler(deps_),
		SET_PRICES_CMD: NewSetPricesHandler(deps_),
		WELCOME_CMD:    NewWelcomeHandler(deps_),
		SUPPORT_CMD:    NewSupportHandler(deps_),
	}}
}

func (uc useCase) Run(bot bin.Interface, payload *domain.Payload) (err error) {
	platform := bot.GetPlatform()
	sender := payload.Message.Sender()
	text := payload.Message.Text()

	// bot.SendMessage(payload.Message.Chat(), "У Вас работает сервис?", &types.Keyboard{
	// 	ButtonRows: [][]types.Button{
	// 		{
	// 			{
	// 				Text: "🎾 Да, всё отлично!",
	// 				Data: "tennis",
	// 			},
	// 		},
	// 		{
	// 			{
	// 				Text: "🏐 Нет, нужна помощь!",
	// 				Data: "support",
	// 			},
	// 		},
	// 		{
	// 			{
	// 				Text: "🆘 Есть проблемы",
	// 				Data: "support",
	// 			},
	// 		},
	// 	}})

	// member, err := bot.GetAPI().(*telegram.Bot).GetChatMember(context.TODO(), &telegram.GetChatMemberParams{
	// 	ChatID: -1003309480333,
	// 	UserID: int64(sender),
	// })
	// if err != nil {
	// 	return err
	// }

	// if member.Left != nil && member.Left.Status == models.ChatMemberTypeLeft {
	// 	attachments := types.NewAttachments()
	// 	attachments.AddURL("https://vkplay.ru/pre_0x736_resize/hotbox/content_files/Stories/2024/03/27/cc906d2ed79d448db77d50cf47424704.jpg")
	// 	return bot.SendMessage(payload.Message.Chat(),
	// 		"Подпишитесь на канал, чтобы пользоваться ботом!",
	// 		attachments, &types.Keyboard{ButtonRows: [][]types.Button{
	// 			{{
	// 				Text: "🗒 Подписаться",
	// 				URL:  "https://t.me/beyondsecurebot",
	// 			}},
	// 			{{
	// 				Text: "🎾 Подписался",
	// 				Data: "check_subscription",
	// 			}},
	// 		}})
	// }

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

			uc.Logger.Info("new account",
				zap.Int("platform", int(platform)),
				zap.Int("sender", payload.Message.Sender()))

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
		uc.Logger.Info("new command",
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
	sender int,
	args []string,
	cmd string,
) (promo sharedDomain.Promo, discount int, err error) {
	if cmd == START_CMD && len(args) >= 2 {
		name := strings.ToLower(args[1])
		promo, discount, err = uc.promoUseCase.Get(name)
		if err != nil {
			uc.Logger.Debug("promo activated for user",
				zap.Int("platform", int(platform)),
				zap.Int("sender", sender),
				zap.String("promo_name", name))
		}
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	var nullablePromoName *string
	if promo.Name != "" {
		nullablePromoName = &promo.Name
	}

	trialTime := time.Now().Add(time.Hour * 72)
	newAccountID, keyID, err := uc.AccountService.Create(ctx, platform, sender, trialTime, nullablePromoName, discount)
	if err != nil {
		uc.Logger.Error("failed to create account", zap.Error(err))
		return promo, discount, err
	}

	if err := uc.XRayManagerRepo.AddClient(ctx, 1, dto.NodeTypeBlacklist, keyID, email.NewAccount(newAccountID)); err != nil {
		uc.Logger.Error("failed to add user", zap.Error(err))
		return promo, discount, err
	}

	if err := bot.SendMessage(update.Chat{ID: 924536264}, fmt.Sprintf("Registered <a href=\"tg://user?id=%d\">new user</a>", sender)); err != nil {
		return promo, discount, err
	}

	if nullablePromoName != nil {
		if err := uc.promoUseCase.RegisterReferral(
			bot.GetAPI().(*telegram.Bot),
			sender, newAccountID, promo,
		); err != nil {
			return promo, discount, err
		}
	}

	return promo, discount, nil
}
