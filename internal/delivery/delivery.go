package delivery

import (
	"github.com/quickpowered/frilly/internal/delivery/discord"
	"github.com/quickpowered/frilly/internal/delivery/telegram"
	"github.com/quickpowered/frilly/internal/delivery/vk"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/use-cases/commands"
	"github.com/quickpowered/frilly/internal/use-cases/invoices"
	"github.com/quickpowered/frilly/pkg/consts"
	"go.uber.org/zap"
)

type Interface interface {
	Listen() error
}

func New(bot bin.Interface, commands commands.Commands, invoices invoices.Interface, logger *zap.Logger) Interface {
	switch bot.GetPlatform() {
	case consts.PlatformTelegram:
		return telegram.New(bot, commands, invoices, logger)
	case consts.PlatformDiscord:
		return discord.New(bot, commands, logger)
	case consts.PlatformVK:
		return vk.New(bot, commands, logger)
	}
	return nil
}
