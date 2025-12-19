package commands

import (
	"fmt"
	"strconv"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/telegram"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	"github.com/quickpowered/thebeyond-master/configs/currency"
	"github.com/quickpowered/thebeyond-master/configs/language"
	"github.com/quickpowered/thebeyond-master/pkg/consts"
)

const RENEW_CMD = "renew"

type renewHandler struct {
	deps.Dependencies
}

func NewRenewHandler(deps deps.Dependencies) renewHandler {
	return renewHandler{deps}
}

var familyPrices = map[string]float64{
	"7":   20,
	"14":  40,
	"30":  80,
	"90":  190,
	"180": 370,
	"365": 736,
}

func (h renewHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	msg := i18n.RenewMessages[language.Language(p.Account.Language)]
	opts := []any{deps.ToForward(bot, p)}

	switch {
	// case 2:
	// 	opts = append(opts, &types.Keyboard{
	// 		ButtonRows: [][]types.Button{
	// 			{{Text: "🚼 Себе", Data: "renew " + p.Args[1] + " self"}},
	// 			{{
	// 				Text: fmt.Sprintf("👨‍👩‍👦‍👦 Семье | +%.2f₽ ", familyPrices[p.Args[1]]),
	// 				Data: "renew " + p.Args[1] + " family",
	// 			}},
	// 			{{Text: "◀️ Назад", Data: RENEW_CMD}},
	// 		},
	// 	})

	// 	return bot.SendMessage(p.Message.Chat(), "Выберите кому подписка:", opts...)
	// case 3:
	// 	prevArgs := strings.Join(p.Args[1:3], " ")
	// 	opts = append(opts, &types.Keyboard{
	// 		ButtonRows: [][]types.Button{
	// 			{
	// 				{Text: "⭐️ Stars", Data: fmt.Sprintf("renew %s stars", prevArgs)},
	// 				{Text: "💳 Карта", Data: fmt.Sprintf("renew %s card", prevArgs)},
	// 			},
	// 			{{Text: "📳 Система быстрых платежей", Data: fmt.Sprintf("renew %s sbp", prevArgs)}},
	// 			{{Text: "◀️ Назад", Data: RENEW_CMD + " " + p.Args[1]}},
	// 		},
	// 	})

	// 	return bot.SendMessage(p.Message.Chat(), "Выберите способ оплаты:", opts...)
	case len(p.Args) >= 2:
		devices := 1

		days, err := strconv.Atoi(p.Args[1])
		if err != nil {
			return err
		}

		subscription := h.SubscriptionsRepo.GetByDays(days)
		if subscription.Emoji == "" {
			return bot.SendMessage(p.Message.Chat(), "Subscription not found")
		}

		if bot.GetPlatform() != consts.PlatformTelegram {
			return bot.SendMessage(p.Message.Chat(), "This command is only available in Telegram", opts...)
		}

		if tgBot, ok := bot.(*telegram.Adapter); ok {
			price := subscription.Prices["stars"]
			price = price - price*float64(p.Account.Discount)/100

			// if len(p.Args) >= 3 {
			// 	if p.Args[2] == "family" {
			// 		devices = 5
			// 		price += familyPrices[strconv.Itoa(days)]
			// 	}

			if err := tgBot.SendInvoice(
				p.Message.Chat(),
				h.newTitle(msg, days),
				"Оплачивая, вы соглашаетесь с правилами сервиса, включая политику возвратов. Доступ предоставляется сразу после успешной оплаты. Спасибо за выбор!",
				fmt.Sprintf("d:%d;devices:%d", days, devices),
				"XTR", int(price),
			); err != nil {
				return err
			}

			return bot.SendMessage(p.Message.Chat(), "Return to main menu:", &types.Keyboard{
				ButtonRows: [][]types.Button{
					{{Text: "🔄 Menu", Data: START_CMD}},
				},
			})
		}
	}

	var totalDiscount int
	targets := h.SubscriptionsRepo.GetTargets()

	buttonRows := make([][]types.Button, len(targets)+2)
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
			if p.Account.NetworkType == "mobile" {
				price *= 1.7
			}

			label := fmt.Sprintf("%s %s — %.2f %s",
				subscription.Emoji,
				h.newTitle(msg, days),
				price,
				currency.Emoji,
			)

			if subscription.Discount > 0 {
				label += fmt.Sprintf(" (%d)", subscription.Discount)
			}

			buttonRows[i+1] = append(buttonRows[i+1], types.Button{
				Text: label,
				Data: fmt.Sprintf("renew %d", days),
			})
		}
	}

	buttonRows[len(buttonRows)-1] = append(buttonRows[len(buttonRows)-1], types.Button{
		Text: "◀️ Назад",
		Data: MENU_CMD,
	})

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
