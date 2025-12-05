package web

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/quickpowered/frilly/internal/repositories/web/dto"
)

type ExchangeRatesInterface interface {
	Get() (map[string]float64, error)
}

type exchangeRates struct {
	httpClient *http.Client
}

func NewExchangeRates() *exchangeRates {
	httpClient := &http.Client{Timeout: 10 * time.Second}
	return &exchangeRates{httpClient}
}

func (e *exchangeRates) Get() (map[string]float64, error) {
	res, err := e.httpClient.Get("https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/rub.json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	data := dto.ExchangeRatesResponse{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data.RUB, nil
}
