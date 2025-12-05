package i18n

import "github.com/quickpowered/frilly/config/language"

type LanguageLocale struct {
	ChooseLanguage    string
	LanguageChangedTo string
}

var LanguageMessages = map[language.Language]LanguageLocale{
	language.English: {
		ChooseLanguage:    "Choose the language you understand",
		LanguageChangedTo: "Language changed to",
	},
	language.Deutsch: {
		ChooseLanguage:    "Wählen Sie die Sprache, die Sie verstehen",
		LanguageChangedTo: "Sprache geändert zu",
	},
	language.Nederlands: {
		ChooseLanguage:    "Kies de taal die u begrijpt",
		LanguageChangedTo: "Taal gewijzigd naar",
	},
	language.Svenska: {
		ChooseLanguage:    "Välj det språk du förstår",
		LanguageChangedTo: "Språk ändrat till",
	},
	language.Norsk: {
		ChooseLanguage:    "Velg språket du forstår",
		LanguageChangedTo: "Språk endret til",
	},
	language.Dansk: {
		ChooseLanguage:    "Vælg det sprog, du forstår",
		LanguageChangedTo: "Sprog ændret til",
	},
	language.Español: {
		ChooseLanguage:    "Elija el idioma que entiende",
		LanguageChangedTo: "Idioma cambiado a",
	},
	language.Français: {
		ChooseLanguage:    "Choisissez la langue que vous comprenez",
		LanguageChangedTo: "Langue changée en",
	},
	language.Português: {
		ChooseLanguage:    "Escolha o idioma que você entende",
		LanguageChangedTo: "Idioma alterado para",
	},
	language.Italiana: {
		ChooseLanguage:    "Scegli la lingua che capisci",
		LanguageChangedTo: "Lingua cambiata in",
	},
	language.Русский: {
		ChooseLanguage:    "Выберите язык, который вы понимаете",
		LanguageChangedTo: "Язык изменён на",
	},
	language.Українська: {
		ChooseLanguage:    "Оберіть мову, яку ви розумієте",
		LanguageChangedTo: "Мову змінено на",
	},
	language.Polski: {
		ChooseLanguage:    "Wybierz język, który rozumiesz",
		LanguageChangedTo: "Język zmieniony na",
	},
	language.Ceština: {
		ChooseLanguage:    "Vyberte jazyk, který rozumíte",
		LanguageChangedTo: "Jazyk změněn na",
	},
	language.Български: {
		ChooseLanguage:    "Изберете езика, който разбирате",
		LanguageChangedTo: "Езикът е променен на",
	},
	language.Српски: {
		ChooseLanguage:    "Изаберите језик који разумете",
		LanguageChangedTo: "Језик промењен у",
	},
	language.Hrvatski: {
		ChooseLanguage:    "Odaberite jezik koji razumijete",
		LanguageChangedTo: "Jezik promijenjen u",
	},
	language.Slovenčina: {
		ChooseLanguage:    "Vyberte jazyk, ktorý rozumiete",
		LanguageChangedTo: "Jazyk zmenený na",
	},
	language.Slovenski: {
		ChooseLanguage:    "Izberite jezik, ki ga razumete",
		LanguageChangedTo: "Jezik spremenjen v",
	},
	language.Lietùvių: {
		ChooseLanguage:    "Pasirinkite kalbą, kurią suprantate",
		LanguageChangedTo: "Kalba pakeista į",
	},
	language.Latviešu: {
		ChooseLanguage:    "Izvēlieties valodu, kuru saprotat",
		LanguageChangedTo: "Valoda mainīta uz",
	},
	language.Eesti: {
		ChooseLanguage:    "Valige keel, mida mõistate",
		LanguageChangedTo: "Keel muudetud",
	},
	language.Suomi: {
		ChooseLanguage:    "Valitse kieli, jota ymmärrät",
		LanguageChangedTo: "Kieli muutettu",
	},
	language.Ελληνικά: {
		ChooseLanguage:    "Επιλέξτε τη γλώσσα που καταλαβαίνετε",
		LanguageChangedTo: "Η γλώσσα άλλαξε σε",
	},
	language.Română: {
		ChooseLanguage:    "Alegeți limba pe care o înțelegeți",
		LanguageChangedTo: "Limba schimbată în",
	},
	language.Magyar: {
		ChooseLanguage:    "Válassza ki a nyelvet, amit ért",
		LanguageChangedTo: "Nyelv megváltoztatva",
	},
	language.Arabic: {
		ChooseLanguage:    "اختر اللغة التي تفهمها",
		LanguageChangedTo: "تم تغيير اللغة إلى",
	},
	language.Farsi: {
		ChooseLanguage:    "زبانی را که می‌فهمید انتخاب کنید",
		LanguageChangedTo: "زبان تغییر یافت به",
	},
	language.Türkçe: {
		ChooseLanguage:    "Anladığınız dili seçin",
		LanguageChangedTo: "Dil değiştirildi",
	},
	language.Hebrew: {
		ChooseLanguage:    "בחר את השפה שאתה מבין",
		LanguageChangedTo: "השפה שונתה ל",
	},
	language.ZH中文: {
		ChooseLanguage:    "选择您理解的语言",
		LanguageChangedTo: "语言更改为",
	},
	language.JA日本語: {
		ChooseLanguage:    "理解できる言語を選択してください",
		LanguageChangedTo: "言語が変更されました",
	},
	language.KO한국어: {
		ChooseLanguage:    "이해하는 언어를 선택하세요",
		LanguageChangedTo: "언어가 변경되었습니다",
	},
	language.TiếngViệt: {
		ChooseLanguage:    "Chọn ngôn ngữ bạn hiểu",
		LanguageChangedTo: "Ngôn ngữ đã thay đổi thành",
	},
	language.THภาษาไทย: {
		ChooseLanguage:    "เลือกภาษาที่คุณเข้าใจ",
		LanguageChangedTo: "ภาษาเปลี่ยนเป็น",
	},
	language.BahasaIndonesia: {
		ChooseLanguage:    "Pilih bahasa yang Anda pahami",
		LanguageChangedTo: "Bahasa diubah menjadi",
	},
	language.BahasaMelayu: {
		ChooseLanguage:    "Pilih bahasa yang anda fahami",
		LanguageChangedTo: "Bahasa ditukar kepada",
	},
	language.Tagalog: {
		ChooseLanguage:    "Piliin ang wika na iyong nauunawaan",
		LanguageChangedTo: "Wika ay binago sa",
	},
	language.Hindi: {
		ChooseLanguage:    "वह भाषा चुनें जो आप समझते हैं",
		LanguageChangedTo: "भाषा बदल दी गई",
	},
	language.URاردو: {
		ChooseLanguage:    "وہ زبان منتخب کریں جو آپ سمجھتے ہیں",
		LanguageChangedTo: "زبان تبدیل کر دی گئی",
	},
	language.Bengali: {
		ChooseLanguage:    "আপনি যে ভাষা বুঝতে পারেন তা নির্বাচন করুন",
		LanguageChangedTo: "ভাষা পরিবর্তিত হয়েছে",
	},
	language.Tamiḻ: {
		ChooseLanguage:    "நீங்கள் புரிந்து கொள்ளும் மொழியைத் தேர்வு செய்க",
		LanguageChangedTo: "மொழி மாற்றப்பட்டது",
	},
	language.Telugu: {
		ChooseLanguage:    "మీరు అర్థం చేసుకునే భాషను ఎంచుకోండి",
		LanguageChangedTo: "భాష మార్చబడింది",
	},
	language.Marathi: {
		ChooseLanguage:    "तुम्ही समजणारी भाषा निवडा",
		LanguageChangedTo: "भाषा बदलली",
	},
}
