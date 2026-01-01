package commands

import (
	"errors"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/telegram"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/configs"
)

func getButtons(cmd string, limitedTarget [][]string, getter func(string) string) ([][]types.Button, error) {
	buttons := make([][]types.Button, len(limitedTarget))
	for idx, rows := range limitedTarget {
		for _, code := range rows {
			btn := types.NewCallbackButton(getter(code), cmd+" "+code)
			buttons[idx] = append(buttons[idx], btn)
		}
	}
	return buttons, nil
}

func setValue(bot bin.Interface, p *domain.Payload, changedToText string, getter func() (configs.ItemInfo, bool), setter func() error) error {
	tgBot, ok := bot.(*telegram.Adapter)
	if !ok {
		return errors.New("failed to cast bot to telegram adapter")
	}

	value, ok := getter()
	if !ok {
		return tgBot.AnswerCallbackQuery(p.Message.CallbackQueryID(), "Invalid code")
	}

	if err := setter(); err != nil {
		return err
	}

	return tgBot.AnswerCallbackQuery(p.Message.CallbackQueryID(), changedToText+" "+value.Flag+" "+value.Name)
}
