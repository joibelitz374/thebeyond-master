package telegram

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	sharedUpdate "github.com/quickpowered/thebeyond-master/cmd/bot/internal/types/update"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/invoices"
	"github.com/quickpowered/thebeyond-master/pkg/consts"
	"go.uber.org/zap"
)

type delivery struct {
	bot      bin.Interface
	commands commands.UseCase
	invoices invoices.UseCase
	Logger   *zap.Logger
}

func New(bot bin.Interface, commands commands.UseCase, invoices invoices.UseCase, logger *zap.Logger) *delivery {
	return &delivery{bot, commands, invoices, logger}
}

func (s *delivery) Listen() error {
	ctx, cancel := signal.NotifyContext(context.TODO(), os.Interrupt)
	defer cancel()

	matchFunc := func(update *models.Update) bool {
		s.Logger.Debug("new update", zap.Any("update", update))
		return true
	}

	b := s.bot.GetAPI().(*bot.Bot)
	b.RegisterHandlerMatchFunc(matchFunc, s.handler)
	b.Start(ctx)

	return nil
}

func (s *delivery) handler(ctx context.Context, tgBot *bot.Bot, update *models.Update) {
	parsed, ok := s.parseUpdate(ctx, tgBot, update)
	if !ok {
		return
	}

	if parsed.typ != models.ChatTypePrivate {
		s.Logger.Debug("not private chat",
			zap.String("type", string(parsed.typ)),
			zap.Int("chat_id", parsed.chatID),
			zap.Int("thread_id", parsed.threadID),
			zap.Int("user_id", parsed.userID))
		return
	}

	if err := s.commands.Run(s.bot, &domain.Payload{
		Message: sharedUpdate.NewMessage(
			consts.PlatformTelegram,
			int(parsed.messageID),
			sharedUpdate.Chat{
				ID:       int(parsed.chatID),
				ThreadID: int(parsed.threadID),
			},
			parsed.userID,
			parsed.callbackQueryID,
			parsed.text,
			nil,
		),
	}); err != nil {
		s.Logger.Error("command failed", zap.Error(err),
			zap.Int("chat_id", parsed.chatID),
			zap.Int("user_id", parsed.userID),
			zap.String("text", parsed.text))
	}
}
