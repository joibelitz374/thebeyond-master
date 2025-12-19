package vk

import (
	"github.com/aejoy/vkgo/api"
	vkUpdate "github.com/aejoy/vkgo/update"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types/update"
	"github.com/quickpowered/thebeyond-master/pkg/consts"
	"go.uber.org/zap"
)

func (s *delivery) onMessage(bot *api.API, message vkUpdate.Message) {
	text := message.Text
	if message.Payload != nil {
		if command, ok := message.Payload["command"]; ok {
			text = command
		}
	}

	if err := s.commands.Run(s.bot, &domain.Payload{
		Message: update.NewMessage(
			consts.PlatformVK,
			message.ChatMessageID,
			update.Chat{ID: message.ChatID},
			int(message.UserID),
			"",
			text,
			nil,
			bot,
		),
	}); err != nil {
		s.Logger.Error("command failed", zap.Error(err))
	}
}
