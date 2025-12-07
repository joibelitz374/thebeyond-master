package update

import (
	"github.com/aejoy/vkgo/api"
	"github.com/quickpowered/frilly/pkg/consts"
)

type NewMessageInterface interface {
	ID() int
	Chat() ChatInterface
	Sender() int
	Text() string
	Reply() *Message
}

type Message struct {
	platform  consts.Platform
	vkBot     *api.API
	messageID int
	chat      ChatInterface
	senderID  int
	text      string
	reply     *Message
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
		platform:  platform,
		messageID: id,
		chat:      chat,
		senderID:  senderID,
		text:      text,
		reply:     reply,
	}

	for _, bot := range bots {
		switch bot.(type) {
		case *api.API:
			message.vkBot = bot.(*api.API)
		}
	}

	return message
}

func (m Message) ID() int {
	return m.messageID
}

func (m Message) Chat() ChatInterface {
	return m.chat
}

func (m Message) Sender() int {
	return m.senderID
}

func (m Message) Text() string {
	return m.text
}

func (m Message) Reply() *Message {
	return m.reply
}
