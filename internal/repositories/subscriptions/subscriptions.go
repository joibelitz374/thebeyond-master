package subscriptions

import (
	"fmt"
	"os"
	"sort"

	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/repositories/web"
	"github.com/spf13/viper"
	"go.yaml.in/yaml/v3"
)

type Repository interface {
	GetByDays(days int) domain.Subscription
	GetTargets() []int
	IncreasePrices(miltiplier float64) error
	DecreasePrices(miltiplier float64) error
	Save() error
}

type subscriptions struct {
	path          string
	Subscriptions map[int]domain.Subscription
	Targets       []int
}

func New(path string, exchangeRates web.ExchangeRatesInterface) (*subscriptions, error) {
	vp := viper.New()
	vp.SetConfigFile(path)
	vp.SetConfigType("yaml")

	if err := vp.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	raw := map[string]domain.Subscription{}
	if err := vp.Unmarshal(&raw); err != nil {
		return nil, err
	}

	subs := &subscriptions{
		path:          path,
		Subscriptions: make(map[int]domain.Subscription, len(raw)),
	}

	if err := subs.convertKeys(raw); err != nil {
		return nil, err
	}

	if err := subs.updatePrices(exchangeRates); err != nil {
		return nil, err
	}

	if err := subs.Save(); err != nil {
		return nil, err
	}

	subs.Targets = make([]int, 0, len(subs.Subscriptions))
	for days := range subs.Subscriptions {
		subs.Targets = append(subs.Targets, days)
	}

	sort.Slice(subs.Targets, func(i, j int) bool {
		return subs.Targets[i] < subs.Targets[j]
	})

	return subs, nil
}

func (s *subscriptions) GetByDays(days int) domain.Subscription {
	return s.Subscriptions[days]
}

func (s *subscriptions) GetTargets() []int {
	return s.Targets
}

func (s *subscriptions) IncreasePrices(miltiplier float64) error {
	for _, sub := range s.Subscriptions {
		for currency := range sub.Prices {
			sub.Prices[currency] *= miltiplier
		}
	}
	return s.Save()
}

func (s *subscriptions) DecreasePrices(miltiplier float64) error {
	for _, sub := range s.Subscriptions {
		for currency := range sub.Prices {
			sub.Prices[currency] /= miltiplier
		}
	}
	return s.Save()
}

func (s *subscriptions) Save() error {
	data, err := yaml.Marshal(s.Subscriptions)
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, data, 0666)
}
