package tariffs

import (
	"fmt"
	"math"
	"time"

	"github.com/quickpowered/thebeyond-master/configs/currency"
	"github.com/quickpowered/thebeyond-master/internal/domain"
	"github.com/quickpowered/thebeyond-master/internal/repositories/tariffs/dto"
	"github.com/quickpowered/thebeyond-master/internal/repositories/web"
	decimal "github.com/shopspring/decimal"
	"github.com/spf13/viper"
)

type Repository interface {
	Get(id int) (domain.Tariff, bool)
	GetPeriod(days int) (domain.Period, bool)
	GetTargets() []int
	GetPeriodTargets() [][]int
	CalculateRenewal(account domain.Account, currency currency.Currency, targetTarrif domain.Tariff, targetTarrifID, days int) (finalPrice float64, extraDays int)
}

type repository struct {
	path           string
	Tariffs        map[int]domain.Tariff
	Periods        map[int]domain.Period
	tarrifsTargets []int
	periodsTargets [][]int
	exchangeRates  web.ExchangeRatesInterface
}

func New(path string, exchangeRates web.ExchangeRatesInterface) (*repository, error) {
	vp := viper.New()
	vp.SetConfigFile(path)
	vp.SetConfigType("yaml")

	if err := vp.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	cfg := dto.Config{}
	if err := vp.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	repository := new(repository)
	repository.path = path
	repository.Tariffs = cfg.Tariffs
	repository.Periods = cfg.Periods
	repository.exchangeRates = exchangeRates

	repository.tarrifsTargets = make([]int, len(cfg.Tariffs))
	for i := range repository.Tariffs {
		repository.tarrifsTargets[i] = i
	}
	repository.periodsTargets = cfg.PeriodsTargets

	return repository, nil
}

func (r *repository) Get(id int) (domain.Tariff, bool) {
	tariff, ok := r.Tariffs[id]
	return tariff, ok
}

func (r *repository) GetPeriod(days int) (domain.Period, bool) {
	period, ok := r.Periods[days]
	return period, ok
}

func (r *repository) GetTargets() []int {
	return r.tarrifsTargets
}

func (r *repository) GetPeriodTargets() [][]int {
	return r.periodsTargets
}

func (r *repository) CalculateRenewal(account domain.Account, currency currency.Currency, targetTariff domain.Tariff, targetTariffID, days int) (finalPrice float64, extraDays int) {
	period, exists := r.Periods[days]
	if !exists {
		return 0, 0
	}

	pricePerMonth := decimal.NewFromFloat(targetTariff.Prices[currency])
	if pricePerMonth.IsZero() {
		return 0, 0
	}

	dailyPrice := pricePerMonth.Div(decimal.NewFromInt(30))
	discount := decimal.NewFromFloat(period.Discount)

	targetPrice := dailyPrice.Mul(decimal.NewFromInt(int64(days))).Mul(decimal.New(1, 0).Sub(discount))
	targetPrice = targetPrice.Round(2)

	finalPriceDec := targetPrice
	finalPrice = finalPriceDec.InexactFloat64()

	currentTariff, found := r.Tariffs[account.Tariff]
	isSwap := found && account.IsActive() && account.Tariff != targetTariffID

	if isSwap {
		targetPrice = dailyPrice.Mul(decimal.NewFromInt(int64(days)))
		targetPrice = targetPrice.Round(2)

		daysRemaining := decimal.NewFromFloat(math.Max(0, time.Until(*account.SubscriptionExpiresAt).Hours()/24))
		currentDailyPrice := decimal.NewFromFloat(currentTariff.Prices[currency]).Div(decimal.NewFromInt(30))
		residualValue := daysRemaining.Mul(currentDailyPrice).Round(2)

		if residualValue.GreaterThanOrEqual(targetPrice) {
			finalPriceDec = decimal.Zero
			finalPrice = 0
			surplus := residualValue.Sub(targetPrice)
			extraDaysDec := surplus.Div(dailyPrice).Round(0)
			extraDays = int(extraDaysDec.IntPart())
		} else {
			finalPriceDec = targetPrice.Sub(residualValue).Round(2)
			finalPrice = finalPriceDec.InexactFloat64()
		}
	}

	return finalPrice, extraDays
}
