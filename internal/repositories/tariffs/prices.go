package tariffs

import (
	"os"

	"github.com/quickpowered/thebeyond-master/configs/currency"
	"go.yaml.in/yaml/v3"
)

func (r *repository) updatePrices() error {
	usdExchangeRates, err := r.exchangeRates.Get()
	if err != nil {
		return err
	}

	for tariffID := range r.Tariffs {
		tariff := r.Tariffs[tariffID]
		for currencyAbbr, multiply := range usdExchangeRates {
			usdPrice := tariff.Prices[currency.USD]
			if ok := tariff.Prices != nil; !ok {
				tariff.Prices = make(map[currency.Currency]float64)
			}
			tariff.Prices[currency.Currency(currencyAbbr)] = usdPrice * multiply
		}
	}

	data, err := yaml.Marshal(r)
	if err != nil {
		return err
	}

	return os.WriteFile(r.path, data, 0666)
}
