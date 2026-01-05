package renew

import (
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	"github.com/quickpowered/thebeyond-master/pkg/utils"
)

const CMD = "renew"

type handler struct {
	deps.Dependencies
}

func NewHandler(deps deps.Dependencies) handler {
	return handler{deps}
}

func (h handler) Execute(bot bin.Interface, p *domain.Payload) (err error) {
	if len(p.Args) < 2 || p.Args[1] == "" {
		p.Args = append(p.Args, "t:0")
	}

	if p.Args[1] == "check-subscription" {
		return h.checkSubscription(bot, p)
	}

	controlMsg := i18n.ControlMessages[p.Account.Language]

	tariffID, days, err := utils.GetTariffPayloadData(p.Args[1])
	if err != nil {
		return bot.SendMessage(p.Message.Chat(), "Invalid payload", types.NewKeyboard().
			NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, CMD)))
	}

	tariff, ok := h.TariffsRepo.Get(tariffID)
	if !ok {
		return bot.SendMessage(p.Message.Chat(), "No tariff found", types.NewKeyboard().
			NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, CMD)))
	}

	if tariffID != 0 && days != 0 {
		if days == -1 {
			return h.showPeriods(bot, p, tariffID, tariff)
		}
		return h.purchaseOffer(bot, p, tariffID, tariff, days)
	}

	return h.showTariffs(bot, p, tariffID, tariff)
}
