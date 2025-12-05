package domain

import (
	"github.com/quickpowered/frilly/internal/types/update"
)

type Payload struct {
	Args    []string
	Message update.NewMessageInterface
	Account Account
}
