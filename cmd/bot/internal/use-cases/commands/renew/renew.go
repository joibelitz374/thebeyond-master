package renew

import (
	"strconv"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
)

const CMD = "renew"

type handler struct {
	deps.Dependencies
}

func NewHandler(deps deps.Dependencies) handler {
	return handler{deps}
}

func (h handler) Execute(bot bin.Interface, p *domain.Payload) error {
	msg := i18n.RenewMessages[p.Account.Language]
	switch len(p.Args) {
	case 2:
		return h.firstStep(bot, p, msg)
	case 3:
		return h.purchaseOffer(bot, p, msg)
	}
	return h.showTariffs(bot, p)
}

func (h handler) firstStep(bot bin.Interface, p *domain.Payload, msg i18n.RenewLocale) error {
	controlMsg := i18n.ControlMessages[p.Account.Language]

	if p.Args[1] == "check-subscription" {
		return h.checkSubscription(bot, p)
	}

	tariffID, err := strconv.Atoi(p.Args[1])
	if err != nil {
		return bot.SendMessage(p.Message.Chat(), "Invalid traffic ID", types.NewKeyboard().
			NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, CMD)))
	}

	tariff, ok := h.TariffsRepo.Get(tariffID)
	if !ok {
		return bot.SendMessage(p.Message.Chat(), "No traffic found", types.NewKeyboard().
			NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, CMD)))
	}

	switch tariffID {
	case 0:
		return h.checkSubscription(bot, p)
	default:
		return h.showPeriods(bot, p, tariff, tariffID, msg)
	}
}
func (h handler) getTariffPeriodName(msg i18n.RenewLocale, days int) (title string) {
	switch days {
	case 7:
		title = msg.Week
	case 14:
		title = msg.TwoWeeks
	case 30:
		title = msg.Month
	case 90:
		title = msg.Season
	case 180:
		title = msg.HalfYear
	case 365:
		title = msg.Year
	}
	return title
}
