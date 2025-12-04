package commands

import (
	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/types"
	"github.com/quickpowered/frilly/internal/use-cases/commands/tools"
	"github.com/quickpowered/frilly/internal/use-cases/components"
)

const REGION_CMD = "region"

var regions = map[string][2]string{
	"cn": {"🇨🇳", "中国"},
	"ru": {"🇷🇺", "Россия"},
	"ir": {"🇮🇷", "جمهوری اسلامی ایران"},
}

var regionsOrder = [][]string{
	{"cn", "ru"},
	{"ir"},
}

type RegionCmd struct {
	tools.Modules
	component components.Component
}

func NewRegionCmd(modules tools.Modules) *RegionCmd {
	return &RegionCmd{modules, components.NewRegionComponent()}
}

func (c *RegionCmd) Execute(bot bin.Interface, payload *domain.Payload) error {
	componentText := c.component.Text("en")
	opts := []any{tools.ToForward(bot, payload), types.DisableMentions}

	if len(payload.NodeRoute) >= 2 {
		// ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
		// defer cancel()

		region := payload.NodeRoute[1]
		// if err := c.UserService.SetRegion(ctx, payload.User.ID, region); err != nil {
		// 	return err
		// }

		return bot.SendMessage(payload.Message.GetChat(), componentText[1]+" "+region, opts...)
	}

	var idx int
	buttonRows := make([][]types.Button, len(regionsOrder))
	for _, rows := range regionsOrder {
		for _, region := range rows {
			buttonRows[idx] = append(buttonRows[idx], types.Button{
				Text: regions[region][0] + " " + regions[region][1],
				Data: "region " + region,
			})
		}
		idx++
	}

	opts = append(opts, &types.Keyboard{ButtonRows: buttonRows})
	return bot.SendMessage(payload.Message.GetChat(), componentText[0]+":", opts...)
}
