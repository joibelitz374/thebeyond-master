package interfaces

import (
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
)

type Command interface {
	Execute(bot bin.Interface, payload *domain.Payload) error
}
