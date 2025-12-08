package middlewares

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

type contextKey string

const InitDataKey contextKey = "init-data"

type auth struct {
	token string
}

func NewAuth(token string) auth {
	return auth{token}
}

func (a auth) Init() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		parts := strings.Split(ctx.Get("Authorization"), " ")
		if len(parts) != 2 {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		authType, authData := parts[0], parts[1]

		switch authType {
		case "tma":
			if err := initdata.Validate(authData, a.token, time.Hour); err != nil {
				return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

			initData, err := initdata.Parse(authData)
			if err != nil {
				return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

			ctx.SetContext(context.WithValue(ctx, InitDataKey, initData))
		}

		return ctx.Next()
	}
}
