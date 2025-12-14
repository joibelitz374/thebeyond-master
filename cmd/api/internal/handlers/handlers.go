package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
	"github.com/quickpowered/thebeyond-master/cmd/api/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/api/internal/middlewares"
	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager"
	"github.com/quickpowered/thebeyond-master/internal/services/account"
	"github.com/quickpowered/thebeyond-master/pkg/consts"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

type subscription struct {
	accountService  account.Interface
	xraymanagerRepo xraymanager.Repository
}

func NewSubscription(accountService account.Interface, xraymanagerRepo xraymanager.Repository) subscription {
	return subscription{accountService, xraymanagerRepo}
}

func (h subscription) Auth(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"access_token": "",
	})
}

func (h subscription) Get(c fiber.Ctx) error {
	initData, ok := c.Context().Value(middlewares.InitDataKey).(initdata.InitData)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Init data not found",
		})
	}

	telegramUserID := int(initData.User.ID)

	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	account, err := h.accountService.Get(ctx, consts.PlatformTelegram, telegramUserID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Account not found",
			})
		} else {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
	}

	return c.JSON(domain.User{
		ID:     telegramUserID,
		Name:   initData.User.FirstName,
		Avatar: initData.User.PhotoURL,
		Subscription: domain.Subscription{
			ID:        account.KeyID,
			ExpiresAt: int(account.SubscriptionExpiresAt.UnixMilli()),
		},
		Language: account.Language,
		Currency: account.Currency,
	})
}
