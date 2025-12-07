package deps

import (
	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/types"
	"github.com/quickpowered/frilly/pkg/consts"
)

func ToForward(bot bin.Interface, payload *domain.Payload) any {
	if platform := bot.GetPlatform(); platform == consts.PlatformDiscord {
		return types.NewForward(payload.Message.Chat().GetThreadID(), payload.Message.ID())
	}

	return types.NewForward(payload.Message.Chat().GetID(), payload.Message.ID())
}
