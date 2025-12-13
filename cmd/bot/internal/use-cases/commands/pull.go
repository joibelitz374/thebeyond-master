package commands

import (
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
)

const PULL_CMD = "pull"

type pullHandler struct {
	deps.Dependencies
}

func NewPullHandler(deps deps.Dependencies) pullHandler {
	return pullHandler{deps}
}

func (h pullHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	return bot.SendMessage(p.Message.Chat(), "<b>Что ещё за крутки?</b>\nЗа КАЖДУЮ подписку Вы получаете +1 крутку, которую можно запустить и выиграть следующие призы:\n• от 3 до 12 дней подписки\n• скидка продления от 13 до 32%")
}
