package telegram

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/types/update"
	"github.com/quickpowered/frilly/internal/use-cases/commands"
	"github.com/quickpowered/frilly/internal/use-cases/commands/tools"
	"github.com/quickpowered/frilly/internal/use-cases/invoices"
	"github.com/quickpowered/frilly/pkg/consts"
	"go.uber.org/zap"
)

type delivery struct {
	bot      bin.Interface
	commands commands.Commands
	invoices invoices.Interface
	Logger   *zap.Logger
}

func New(bot bin.Interface, commands commands.Commands, invoices invoices.Interface, logger *zap.Logger) *delivery {
	return &delivery{bot, commands, invoices, logger}
}

func (s *delivery) Listen() error {
	ctx, cancel := signal.NotifyContext(context.TODO(), os.Interrupt)
	defer cancel()

	matchFunc := func(update *models.Update) bool {
		return true
	}

	b := s.bot.GetAPI().(*bot.Bot)
	b.RegisterHandlerMatchFunc(matchFunc, s.handler)
	b.Start(ctx)

	return nil
}

func (s *delivery) handler(ctx context.Context, tgBot *bot.Bot, tgUpdate *models.Update) {
	var userID, chatID, threadID int
	var text string

	switch {
	case tgUpdate.PreCheckoutQuery != nil:
		ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
		defer cancel()

		if err := s.invoices.HandlePreCheckoutQuery(ctx, tgBot, tgUpdate); err != nil {
			s.Logger.Error("pre checkout query failed", zap.Error(err))
		}
		return
	case tgUpdate.Message != nil:
		userID = int(tgUpdate.Message.From.ID)
		chatID = int(tgUpdate.Message.Chat.ID)
		threadID = int(tgUpdate.Message.MessageThreadID)
		text = tgUpdate.Message.Text
	case tgUpdate.CallbackQuery != nil:
		userID = int(tgUpdate.CallbackQuery.From.ID)
		chatID = int(tgUpdate.CallbackQuery.Message.Message.Chat.ID)
		threadID = int(tgUpdate.CallbackQuery.Message.Message.MessageThreadID)
		text = tgUpdate.CallbackQuery.Data
	default:
		return
	}

	err := s.commands.Run(s.bot, &domain.Payload{
		Message: update.NewMessage(
			consts.PlatformTelegram,
			int(tgUpdate.ID),
			update.Chat{ID: int(chatID), ThreadID: int(threadID)},
			userID,
			text,
			nil,
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
