package commands

import (
	"fmt"
	"math"

	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/repositories/bot/telegram"
	"github.com/quickpowered/frilly/internal/repositories/web"
	"github.com/quickpowered/frilly/internal/types"
	"github.com/quickpowered/frilly/internal/use-cases/commands/tools"
	"github.com/quickpowered/frilly/internal/use-cases/components"
	"github.com/quickpowered/frilly/pkg/consts"
	"github.com/quickpowered/frilly/pkg/values"
)

const RENEW_CMD = "renew"

func updatePrices(subs map[string]*subscriptionPeriod, exchangeRates web.ExchangeRatesInterface) error {
	rates, err := exchangeRates.Get()
	if err != nil {
		return err
	}

	for _, period := range subs {
		p := period.prices
		for currency := range values.GetCurrencies() {
			p[currency] = round(period.prices["usd"]*rates[currency], 2)
		}

		period.prices = p
	}

	return nil
}

func round(val float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Round(val*shift) / shift
}

type subscriptionPeriod struct {
	index    int
	emoji    string
	prices   map[string]float64
	discount int
}

var SubscriptionPeriods = map[string]*subscriptionPeriod{
	"3": {
		index: 0,
		emoji: "💊",
		prices: map[string]float64{
			"stars": 25,
			"usd":   0.36,
		},
	},
	"7": {
		index: 1,
		emoji: "🍿",
		prices: map[string]float64{
			"stars": 57,
			"usd":   0.81,
		},
	},
	"14": {
		index: 2,
		emoji: "💖",
		prices: map[string]float64{
			"stars": 109,
			"usd":   1.57,
		},
	},
	"30": {
		index: 3,
		emoji: "🌕",
		prices: map[string]float64{
			"stars": 230,
			"usd":   3.29, // 2.99
		},
	},
	"90": {
		index: 4,
		emoji: "🌷",
		prices: map[string]float64{
			"stars": 576,
			"usd":   8.99, // 7.49
		},
		discount: 9,
		// discount: 16,
	},
	"180": {
		index: 5,
		emoji: "🍀",
		prices: map[string]float64{
			"stars": 999,
			"usd":   17.49, // 12.99
		},
		discount: 11,
		// discount: 28,
	},
	"365": {
		index: 6,
		emoji: "🌈",
		prices: map[string]float64{
			"stars": 1538,
			"usd":   31.49, // 19.99
		},
		discount: 20,
		// discount: 44,
	},
	"730": {
		index: 7,
		emoji: "⭐️",
		prices: map[string]float64{
			"stars": 2307,
			"usd":   54.99, // 29.99
		},
		discount: 30,
		// discount: 58,
	},
}

func (period *subscriptionPeriod) GetPrice(currency string) float64 {
	return period.prices[currency]
}

var periodsOrder = [][]string{
	{"3"},
	{"7", "14"},
	{"30", "90"},
	{"180"},
	{"365"},
	{"730"},
}

type RenewCmd struct {
	tools.Modules
	component     components.Component
	exchangeRates web.ExchangeRatesInterface
}

func NewRenewCmd(modules tools.Modules, exchangeRates web.ExchangeRatesInterface) *RenewCmd {
	if err := updatePrices(SubscriptionPeriods, exchangeRates); err != nil {
		// log.Fatalf("Ошибка обновления курсов: %v", err)
	}

	return &RenewCmd{modules, components.NewRenewComponent(), exchangeRates}
}

func (c *RenewCmd) Execute(bot bin.Interface, payload *domain.Payload) error {
	componentText := c.component.Text(payload.Account.Language)
	opts := []any{tools.ToForward(bot, payload)}

	if len(payload.NodeRoute) >= 2 {
		days := payload.NodeRoute[1]
		period := SubscriptionPeriods[days]
		description := "Support: @quickworkshop"
		if bot.GetPlatform() != consts.PlatformTelegram {
			return bot.SendMessage(payload.Message.GetChat(), "This command is only available in Telegram", opts...)
		}
		tgBot := bot.(*telegram.Adapter)
		return tgBot.SendInvoice(payload.Message.GetChat(), componentText[period.index+1], description, "d:"+days, "XTR", int(period.GetPrice("stars")))
	}

	buttonRows := make([][]types.Button, len(periodsOrder))
	for i, periods := range periodsOrder {
		for j, periodKey := range periods {
			period := SubscriptionPeriods[periodKey]
			currency, ok := values.GetCurrency(payload.Account.Currency)
			if !ok {
				return bot.SendMessage(payload.Message.GetChat(), "Invalid currency code", opts...)
			}

			buttonRows[i] = append(buttonRows[i], types.Button{
				Text: fmt.Sprintf("%s %s — %.2f %s",
					period.emoji, componentText[period.index+1],
					period.GetPrice(payload.Account.Currency),
					currency.Emoji,
				),
				Data: "renew " + periodKey,
			})

			if period.discount > 0 {
				buttonRows[i][j].Text += fmt.Sprintf(" (-%d%%)", period.discount)
			}
		}
	}

	opts = append(opts, &types.Keyboard{ButtonRows: buttonRows})
	return bot.SendMessage(payload.Message.GetChat(), componentText[0], opts...)
}
