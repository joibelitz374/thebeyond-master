package commands

import (
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
)

const REGION_CMD = "region"

var regions = map[string][2]string{
	"cn":    {"🇨🇳", "中国"},
	"ru":    {"🇷🇺", "Россия"},
	"ir":    {"🇮🇷", "جمهوری اسلامی ایران"},
	"eu_av": {"🇬🇧🇫🇷🇪🇸🇮🇹🇩🇰🇬🇷🍓", "EU Anti-AV"},
}

var regionsOrder = [][]string{
	{"cn", "ru"},
	{"ir"},
	{"eu_av"},
}

type regionHandler struct {
	deps.Dependencies
}

func NewRegionHandler(deps deps.Dependencies) regionHandler {
	return regionHandler{deps}
}

func (h regionHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	opts := []any{deps.ToForward(bot, p), types.DisableMentions}

	var idx int
	buttonRows := make([][]types.Button, len(regionsOrder))
	for _, rows := range regionsOrder {
		for _, region := range rows {
			regionName := regions[region][0] + " " + regions[region][1]
			if region != "ru" {
				regionName += " (soon)"
			}

			buttonRows[idx] = append(buttonRows[idx], types.Button{
				Text: regionName,
				Data: "region " + region,
			})
		}
		idx++
	}

	opts = append(opts, &types.Keyboard{ButtonRows: buttonRows})
	return bot.SendMessage(p.Message.Chat(), "Выберите ваш регион:\n\nОт региона зависит список серверов и их настройки. Выбирайте тот регион, где вы находитесь и для каких целей используете наш сервис!", opts...)
}
