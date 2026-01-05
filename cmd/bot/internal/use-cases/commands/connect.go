package commands

import (
	"fmt"
	"os"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	"github.com/quickpowered/thebeyond-master/configs/application"
)

const CONNECT_CMD = "connect"

type connectHandler struct {
	deps.Dependencies
	projectDomain string
}

func NewConnectHandler(deps deps.Dependencies) connectHandler {
	return connectHandler{deps, os.Getenv("PROJECT_DOMAIN")}
}

func (h connectHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	msg := i18n.ConnectMessages[p.Account.Language]
	controlMsg := i18n.ControlMessages[p.Account.Language]
	if len(p.Args) > 1 {
		return h.connectViaApp(bot, p, msg, controlMsg)
	}

	subscriptionURL := fmt.Sprintf("https://%s/sub/%s/smart/%s", h.projectDomain, p.Account.KeyID, p.Account.Region)
	redirectToAddURL := fmt.Sprintf("https://%s/sub/r?url=happ://add/%s", h.projectDomain, subscriptionURL)
	return bot.SendMessage(p.Message.Chat(), msg.DoYouHaveAnApp,
		types.NewKeyboard().
			NewRow(types.NewURLButton("🎾 "+msg.IHave, redirectToAddURL)).
			NewRow(types.NewCallbackButton("📥 "+msg.DownloadApp, CONNECT_CMD+" download")).
			NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, MENU_CMD)))
}

func (h connectHandler) connectViaApp(bot bin.Interface, p *domain.Payload, msg i18n.ConnectLocale, controlMsg i18n.ControlLocale) error {
	if p.Args[1] == "download" {
		return bot.SendMessage(p.Message.Chat(), msg.InstallationDevice, types.NewKeyboard().
			NewRow(types.NewCallbackButton("🖼 Windows", CONNECT_CMD+" windows")).
			NewRow(types.NewCallbackButton("🐧 Linux", CONNECT_CMD+" linux")).
			NewRow(types.NewCallbackButton("🙂 MacOS", CONNECT_CMD+" macos")).
			NewRow(types.NewCallbackButton("🤖 Android", CONNECT_CMD+" android")).
			NewRow(types.NewCallbackButton("🍎 iOS", CONNECT_CMD+" ios")).
			NewRow(types.NewCallbackButton("🖥 TV", CONNECT_CMD+" tv")).
			NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, CONNECT_CMD)))
	}

	urls, ok := application.URLs[p.Args[1]]
	if !ok {
		return bot.SendMessage(p.Message.Chat(), "Device not found")
	}

	keyboard := types.NewKeyboard()
	for _, url := range urls {
		keyboard.NewRow(types.NewURLButton(url.Name, url.URL))
	}

	keyboard.NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, CONNECT_CMD+" download"))
	return bot.SendMessage(p.Message.Chat(), msg.InstallationMethod, keyboard)
}
