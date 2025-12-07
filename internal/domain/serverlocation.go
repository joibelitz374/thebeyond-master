package domain

type ServerLocation struct {
	Emoji    string             `mapstructure:"emoji" yaml:"emoji"`
	Prices   map[string]float64 `mapstructure:"prices" yaml:"prices"`
	Discount int                `mapstructure:"discount" yaml:"discount"`
}
