package invoices

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/quickpowered/frilly/internal/repositories/xray"
	"github.com/quickpowered/frilly/internal/services/account"
	"github.com/quickpowered/frilly/internal/services/payment"
	"github.com/quickpowered/frilly/internal/use-cases/commands"
	"github.com/quickpowered/frilly/pkg/consts"
)

type Interface interface {
	HandlePreCheckoutQuery(ctx context.Context, botAPI *bot.Bot, tgUpdate *models.Update) error
}

type UseCase struct {
	paymentService payment.Interface
	accountService account.Interface
	xrayClient     xray.Interface
}

func NewUseCase(paymentService payment.Interface, accountService account.Interface, xrayClient xray.Interface) *UseCase {
	return &UseCase{paymentService, accountService, xrayClient}
}

func (uc *UseCase) HandlePreCheckoutQuery(ctx context.Context, botAPI *bot.Bot, tgUpdate *models.Update) error {
	if tgUpdate.PreCheckoutQuery.Currency != "XTR" {
		return errors.New("invalid currency")
	}

	payloadParts := strings.Split(tgUpdate.PreCheckoutQuery.InvoicePayload, ":")
	if len(payloadParts) != 2 || payloadParts[0] != "d" {
		return errors.New("invalid payload")
	}

	days := payloadParts[1]
	subscriptionPeriod := commands.SubscriptionPeriods[days]
	if subscriptionPeriod == nil {
		return errors.New("invalid subscription period")
	}

	price := subscriptionPeriod.GetPrice("stars")
	if price != float64(tgUpdate.PreCheckoutQuery.TotalAmount) {
		return errors.New("invalid price")
	}

	account, err := uc.accountService.Get(ctx, consts.PlatformTelegram, int(tgUpdate.PreCheckoutQuery.From.ID))
	if err != nil {
		return err
	}

	daysCount, err := strconv.Atoi(days)
	if err != nil {
		return err
	}

	if err := uc.paymentService.Create(ctx, account.ID, int(price), time.Now().Add(time.Duration(daysCount)*time.Hour*24)); err != nil {
		return err
	}

	var expiresAt time.Time
	if account.SubscriptionExpiresAt != nil && !account.SubscriptionExpiresAt.Before(time.Now()) {
		expiresAt = *account.SubscriptionExpiresAt
	} else {
		if err := uc.xrayClient.AddUser(fmt.Sprintf("id%d@user", account.ID), account.KeyID); err != nil {
			return err
		}
		expiresAt = time.Now()
	}

	expiresAt = expiresAt.Add(time.Duration(daysCount) * time.Hour * 24)
	if err := uc.accountService.SetSubscriptionExpiresAt(ctx, account.ID, &expiresAt); err != nil {
		return err
	}

	botAPI.AnswerPreCheckoutQuery(ctx, &bot.AnswerPreCheckoutQueryParams{
		PreCheckoutQueryID: tgUpdate.PreCheckoutQuery.ID,
		OK:                 true,
	})

	return nil
}
