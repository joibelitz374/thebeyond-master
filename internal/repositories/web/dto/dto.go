package dto

type ExchangeRatesResponse struct {
	USD map[string]float64 `json:"usd"`
}
