package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types/update"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	"github.com/quickpowered/thebeyond-master/configs/language"
	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager/dto"
	"github.com/quickpowered/thebeyond-master/pkg/email"
	"go.uber.org/zap"
)

const REFUND_CMD = "refund"

type refundHandler struct {
	deps.Dependencies
}

func NewRefundHandler(deps deps.Dependencies) refundHandler {
	return refundHandler{deps}
}

func (h refundHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	msg := i18n.RefundMessages[language.Language(p.Account.Language)]
	opts := []any{deps.ToForward(bot, p), types.DisableMentions}

	if len(p.Args) >= 2 {
		if p.Args[1] == "confirm" {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			payments, err := h.PaymentService.Get(ctx, p.Account.ID, time.Now())
			if err != nil {
				h.Logger.Error("failed to get payment", zap.Error(err))
				return err
			}

			if len(payments) == 0 {
				return bot.SendMessage(p.Message.Chat(), msg.NoPayments, opts...)
			}

			refundAmount := 0.0
			now := time.Now()
			for _, payment := range payments {
				if payment.ExpiresAt.After(now) {
					totalDuration := payment.ExpiresAt.Sub(payment.CreatedAt)
					remainingDuration := payment.ExpiresAt.Sub(now)

					if totalDuration > 0 {
						ratio := float64(remainingDuration) / float64(totalDuration)
						refundAmount += float64(payment.Amount) * ratio
					}
				}
			}

			var durationToRemove time.Duration
			for _, payment := range payments {
				if payment.ExpiresAt.After(now) {
					remaining := payment.ExpiresAt.Sub(now)
					if remaining > 0 {
						durationToRemove += remaining
					}
				}
			}

			remaining := time.Until(*p.Account.SubscriptionExpiresAt)
			if remaining < 0 {
				remaining = 0
			}

			if durationToRemove > remaining {
				durationToRemove = remaining
			}

			if durationToRemove > 0 {
				if err := h.AccountService.RemoveSubscriptionExpiresAt(ctx, p.Account.ID, durationToRemove); err != nil {
					return err
				}
			}

			if err := h.PaymentService.Delete(ctx, p.Account.ID); err != nil {
				return err
			}

			if err := h.XRayManagerRepo.RemoveClient(ctx, dto.ClusterID(p.Account.ClusterID), dto.NodeTypeBlacklist, email.NewAccount(p.Account.ID)); err != nil {
				return err
			}

			if err := bot.SendMessage(update.Chat{ID: 924536264}, fmt.Sprintf("<b>Заявка на возврат</b> от <a href=\"tg://user?id=%d\">пользователя</a>\n<b>Сумма</b>: %.2f\n<b>Дата истечения</b>: %s", p.Message.Sender(), refundAmount, p.Account.SubscriptionExpiresAt.Format("02/01/2006 15:04"))); err != nil {
				return err
			}

			return bot.SendMessage(p.Message.Chat(), msg.SuccessMessage, opts...)
		}
	}

	opts = append(opts, &types.Keyboard{
		ButtonRows: [][]types.Button{
			{{Text: "🎾 Confirm", Data: "refund confirm"}},
		},
	})

	return bot.SendMessage(p.Message.Chat(), msg.InfoText, opts...)
}
