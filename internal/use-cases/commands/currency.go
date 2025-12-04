package commands

import (
	"context"
	"strconv"
	"time"

	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/types"
	"github.com/quickpowered/frilly/internal/use-cases/commands/tools"
	"github.com/quickpowered/frilly/internal/use-cases/components"
	"github.com/quickpowered/frilly/pkg/consts"
	"github.com/quickpowered/frilly/pkg/values"
)

const CURRENCY_CMD = "currency"

var currenciesOrder = [][]string{
	{"usd", "eur", "gbp"},
	{"rub", "uah", "pln"},
	{"czk", "ron", "huf"},
	{"bgn", "rsd"},
	{"sek", "nok", "dkk"},
	{"inr", "cny", "jpy"},
	{"krw", "vnd", "thb"},
	{"idr", "myr", "php"},
	{"sgd"},
	{"pkr", "bdt"},
	{"brl", "mxn", "cad"},
	{"try", "irr", "ils"},
	{"sar", "aed"},
}

var vkCurrenciesPages = [][][]string{
	{
		{"usd", "eur"},
		{"gbp", "rub"},
		{"sek", "nok"},
		{"dkk", "try"},
	},
	{
		{"uah", "pln"},
		{"czk", "huf"},
		{"ron", "bgn"},
		{"rsd", "ils"},
	},
	{
		{"cny", "jpy"},
		{"krw", "inr"},
		{"pkr", "bdt"},
	},
	{
		{"sgd", "thb"},
		{"vnd", "idr"},
		{"myr", "php"},
	},
	{
		{"cad", "brl"},
		{"mxn", "sar"},
		{"aed", "irr"},
	},
}

type CurrencyCmd struct {
	tools.Modules
	component components.Component
}

func NewCurrencyCmd(modules tools.Modules) *CurrencyCmd {
	return &CurrencyCmd{modules, components.NewCurrencyComponent()}
}

func (c *CurrencyCmd) Execute(bot bin.Interface, payload *domain.Payload) error {
	componentText := c.component.Text(payload.Account.Language)
	opts := []any{tools.ToForward(bot, payload)}

	var page int
	if len(payload.NodeRoute) >= 2 {
		arg, err := strconv.Atoi(payload.NodeRoute[1])
		if err != nil {
			payload.Account.Currency = payload.NodeRoute[1]
			getter := func() (values.Value, bool) { return values.GetCurrency(payload.Account.Currency) }
			setter := func() error {
				ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
				defer cancel()
				return c.AccountService.SetCurrency(ctx, payload.Account.ID, payload.Account.Currency)
			}

			if err := setValue(bot, payload, componentText, getter, setter, opts...); err != nil {
				return err
			}
			return nil
		} else {
			page = arg
			if page < 1 || page >= len(vkCurrenciesPages) {
				return bot.SendMessage(payload.Message.GetChat(), "Invalid page number", opts...)
			}
		}
	}

	getter := func(code string) string {
		currency, ok := values.GetCurrency(code)
		if !ok {
			return ""
		}
		return currency.Flag + " " + currency.Name
	}

	buttonRows, err := getButtons(bot, page, "currency", vkCurrenciesPages, currenciesOrder, getter)
	if err != nil {
		bot.SendMessage(payload.Message.GetChat(), "Invalid currency code", opts...)
		return err
	}

	if bot.GetPlatform() != consts.PlatformTelegram {
		buttonRows = append(buttonRows, addNavigationButtons(bot, page)...)
	}

	opts = append(opts, &types.Keyboard{ButtonRows: buttonRows})
	return bot.SendMessage(payload.Message.GetChat(), componentText[0]+":", opts...)
}

func setValue(
	bot bin.Interface,
	payload *domain.Payload, componentText []string,
	getter func() (values.Value, bool),
	setter func() error,
	opts ...any,
) error {
	val, ok := getter()
	if !ok {
		return bot.SendMessage(payload.Message.GetChat(), "Invalid code", opts...)
	}

	if err := setter(); err != nil {
		return err
	}

	return bot.SendMessage(payload.Message.GetChat(), componentText[1]+" "+val.Flag+" "+val.Name, opts...)
}
