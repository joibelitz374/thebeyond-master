package discord

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/types/update"
	"github.com/quickpowered/frilly/internal/use-cases/commands"
	"github.com/quickpowered/frilly/internal/use-cases/commands/tools"
	"github.com/quickpowered/frilly/pkg/consts"
	"go.uber.org/zap"
)

type delivery struct {
	bot      bin.Interface
	commands commands.Commands
	Logger   *zap.Logger
}

func New(bot bin.Interface, commands commands.Commands, logger *zap.Logger) *delivery {
	return &delivery{bot, commands, logger}
}

func (s *delivery) Listen() error {
	dg := s.bot.GetAPI().(*discordgo.Session)

	dg.AddHandler(s.handler)

	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentMessageContent

	err := dg.Open()
	if err != nil {
		return fmt.Errorf("error opening connection, %w", err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()

	return nil
}

func (s *delivery) handler(session *discordgo.Session, message *discordgo.MessageCreate) {
	return

	messageID, _ := strconv.Atoi(message.ID)
	guildID, _ := strconv.Atoi(message.GuildID)
	channelID, _ := strconv.Atoi(message.ChannelID)
	senderID, _ := strconv.Atoi(message.Author.ID)

	var replyMessage *update.Message
	if reference := message.Reference(); reference != nil {
		originalMsg, err := session.ChannelMessage(message.ChannelID, reference.MessageID)
		if err != nil {
			return
		}

		refSenderID, _ := strconv.Atoi(originalMsg.Author.ID)

		replyMessage = &update.Message{
			ID: messageID,
			Chat: update.Chat{
				ID:       guildID,
				ThreadID: channelID,
			},
			SenderID: refSenderID,
			Text:     originalMsg.Content,
		}
	}

	err := s.commands.Run(s.bot, &domain.Payload{
		Message: update.NewMessage(
			consts.PlatformTelegram,
			messageID,
			update.Chat{
				ID:       guildID,
				ThreadID: channelID,
			},
			senderID,
			message.Content,
			replyMessage,
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
