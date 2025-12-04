package telegram

import (
	"context"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/types"
	"github.com/quickpowered/frilly/internal/types/update"
	"github.com/quickpowered/frilly/pkg/consts"
)

type Adapter struct {
	bot *bot.Bot
}

func NewBot(bot *bot.Bot) bin.Interface {
	return &Adapter{bot}
}

func (a *Adapter) GetSelfID() int {
	return int(a.bot.ID())
}

func (a *Adapter) GetAPI() any {
	return a.bot
}

func (a *Adapter) GetPlatform() consts.Platform {
	return consts.PlatformTelegram
}

func (a *Adapter) SendMessage(chatID update.ChatInterface, text string, opts ...any) (err error) {
	var photo string

	for _, param := range opts {
		switch val := param.(type) {
		case *types.Attachments:
			if attachments := val; attachments != nil {
				for _, attachment := range *attachments {
					photo = attachment
					break
				}
			}
		}
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	if photo != "" {
		_, err = a.bot.SendPhoto(ctx, &bot.SendPhotoParams{
			ChatID: chatID.GetID(),
			Photo:  &models.InputFileString{photo},
		})
	} else {
		sendMessage := &bot.SendMessageParams{
			ChatID:          chatID.GetID(),
			MessageThreadID: chatID.GetThreadID(),
			Text:            text,
			ParseMode:       models.ParseModeHTML,
		}

		for _, param := range opts {
			switch val := param.(type) {
			case *types.Keyboard:
				inlineKeyboardButtons := [][]models.InlineKeyboardButton{}
				if keyboard := val; keyboard != nil {
					for _, buttonRow := range keyboard.ButtonRows {
						inlineKeyboardButtons = append(inlineKeyboardButtons, []models.InlineKeyboardButton{})
						for _, button := range buttonRow {
							inlineKeyboardButtons[len(inlineKeyboardButtons)-1] = append(inlineKeyboardButtons[len(inlineKeyboardButtons)-1], models.InlineKeyboardButton{
								Text:         button.Text,
								CallbackData: button.Data,
							})
						}
					}

					sendMessage.ReplyMarkup = &models.InlineKeyboardMarkup{
						InlineKeyboard: inlineKeyboardButtons,
					}
				}
			}
		}

		_, err = a.bot.SendMessage(ctx, sendMessage)
	}

	return err
}

func (a *Adapter) SendInvoice(chatID update.ChatInterface, title, description, payload, currency string, amount int) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	_, err := a.bot.SendInvoice(ctx, &bot.SendInvoiceParams{
		ChatID:          chatID.GetID(),
		MessageThreadID: chatID.GetThreadID(),
		Title:           title,
		Description:     description,
		Payload:         payload,
		Currency:        currency,
		Prices: []models.LabeledPrice{
			{
				Label:  "Stars",
				Amount: amount,
			},
		},
	})

	return err
}

func (a *Adapter) DeleteMessage(chatID update.ChatInterface, messageID int) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	_, err := a.bot.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    chatID.GetID(),
		MessageID: messageID,
	})
	return err
}
