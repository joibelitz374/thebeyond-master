package domain

import "github.com/quickpowered/thebeyond-master/configs/currency"

type Tariff struct {
	Emoji    string                        `mapstructure:"emoji" yaml:"emoji"`
	Prices   map[currency.Currency]float64 `mapstructure:"prices" yaml:"prices"`
	Traffic  int                           `mapstructure:"traffic" yaml:"traffic"`
	Devices  int                           `mapstructure:"devices" yaml:"devices"`
	Discount string                        `mapstructure:"discount" yaml:"discount"`
}

type Period struct {
	Emoji    string  `mapstructure:"emoji" yaml:"emoji"`
	Discount float64 `mapstructure:"discount" yaml:"discount"`
}
