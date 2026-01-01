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
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/promo"
	ccurrency "github.com/quickpowered/thebeyond-master/configs/currency"
	"github.com/quickpowered/thebeyond-master/internal/repositories/tariffs"
	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager/dto"
	"github.com/quickpowered/thebeyond-master/internal/services/account"
	"github.com/quickpowered/thebeyond-master/internal/services/payment"
	"github.com/quickpowered/thebeyond-master/internal/services/subscription"
	"github.com/quickpowered/thebeyond-master/internal/services/xraymanager"
	"github.com/quickpowered/thebeyond-master/pkg/consts"
	"github.com/quickpowered/thebeyond-master/pkg/utils"
	"go.uber.org/zap"
)

type UseCase interface {
	HandlePreCheckoutQuery(botAPI *bot.Bot, tgUpdate *models.Update) error
}

type useCase struct {
	accountService      account.Interface
	subscriptionService subscription.Interface
	paymentService      payment.Interface
	promoUseCase        promo.UseCase
	tariffsRepo         tariffs.Repository
	xraymanagerService  xraymanager.Interface
	logger              *zap.Logger
}

func NewUseCase(
	accountService account.Interface,
	subscriptionService subscription.Interface,
	paymentService payment.Interface,
	promoUseCase promo.UseCase,
	tariffsRepo tariffs.Repository,
	xraymanagerService xraymanager.Interface,
	logger *zap.Logger,
) useCase {
	return useCase{
		accountService, subscriptionService,
		paymentService, promoUseCase,
		tariffsRepo, xraymanagerService,
		logger,
	}
}

func (uc useCase) HandlePreCheckoutQuery(tgBot *bot.Bot, tgUpdate *models.Update) error {
	tgUserID := int(tgUpdate.PreCheckoutQuery.From.ID)
	currency := tgUpdate.PreCheckoutQuery.Currency
	payload := tgUpdate.PreCheckoutQuery.InvoicePayload

	uc.logger.Debug("pre checkout query", zap.String("currency", currency), zap.String("payload", payload))

	if currency != "XTR" {
		uc.logger.Error("invalid currency", zap.String("currency", currency), zap.String("payload", payload))
		return errors.New("invalid currency")
	}

	tariffID, days, err := uc.getPayloadData(payload)
	if err != nil {
		uc.logger.Error("invalid subscription payload", zap.Int("days", days))
		return errors.New("invalid payload")
	}

	period, exists := uc.tariffsRepo.GetPeriod(days)
	if !exists {
		uc.logger.Error("invalid subscription period", zap.Int("days", days))
		return errors.New("invalid subscription")
	}

	uc.logger.Debug("subscription found",
		zap.Int("tariff_id", tariffID),
		zap.Int("days", days),
		zap.String("emoji", period.Emoji))

	tariff, exists := uc.tariffsRepo.Get(tariffID)
	if !exists {
		uc.logger.Error("invalid tariff", zap.Int("tariff_id", tariffID))
		return errors.New("invalid tariff")
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()

	account, err := uc.accountService.Get(ctx, consts.PlatformTelegram, tgUserID)
	if err != nil {
		uc.logger.Error("get account failed", zap.Error(err), zap.Int("user_id", tgUserID))
		return err
	}

	finalPrice, extraDays := uc.tariffsRepo.CalculateRenewal(account, ccurrency.XTR, tariff, tariffID, days)
	price := float64(finalPrice)
	price = math.Round(price - price*float64(account.Discount)/100)

	if err := uc.subscriptionService.SetTariff(ctx, account.ID, tariffID); err != nil {
		uc.logger.Error("set tariff failed", zap.Error(err), zap.Int("account_id", account.ID))
		return err
	}

	uc.logger.Debug("subscription price",
		zap.Float64("price", price),
		zap.Int("account_discount", account.Discount),
		zap.Int("extra_days", extraDays))

	totalDays := days + extraDays
	paymentExpiresAt := time.Now().Add(time.Duration(totalDays) * time.Hour * 24)
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

	if !account.IsActive() {
		uc.logger.Debug("account has no subscription expires at",
			zap.Int("account_id", account.ID))

		if err := uc.xraymanagerService.AddClient(ctx, dto.RegionRussia, account.KeyID, utils.NewEmail(account.ID)); err != nil {
			uc.logger.Error("add user failed", zap.Error(err),
				zap.Int("account_id", account.ID))
			return err
		}
	}

	duration := time.Duration(totalDays) * time.Hour * 24
	if err := uc.subscriptionService.AddExpiresAt(ctx, account.ID, duration); err != nil {
		uc.logger.Error("add subscription expires at failed", zap.Error(err),
			zap.Int("account_id", account.ID),
			zap.Duration("duration", duration))
		return err
	}

	uc.logger.Debug("account subscription expires at added",
		zap.Int("account_id", account.ID),
		zap.Duration("duration", duration))

	if err := uc.subscriptionService.ResetDiscount(ctx, account.ID); err != nil {
		return err
	}

	tgBot.AnswerPreCheckoutQuery(ctx, &bot.AnswerPreCheckoutQueryParams{
		PreCheckoutQueryID: tgUpdate.PreCheckoutQuery.ID,
		OK:                 true,
	})

	return nil
}

func (uc useCase) getPayloadData(payload string) (tariffID, days int, err error) {
	tariffID = 1
	for _, part := range strings.Split(payload, ";") {
		key, val, ok := strings.Cut(part, ":")
		if !ok {
			return tariffID, days, fmt.Errorf("invalid payload part: %s", part)
		}

		value, err := strconv.Atoi(val)
		if err != nil {
			return tariffID, days, err
		}

		switch key {
		case "d":
			days = value
		case "t":
			tariffID = value
		default:
			return tariffID, days, fmt.Errorf("unknown key: %s", part)
		}
	}
	return
}
