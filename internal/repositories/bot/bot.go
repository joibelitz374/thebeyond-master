package bot

import (
	"fmt"

	"github.com/aejoy/vkgo/api"
	"github.com/bwmarrin/discordgo"
	"github.com/go-telegram/bot"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/repositories/bot/discord"
	"github.com/quickpowered/frilly/internal/repositories/bot/telegram"
	"github.com/quickpowered/frilly/internal/repositories/bot/vk"
	"go.uber.org/zap"
)

func New(platform, token string) bin.Interface {
	var bin bin.Interface
	switch platform {
	case "telegram":
		b, err := bot.New(token)
		if err != nil {
			panic(err)
		}

		bin = telegram.NewBot(b)
	case "discord":
		dg, err := discordgo.New("Bot " + token)
		if err != nil {
			panic(fmt.Errorf("error creating Discord session: %w", err))
		}

		bin = discord.NewBot(dg)
	case "vk":
		vkapi, err := api.New(token, zap.NewNop())
		if err != nil {
			panic(err)
		}

		bin = vk.NewBot(vkapi)
	default:
		panic("unknown platform")
	}

	return bin
}
