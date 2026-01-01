package dto

import "github.com/quickpowered/thebeyond-master/internal/domain"

type Config struct {
	Tariffs map[int]domain.Tariff `mapstructure:"tariffs" yaml:"tariffs"`
	Periods map[int]domain.Period `mapstructure:"periods" yaml:"periods"`
}
