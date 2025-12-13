package commands

import (
	"fmt"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
)

const DEVICES_CMD = "devices"

type DeviceLevel struct {
	Name         string
	Count, Price int
}

type devicesHandler struct {
	deps.Dependencies
}

func NewDevicesHandler(deps deps.Dependencies) devicesHandler {
	return devicesHandler{deps}
}

func (h devicesHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	devices := []DeviceLevel{
		{Name: "Семье", Count: 4, Price: (4 * 30) - 60},
		{Name: "Команде", Count: 7, Price: (7 * 30) - 60},
	}

	buttonRows := make([][]types.Button, 0, len(devices))
	for _, device := range devices {
		buttonRows = append(buttonRows, []types.Button{
			{
				Text: fmt.Sprintf("%s (%d шт.) — %d", device.Name, device.Count, device.Price),
				Data: fmt.Sprintf("device %d", device.Count),
			},
		})
	}

	return bot.SendMessage(p.Message.Chat(), "Ваши устройства:", &types.Keyboard{ButtonRows: buttonRows})
}
