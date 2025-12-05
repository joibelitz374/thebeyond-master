package commands

import (
	"github.com/quickpowered/frilly/config/language"
	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/i18n"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/types"
	"github.com/quickpowered/frilly/internal/use-cases/commands/deps"
)

const WELCOME_CMD = "welcome"

type welcomeHandler struct {
	deps.Dependencies
}

func NewWelcomeHandler(deps deps.Dependencies) *welcomeHandler {
	return &welcomeHandler{deps}
}

func (c *welcomeHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	msg := i18n.WelcomeMessages[language.Language(p.Account.Language)]
	opts := []any{
		deps.ToForward(bot, p),
		types.NewAttachments().AddURL("https://i.ibb.co/Y7RpD8rc/image-40.png"),
	}
	return bot.SendMessage(p.Message.Chat(), msg.Welcome, opts...)
}
