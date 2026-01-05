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
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/configs/language"
	promoCfg "github.com/quickpowered/thebeyond-master/configs/promo"
	"github.com/quickpowered/thebeyond-master/internal/domain"
	"github.com/quickpowered/thebeyond-master/internal/services/account"
	"github.com/quickpowered/thebeyond-master/internal/services/promo"
	"github.com/quickpowered/thebeyond-master/internal/services/subscription"
	"go.uber.org/zap"
)

type UseCase interface {
	Get(promoName string) (promo domain.Promo, discount int, err error)
	RegisterReferral(botAPI *bot.Bot, senderID, accountID int, promo domain.Promo) error
}

type useCase struct {
	accountService      account.Interface
	subscriptionService subscription.Interface
	promoService        promo.Interface
	logger              *zap.Logger
}

func NewUseCase(
	accountService account.Interface,
	subscriptionService subscription.Interface,
	promoService promo.Interface,
	logger *zap.Logger,
) UseCase {
	return useCase{accountService, subscriptionService, promoService, logger}
}

func (uc useCase) Get(promoName string) (domain.Promo, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	uc.logger.Debug("getting promo",
		zap.String("promo_name", promoName))

	promo, err := uc.promoService.Get(ctx, promoName)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			uc.logger.Debug("promo not found", zap.String("promo_name", promoName))
			return domain.Promo{}, 0, nil
		}

		uc.logger.Error("failed to get promo", zap.String("name", promoName), zap.Error(err))
		return domain.Promo{}, 0, err
	}

	uc.logger.Debug("promo found",
		zap.String("promo_name", promo.Name),
		zap.Int("promo_level", promo.Level),
		zap.Int("promo_creator", promo.Creator),
		zap.Int("promo_clients", promo.Clients))

	return promo, promoCfg.LevelDiscounts[promo.Level-1], nil
}

func (uc useCase) RegisterReferral(botAPI *bot.Bot, senderID, accountID int, promo domain.Promo) (err error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()

	uc.logger.Debug("registering referral",
		zap.Int("account_id", accountID),
		zap.String("promo_name", promo.Name),
		zap.Int("promo_level", promo.Level),
		zap.Int("promo_creator", promo.Creator),
		zap.Int("promo_clients", promo.Clients))

	if err := uc.promoService.IncreaseClients(ctx, promo.Name); err != nil {
		uc.logger.Error("failed to increase client", zap.String("promo_name", promo.Name), zap.Error(err))
		return err
	}
	promo.Clients++

	uc.logger.Debug("promo updated",
		zap.String("name", promo.Name),
		zap.Int("clients", promo.Clients))

	creatorAccount, err := uc.accountService.GetByAccountID(ctx, promo.Creator)
	if err != nil {
		uc.logger.Error("failed to get creator account", zap.Int("promo_creator", promo.Creator), zap.Error(err))
		return err
	}

	uc.logger.Debug("creator account found",
		zap.Int("promo_creator", promo.Creator),
		zap.Int("creator_id", creatorAccount.ID),
		zap.String("creator_language", string(creatorAccount.Language)))

	if creatorAccount.ID == 0 {
		uc.logger.Error("creator account not found", zap.Int("promo_creator", promo.Creator))
		return errors.New("creator account not found")
	}

	promoGoals := promoCfg.RewardLevels[promo.Level]
	promoMsg := i18n.PromoMessages[language.Language(creatorAccount.Language)]

	uc.logger.Debug("promo goals found",
		zap.String("name", promo.Name),
		zap.Int("level", promo.Level),
		zap.Int("clients", promo.Clients),
		zap.Int("referral_bonus_days", promoGoals.ReferralBonusDays))

	text := fmt.Sprintf("%s %s %s <a href=\"tg://user?id=%d\">%s</a>\n%s: %d",
		promoMsg.Promo, strings.ToUpper(promo.Name), promoMsg.Attracted,
		senderID, promoMsg.NewClient,
		promoMsg.AttractedClients, promo.Clients)

	reward, hasReward := promoGoals.ClientRewards[promo.Clients]
	if hasReward {
		duration := time.Duration(reward.Days) * 24 * time.Hour
		if err := uc.subscriptionService.AddExpiresAt(ctx, promo.Creator, duration); err != nil {
			uc.logger.Error("failed to add subscription expires at", zap.Error(err))
		} else {
			text += fmt.Sprintf("\n%s: %d", promoMsg.ReceivedDays, reward.Days)
		}

		if reward.Discount > 0 {
			if err := uc.subscriptionService.SetDiscount(ctx, promo.Creator, reward.Discount); err != nil {
				uc.logger.Error("failed to set discount", zap.Error(err))
			} else {
				text += fmt.Sprintf("\n%s: %d%%", promoMsg.Discount, reward.Discount)
			}
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

	clientsTarget := promoGoals.ClientTargets[len(promoGoals.ClientTargets)-1]
	uc.logger.Debug("promo level check", zap.Int("clients", promo.Clients), zap.Int("clients_target", clientsTarget))

	if promo.Level < 2 && promo.Clients >= clientsTarget {
		if err := uc.upgradePromoLevel(botAPI, promo.Name, externalAccountID, promoGoals.ReferralBonusDays, promo, promoMsg); err != nil {
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

	if err := uc.subscriptionService.AddExpiresAt(ctx, promo.Creator, duration); err != nil {
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
