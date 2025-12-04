package discord

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/types"
	"github.com/quickpowered/frilly/internal/types/update"
	"github.com/quickpowered/frilly/pkg/consts"
)

type adapter struct {
	bot *discordgo.Session
}

func NewBot(bot *discordgo.Session) bin.Interface {
	return &adapter{bot}
}

func (a *adapter) GetSelfID() int {
	selfBotID, _ := strconv.Atoi(a.bot.State.User.ID)
	return selfBotID
}

func (a *adapter) GetAPI() any {
	return a.bot
}

func (a *adapter) GetPlatform() consts.Platform {
	return consts.PlatformDiscord
}

func (a *adapter) SendMessage(chatID update.ChatInterface, text string, opts ...any) error {
	sendMessage := &discordgo.MessageSend{Content: text}

	for _, param := range opts {
		switch val := param.(type) {
		case types.Mentions:
			if !bool(val) {
				sendMessage.Flags = discordgo.MessageFlagsSuppressNotifications
				sendMessage.AllowedMentions = &discordgo.MessageAllowedMentions{
					Users: []string{},
				}
			}
		case *types.RichEmbed:
			if embed := val; embed != nil {
				richEmbed := &discordgo.MessageEmbed{
					Type:        discordgo.EmbedTypeRich,
					Title:       embed.Title,
					Color:       embed.Color,
					Description: embed.Description,
				}

				if embed.Author != nil {
					richEmbed.Author = &discordgo.MessageEmbedAuthor{
						Name:    embed.Author.Name,
						IconURL: embed.Author.Icon,
					}
				}
				sendMessage.Embeds = append(sendMessage.Embeds, richEmbed)
			}
		case *types.Keyboard:
			if keyboard := val; keyboard != nil {
				components := []discordgo.MessageComponent{}

				for _, buttonRow := range keyboard.ButtonRows {
					actionRows := discordgo.ActionsRow{}

					for _, button := range buttonRow {
						actionRows.Components = append(actionRows.Components, discordgo.Button{
							Label:    button.Text,
							Style:    discordgo.SecondaryButton,
							CustomID: button.Data,
						})
					}

					components = append(components, actionRows)
				}

				sendMessage.Components = components
			}
		case *types.Forward:
			if reference := val; reference != nil {
				sendMessage.Reference = &discordgo.MessageReference{
					MessageID: fmt.Sprint(reference.MessageID),
					ChannelID: fmt.Sprint(reference.ChatID),
				}
			}
		}
	}

	_, err := a.bot.ChannelMessageSendComplex(fmt.Sprint(chatID.GetThreadID()), sendMessage)
	return err
}

func (a *adapter) DeleteMessage(chatID update.ChatInterface, messageID int) error {
	return a.bot.ChannelMessageDelete(fmt.Sprint(chatID.GetID()), fmt.Sprint(messageID))
}
