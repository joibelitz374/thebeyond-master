package db

import (
	"github.com/quickpowered/thebeyond-master/internal/domain"
	"github.com/quickpowered/thebeyond-master/pkg/sqlc"
)

func ToDomain(account sqlc.Account) domain.Account {
	var promo *string
	if account.Promo.Valid {
		promo = &account.Promo.String
	}

	return domain.Account{
		ID:                    int(account.ID),
		KeyID:                 account.KeyID,
		ShortID:               account.ShortID,
		Region:                account.Region,
		Language:              account.Language,
		Currency:              account.Currency,
		Protocol:              account.Protocol,
		SubscriptionExpiresAt: &account.SubscriptionExpiresAt.Time,
		LastKeyRefreshAt:      &account.LastKeyRefreshAt.Time,
		Promo:                 promo,
		Discount:              int(account.Discount.Int32),
	}
}
