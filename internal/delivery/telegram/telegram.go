package telegram

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	sharedUpdate "github.com/quickpowered/frilly/internal/types/update"
	"github.com/quickpowered/frilly/internal/use-cases/commands"
	"github.com/quickpowered/frilly/internal/use-cases/invoices"
	"github.com/quickpowered/frilly/pkg/consts"
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
	var typ models.ChatType
	var chatID, threadID, userID int
	var text string

	switch {
	case update.PreCheckoutQuery != nil:
		typ = models.ChatTypePrivate
		if err := s.invoices.HandlePreCheckoutQuery(tgBot, update); err != nil {
			s.Logger.Error("pre checkout query failed", zap.Error(err))
		}
		return
	case update.Message != nil:
		typ = update.Message.Chat.Type
		userID = int(update.Message.From.ID)
		chatID = int(update.Message.Chat.ID)
		threadID = int(update.Message.MessageThreadID)
		text = update.Message.Text
	case update.CallbackQuery != nil:
		message := update.CallbackQuery.Message.Message
		if message == nil {
			return
		}
		typ = message.Chat.Type
		userID = int(update.CallbackQuery.From.ID)
		chatID = int(message.Chat.ID)
		threadID = int(message.MessageThreadID)
		text = update.CallbackQuery.Data
	default:
		return
	}

	if typ != models.ChatTypePrivate {
		s.Logger.Debug("not private chat",
			zap.String("type", string(typ)),
			zap.Int("chat_id", chatID),
			zap.Int("thread_id", threadID),
			zap.Int("user_id", userID))
		return
	}

	if err := s.commands.Run(s.bot, &domain.Payload{
		Message: sharedUpdate.NewMessage(
			consts.PlatformTelegram,
			int(update.ID),
			sharedUpdate.Chat{
				ID:       int(chatID),
				ThreadID: int(threadID),
			},
			userID,
			text,
			nil,
		),
	}); err != nil {
		s.Logger.Error("command failed", zap.Error(err),
			zap.Int("chat_id", chatID),
			zap.Int("user_id", userID),
			zap.String("text", text))
	}
}
