package delivery

import (
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/delivery/discord"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/delivery/telegram"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/delivery/vk"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/invoices"
	"github.com/quickpowered/thebeyond-master/pkg/consts"
	"go.uber.org/zap"
)

type Interface interface {
	Listen() error
}

func New(bot bin.Interface, commands commands.UseCase, invoices invoices.UseCase, logger *zap.Logger) Interface {
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
