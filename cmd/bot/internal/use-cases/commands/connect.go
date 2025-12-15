package commands

import (
	"os"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	"github.com/quickpowered/thebeyond-master/configs/language"
)

const CONNECT_CMD = "connect"

type connectHandler struct {
	deps.Dependencies
}

func NewConnectHandler(deps deps.Dependencies) connectHandler {
	return connectHandler{deps}
}

func (h connectHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	msg := i18n.ConnectMessages[language.Language(p.Account.Language)]

	if len(p.Args) > 1 {
		return h.connectViaApp(bot, msg, p)
	}

	projectDomain := os.Getenv("PROJECT_DOMAIN")
	subscriptionURL := "https://" + projectDomain + "/sub/" + p.Account.KeyID + "/smart/ru"

	return bot.SendMessage(p.Message.Chat(), msg.DoYouHaveAnApp, &types.Keyboard{
		ButtonRows: [][]types.Button{
			{{Text: "🎾 " + msg.IHave, URL: "https://" + projectDomain + "/sub/r?url=happ://add/" + subscriptionURL}},
			{{Text: "📥 " + msg.DownloadApp, Data: "connect download"}},
		},
	})
}

type DownloadURL struct {
	Name string
	URL  string
}

var downloadURLs = map[string][]DownloadURL{
	"windows": {
		{
			Name: "x64 Installer",
			URL:  "https://github.com/Happ-proxy/happ-desktop/releases/latest/download/setup-Happ.x64.exe",
		},
	},
	"linux": {
		{
			Name: "x64 Deb",
			URL:  "https://github.com/Happ-proxy/happ-desktop/releases/latest/download/Happ.linux.x64.deb",
		},
	},
	"macos": {
		{
			Name: "App Store (Global)",
			URL:  "https://apps.apple.com/us/app/happ-proxy-utility/id6504287215",
		},
		{
			Name: "App Store (Rus)",
			URL:  "https://apps.apple.com/ru/app/happ-proxy-utility-plus/id6746188973",
		},
		{
			Name: "Universal (Dmg)",
			URL:  "https://github.com/Happ-proxy/happ-desktop/releases/latest/download/Happ.macOS.universal.dmg",
		},
	},
	"android": {
		{
			Name: "Google Play",
			URL:  "https://play.google.com/store/apps/details?id=com.happproxy",
		},
		{
			Name: "APK",
			URL:  "https://github.com/Happ-proxy/happ-android/releases/latest/download/Happ.apk",
		},
	},
	"ios": {
		{
			Name: "App Store (Global)",
			URL:  "https://apps.apple.com/us/app/happ-proxy-utility/id6504287215",
		},
		{
			Name: "App Store (Rus)",
			URL:  "https://apps.apple.com/ru/app/happ-proxy-utility-plus/id6746188973",
		},
	},
	"tv": {
		{
			Name: "Android TV",
			URL:  "https://play.google.com/store/apps/details?id=com.happproxy",
		},
		{
			Name: "Android APK",
			URL:  "https://github.com/Happ-proxy/happ-android/releases/latest/download/Happ.apk",
		},
		{
			Name: "Apple TV",
			URL:  "https://apps.apple.com/us/app/happ-proxy-utility-for-tv/id6748297274",
		},
	},
}

func (h connectHandler) connectViaApp(bot bin.Interface, msg i18n.ConnectLocale, p *domain.Payload) error {
	switch device := p.Args[1]; device {
	case "download":
		return bot.SendMessage(p.Message.Chat(), msg.InstallationDevice, &types.Keyboard{
			ButtonRows: [][]types.Button{
				{{Text: "🖼 Windows", Data: "connect windows"}},
				{{Text: "🐧 Linux", Data: "connect linux"}},
				{{Text: "🙂 MacOS", Data: "connect macos"}},
				{{Text: "🤖 Android", Data: "connect android"}},
				{{Text: "🍎 iOS", Data: "connect ios"}},
				{{Text: "🖥 TV", Data: "connect tv"}},
			},
		})
	default:
		urls, ok := downloadURLs[device]
		if !ok {
			return bot.SendMessage(p.Message.Chat(), "Device not found")
		}

		buttonRows := [][]types.Button{}
		for _, url := range urls {
			buttonRows = append(buttonRows, []types.Button{{Text: url.Name, URL: url.URL}})
		}

		return bot.SendMessage(p.Message.Chat(), msg.InstallationMethod, &types.Keyboard{ButtonRows: buttonRows})
	}
}
