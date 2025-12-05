package commands

import (
	"fmt"
	"log"
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
			p[currency] = round(period.prices["rub"]*rates[currency], 2)
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
	"7": {
		index: 1,
		emoji: "🍿",
		prices: map[string]float64{
			"stars": 24,
			"rub":   44,
		},
	},
	"14": {
		index: 2,
		emoji: "💖",
		prices: map[string]float64{
			"stars": 48,
			"rub":   88,
		},
	},
	"30": {
		index: 3,
		emoji: "🌕",
		prices: map[string]float64{
			"stars": 100,
			"rub":   179,
		},
	},
	"90": {
		index: 4,
		emoji: "🌷",
		prices: map[string]float64{
			"stars": 250,
			"rub":   424,
		},
	},
	"180": {
		index: 5,
		emoji: "🌈", // 🍀
		prices: map[string]float64{
			"stars": 500,
			"rub":   829,
		},
	},
	"365": {
		index: 6,
		emoji: "⭐️",
		prices: map[string]float64{
			"stars": 1000,
			"rub":   1649,
		},
	},
}

// func applyDiscounts() error {
// 	monthly, ok := SubscriptionPeriods["30"]
// 	if !ok {
// 		return fmt.Errorf("нет месячного периода (ключ \"30\")")
// 	}

// 	monthlyPrice, ok := monthly.prices["rub"]
// 	if !ok {
// 		return fmt.Errorf("в месячном периоде нет цены в рублях (ключ \"rub\")")
// 	}

// 	type item struct {
// 		days string
// 		sp   *subscriptionPeriod
// 	}
// 	var list []item
// 	for days, sp := range SubscriptionPeriods {
// 		list = append(list, item{days: days, sp: sp})
// 	}
// 	sort.Slice(list, func(i, j int) bool {
// 		return list[i].sp.index < list[j].sp.index
// 	})

// 	fmt.Printf("%6s | %10s | %16s | %12s | %9s | %16s\n", "Days", "Price (RUB)", "Month-equiv (RUB)", "Discount (RUB)", "Disc %", "Eff. month (RUB)")
// 	fmt.Println("------+------------+------------------+--------------+-----------+------------------")

// 	for _, it := range list {
// 		var daysFloat float64
// 		_, err := fmt.Sscanf(it.days, "%f", &daysFloat)
// 		if err != nil || daysFloat <= 0 {
// 			continue
// 		}

// 		price, ok := it.sp.prices["rub"]
// 		if !ok {
// 			continue
// 		}

// 		fullEquiv := monthlyPrice * (daysFloat / 30.0)
// 		discountRub := fullEquiv - price

// 		var discountPct float64
// 		if fullEquiv != 0 {
// 			discountPct = discountRub / fullEquiv * 100.0
// 		} else {
// 			discountPct = 0
// 		}

// 		it.sp.discount = int(math.Round(discountPct))

// 		effMonth := price * (30.0 / daysFloat)

// 		fmt.Printf("%6s | %10.2f | %16.2f | %12.2f | %8.2f%% | %16.2f\n",
// 			it.days, price, fullEquiv, discountRub, discountPct, effMonth)
// 	}

// 	return nil
// }

func (period *subscriptionPeriod) GetPrice(currency string) float64 {
	return period.prices[currency]
}

var periodsOrder = [][]string{
	{"7", "14"},
	{"30", "90"},
	{"180"},
	{"365"},
}

type RenewCmd struct {
	tools.Modules
	component     components.Component
	exchangeRates web.ExchangeRatesInterface
}

func NewRenewCmd(modules tools.Modules, exchangeRates web.ExchangeRatesInterface) *RenewCmd {
	if err := updatePrices(SubscriptionPeriods, exchangeRates); err != nil {
		log.Fatalf("Ошибка обновления курсов: %v", err)
	}

	// if err := applyDiscounts(); err != nil {
	// 	log.Fatalf("Ошибка применения скидок: %v", err)
	// }

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
