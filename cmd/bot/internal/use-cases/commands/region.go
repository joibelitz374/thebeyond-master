package commands

import (
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
)

const REGION_CMD = "region"

var regions = map[string][2]string{
	"us": {"🇺🇸", "США"},
	"eu": {"🇪🇺", "ЕС"},
	"ru": {"🇷🇺", "Россия"},
	"ir": {"🇮🇷", "Иран"},
	"tr": {"🇹🇷", "Турция"},
	"cn": {"🇨🇳", "Китай"},
	"in": {"🇮🇳", "Индия"},
}

var regionsOrder = [][]string{
	{"cn", "ru"},
	{"ir"},
	{"eu", "us"},
	{"in", "tr"},
}

type regionHandler struct {
	deps.Dependencies
}

func NewRegionHandler(deps deps.Dependencies) regionHandler {
	return regionHandler{deps}
}

func (h regionHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	controlMsg := i18n.ControlMessages[p.Account.Language]
	keyboard := types.NewKeyboard()

	for _, region := range regionsOrder {
		buttons := make([]types.Button, len(region))
		for i, region := range region {
			regionName := regions[region][0] + " " + regions[region][1]
			if region != "ru" {
				regionName += " (soon)"
			}

			buttons[i] = types.NewCallbackButton(regionName, REGION_CMD+" "+region)
		}
		keyboard.NewRow(buttons...)
	}

	keyboard.NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, SETTINGS_CMD))
	return bot.SendMessage(p.Message.Chat(), "Выберите регион:", keyboard)
}
