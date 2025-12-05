package dto

type ExchangeRatesResponse struct {
	RUB map[string]float64 `json:"rub"`
}
