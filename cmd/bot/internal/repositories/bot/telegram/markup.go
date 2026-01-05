package telegram

import (
	"github.com/go-telegram/bot/models"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
)

func inlineKeyboardMarkup(keyboard *types.Keyboard) *models.InlineKeyboardMarkup {
	inlineKeyboardButtons := [][]models.InlineKeyboardButton{}
	if keyboard != nil {
		for _, buttonRow := range keyboard.ButtonRows {
			inlineKeyboardButtons = append(inlineKeyboardButtons, []models.InlineKeyboardButton{})
			for _, button := range buttonRow {
				inlineButton := models.InlineKeyboardButton{}
				switch {
				case button.Pay:
					inlineButton = models.InlineKeyboardButton{
						Text: button.Text,
						Pay:  true,
					}
				case button.Data != "":
					inlineButton = models.InlineKeyboardButton{
						Text:         button.Text,
						CallbackData: button.Data,
					}
				case button.URL != "":
					inlineButton = models.InlineKeyboardButton{
						Text: button.Text,
						URL:  button.URL,
					}
				}
				buttonsLength := len(inlineKeyboardButtons) - 1
				inlineKeyboardButtons[buttonsLength] = append(inlineKeyboardButtons[buttonsLength], inlineButton)
			}
		}
	}

	return &models.InlineKeyboardMarkup{
		InlineKeyboard: inlineKeyboardButtons,
	}
}
