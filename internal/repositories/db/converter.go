package db

import (
	"github.com/quickpowered/thebeyond-master/configs/currency"
	"github.com/quickpowered/thebeyond-master/configs/language"
	"github.com/quickpowered/thebeyond-master/internal/domain"
	"github.com/quickpowered/thebeyond-master/pkg/sqlc"
)

func AccountByIDToDomain(account sqlc.GetAccountByIDRow) domain.Account {
	var promo *string
	if account.Promo.Valid {
		promo = &account.Promo.String
	}

	return domain.Account{
		ID:                    int(account.ID),
		ClusterID:             int(account.ClusterID.Int16),
		KeyID:                 account.KeyID,
		ShortID:               account.ShortID,
		Tariff:                int(account.Tariff),
		Discount:              int(account.Discount.Int32),
		Promo:                 promo,
		FreemiumStatus:        string(account.FreemiumStatus),
		SubscriptionStatus:    string(account.SubscriptionStatus),
		SubscriptionExpiresAt: &account.SubscriptionExpiresAt.Time,
		Devices:               int(account.Devices.Int16),
		UsedUplink:            account.UsedUplink,
		UsedDownlink:          account.UsedDownlink,
		Protocol:              account.Protocol,
		Region:                account.Region,
		Language:              language.Language(account.Language),
		Currency:              currency.Currency(account.Currency),
		LastKeyRefreshAt:      &account.LastKeyRefreshAt.Time,
		LastTrafficRefreshAt:  &account.LastTrafficRefreshAt.Time,
	}
}

func AccountByKeyIDToDomain(account sqlc.GetAccountByKeyIDRow) domain.Account {
	var promo *string
	if account.Promo.Valid {
		promo = &account.Promo.String
	}

	return domain.Account{
		ID:                    int(account.ID),
		ClusterID:             int(account.ClusterID.Int16),
		KeyID:                 account.KeyID,
		ShortID:               account.ShortID,
		Tariff:                int(account.Tariff),
		Discount:              int(account.Discount.Int32),
		Promo:                 promo,
		FreemiumStatus:        string(account.FreemiumStatus),
		SubscriptionStatus:    string(account.SubscriptionStatus),
		SubscriptionExpiresAt: &account.SubscriptionExpiresAt.Time,
		Devices:               int(account.Devices.Int16),
		UsedUplink:            account.UsedUplink,
		UsedDownlink:          account.UsedDownlink,
		Protocol:              account.Protocol,
		Region:                account.Region,
		Language:              language.Language(account.Language),
		Currency:              currency.Currency(account.Currency),
		LastKeyRefreshAt:      &account.LastKeyRefreshAt.Time,
		LastTrafficRefreshAt:  &account.LastTrafficRefreshAt.Time,
	}
}
