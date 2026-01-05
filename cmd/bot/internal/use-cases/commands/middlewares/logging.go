package middlewares

import (
	"strings"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/interfaces"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"go.uber.org/zap"
)

func WithLogging(inner interfaces.Command, logger *zap.Logger, cmdName string) interfaces.Command {
	return &logging{inner, logger, cmdName}
}

type logging struct {
	inner   interfaces.Command
	logger  *zap.Logger
	cmdName string
}

func (m *logging) Execute(bot bin.Interface, p *domain.Payload) (err error) {
	sender := p.Message.Sender()
	msgID := p.Message.ID()
	args := strings.Join(p.Args[1:], " ")

	defer func() {
		if err != nil {
			m.logger.Error("command failed",
				zap.String("command", m.cmdName),
				zap.Int("sender", sender),
				zap.Int("message_id", msgID),
				zap.String("args", args),
				zap.Error(err))
		} else {
			m.logger.Info("command executed",
				zap.String("command", m.cmdName),
				zap.Int("sender", sender),
				zap.Int("message_id", msgID),
				zap.String("args", args))
		}
	}()

	return m.inner.Execute(bot, p)
}
