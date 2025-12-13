package language

import (
	"github.com/quickpowered/thebeyond-master/configs"
)

type Language string

const (
	English         Language = "en"
	Deutsch         Language = "de"
	Nederlands      Language = "nl"
	Svenska         Language = "sv"
	Norsk           Language = "no"
	Dansk           Language = "da"
	Español         Language = "es"
	Français        Language = "fr"
	Português       Language = "pt"
	Italiana        Language = "it"
	Русский         Language = "ru"
	Українська      Language = "ua"
	Polski          Language = "pl"
	Ceština         Language = "cs"
	Български       Language = "bg"
	Српски          Language = "sr"
	Hrvatski        Language = "hr"
	Slovenčina      Language = "sk"
	Slovenski       Language = "sl"
	Lietùvių        Language = "lt"
	Latviešu        Language = "lv"
	Eesti           Language = "et"
	Suomi           Language = "fi"
	Ελληνικά        Language = "el"
	Română          Language = "ro"
	Magyar          Language = "hu"
	Arabic          Language = "ar"
	Farsi           Language = "fa"
	Türkçe          Language = "tr"
	Hebrew          Language = "he"
	ZH中文            Language = "zh"
	JA日本語           Language = "ja"
	KO한국어           Language = "ko"
	TiếngViệt       Language = "vi"
	THภาษาไทย       Language = "th"
	BahasaIndonesia Language = "id"
	BahasaMelayu    Language = "ms"
	Tagalog         Language = "tl"
	Hindi           Language = "hi"
	URاردو          Language = "ur"
	Bengali         Language = "bn"
	Tamiḻ           Language = "ta"
	Telugu          Language = "te"
	Marathi         Language = "mr"
)

var Languages = map[Language]configs.ItemInfo{
	English:         {Flag: "🇺🇸", Emoji: "", Name: "English"},
	Deutsch:         {Flag: "🇩🇪", Emoji: "", Name: "Deutsch"},
	Nederlands:      {Flag: "🇳🇱", Emoji: "", Name: "Nederlands"},
	Svenska:         {Flag: "🇸🇪", Emoji: "", Name: "Svenska"},
	Norsk:           {Flag: "🇳🇴", Emoji: "", Name: "Norsk"},
	Dansk:           {Flag: "🇩🇰", Emoji: "", Name: "Dansk"},
	Español:         {Flag: "🇪🇸", Emoji: "", Name: "Español"},
	Français:        {Flag: "🇫🇷", Emoji: "", Name: "Français"},
	Português:       {Flag: "🇧🇷", Emoji: "", Name: "Português"},
	Italiana:        {Flag: "🇮🇹", Emoji: "", Name: "Italiano"},
	Русский:         {Flag: "🇷🇺", Emoji: "", Name: "Русский"},
	Українська:      {Flag: "🇺🇦", Emoji: "", Name: "Українська"},
	Polski:          {Flag: "🇵🇱", Emoji: "", Name: "Polski"},
	Ceština:         {Flag: "🇨🇿", Emoji: "", Name: "Čeština"},
	Български:       {Flag: "🇧🇬", Emoji: "", Name: "Български"},
	Српски:          {Flag: "🇷🇸", Emoji: "", Name: "Српски"},
	Hrvatski:        {Flag: "🇭🇷", Emoji: "", Name: "Hrvatski"},
	Slovenčina:      {Flag: "🇸🇰", Emoji: "", Name: "Slovenčina"},
	Slovenski:       {Flag: "🇸🇮", Emoji: "", Name: "Slovenščina"},
	Lietùvių:        {Flag: "🇱🇹", Emoji: "", Name: "Lietuvių"},
	Latviešu:        {Flag: "🇱🇻", Emoji: "", Name: "Latviešu"},
	Eesti:           {Flag: "🇪🇪", Emoji: "", Name: "Eesti"},
	Suomi:           {Flag: "🇫🇮", Emoji: "", Name: "Suomi"},
	Ελληνικά:        {Flag: "🇬🇷", Emoji: "", Name: "Ελληνικά"},
	Română:          {Flag: "🇷🇴", Emoji: "", Name: "Română"},
	Magyar:          {Flag: "🇭🇺", Emoji: "", Name: "Magyar"},
	Arabic:          {Flag: "🇦🇪", Emoji: "", Name: "العربية"},
	Farsi:           {Flag: "🇮🇷", Emoji: "", Name: "فارسی"},
	Türkçe:          {Flag: "🇹🇷", Emoji: "", Name: "Türkçe"},
	Hebrew:          {Flag: "🇮🇱", Emoji: "", Name: "עברית"},
	ZH中文:            {Flag: "🇨🇳", Emoji: "", Name: "中文"},
	JA日本語:           {Flag: "🇯🇵", Emoji: "", Name: "日本語"},
	KO한국어:           {Flag: "🇰🇷", Emoji: "", Name: "한국어"},
	TiếngViệt:       {Flag: "🇻🇳", Emoji: "", Name: "Tiếng Việt"},
	THภาษาไทย:       {Flag: "🇹🇭", Emoji: "", Name: "ไทย"},
	BahasaIndonesia: {Flag: "🇮🇩", Emoji: "", Name: "Bahasa Indonesia"},
	BahasaMelayu:    {Flag: "🇲🇾", Emoji: "", Name: "Bahasa Melayu"},
	Tagalog:         {Flag: "🇵🇭", Emoji: "", Name: "Filipino"},
	Hindi:           {Flag: "🇮🇳", Emoji: "", Name: "हिन्दी"},
	URاردو:          {Flag: "🇵🇰", Emoji: "", Name: "اردو"},
	Bengali:         {Flag: "🇧🇩", Emoji: "", Name: "বাংলা"},
	Tamiḻ:           {Flag: "🇮🇳", Emoji: "", Name: "தமிழ்"},
	Telugu:          {Flag: "🇮🇳", Emoji: "", Name: "తెలుగు"},
	Marathi:         {Flag: "🇮🇳", Emoji: "", Name: "मराठी"},
}

func Get(id Language) (configs.ItemInfo, bool) {
	currency, ok := Languages[id]
	return currency, ok
}
