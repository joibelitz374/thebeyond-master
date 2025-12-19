package telegram

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types/update"
	"github.com/quickpowered/thebeyond-master/pkg/consts"
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
	var photoURL string
	var photoData io.Reader

	for _, param := range opts {
		switch val := param.(type) {
		case *types.Attachments:
			if attachments := val; attachments != nil {
				if len(attachments.Urls) > 0 {
					for _, attachment := range attachments.Urls {
						photoURL = attachment
						break
					}
				} else if len(attachments.Files) > 0 {
					for _, attachment := range attachments.Files {
						photoData = attachment
						break
					}
				}
			}
		}
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	var isPhoto bool
	if photoURL != "" || photoData != nil {
		isPhoto = true
	}

	if isPhoto {
		sendPhotoParams := &bot.SendPhotoParams{
			ChatID:    chatID.GetID(),
			Caption:   text,
			ParseMode: models.ParseModeHTML,
		}

		if photoURL != "" {
			sendPhotoParams.Photo = &models.InputFileString{Data: photoURL}
		} else if photoData != nil {
			sendPhotoParams.Photo = &models.InputFileUpload{Filename: "photo.png", Data: photoData}
		}

		for _, param := range opts {
			switch val := param.(type) {
			case *types.Keyboard:
				sendPhotoParams.ReplyMarkup = inlineKeyboardMarkup(val)
			}
		}

		_, err = a.bot.SendPhoto(ctx, sendPhotoParams)
		return err
	}

	sendMessageParams := &bot.SendMessageParams{
		ChatID:          chatID.GetID(),
		MessageThreadID: chatID.GetThreadID(),
		Text:            text,
		ParseMode:       models.ParseModeHTML,
	}

	for _, param := range opts {
		switch val := param.(type) {
		case *types.Keyboard:
			sendMessageParams.ReplyMarkup = inlineKeyboardMarkup(val)
		}
	}

	_, err = a.bot.SendMessage(ctx, sendMessageParams)

	return err
}

func (a *Adapter) ForwardMessage(chatID update.ChatInterface, fromChatID, fromMessageID int) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	fmt.Println(chatID.GetID(), fromChatID, fromMessageID)
	_, err := a.bot.ForwardMessage(ctx, &bot.ForwardMessageParams{
		ChatID:              chatID.GetID(),
		FromChatID:          fromChatID,
		MessageID:           fromMessageID,
		DisableNotification: true,
	})
	return err
}

func (a *Adapter) AnswerCallbackQuery(callbackQueryID, text string) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	_, err := a.bot.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: callbackQueryID,
		Text:            text,
	})
	return err
}

func (a *Adapter) SendInvoice(chatID update.ChatInterface, title, description, payload, currency string, amount int, opts ...any) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	invoiceParams := &bot.SendInvoiceParams{
		ChatID:          chatID.GetID(),
		MessageThreadID: chatID.GetThreadID(),
		Title:           title,
		Description:     description,
		Payload:         payload,
		Currency:        currency,
		Prices: []models.LabeledPrice{
			{Label: "Stars", Amount: amount},
		},
	}

	_, err := a.bot.SendInvoice(ctx, invoiceParams)
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
