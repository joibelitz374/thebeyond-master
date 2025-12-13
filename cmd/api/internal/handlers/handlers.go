package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
	"github.com/quickpowered/thebeyond-master/cmd/api/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/api/internal/middlewares"
	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager"
	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager/dto"
	"github.com/quickpowered/thebeyond-master/internal/services/account"
	"github.com/quickpowered/thebeyond-master/pkg/consts"
	errorsvar "github.com/quickpowered/thebeyond-master/pkg/errors"
	"github.com/quickpowered/thebeyond-master/pkg/utils"
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

func (h subscription) Info(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	account, err := h.accountService.GetByKeyID(ctx, c.Params("sub_id"))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Subscription not found",
			})
		} else {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
	}

	expire := time.Now().Add(32 * time.Hour).Unix()
	subscription := "#profile-title: base64:8J+SmyBCZXlvbmQgU2VjdXJl" +
		"\n#profile-update-interval: 12" +
		"\n#profile-web-page-url: https://t.me/beyondsecurenews" +
		"\n#support-url: https://t.me/beyondsecurenews?direct" +
		"\n#notification-subs-expire: 1" +
		"\n#announce: base64:8J+XkiDigJQgV2hpdGVsaXN0CvCfp6Ag4oCUIFNtYXJ0CvCfjZQg4oCUIFdBUlA=" +
		fmt.Sprintf("\n#subscription-userinfo: expire=%v", expire)

	lists := []dto.ListType{dto.NodeTypeBlacklist}
	if account.WhitelistExpiresAt != nil {
		lists = append(lists, dto.NodeTypeWhitelist)
	}

	for _, list := range lists {
		nodes, err := h.xraymanagerRepo.GetNodes(ctx, dto.ClusterID(account.ClusterID), list)
		if err != nil {
			if errors.Is(err, errorsvar.ErrListNodesNotFound) {
				continue
			}

			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}

		for _, node := range nodes {
			subscription += "\n" + utils.GenerateVLESSURI(node, account.KeyID, node.PublicKey, account.ShortID)
		}
	}

	return c.SendString(subscription)
}
