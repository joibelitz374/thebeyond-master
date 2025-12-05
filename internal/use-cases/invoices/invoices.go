package invoices

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/quickpowered/frilly/internal/repositories/subscriptions"
	"github.com/quickpowered/frilly/internal/repositories/xray"
	"github.com/quickpowered/frilly/internal/services/account"
	"github.com/quickpowered/frilly/internal/services/payment"
	"github.com/quickpowered/frilly/internal/use-cases/promo"
	"github.com/quickpowered/frilly/pkg/consts"
	"go.uber.org/zap"
)

type UseCase interface {
	HandlePreCheckoutQuery(botAPI *bot.Bot, tgUpdate *models.Update) error
}

type useCase struct {
	accountService    account.Interface
	paymentService    payment.Interface
	promoUseCase      promo.UseCase
	subscriptionsRepo subscriptions.Repository
	xrayClient        xray.Interface
	logger            *zap.Logger
}

func NewUseCase(
	accountService account.Interface,
	paymentService payment.Interface,
	promoUseCase promo.UseCase,
	subscriptionsRepo subscriptions.Repository,
	xrayClient xray.Interface,
	logger *zap.Logger,
) useCase {
	return useCase{
		accountService,
		paymentService,
		promoUseCase,
		subscriptionsRepo,
		xrayClient,
		logger,
	}
}

func (uc useCase) HandlePreCheckoutQuery(tgBot *bot.Bot, tgUpdate *models.Update) error {
	tgUserID := int(tgUpdate.PreCheckoutQuery.From.ID)
	currency := tgUpdate.PreCheckoutQuery.Currency
	payload := tgUpdate.PreCheckoutQuery.InvoicePayload

	uc.logger.Debug("pre checkout query",
		zap.String("currency", currency),
		zap.String("payload", payload))

	if currency != "XTR" {
		uc.logger.Error("invalid currency",
			zap.String("currency", currency),
			zap.String("payload", payload))
		return errors.New("invalid currency")
	}

	payloadParts := strings.Split(payload, ":")
	if len(payloadParts) != 2 || payloadParts[0] != "d" {
		uc.logger.Error("invalid payload",
			zap.String("payload", payload),
			zap.String("payload_parts", strings.Join(payloadParts, ",")))
		return errors.New("invalid payload")
	}

	days, err := strconv.Atoi(payloadParts[1])
	if err != nil {
		uc.logger.Error("invalid days",
			zap.Error(err),
			zap.String("days", payloadParts[1]))
		return err
	}

	subscription := uc.subscriptionsRepo.GetByDays(days)
	if subscription.Emoji == "" {
		uc.logger.Error("invalid subscription",
			zap.String("days", payloadParts[1]))
		return errors.New("invalid subscription")
	}

	uc.logger.Debug("subscription found",
		zap.String("days", payloadParts[1]),
		zap.String("emoji", subscription.Emoji))

	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()

	account, err := uc.accountService.Get(ctx, consts.PlatformTelegram, tgUserID)
	if err != nil {
		uc.logger.Error("get account failed", zap.Error(err),
			zap.Int("user_id", tgUserID))
		return err
	}

	price := subscription.Prices["stars"]
	price = math.Round(price - price*float64(account.Discount)/100)

	uc.logger.Debug("subscription price",
		zap.Float64("price", price),
		zap.Int("account_discount", account.Discount),
		zap.Int("subscription_discount", subscription.Discount))

	paymentExpiresAt := time.Now().Add(time.Duration(days) * time.Hour * 24)
	if err := uc.paymentService.Create(ctx, account.ID, int(price), paymentExpiresAt); err != nil {
		uc.logger.Error("create payment failed", zap.Error(err),
			zap.Int("account_id", account.ID),
			zap.Float64("price", price),
			zap.Time("payment_expires_at", paymentExpiresAt))
		return err
	}

	uc.logger.Debug("payment created",
		zap.Int("account_id", account.ID),
		zap.Float64("price", price),
		zap.Time("payment_expires_at", paymentExpiresAt))

	email := fmt.Sprintf("id%d@user", account.ID)

	if account.SubscriptionExpiresAt == nil {
		uc.logger.Debug("account has no subscription expires at",
			zap.Int("account_id", account.ID))
		if err := uc.xrayClient.AddUser(email, account.KeyID); err != nil {
			uc.logger.Error("add user failed", zap.Error(err),
				zap.Int("account_id", account.ID))
			return err
		}
	}

	duration := time.Duration(days) * time.Hour * 24
	if err := uc.accountService.AddSubscriptionExpiresAt(ctx, account.ID, duration); err != nil {
		uc.logger.Error("add subscription expires at failed", zap.Error(err),
			zap.Int("account_id", account.ID),
			zap.Duration("duration", duration))
		return err
	}

	uc.logger.Debug("account subscription expires at added",
		zap.Int("account_id", account.ID),
		zap.Duration("duration", duration))

	if account.Promo != nil && price > 100 {
		promo, _, err := uc.promoUseCase.Get(*account.Promo)
		if err != nil {
			uc.logger.Error("get promo failed", zap.Error(err),
				zap.Int("account_id", account.ID))
			return err
		}

		if err := uc.promoUseCase.RegisterReferral(tgBot, tgUserID, account.ID, promo, consts.PromoTargetBuyer); err != nil {
			uc.logger.Error("register referral failed", zap.Error(err),
				zap.Int("account_id", account.ID))
			return err
		}
	}

	tgBot.AnswerPreCheckoutQuery(ctx, &bot.AnswerPreCheckoutQueryParams{
		PreCheckoutQueryID: tgUpdate.PreCheckoutQuery.ID,
		OK:                 true,
	})

	return nil
}
