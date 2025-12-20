package kassa

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/quickpowered/thebeyond-master/internal/repositories/kassa/lava/dto"
)

type Repository interface {
	GetAvailableTariffs() (availableTariffs dto.AvailableTariffs, err error)
	CreateInvoice(amount int) (invoice dto.Invoice, err error)
}

type kassa struct {
	secretKey, shopID string
	client            *http.Client
}

func NewRepository(secretKey, shopID string) Repository {
	return kassa{secretKey, shopID, http.DefaultClient}
}

func (k kassa) GetAvailableTariffs() (availableTariffs dto.AvailableTariffs, err error) {
	body, err := json.Marshal(map[string]any{
		"shopId": k.shopID,
	})
	if err != nil {
		return availableTariffs, err
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.lava.ru/business/invoice/get-available-tariffs", bytes.NewReader(body))
	if err != nil {
		return availableTariffs, err
	}

	k.setHeaders(req, k.sign(body))

	res, err := k.client.Do(req)
	if err != nil {
		return availableTariffs, err
	}

	if err := json.NewDecoder(res.Body).Decode(&availableTariffs); err != nil {
		return availableTariffs, err
	}

	return availableTariffs, nil
}

func (k kassa) CreateInvoice(amount int) (invoice dto.Invoice, err error) {
	orderId, err := gonanoid.New(21)
	if err != nil {
		return invoice, err
	}

	body, err := json.Marshal(map[string]any{
		"shopId":         k.shopID,
		"sum":            amount,
		"orderId":        orderId,
		"hookUrl":        "https://lava.ru",
		"successUrl":     "https://lava.ru",
		"failUrl":        "https://lava.ru",
		"expire":         300,
		"customFields":   "test",
		"comment":        "test",
		"includeService": []string{"lava_pay_in"},
	})
	if err != nil {
		return invoice, err
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.lava.ru/business/invoice/create", bytes.NewReader(body))
	if err != nil {
		return invoice, err
	}

	k.setHeaders(req, k.sign(body))

	res, err := k.client.Do(req)
	if err != nil {
		return invoice, err
	}

	if err := json.NewDecoder(res.Body).Decode(&invoice); err != nil {
		return invoice, err
	}

	return invoice, nil
}

func (k kassa) setHeaders(req *http.Request, signature string) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Signature", signature)
}

func (k kassa) sign(data []byte) string {
	mac := hmac.New(sha256.New, []byte(k.secretKey))
	mac.Write(data)
	signatureBytes := mac.Sum(nil)
	signature := hex.EncodeToString(signatureBytes)
	return signature
}
