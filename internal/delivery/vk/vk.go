package vk

import (
	"github.com/aejoy/vkgo/api"
	"github.com/aejoy/vkgo/longpoll"
	"github.com/aejoy/vkgo/scene"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/use-cases/commands"
	"go.uber.org/zap"
)

type delivery struct {
	bot      bin.Interface
	commands commands.Commands
	Logger   *zap.Logger
}

func New(bot bin.Interface, commands commands.Commands, logger *zap.Logger) *delivery {
	return &delivery{bot, commands, logger}
}

func (s *delivery) Listen() error {
	scenes := scene.Message(s.onMessage)
	return longpoll.Start(s.bot.GetAPI().(*api.API), scenes)
}
