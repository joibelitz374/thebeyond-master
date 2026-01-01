package web

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/quickpowered/thebeyond-master/internal/repositories/web/dto"
)

type ExchangeRatesInterface interface {
	Get() (map[string]float64, error)
}

type repository struct {
	httpClient *http.Client
}

func NewExchangeRates() *repository {
	return &repository{&http.Client{Timeout: 10 * time.Second}}
}

func (e *repository) Get() (map[string]float64, error) {
	res, err := e.httpClient.Get("https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/usd.json")
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

	return data.USD, nil
}
