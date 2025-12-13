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

const HOWTOUSE_CMD = "howtouse"

type howToUseHandler struct {
	deps.Dependencies
}

func NewHowToUseHandler(deps deps.Dependencies) *howToUseHandler {
	return &howToUseHandler{deps}
}

func (c *howToUseHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	msg := i18n.HowToUseMessages[language.Language(p.Account.Language)]

	file, err := os.ReadFile(os.Getenv("HOWTOUSE_IMAGE_PATH"))
	if err != nil {
		return err
	}

	opts := []any{deps.ToForward(bot, p), types.NewAttachments().AddFile(bytes.NewReader(file))}
	return bot.SendMessage(p.Message.Chat(), msg.Instruction, opts...)
}
