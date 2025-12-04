package telegram

import (
	"github.com/go-telegram/bot/models"
	"github.com/quickpowered/frilly/internal/types"
)

func inlineKeyboardMarkup(keyboard *types.Keyboard) *models.InlineKeyboardMarkup {
	inlineKeyboardButtons := [][]models.InlineKeyboardButton{}
	if keyboard != nil {
		for _, buttonRow := range keyboard.ButtonRows {
			inlineKeyboardButtons = append(inlineKeyboardButtons, []models.InlineKeyboardButton{})
			for _, button := range buttonRow {
				inlineKeyboardButtons[len(inlineKeyboardButtons)-1] = append(inlineKeyboardButtons[len(inlineKeyboardButtons)-1], models.InlineKeyboardButton{
					Text:         button.Text,
					CallbackData: button.Data,
				})
			}
		}
	}

	return &models.InlineKeyboardMarkup{
		InlineKeyboard: inlineKeyboardButtons,
	}
}
