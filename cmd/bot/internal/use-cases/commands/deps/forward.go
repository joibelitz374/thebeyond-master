package deps

import (
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/pkg/consts"
)

func ToForward(bot bin.Interface, payload *domain.Payload) any {
	if platform := bot.GetPlatform(); platform == consts.PlatformDiscord {
		return types.NewForward(payload.Message.Chat().GetThreadID(), payload.Message.ID())
	}

	return types.NewForward(payload.Message.Chat().GetID(), payload.Message.ID())
}
