package commands

import (
	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/types"
	"github.com/quickpowered/frilly/internal/use-cases/commands/tools"
	"github.com/quickpowered/frilly/pkg/consts"
)

const LOCATION_CMD = "location"

var locations = map[string]domain.Location{
	"de": {
		Flag: "🩷",
		Name: "Frankfurt 🇩🇪",
		IP:   "77.221.157.159:443",
	},
}

var locationsOrder = [][]string{
	{"de"},
}

type LocationCmd struct {
	tools.Modules
}

func NewLocationCmd(modules tools.Modules) *LocationCmd {
	return &LocationCmd{modules}
}

func (c *LocationCmd) Execute(bot bin.Interface, payload *domain.Payload) error {
	opts := []any{tools.ToForward(bot, payload)}

	buttonRows, err := getButtons(bot, 0, "location", [][][]string{locationsOrder}, locationsOrder, func(code string) string {
		location := locations[code]
		return location.Flag + " " + location.Name
	})
	if err != nil {
		bot.SendMessage(payload.Message.GetChat(), "Invalid location code", opts...)
		return err
	}

	if bot.GetPlatform() != consts.PlatformTelegram {
		buttonRows = append(buttonRows, addNavigationButtons(bot, 0)...)
	}

	opts = append(opts, &types.Keyboard{ButtonRows: buttonRows})
	return bot.SendMessage(payload.Message.GetChat(), "Choose location:", opts...)
}
