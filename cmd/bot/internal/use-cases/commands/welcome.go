package commands

import (
	"bytes"
	"os"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	"github.com/quickpowered/thebeyond-master/configs/language"
)

const WELCOME_CMD = "welcome2welcome"

type welcomeHandler struct {
	deps.Dependencies
}

func NewWelcomeHandler(deps deps.Dependencies) *welcomeHandler {
	return &welcomeHandler{deps}
}

func (c *welcomeHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	msg := i18n.WelcomeMessages[language.Language(p.Account.Language)]

	file, err := os.ReadFile(os.Getenv("WELCOME_IMAGE_PATH"))
	if err != nil {
		return err
	}

	opts := []any{
		deps.ToForward(bot, p),
		types.NewAttachments().AddFile(bytes.NewReader(file)),
	}
	return bot.SendMessage(p.Message.Chat(), msg.Welcome, opts...)
}
