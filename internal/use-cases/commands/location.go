package commands

import (
	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/use-cases/commands/deps"
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

type locationHandler struct {
	deps.Dependencies
}

func NewLocationCmd(deps deps.Dependencies) locationHandler {
	return locationHandler{deps}
}

func (h locationHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	opts := []any{deps.ToForward(bot, p)}

	keyboard, err := buildKeyboardList(bot.GetPlatform(), 0, LOCATION_CMD, [][][]string{locationsOrder}, locationsOrder, func(code string) string {
		location := locations[code]
		return location.Flag + " " + location.Name
	})
	if err != nil {
		return err
	}

	opts = append(opts, keyboard)
	return bot.SendMessage(p.Message.Chat(), "Choose location:", opts...)
}
