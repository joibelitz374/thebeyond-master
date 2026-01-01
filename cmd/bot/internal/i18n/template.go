package i18n

import "github.com/quickpowered/thebeyond-master/configs/language"

type TemplateLocale struct {
	Text string
}

var TemplateMessages = map[language.Language]TemplateLocale{
	language.English:         {},
	language.Deutsch:         {},
	language.Nederlands:      {},
	language.Svenska:         {},
	language.Norsk:           {},
	language.Dansk:           {},
	language.Español:         {},
	language.Français:        {},
	language.Português:       {},
	language.Italiana:        {},
	language.Русский:         {},
	language.Українська:      {},
	language.Polski:          {},
	language.Ceština:         {},
	language.Български:       {},
	language.Српски:          {},
	language.Hrvatski:        {},
	language.Slovenčina:      {},
	language.Slovenski:       {},
	language.Lietùvių:        {},
	language.Latviešu:        {},
	language.Eesti:           {},
	language.Suomi:           {},
	language.Ελληνικά:        {},
	language.Română:          {},
	language.Magyar:          {},
	language.Arabic:          {},
	language.Farsi:           {},
	language.Türkçe:          {},
	language.Hebrew:          {},
	language.ZH中文:            {},
	language.JA日本語:           {},
	language.KO한국어:           {},
	language.TiếngViệt:       {},
	language.THภาษาไทย:       {},
	language.BahasaIndonesia: {},
	language.BahasaMelayu:    {},
	language.Tagalog:         {},
	language.Hindi:           {},
	language.URاردو:          {},
	language.Bengali:         {},
	language.Tamiḻ:           {},
	language.Telugu:          {},
	language.Marathi:         {},
}
