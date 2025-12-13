package subscriptions

import (
	"github.com/quickpowered/thebeyond-master/configs/currency"
	"github.com/quickpowered/thebeyond-master/internal/repositories/web"
	"github.com/quickpowered/thebeyond-master/pkg/utils"
)

func (s *subscriptions) updatePrices(exchangeRates web.ExchangeRatesInterface) error {
	rates, err := exchangeRates.Get()
	if err != nil {
		return err
	}

	for days, subscription := range s.Subscriptions {
		if subscription.Prices == nil {
			subscription.Prices = make(map[currency.Currency]float64)
		}

		rub, ok := subscription.Prices["rub"]
		if !ok {
			continue
		}

		for currency := range currency.Currencies {
			rate, ok := rates[string(currency)]
			if !ok {
				continue
			}
			subscription.Prices[currency] = utils.Round(rub*rate, 2)
		}

		s.Subscriptions[days] = subscription
	}

	return nil
}
