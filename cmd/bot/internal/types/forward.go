package types

type Forward struct {
	ChatID    any
	MessageID any
}

func NewForward(chatID, messageID any) *Forward {
	return &Forward{chatID, messageID}
}
