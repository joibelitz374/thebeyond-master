package domain

import (
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types/update"
	"github.com/quickpowered/thebeyond-master/internal/domain"
)

type Payload struct {
	Args    []string
	Message update.NewMessageInterface
	Account domain.Account
}
