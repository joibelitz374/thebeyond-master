package domain

import (
	"time"

	"github.com/quickpowered/thebeyond-master/configs/currency"
	"github.com/quickpowered/thebeyond-master/configs/language"
)

type Account struct {
	ID                    int
	ClusterID             int
	KeyID                 string
	ShortID               string
	Tariff                int
	Discount              int
	Promo                 *string
	FreemiumStatus        string
	SubscriptionStatus    string
	SubscriptionExpiresAt *time.Time
	Devices               int
	UsedUplink            int64
	UsedDownlink          int64
	Protocol              string
	Region                string
	Language              language.Language
	Currency              currency.Currency
	LastKeyRefreshAt      *time.Time
	LastTrafficRefreshAt  *time.Time
	// WhitelistExpiresAt    *time.Time
}

type ExternalAccount struct {
	ID                int
	ExternalAccountID int
}

func (a Account) IsActive() (active bool) {
	if a.SubscriptionStatus != "unavailable" || a.FreemiumStatus != "unavailable" {
		active = true
	}
	return
}

func (a Account) IsNewcomer() (newcomer bool) {
	if a.Language == "" || a.Currency == "" {
		newcomer = true
	}
	return
}
