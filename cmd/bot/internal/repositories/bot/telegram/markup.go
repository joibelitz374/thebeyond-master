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
				switch {
				case button.Data != "":
					inlineKeyboardButtons[len(inlineKeyboardButtons)-1] = append(inlineKeyboardButtons[len(inlineKeyboardButtons)-1], models.InlineKeyboardButton{
						Text:         button.Text,
						CallbackData: button.Data,
					})
				case button.URL != "":
					inlineKeyboardButtons[len(inlineKeyboardButtons)-1] = append(inlineKeyboardButtons[len(inlineKeyboardButtons)-1], models.InlineKeyboardButton{
						Text: button.Text,
						URL:  button.URL,
					})
				}
			}
		}
	}

	return &models.InlineKeyboardMarkup{
		InlineKeyboard: inlineKeyboardButtons,
	}
}
