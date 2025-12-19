package vk

import (
	"github.com/aejoy/vkgo/api"
	vkUpdate "github.com/aejoy/vkgo/update"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types/update"
	"github.com/quickpowered/thebeyond-master/pkg/consts"
	"go.uber.org/zap"
)

func (s *delivery) onCallback(bot *api.API, callback vkUpdate.Callback) {
	command, ok := callback.Payload["command"]
	if ok {
		if err := s.commands.Run(s.bot, &domain.Payload{
			Message: update.NewMessage(
				consts.PlatformVK,
				callback.ChatMessageID,
				update.Chat{ID: callback.ChatID},
				int(callback.UserID),
				"",
				command,
				nil,
				bot,
			),
		}); err != nil {
			s.Logger.Error("command failed", zap.Error(err))
		}
	}
}
