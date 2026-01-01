package commands

import (
	"fmt"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types/update"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	"github.com/quickpowered/thebeyond-master/configs/language"
)

const SUPPORT_CMD = "support"

type supportHandler struct {
	deps.Dependencies
}

func NewSupportHandler(deps deps.Dependencies) *supportHandler {
	return &supportHandler{deps}
}

func (c *supportHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	msg := i18n.SupportMessages[language.Language(p.Account.Language)]
	if len(p.Args) >= 2 && p.Args[1] == "ok" {
		if err := bot.SendMessage(update.Chat{ID: 924536264}, fmt.Sprintf("У <a href=\"tg://user?id=%d\">пользователя</a> всё отлично!", p.Message.Sender())); err != nil {
			return err
		}
		return bot.SendMessage(p.Message.Chat(), "Thank you for your feedback! ✌️😊")
	}

	return bot.SendMessage(p.Message.Chat(), msg.NeedHelp+" "+msg.WriteToUs+"\n\n"+msg.PersonalMessages+"\n"+msg.Chat+"\n"+msg.Creator)
}
