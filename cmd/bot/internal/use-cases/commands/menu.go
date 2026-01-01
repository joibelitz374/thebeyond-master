package commands

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/renew"
	"go.uber.org/zap"
)

const MENU_CMD = "start"

type menuHandler struct {
	deps deps.Dependencies
	path string
}

func NewMenuHandler(deps deps.Dependencies) menuHandler {
	path := os.Getenv("MENU_IMAGE_PATH")
	if path == "" {
		path = "./assets/menu.png"
	}
	return menuHandler{deps, path}
}

func (h menuHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	msg := i18n.AccountMessages[p.Account.Language]
	tariffsMsg := i18n.TariffsMessages[p.Account.Language]

	tariff, exists := h.deps.TariffsRepo.Get(p.Account.Tariff)
	if !exists && p.Account.Tariff > 0 {
		h.deps.Logger.Error("invalid tariff", zap.Int("tariff_id", p.Account.Tariff))
		return errors.New("invalid tariff")
	}

	var text string
	if p.Account.IsActive() {
		text += fmt.Sprintf("%s %s: %s\n", tariff.Emoji, msg.Tariff, tariffsMsg[p.Account.Tariff])

		if p.Account.Tariff > 0 && p.Account.FreemiumStatus == "available" {
			tariff.Traffic += 10
		}

		if tariff.Traffic < 5000 {
			traffic := fmt.Sprintf("%d GB", tariff.Traffic)
			usedGB := float64(p.Account.UsedUplink+p.Account.UsedDownlink) / (1024 * 1024 * 1024)
			if traffic != "" {
				text += fmt.Sprintf("🚌 %s: %.1f/%s\n", msg.Traffic, usedGB, traffic)
			}
		}

		text += fmt.Sprintf("🗺 %s: %d\n", msg.Connections, tariff.Devices)
		if daysSinceActivation := time.Until(*p.Account.SubscriptionExpiresAt).Hours() / 24; daysSinceActivation > 0 {
			text += fmt.Sprintf("🕒 %s: %.0f", msg.Remaining, daysSinceActivation)
		}
	} else {
		text += "🗒 " + msg.SubscriptionExpired
	}

	image, err := os.ReadFile(h.path)
	if err != nil {
		return err
	}

	return bot.SendMessage(p.Message.Chat(), text,
		types.NewAttachments().AddFile(bytes.NewReader(image)),
		types.NewKeyboard().
			NewRow(types.NewCallbackButton("🛜 "+msg.Connect, CONNECT_CMD)).
			NewRow(
				types.NewCallbackButton("🛍 "+msg.Renew, renew.CMD),
				types.NewCallbackButton("🔄 "+msg.NewKey, NEWKEY_CMD),
			).
			NewRow(types.NewCallbackButton("🌈 "+msg.FREEDAYS+" 🔥", PROMO_CMD)).
			NewRow(types.NewCallbackButton("ℹ️ "+msg.AboutServiceAndTerms, ABOUT_CMD)).
			NewRow(types.NewCallbackButton("💊 "+msg.Settings, SETTINGS_CMD)))
}
