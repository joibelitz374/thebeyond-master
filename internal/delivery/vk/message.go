package vk

import (
	"github.com/aejoy/vkgo/api"
	vkUpdate "github.com/aejoy/vkgo/update"
	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/types/update"
	"github.com/quickpowered/frilly/internal/use-cases/commands/tools"
	"github.com/quickpowered/frilly/pkg/consts"
	"go.uber.org/zap"
)

func (s *delivery) onMessage(bot *api.API, message vkUpdate.Message) {
	text := message.Text
	if message.Payload != nil {
		if command, ok := message.Payload["command"]; ok {
			text = command
		}
	}

	err := s.commands.Run(s.bot, &domain.Payload{
		Message: update.NewMessage(
			consts.PlatformVK,
			message.ChatMessageID,
			update.Chat{ID: message.ChatID},
			int(message.UserID),
			text,
			nil,
			bot,
		),
	})
	if err != nil {
		if cerr, ok := err.(tools.CommandError); ok {
			s.Logger.Error("command failed",
				zap.String("cmd", cerr.Command),
				zap.String("service", cerr.Service),
				zap.String("method", cerr.Method),
				zap.String("chat_id", cerr.ChatID),
				zap.Int("member_id", cerr.MemberID),
				zap.Error(err),
			)
		} else {
			s.Logger.Error("command failed", zap.Error(err))
		}
	}
}
