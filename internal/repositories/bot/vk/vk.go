package vk

import (
	"fmt"

	"github.com/aejoy/vkgo/api"
	"github.com/aejoy/vkgo/keyboard"
	vkTypes "github.com/aejoy/vkgo/types"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/types"
	"github.com/quickpowered/frilly/internal/types/update"
	"github.com/quickpowered/frilly/pkg/consts"
)

type adapter struct {
	bot *api.API
}

func NewBot(bot *api.API) bin.Interface {
	return &adapter{bot}
}

func (a *adapter) GetSelfID() int {
	return a.bot.ID
}

func (a *adapter) GetAPI() any {
	return a.bot
}

func (a *adapter) GetPlatform() consts.Platform {
	return consts.PlatformVK
}

func (a *adapter) SendMessage(chatID update.ChatInterface, text string, opts ...any) error {
	sendMessage := vkTypes.SendMessage{ChatID: chatID.GetID(), Text: text}

	for _, param := range opts {
		switch val := param.(type) {
		case types.Mentions:
			sendMessage.DisableMentions = bool(val)
		case *types.Attachments:
			if attachments := val; attachments != nil {
				for _, url := range attachments.Urls {
					sendMessage.Attachments = append(sendMessage.Attachments, url)
				}
			}
		case *types.Keyboard:
			if val := val; val != nil {
				kb := keyboard.New().Inline()
				for _, buttonRow := range val.ButtonRows {
					row := kb.Row()
					for _, button := range buttonRow {
						row.Text(button.Text, fmt.Sprintf("{\"command\": \"%s\"}", button.Data), keyboard.White)
					}
				}
				sendMessage.Keyboard = kb.JSON()
			}
		case *types.Forward:
			if forward := val; forward != nil {
				fwdMessage, err := vkTypes.NewForward(forward.ChatID.(int), forward.MessageID.(int), true)
				if err != nil {
					continue
				}
				sendMessage.Forward = fwdMessage
			}
		}
	}

	_, err := a.bot.SendMessage(sendMessage)
	return err
}

func (a *adapter) DeleteMessage(chatID update.ChatInterface, messageID int) error {
	_, err := a.bot.DeleteMessage(chatID.GetID(), messageID, 1)
	return err
}
