package domain

import "time"

type Account struct {
	ID                    int
	ClusterID             int
	KeyID                 string
	ShortID               string
	NetworkType           string
	Devices               int
	Region                string
	Language              string
	Currency              string
	Protocol              string
	SubscriptionExpiresAt *time.Time
	WhitelistExpiresAt    *time.Time
	LastKeyRefreshAt      *time.Time
	Promo                 *string
	Discount              int
}

func (a Account) IsActive() (active bool) {
	if a.SubscriptionExpiresAt != nil && a.SubscriptionExpiresAt.After(time.Now()) {
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
