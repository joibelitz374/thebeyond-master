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
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
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
		MENU_CMD:     NewMenuHandler(deps_),
		ABOUT_CMD:    NewAboutHandler(deps_),
		SETTINGS_CMD: NewSettingsHandler(deps_),
		TOS_CMD:      NewTosHandler(deps_),
		PRIVACY_CMD:  NewPrivacyHandler(deps_),
		REFUND_CMD:   NewRefundHandler(deps_),
		CONNECT_CMD:  NewConnectHandler(deps_),
		PROMO_CMD:    NewPromoCmd(deps_),
		RENEW_CMD:    NewRenewHandler(deps_),
		NEWKEY_CMD:   NewNewKeyHandler(deps_),
		REGION_CMD:   NewRegionHandler(deps_),
		// NETWORK_CMD:  NewNetworkHandler(deps_),
		LANGUAGE_CMD: NewLanguageHandler(deps_),
		CURRENCY_CMD: NewCurrencyHandler(deps_),
		// WHITELIST_RENEW_CMD: NewWhitelistRenewHandler(deps_),
		// PULL_CMD:       NewPullHandler(deps_),
		// PROTOCOL_CMD:   NewProtocolHandler(deps_),
		SUPPORT_CMD:    NewSupportHandler(deps_),
		REPOST_CMD:     NewRepostHandler(deps_),
		SET_PRICES_CMD: NewSetPricesHandler(deps_),
	}}
}

var commandsWithDeferredDeletion = map[string]struct{}{
	LANGUAGE_CMD: {},
	CURRENCY_CMD: {},
	PROMO_CMD:    {},
}

func (uc useCase) Run(bot bin.Interface, p *domain.Payload) (err error) {
	platform := bot.GetPlatform()
	sender := p.Message.Sender()
	text := p.Message.Text()

	if len(text) == 0 {
		return nil
	}

	if text[0] == '/' {
		text = text[1:]
	}

	if len(text) == 0 {
		return nil
	}

	p.Args = strings.Split(text, " ")
	p.Args[0] = strings.ToLower(p.Args[0])

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	p.Account, err = uc.AccountService.Get(ctx, platform, sender)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			if p.Account.Language == "" {
				p.Account.Language = "en"
			}

			uc.Logger.Info("new account",
				zap.Int("platform", int(platform)),
				zap.Int("sender", p.Message.Sender()))

			p.Account.Region = "ru"

			promo, discount, err := uc.createAccount(bot, platform, sender, p.Account.Region, p.Args)
			if err != nil {
				return err
			}

			targetPromoMsg := i18n.PromoMessages[language.Language(p.Account.Language)]
			if promo.Name != "" {
				bot.SendMessage(p.Message.Chat(), fmt.Sprintf("🎫 %s %s %s!\n%s: %d%%",
					targetPromoMsg.Promo, strings.ToUpper(promo.Name), targetPromoMsg.Activated,
					targetPromoMsg.Discount, discount))
			}

			p.Args[0] = LANGUAGE_CMD
			return uc.commands[LANGUAGE_CMD].Execute(bot, p)
		} else {
			uc.Logger.Error("failed to get account", zap.Error(err))
			return err
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

	if cmd, ok := uc.commands[p.Args[0]]; ok {
		uc.Logger.Info("new command",
			zap.Int("platform", int(platform)),
			zap.Int("chat_id", p.Message.Chat().GetID()),
			zap.Int("thread_id", p.Message.Chat().GetThreadID()),
			zap.Int("message_id", p.Message.ID()),
			zap.Int("sender", sender),
			zap.String("text", text))

		if _, ok := commandsWithDeferredDeletion[p.Args[0]]; !ok || (ok && len(p.Args) < 2) {
			if err := bot.DeleteMessage(p.Message.Chat(), p.Message.ID()); err != nil {
				uc.Logger.Error("failed to delete message", zap.Error(err))
			}
		} else {
			defer func() {
				if err := bot.DeleteMessage(p.Message.Chat(), p.Message.ID()); err != nil {
					uc.Logger.Error("failed to delete message", zap.Error(err))
				}
			}()
		}

		if err := cmd.Execute(bot, p); err != nil {
			if err := bot.SendMessage(p.Message.Chat(), "Internal error", &types.Keyboard{
				ButtonRows: [][]types.Button{{
					{Text: "🔄 Menu", Data: MENU_CMD},
				}},
			}); err != nil {
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
	region string,
	args []string,
) (promo sharedDomain.Promo, discount int, err error) {
	if args[0] == START_CMD && len(args) >= 2 {
		name := strings.ToLower(args[1])
		promo, discount, err = uc.promoUseCase.Get(name)
		if err != nil {
			uc.Logger.Debug("promo activated for user",
				zap.Int("platform", int(platform)),
				zap.Int("sender", sender),
				zap.String("promo_name", name))
		}
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
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

	regions := []dto.Region{}
	switch region {
	case "ru":
		regions = append(regions, dto.RegionRussia)
	}

	for _, region := range regions {
		if err := uc.XRayManagerRepo.AddClient(ctx, region, keyID, email.NewAccount(newAccountID)); err != nil {
			uc.Logger.Error("failed to add user", zap.Error(err))
			return promo, discount, err
		}
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
