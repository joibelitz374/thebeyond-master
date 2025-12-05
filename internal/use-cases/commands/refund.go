package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/types"
	"github.com/quickpowered/frilly/internal/types/update"
	"github.com/quickpowered/frilly/internal/use-cases/commands/tools"
	"github.com/quickpowered/frilly/internal/use-cases/components"
	"go.uber.org/zap"
)

const REFUND_CMD = "refund"

type RefundCmd struct {
	tools.Modules
	component components.Component
}

func NewRefundCmd(modules tools.Modules) *RefundCmd {
	return &RefundCmd{modules, components.NewRefundComponent()}
}

func (c *RefundCmd) Execute(bot bin.Interface, payload *domain.Payload) error {
	componentText := c.component.Text(payload.Account.Language)
	opts := []any{tools.ToForward(bot, payload), types.DisableMentions}

	if len(payload.NodeRoute) >= 2 {
		if payload.NodeRoute[1] == "confirm" {
			// now := time.Now()
			// expiresAt := payload.Account.SubscriptionExpiresAt.In(now.Location())

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			payments, err := c.PaymentService.Get(ctx, payload.Account.ID, time.Now())
			if err != nil {
				c.Logger.Error("failed to get payment", zap.Error(err))
				return err
			}

			if len(payments) == 0 {
				return bot.SendMessage(payload.Message.GetChat(), "Payments not found", opts...)
			}

			refundAmount := 0.0
			now := time.Now()
			for _, p := range payments {
				if p.ExpiresAt.After(now) {
					totalDuration := p.ExpiresAt.Sub(p.CreatedAt)
					remainingDuration := p.ExpiresAt.Sub(now)

					if totalDuration > 0 {
						ratio := float64(remainingDuration) / float64(totalDuration)
						refundAmount += float64(p.Amount) * ratio
					}
				}
			}

			if err := bot.SendMessage(update.Chat{ID: 924536264}, fmt.Sprintf("<b>Заявка на возврат</b> от <a href=\"tg://user?id=%d\">пользователя</a>\n<b>Сумма</b>: %.2f\n<b>Дата истечения</b>: %s", payload.Message.GetSender(), refundAmount, payload.Account.SubscriptionExpiresAt.Format("02/01/2006 15:04"))); err != nil {
				return err
			}

			if err := c.AccountService.SetSubscriptionExpiresAt(ctx, payload.Account.ID, nil); err != nil {
				return err
			}

			if err := c.PaymentService.Delete(ctx, payload.Account.ID); err != nil {
				return err
			}

			if err := c.XRayClient.RemoveUser(fmt.Sprintf("id%d@user", payload.Account.ID)); err != nil {
				return err
			}

			return bot.SendMessage(payload.Message.GetChat(), componentText[1], opts...)
		}
	}

	opts = append(opts, &types.Keyboard{
		ButtonRows: [][]types.Button{
			{{Text: "🎾 Confirm", Data: "refund confirm"}},
		},
	})

	return bot.SendMessage(payload.Message.GetChat(), componentText[0], opts...)
}
