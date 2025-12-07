package domain

import "time"

type Payment struct {
	Amount    int
	CreatedAt time.Time
	ExpiresAt time.Time
}
