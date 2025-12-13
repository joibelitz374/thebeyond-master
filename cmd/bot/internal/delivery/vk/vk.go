package vk

import (
	"github.com/aejoy/vkgo/api"
	"github.com/aejoy/vkgo/longpoll"
	"github.com/aejoy/vkgo/scene"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands"
	"go.uber.org/zap"
)

type delivery struct {
	bot      bin.Interface
	commands commands.UseCase
	Logger   *zap.Logger
}

func New(bot bin.Interface, commands commands.UseCase, logger *zap.Logger) *delivery {
	return &delivery{bot, commands, logger}
}

func (s *delivery) Listen() error {
	scenes := scene.Message(s.onMessage)
	return longpoll.Start(s.bot.GetAPI().(*api.API), scenes)
}
