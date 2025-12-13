package commands

import (
	"fmt"
	"strconv"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/configs"
	"github.com/quickpowered/thebeyond-master/pkg/consts"
)

func buildKeyboardList(
	platform consts.Platform,
	page int,
	cmd string,
	pages [][][]string,
	limitedTarget [][]string,
	getter func(code string) string,
) (*types.Keyboard, error) {
	buttonRows, err := getButtons(platform, page, cmd, pages, limitedTarget, getter)
	if err != nil {
		return nil, fmt.Errorf("invalid currency code")
	}

	if platform != consts.PlatformTelegram {
		buttonRows = append(buttonRows, addNavigationButtons(cmd, page, limitedTarget)...)
	}

	return &types.Keyboard{ButtonRows: buttonRows}, nil
}

func getButtons(
	platform consts.Platform,
	page int, cmd string,
	pages [][][]string, limitedTarget [][]string,
	getter func(string) string,
) ([][]types.Button, error) {
	if platform != consts.PlatformTelegram {
		limitedTarget = pages[page]
	}

	buttonRows, idx := make([][]types.Button, len(limitedTarget)), 0
	for _, rows := range limitedTarget {
		for _, code := range rows {
			buttonRows[idx] = append(buttonRows[idx], types.Button{
				Text: getter(code),
				Data: cmd + " " + code,
			})
		}
		idx++
	}

	return buttonRows, nil
}

func addNavigationButtons(cmd string, page int, limitedTarget [][]string) [][]types.Button {
	buttons := make([][]types.Button, 1)
	if page > 0 {
		buttons[0] = append(buttons[0], types.Button{
			Text: "⬅️",
			Data: cmd + " " + strconv.Itoa(page-1),
		})
	}
	if page < len(limitedTarget)-1 {
		buttons[0] = append(buttons[0], types.Button{
			Text: "➡️",
			Data: cmd + " " + strconv.Itoa(page+1),
		})
	}

	return buttons
}

func setValue(
	bot bin.Interface,
	p *domain.Payload, changedToText string,
	getter func() (configs.ItemInfo, bool),
	setter func() error,
	opts ...any,
) error {
	val, ok := getter()
	if !ok {
		if err := bot.SendMessage(p.Message.Chat(), "Invalid code", opts...); err != nil {
			return err
		}

		return fmt.Errorf("invalid code")
	}

	if err := setter(); err != nil {
		return err
	}

	return bot.SendMessage(p.Message.Chat(), changedToText+" "+val.Flag+" "+val.Name, opts...)
}
