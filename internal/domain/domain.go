package domain

import (
	"time"

	"github.com/quickpowered/frilly/internal/types/update"
)

type Payload struct {
	NodeRoute []string
	Message   update.NewMessageInterface
	Account   Account
}

type Account struct {
	ID                    int
	KeyID                 string
	ShortID               string
	SubscriptionExpiresAt *time.Time
	Region                string
	Language              string
	Currency              string
	Protocol              string
	ServerLocation        string
	LastKeyRefreshAt      *time.Time
}

type Location struct {
	Flag string
	Name string
	IP   string
}

type Payment struct {
	Amount    int
	CreatedAt time.Time
	ExpiresAt time.Time
}
