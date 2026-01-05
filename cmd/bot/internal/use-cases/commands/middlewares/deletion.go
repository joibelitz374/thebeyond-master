package middlewares

import (
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/interfaces"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"go.uber.org/zap"
)

func WithDeletion(inner interfaces.Command, logger *zap.Logger) interfaces.Command {
	return &deletion{inner, logger}
}

type deletion struct {
	inner  interfaces.Command
	logger *zap.Logger
}

func (m *deletion) Execute(bot bin.Interface, p *domain.Payload) error {
	err := m.inner.Execute(bot, p)

	if delErr := bot.DeleteMessage(p.Message.Chat(), p.Message.ID()); delErr != nil {
		m.logger.Error("failed to delete message",
			zap.Int("sender", p.Message.Sender()),
			zap.Int("message_id", p.Message.ID()),
			zap.Error(delErr))
	}

	return err
}
