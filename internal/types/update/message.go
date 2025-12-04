package update

import (
	"github.com/aejoy/vkgo/api"
	"github.com/quickpowered/frilly/pkg/consts"
)

type NewMessageInterface interface {
	GetID() int
	GetChat() ChatInterface
	GetSender() int
	GetText() string
	GetReply() *Message
}

type Message struct {
	platform consts.Platform
	vkBot    *api.API
	ID       int
	Chat     ChatInterface
	SenderID int
	Text     string
	Reply    *Message
}

func NewMessage(
	platform consts.Platform,
	id int,
	chat ChatInterface,
	senderID int,
	text string,
	reply *Message,
	bots ...any,
) Message {
	message := Message{
		platform: platform,
		ID:       id,
		Chat:     chat,
		SenderID: senderID,
		Text:     text,
		Reply:    reply,
	}

	for _, bot := range bots {
		switch bot.(type) {
		case *api.API:
			message.vkBot = bot.(*api.API)
		}
	}

	return message
}

func (m Message) GetID() int {
	return m.ID
}

func (m Message) GetChat() ChatInterface {
	return m.Chat
}

func (m Message) GetSender() int {
	return m.SenderID
}

func (m Message) GetText() string {
	return m.Text
}

func (m Message) GetReply() *Message {
	return m.Reply
}
