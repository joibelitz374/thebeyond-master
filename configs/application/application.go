package application

type Application struct {
	Name string
	URL  string
}

func newApplication(name, url string) Application {
	return Application{Name: name, URL: url}
}

const (
	windows = "windows"
	linux   = "linux"
	macos   = "macos"
	android = "android"
	ios     = "ios"
	tv      = "tv"
)

var URLs = map[string][]Application{
	windows: {
		newApplication("x64 Installer", "https://github.com/Happ-proxy/happ-desktop/releases/latest/download/setup-Happ.x64.exe"),
	},
	linux: {
		newApplication("x64 Deb", "https://github.com/Happ-proxy/happ-desktop/releases/latest/download/Happ.linux.x64.deb"),
	},
	macos: {
		newApplication("App Store (Global)", "https://apps.apple.com/us/app/happ-proxy-utility/id6504287215"),
		newApplication("App Store (Rus)", "https://apps.apple.com/ru/app/happ-proxy-utility-plus/id6746188973"),
		newApplication("Universal (Dmg)", "https://github.com/Happ-proxy/happ-desktop/releases/latest/download/Happ.macOS.universal.dmg"),
	},
	android: {
		newApplication("Google Play", "https://play.google.com/store/apps/details?id=com.happproxy"),
		newApplication("APK", "https://github.com/Happ-proxy/happ-android/releases/latest/download/Happ.apk"),
	},
	ios: {
		newApplication("App Store (Global)", "https://apps.apple.com/us/app/happ-proxy-utility/id6504287215"),
		newApplication("App Store (Rus)", "https://apps.apple.com/ru/app/happ-proxy-utility-plus/id6746188973"),
	},
	tv: {
		newApplication("Android TV", "https://play.google.com/store/apps/details?id=com.happproxy"),
		newApplication("Android APK", "https://github.com/Happ-proxy/happ-android/releases/latest/download/Happ.apk"),
		newApplication("Apple TV", "https://apps.apple.com/us/app/happ-proxy-utility-for-tv/id6748297274"),
	},
}
