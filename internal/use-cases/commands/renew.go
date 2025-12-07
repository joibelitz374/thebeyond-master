package commands

import (
	"fmt"
	"math"
	"strconv"

	"github.com/quickpowered/frilly/config/currency"
	"github.com/quickpowered/frilly/config/language"
	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/i18n"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/repositories/bot/telegram"
	"github.com/quickpowered/frilly/internal/types"
	"github.com/quickpowered/frilly/internal/use-cases/commands/deps"
	"github.com/quickpowered/frilly/pkg/consts"
)

const RENEW_CMD = "renew"

type renewHandler struct {
	deps.Dependencies
}

func NewRenewHandler(deps deps.Dependencies) renewHandler {
	return renewHandler{deps}
}

func (h renewHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	msg := i18n.RenewMessages[language.Language(p.Account.Language)]
	opts := []any{deps.ToForward(bot, p)}

	if len(p.Args) >= 2 {
		days, err := strconv.Atoi(p.Args[1])
		if err != nil {
			return err
		}

		subscription := h.SubscriptionsRepo.GetByDays(days)
		if subscription.Emoji == "" {
			return bot.SendMessage(p.Message.Chat(), "Subscription not found")
		}

		description := ">^<"

		if bot.GetPlatform() != consts.PlatformTelegram {
			return bot.SendMessage(p.Message.Chat(), "This command is only available in Telegram", opts...)
		}

		if tgBot, ok := bot.(*telegram.Adapter); ok {
			title := h.newTitle(msg, days)
			price := subscription.Prices["stars"]
			price = price - price*float64(p.Account.Discount)/100
			return tgBot.SendInvoice(p.Message.Chat(), title, description, "d:"+strconv.Itoa(days), "XTR", int(price))
		}
	}

	var totalDiscount int
	targets := h.SubscriptionsRepo.GetTargets()
	buttonRows := make([][]types.Button, len(targets))

	for i, children := range [][]int{targets[0:2], targets[2:4], targets[4:5], targets[5:]} {
		for _, days := range children {
			subscription := h.SubscriptionsRepo.GetByDays(days)
			senderCurrency := currency.Currency(p.Account.Currency)

			currency, ok := currency.Get(senderCurrency)
			if !ok {
				return bot.SendMessage(p.Message.Chat(), "Invalid currency code", opts...)
			}

			price := subscription.Prices[senderCurrency]
			price = price - price*float64(p.Account.Discount)/100

			label := fmt.Sprintf("%s %s — %.2f %s",
				subscription.Emoji,
				h.newTitle(msg, days),
				math.Round(price),
				currency.Emoji,
			)

			if subscription.Discount > 0 {
				label += fmt.Sprintf(" (%d)", subscription.Discount)
			}

			buttonRows[i] = append(buttonRows[i], types.Button{
				Text: label,
				Data: fmt.Sprintf("renew %d", days),
			})
		}
	}

	if p.Account.Discount > 0 {
		totalDiscount += p.Account.Discount
	}

	text := msg.Question
	if totalDiscount > 0 {
		promoMsg := i18n.PromoMessages[language.Language(p.Account.Language)]
		text += fmt.Sprintf("\n🏷️ "+promoMsg.Discount+": %d%%", totalDiscount)
	}

	opts = append(opts, &types.Keyboard{ButtonRows: buttonRows})
	return bot.SendMessage(p.Message.Chat(), text, opts...)
}

func (h renewHandler) newTitle(msg i18n.RenewLocale, days int) (title string) {
	switch days {
	case 7:
		title = msg.Week
	case 14:
		title = msg.TwoWeeks
	case 30:
		title = msg.Month
	case 90:
		title = msg.Season
	case 180:
		title = msg.HalfYear
	case 365:
		title = msg.Year
	}
	return title
}
