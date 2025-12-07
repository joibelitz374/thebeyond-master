package promo

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/jackc/pgx/v5"
	"github.com/quickpowered/frilly/config/language"
	promoCfg "github.com/quickpowered/frilly/config/promo"
	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/i18n"
	"github.com/quickpowered/frilly/internal/repositories/subscriptions"
	"github.com/quickpowered/frilly/internal/repositories/xray"
	"github.com/quickpowered/frilly/internal/services/account"
	"github.com/quickpowered/frilly/internal/services/payment"
	"github.com/quickpowered/frilly/internal/services/promo"
	"github.com/quickpowered/frilly/pkg/consts"
	"go.uber.org/zap"
)

type UseCase interface {
	Get(promoName string) (promo domain.Promo, discount int, err error)
	RegisterReferral(botAPI *bot.Bot, senderID, accountID int, promo domain.Promo, target consts.PromoTarget) error
}

type useCase struct {
	accountService    account.Interface
	paymentService    payment.Interface
	promoService      promo.Interface
	subscriptionsRepo subscriptions.Repository
	logger            *zap.Logger
}

func NewUseCase(
	accountService account.Interface,
	paymentService payment.Interface,
	promoService promo.Interface,
	subscriptionsRepo subscriptions.Repository,
	xrayClient xray.Interface,
	logger *zap.Logger,
) useCase {
	return useCase{accountService, paymentService, promoService, subscriptionsRepo, logger}
}

func (uc useCase) Get(promoName string) (promo domain.Promo, discount int, err error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()

	uc.logger.Debug("getting promo",
		zap.String("promo_name", promoName))

	promo, err = uc.promoService.Get(ctx, promoName)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			uc.logger.Debug("promo not found",
				zap.String("promo_name", promoName))
			return promo, discount, nil
		}
		uc.logger.Error("failed to get promo", zap.Error(err))
		return promo, discount, err
	}

	uc.logger.Debug("promo found",
		zap.String("promo_name", promo.Name),
		zap.Int("promo_level", promo.Level),
		zap.Int("promo_creator", promo.Creator),
		zap.Int("promo_clients", promo.Clients),
		zap.Int("promo_buyers", promo.Buyers),
	)

	return promo, promoCfg.LevelDiscounts[promo.Level-1], nil
}

func (uc useCase) RegisterReferral(botAPI *bot.Bot, senderID, accountID int, promo domain.Promo, target consts.PromoTarget) (err error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()

	uc.logger.Debug("registering referral",
		zap.Int("account_id", accountID),
		zap.String("promo_name", promo.Name),
		zap.Int("promo_level", promo.Level),
		zap.Int("promo_creator", promo.Creator),
		zap.Int("promo_clients", promo.Clients),
		zap.Int("promo_buyers", promo.Buyers),
		zap.String("promo_target", string(target)))

	switch target {
	case consts.PromoTargetClient:
		if err := uc.promoService.IncreaseClients(ctx, promo.Name); err != nil {
			uc.logger.Error("failed to increase client",
				zap.Error(err),
				zap.String("promo_name", promo.Name))
			return err
		}
		promo.Clients++
	case consts.PromoTargetBuyer:
		err = uc.promoService.CheckBuyer(ctx, promo.Name, accountID)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			uc.logger.Error("failed to check buyer",
				zap.Error(err),
				zap.String("promo_name", promo.Name))
			return err
		}

		if err := uc.promoService.IncreaseBuyers(ctx, promo.Name); err != nil {
			uc.logger.Error("failed to increase buyers",
				zap.Error(err),
				zap.String("promo_name", promo.Name))
			return err
		}

		if err := uc.promoService.AddBuyer(ctx, promo.Name, accountID); err != nil {
			uc.logger.Error("failed to add buyer",
				zap.Error(err),
				zap.String("promo_name", promo.Name))
			return err
		}

		promo.Buyers++
	}

	uc.logger.Debug("promo updated",
		zap.String("name", promo.Name),
		zap.Int("clients", promo.Clients),
		zap.Int("buyers", promo.Buyers))

	creatorAccount, err := uc.accountService.GetByAccountID(ctx, promo.Creator)
	if err != nil {
		uc.logger.Error("failed to get creator account",
			zap.Error(err),
			zap.Int("promo_creator", promo.Creator))
		return err
	}

	uc.logger.Debug("creator account found",
		zap.Int("promo_creator", promo.Creator),
		zap.Int("creator_id", creatorAccount.ID),
		zap.String("creator_language", creatorAccount.Language))

	if creatorAccount.ID == 0 {
		uc.logger.Error("creator account not found",
			zap.Int("promo_creator", promo.Creator))
		return errors.New("creator account not found")
	}

	var counter int
	switch target {
	case consts.PromoTargetClient:
		counter = promo.Clients
	case consts.PromoTargetBuyer:
		counter = promo.Buyers
	}

	promoGoals := promoCfg.RewardLevels[promo.Level]
	promoMsg := i18n.PromoMessages[language.Language(creatorAccount.Language)]

	uc.logger.Debug("promo goals found",
		zap.String("name", promo.Name),
		zap.Int("level", promo.Level),
		zap.Int("clients", promo.Clients),
		zap.Int("buyers", promo.Buyers),
		zap.Int("referral_bonus_days", promoGoals.ReferralBonusDays))

	text := fmt.Sprintf("%s %s %s <a href=\"tg://user?id=%d\">%s</a>\n%s: %d",
		promoMsg.Promo, strings.ToUpper(promo.Name), promoMsg.Attracted,
		senderID, promoMsg.NewClient,
		promoMsg.AttractedClients, counter)

	var reward promoCfg.Reward
	var ok bool

	switch target {
	case consts.PromoTargetClient:
		reward, ok = promoGoals.ClientRewards[counter]
	case consts.PromoTargetBuyer:
		reward, ok = promoGoals.BuyerRewards[counter]
	}

	if ok {
		ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
		defer cancel()

		duration := time.Duration(reward.Days) * 24 * time.Hour
		if err := uc.accountService.AddSubscriptionExpiresAt(ctx, promo.Creator, duration); err != nil {
			uc.logger.Error("failed to add subscription expires at", zap.Error(err))
			return err
		}

		text += fmt.Sprintf("\n%s: %d", promoMsg.ReceivedDays, reward.Days)

		if reward.Discount > 0 {
			ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
			defer cancel()

			if err := uc.accountService.SetDiscount(ctx, promo.Creator, reward.Discount); err != nil {
				uc.logger.Error("failed to set discount", zap.Error(err))
				return err
			}

			text += fmt.Sprintf("\n%s: %d%%", promoMsg.Discount, reward.Discount)
		}
	}

	externalAccountID, err := uc.accountService.GetPlatformUserID(ctx, promo.Creator)
	if err != nil {
		uc.logger.Error("failed to get platform user id", zap.Error(err))
		return err
	}

	if _, err := botAPI.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    externalAccountID,
		Text:      text,
		ParseMode: models.ParseModeHTML,
	}); err != nil {
		uc.logger.Error("failed to send message", zap.Error(err))
		return err
	}

	uc.logger.Debug("promo level check",
		zap.Int("clients", promo.Clients),
		zap.Int("buyers", promo.Buyers),
		zap.Int("clients_target", promoGoals.ClientTargets[len(promoGoals.ClientTargets)-1]),
		zap.Int("buyers_target", promoGoals.BuyerTargets[len(promoGoals.BuyerTargets)-1]))

	if promo.Level != 2 &&
		promo.Clients >= promoGoals.ClientTargets[len(promoGoals.ClientTargets)-1] &&
		promo.Buyers >= promoGoals.BuyerTargets[len(promoGoals.BuyerTargets)-1] {
		if err := uc.upgradePromoLevel(
			botAPI, promo.Name,
			externalAccountID,
			promoGoals.ReferralBonusDays,
			promo, promoMsg); err != nil {
			uc.logger.Error("failed to upgrade promo level", zap.Error(err))
			return err
		}
	}

	return nil
}

func (uc useCase) upgradePromoLevel(
	botAPI *bot.Bot,
	promoName string,
	externalAccountID,
	referralBonusDays int,
	promo domain.Promo,
	promoMsg i18n.PromoLocale,
) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	uc.logger.Debug("promo level upgraded",
		zap.String("name", promoName),
		zap.Int("currency_level", promo.Level),
		zap.Int("new_level", promo.Level+1))

	if err := uc.promoService.UpgradeLevel(ctx, promoName); err != nil {
		uc.logger.Error("failed to upgrade promo level", zap.Error(err))
		return err
	}

	duration := time.Duration(referralBonusDays) * 24 * time.Hour
	uc.logger.Debug("adding subscription expires at",
		zap.String("name", promoName),
		zap.Duration("duration", duration))

	if err := uc.accountService.AddSubscriptionExpiresAt(ctx, promo.Creator, duration); err != nil {
		uc.logger.Error("failed to add subscription expires at", zap.Error(err))
		return err
	}

	_, err := botAPI.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: externalAccountID,
		Text: fmt.Sprintf("%s %s %s %d %s\n%s: %d",
			promoMsg.Promo, strings.ToUpper(promoName), promoMsg.Reached, promo.Level+1, promoMsg.Level,
			promoMsg.ReceivedDays, referralBonusDays),
		ParseMode: models.ParseModeHTML,
	})
	return err
}
