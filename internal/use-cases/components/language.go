package components

const LANGUAGE_COMPONENT = "language"

type LanguageComponent struct {
	text map[string][]string
}

func NewLanguageComponent() *LanguageComponent {
	text := map[string][]string{
		"en": {
			"Choose the language you understand",
			"Language changed to",
		},
		"de": {
			"Wählen Sie die Sprache, die Sie verstehen",
			"Sprache geändert zu",
		},
		"nl": {
			"Kies de taal die u begrijpt",
			"Taal gewijzigd naar",
		},
		"sv": {
			"Välj det språk du förstår",
			"Språk ändrat till",
		},
		"no": {
			"Velg språket du forstår",
			"Språk endret til",
		},
		"da": {
			"Vælg det sprog, du forstår",
			"Sprog ændret til",
		},
		"es": {
			"Elija el idioma que entiende",
			"Idioma cambiado a",
		},
		"fr": {
			"Choisissez la langue que vous comprenez",
			"Langue changée en",
		},
		"pt": {
			"Escolha o idioma que você entende",
			"Idioma alterado para",
		},
		"it": {
			"Scegli la lingua che capisci",
			"Lingua cambiata in",
		},
		"ru": {
			"Выберите язык, который вы понимаете",
			"Язык изменён на",
		},
		"ua": {
			"Оберіть мову, яку ви розумієте",
			"Мову змінено на",
		},
		"pl": {
			"Wybierz język, który rozumiesz",
			"Język zmieniony na",
		},
		"cs": {
			"Vyberte jazyk, který rozumíte",
			"Jazyk změněn na",
		},
		"bg": {
			"Изберете езика, който разбирате",
			"Езикът е променен на",
		},
		"sr": {
			"Изаберите језик који разумете",
			"Језик промењен у",
		},
		"hr": {
			"Odaberite jezik koji razumijete",
			"Jezik promijenjen u",
		},
		"sk": {
			"Vyberte jazyk, ktorý rozumiete",
			"Jazyk zmenený na",
		},
		"sl": {
			"Izberite jezik, ki ga razumete",
			"Jezik spremenjen v",
		},
		"lt": {
			"Pasirinkite kalbą, kurią suprantate",
			"Kalba pakeista į",
		},
		"lv": {
			"Izvēlieties valodu, kuru saprotat",
			"Valoda mainīta uz",
		},
		"et": {
			"Valige keel, mida mõistate",
			"Keel muudetud",
		},
		"fi": {
			"Valitse kieli, jota ymmärrät",
			"Kieli muutettu",
		},
		"el": {
			"Επιλέξτε τη γλώσσα που καταλαβαίνετε",
			"Η γλώσσα άλλαξε σε",
		},
		"ro": {
			"Alegeți limba pe care o înțelegeți",
			"Limba schimbată în",
		},
		"hu": {
			"Válassza ki a nyelvet, amit ért",
			"Nyelv megváltoztatva",
		},
		"ar": {
			"اختر اللغة التي تفهمها",
			"تم تغيير اللغة إلى",
		},
		"fa": {
			"زبانی را که می‌فهمید انتخاب کنید",
			"زبان تغییر یافت به",
		},
		"tr": {
			"Anladığınız dili seçin",
			"Dil değiştirildi",
		},
		"he": {
			"בחר את השפה שאתה מבין",
			"השפה שונתה ל",
		},
		"zh": {
			"选择您理解的语言",
			"语言更改为",
		},
		"ja": {
			"理解できる言語を選択してください",
			"言語が変更されました",
		},
		"ko": {
			"이해하는 언어를 선택하세요",
			"언어가 변경되었습니다",
		},
		"vi": {
			"Chọn ngôn ngữ bạn hiểu",
			"Ngôn ngữ đã thay đổi thành",
		},
		"th": {
			"เลือกภาษาที่คุณเข้าใจ",
			"ภาษาเปลี่ยนเป็น",
		},
		"id": {
			"Pilih bahasa yang Anda pahami",
			"Bahasa diubah menjadi",
		},
		"ms": {
			"Pilih bahasa yang anda fahami",
			"Bahasa ditukar kepada",
		},
		"tl": {
			"Piliin ang wika na iyong nauunawaan",
			"Wika ay binago sa",
		},
		"hi": {
			"वह भाषा चुनें जो आप समझते हैं",
			"भाषा बदल दी गई",
		},
		"ur": {
			"وہ زبان منتخب کریں جو آپ سمجھتے ہیں",
			"زبان تبدیل کر دی گئی",
		},
		"bn": {
			"আপনি যে ভাষা বুঝতে পারেন তা নির্বাচন করুন",
			"ভাষা পরিবর্তিত হয়েছে",
		},
		"ta": {
			"நீங்கள் புரிந்து கொள்ளும் மொழியைத் தேர்வு செய்க",
			"மொழி மாற்றப்பட்டது",
		},
		"te": {
			"మీరు అర్థం చేసుకునే భాషను ఎంచుకోండి",
			"భాష మార్చబడింది",
		},
		"mr": {
			"तुम्ही समजणारी भाषा निवडा",
			"भाषा बदलली",
		},
	}

	return &LanguageComponent{text}
}

func (c *LanguageComponent) Text(language string) []string {
	return c.text[language]
}
