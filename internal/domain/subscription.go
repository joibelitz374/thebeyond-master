package domain

import "github.com/quickpowered/frilly/config/currency"

type Subscription struct {
	Emoji    string                        `mapstructure:"emoji" yaml:"emoji"`
	Prices   map[currency.Currency]float64 `mapstructure:"prices" yaml:"prices"`
	Discount int                           `mapstructure:"discount" yaml:"discount"`
}
