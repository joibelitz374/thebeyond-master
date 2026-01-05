package commands

import (
	"context"
	"fmt"
	"strings"
	"time"

	telegram "github.com/go-telegram/bot"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types/update"
	"github.com/quickpowered/thebeyond-master/configs/language"
	sharedDomain "github.com/quickpowered/thebeyond-master/internal/domain"
	"github.com/quickpowered/thebeyond-master/pkg/consts"
	"go.uber.org/zap"
)

func (uc useCase) createAccount(bot bin.Interface, platform consts.Platform, p *domain.Payload) (err error) {
	if p.Account.Language == "" {
		p.Account.Language = language.English
	}

	p.Account.Region = "ru"

	var promo sharedDomain.Promo
	var discount int

	if len(p.Args) >= 2 && strings.EqualFold(p.Args[0], START_CMD) {
		var err error
		promo, discount, err = uc.promoUseCase.Get(strings.ToLower(p.Args[1]))
		if err != nil {
			uc.Logger.Error("failed to get promo",
				zap.Int("sender", p.Message.Sender()),
				zap.String("promo_code", p.Args[1]),
				zap.Error(err))
		}
	}

	var promoName *string
	if promo.Name != "" {
		promoName = &promo.Name
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()

	newAccountID, keyID, err := uc.AccountService.Create(ctx, platform, p.Message.Sender(), promoName, discount)
	if err != nil {
		uc.Logger.Error("failed to create account",
			zap.Int("account_id", newAccountID),
			zap.Int("sender", p.Message.Sender()),
			zap.String("key_id", keyID),
			zap.String("promo_name", promo.Name),
			zap.Int("discount", discount),
			zap.Error(err))
		return err
	}

	uc.Logger.Info("account created successfully",
		zap.Int("account_id", newAccountID),
		zap.Int("sender", p.Message.Sender()),
		zap.String("key_id", keyID),
		zap.String("promo_name", promo.Name),
		zap.Int("discount", discount))

	if promoName != nil {
		tgBot := bot.GetAPI().(*telegram.Bot)
		if err := uc.promoUseCase.RegisterReferral(tgBot, p.Message.Sender(), newAccountID, promo); err != nil {
			uc.Logger.Error("failed to register referral",
				zap.Int("account_id", newAccountID),
				zap.String("promo_name", promo.Name),
				zap.Error(err))
		}

		msg := i18n.PromoMessages[language.Language(p.Account.Language)]
		if promo.Name != "" {
			if err := bot.SendMessage(p.Message.Chat(), fmt.Sprintf("🎫 %s %s %s!\n%s: %d%%",
				msg.Promo, strings.ToUpper(promo.Name), msg.Activated,
				msg.Discount, discount)); err != nil {
				uc.Logger.Error("failed to send promo message to user",
					zap.Int("account_id", newAccountID),
					zap.Error(err))
			}
		}
	}

	if err := bot.SendMessage(update.Chat{ID: 924536264}, "Created new account", types.NewKeyboard().
		NewRow(types.NewURLButton("Profile", fmt.Sprintf("tg://user?id=%d", p.Message.Sender())))); err != nil {
		uc.Logger.Error("failed to notify admin about new account",
			zap.Int("sender", p.Message.Sender()),
			zap.Error(err))
	}

	p.Args[0] = LANGUAGE_CMD
	return uc.commands[LANGUAGE_CMD].Execute(bot, p)
}
