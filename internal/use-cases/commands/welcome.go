package commands

import (
	"bytes"
	"os"

	"github.com/quickpowered/frilly/config/language"
	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/i18n"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/types"
	"github.com/quickpowered/frilly/internal/use-cases/commands/deps"
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

	file, err := os.ReadFile(os.Getenv("WELCOME_PREVIEW_PATH"))
	if err != nil {
		return err
	}

	opts := []any{
		deps.ToForward(bot, p),
		types.NewAttachments().AddFile(bytes.NewReader(file)),
	}
	return bot.SendMessage(p.Message.Chat(), msg.Welcome, opts...)
}
