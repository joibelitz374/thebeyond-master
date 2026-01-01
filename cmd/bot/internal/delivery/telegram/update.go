package telegram

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"go.uber.org/zap"
)

type parsedUpdate struct {
	typ             models.ChatType
	userID          int
	chatID          int
	threadID        int
	messageID       int
	callbackQueryID string
	text            string
}

func (s *delivery) parseUpdate(ctx context.Context, tgBot *bot.Bot, update *models.Update) (*parsedUpdate, bool) {
	parsed := &parsedUpdate{}

	switch {
	case update.PreCheckoutQuery != nil:
		parsed.typ = models.ChatTypePrivate
		if err := s.invoices.HandlePreCheckoutQuery(tgBot, update); err != nil {
			s.Logger.Error("pre checkout query failed", zap.Error(err))
		}
		return nil, false
	case update.Message != nil:
		parsed.typ = update.Message.Chat.Type
		parsed.userID = int(update.Message.From.ID)
		parsed.chatID = int(update.Message.Chat.ID)
		parsed.threadID = int(update.Message.MessageThreadID)
		parsed.messageID = update.Message.ID
		parsed.text = update.Message.Text
	case update.CallbackQuery != nil:
		message := update.CallbackQuery.Message.Message
		if message == nil {
			return nil, false
		}
		parsed.typ = message.Chat.Type
		parsed.userID = int(update.CallbackQuery.From.ID)
		parsed.chatID = int(message.Chat.ID)
		parsed.threadID = int(message.MessageThreadID)
		parsed.messageID = message.ID
		parsed.callbackQueryID = update.CallbackQuery.ID
		parsed.text = update.CallbackQuery.Data
	default:
		return nil, false
	}

	return parsed, true
}
