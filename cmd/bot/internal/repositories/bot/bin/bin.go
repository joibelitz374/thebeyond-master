package bin

import (
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types/update"
	"github.com/quickpowered/thebeyond-master/pkg/consts"
)

type Interface interface {
	GetSelfID() int
	GetAPI() any
	GetPlatform() consts.Platform
	SendMessage(chat update.ChatInterface, text string, opts ...any) error
	DeleteMessage(chat update.ChatInterface, messageID int) error
}
