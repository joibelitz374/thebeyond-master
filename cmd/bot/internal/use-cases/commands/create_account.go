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
	var promo sharedDomain.Promo
	var discount int

	uc.setDefaultLanguage(p)
	p.Account.Region = "ru"

	if p.Args[0] == START_CMD && len(p.Args) >= 2 {
		promo, discount, err = uc.promoUseCase.Get(strings.ToLower(p.Args[1]))
		if err != nil {
			uc.Logger.Error("failed to get promo", zap.Error(err))
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
		uc.Logger.Error("failed to create account", zap.Error(err))
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
			uc.Logger.Error("failed to register referral", zap.Error(err))
		}

		promoMsg := i18n.PromoMessages[language.Language(p.Account.Language)]
		if promo.Name != "" {
			bot.SendMessage(p.Message.Chat(), fmt.Sprintf("🎫 %s %s %s!\n%s: %d%%",
				promoMsg.Promo, strings.ToUpper(promo.Name), promoMsg.Activated,
				promoMsg.Discount, discount))
		}
	}

	profileURI := fmt.Sprintf("tg://user?id=%d", p.Message.Sender())
	if err := bot.SendMessage(
		update.Chat{ID: 924536264}, "Created new account",
		types.NewKeyboard().NewRow(types.NewURLButton("Profile", profileURI))); err != nil {
		return err
	}

	p.Args[0] = LANGUAGE_CMD
	return uc.commands[LANGUAGE_CMD].Execute(bot, p)
}

func (uc useCase) setDefaultLanguage(p *domain.Payload) {
	if p.Account.Language == "" {
		p.Account.Language = "en"
	}
}
