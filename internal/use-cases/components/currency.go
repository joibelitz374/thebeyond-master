package components

const CURRENCY_COMPONENT = "currency"

type CurrencyComponent struct {
	text map[string][]string
}

func NewCurrencyComponent() *CurrencyComponent {
	text := map[string][]string{
		"en": {
			"Choose your preferred currency for pricing",
			"Currency changed to",
		},
		"de": {
			"Wählen Sie Ihre bevorzugte Währung für die Preise",
			"Währung geändert zu",
		},
		"nl": {
			"Kies uw voorkeurvaluta voor prijzen",
			"Valuta gewijzigd naar",
		},
		"sv": {
			"Välj din föredragna valuta för prissättning",
			"Valuta ändrad till",
		},
		"no": {
			"Velg din foretrukne valuta for prising",
			"Valuta endret til",
		},
		"da": {
			"Vælg din foretrukne valuta til prissætning",
			"Valuta ændret til",
		},
		"es": {
			"Elija su moneda preferida para los precios",
			"Moneda cambiada a",
		},
		"fr": {
			"Choisissez votre devise préférée pour les prix",
			"Devise modifiée en",
		},
		"pt": {
			"Escolha sua moeda preferida para precificação",
			"Moeda alterada para",
		},
		"it": {
			"Scegli la tua valuta preferita per i prezzi",
			"Valuta cambiata in",
		},
		"ru": {
			"Выберите предпочитаемую валюту для ценообразования",
			"Валюта изменена на",
		},
		"ua": {
			"Оберіть бажану валюту для ціноутворення",
			"Валюту змінено на",
		},
		"pl": {
			"Wybierz preferowaną walutę dla cen",
			"Waluta zmieniona na",
		},
		"cs": {
			"Vyberte preferovanou měnu pro ceny",
			"Měna změněna na",
		},
		"bg": {
			"Изберете предпочитаната си валута за ценообразуване",
			"Валутата е променена на",
		},
		"sr": {
			"Изаберите жељену валуту за цене",
			"Валута промењена у",
		},
		"hr": {
			"Odaberite željenu valutu za cijene",
			"Valuta promijenjena u",
		},
		"sk": {
			"Vyberte preferovanú menu pre ceny",
			"Mena zmenená na",
		},
		"sl": {
			"Izberite želeno valuto za cene",
			"Valuta spremenjena v",
		},
		"lt": {
			"Pasirinkite pageidaujamą valiutą kainodarai",
			"Valiuta pakeista į",
		},
		"lv": {
			"Izvēlieties vēlamo valūtu cenām",
			"Valūta mainīta uz",
		},
		"et": {
			"Valige eelistatud valuuta hindade jaoks",
			"Valuuta muudetud",
		},
		"fi": {
			"Valitse haluamasi valuutta hinnoittelulle",
			"Valuutta muutettu",
		},
		"el": {
			"Επιλέξτε το προτιμώμενο νόμισμά σας για τιμολόγηση",
			"Νόμισμα άλλαξε σε",
		},
		"ro": {
			"Alegeți moneda preferată pentru prețuri",
			"Moneda schimbată în",
		},
		"hu": {
			"Válassza ki a kívánt pénznemet az árakhoz",
			"Pénznem megváltoztatva",
		},
		"ar": {
			"اختر عملتك المفضلة للتسعير",
			"تم تغيير العملة إلى",
		},
		"fa": {
			"ارز مورد نظر خود را برای قیمت‌گذاری انتخاب کنید",
			"ارز تغییر یافت به",
		},
		"tr": {
			"Fiyatlandırma için tercih ettiğiniz para birimini seçin",
			"Para birimi değiştirildi",
		},
		"he": {
			"בחר את המטבע המועדף שלך לתמחור",
			"מטבע שונה ל",
		},
		"zh": {
			"选择您喜欢的定价货币",
			"货币更改为",
		},
		"ja": {
			"価格設定の優先通貨を選択してください",
			"通貨が変更されました",
		},
		"ko": {
			"가격 책정을 위한 선호 통화를 선택하세요",
			"통화가 변경됨",
		},
		"vi": {
			"Chọn đơn vị tiền tệ ưa thích của bạn cho định giá",
			"Tiền tệ đã thay đổi thành",
		},
		"th": {
			"เลือกสกุลเงินที่ต้องการสำหรับการกำหนดราคา",
			"สกุลเงินเปลี่ยนเป็น",
		},
		"id": {
			"Pilih mata uang pilihan Anda untuk penetapan harga",
			"Mata uang diubah menjadi",
		},
		"ms": {
			"Pilih mata wang pilihan anda untuk penetapan harga",
			"Mata wang diubah kepada",
		},
		"tl": {
			"Piliin ang iyong gustong pera para sa pagpepresyo",
			"Pera binago sa",
		},
		"hi": {
			"मूल्य निर्धारण के लिए अपनी पसंदीदा मुद्रा चुनें",
			"मुद्रा बदलकर",
		},
		"ur": {
			"قیمتوں کے لیے اپنی پسندیدہ کرنسی کا انتخاب کریں",
			"کرنسی تبدیل کر دی گئی",
		},
		"bn": {
			"মূল্য নির্ধারণের জন্য আপনার পছন্দের মুদ্রা নির্বাচন করুন",
			"মুদ্রা পরিবর্তন করা হয়েছে",
		},
		"ta": {
			"விலை நிர்ணயத்திற்கான உங்கள் விருப்பமான நாணயத்தைத் தேர்வுசெய்க",
			"நாணயம் மாற்றப்பட்டது",
		},
		"te": {
			"ధరలకు మీకు ఇష్టమైన కరెన్సీని ఎంచుకోండి",
			"కరెన్సీ మార్చబడింది",
		},
		"mr": {
			"किंमतसाठी आपली पसंतीची चलन निवडा",
			"चलन बदलले",
		},
	}

	return &CurrencyComponent{text}
}

func (c *CurrencyComponent) Text(language string) []string {
	return c.text[language]
}
