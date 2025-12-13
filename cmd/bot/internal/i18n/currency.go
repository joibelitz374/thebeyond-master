package i18n

import "github.com/quickpowered/thebeyond-master/configs/language"

type CurrencyLocale struct {
	ChooseCurrency    string
	CurrencyChangedTo string
}

var CurrencyMessages = map[language.Language]CurrencyLocale{
	language.English: {
		ChooseCurrency:    "Choose your preferred currency for pricing",
		CurrencyChangedTo: "Currency changed to",
	},
	language.Deutsch: {
		ChooseCurrency:    "Wählen Sie Ihre bevorzugte Währung für die Preise",
		CurrencyChangedTo: "Währung geändert zu",
	},
	language.Nederlands: {
		ChooseCurrency:    "Kies uw voorkeurvaluta voor prijzen",
		CurrencyChangedTo: "Valuta gewijzigd naar",
	},
	language.Svenska: {
		ChooseCurrency:    "Välj din föredragna valuta för prissättning",
		CurrencyChangedTo: "Valuta ändrad till",
	},
	language.Norsk: {
		ChooseCurrency:    "Velg din foretrukne valuta for prising",
		CurrencyChangedTo: "Valuta endret til",
	},
	language.Dansk: {
		ChooseCurrency:    "Vælg din foretrukne valuta til prissætning",
		CurrencyChangedTo: "Valuta ændret til",
	},
	language.Español: {
		ChooseCurrency:    "Elija su moneda preferida para los precios",
		CurrencyChangedTo: "Moneda cambiada a",
	},
	language.Français: {
		ChooseCurrency:    "Choisissez votre devise préférée pour les prix",
		CurrencyChangedTo: "Devise modifiée en",
	},
	language.Português: {
		ChooseCurrency:    "Escolha sua moeda preferida para precificação",
		CurrencyChangedTo: "Moeda alterada para",
	},
	language.Italiana: {
		ChooseCurrency:    "Scegli la tua valuta preferita per i prezzi",
		CurrencyChangedTo: "Valuta cambiata in",
	},
	language.Русский: {
		ChooseCurrency:    "Выберите предпочитаемую валюту для ценообразования",
		CurrencyChangedTo: "Валюта изменена на",
	},
	language.Українська: {
		ChooseCurrency:    "Оберіть бажану валюту для ціноутворення",
		CurrencyChangedTo: "Валюту змінено на",
	},
	language.Polski: {
		ChooseCurrency:    "Wybierz preferowaną walutę dla cen",
		CurrencyChangedTo: "Waluta zmieniona na",
	},
	language.Ceština: {
		ChooseCurrency:    "Vyberte preferovanou měnu pro ceny",
		CurrencyChangedTo: "Měna změněna na",
	},
	language.Български: {
		ChooseCurrency:    "Изберете предпочитаната си валута за ценообразуване",
		CurrencyChangedTo: "Валутата е променена на",
	},
	language.Српски: {
		ChooseCurrency:    "Изаберите жељену валуту за цене",
		CurrencyChangedTo: "Валута промењена у",
	},
	language.Hrvatski: {
		ChooseCurrency:    "Odaberite željenu valutu za cijene",
		CurrencyChangedTo: "Valuta promijenjena u",
	},
	language.Slovenčina: {
		ChooseCurrency:    "Vyberte preferovanú menu pre ceny",
		CurrencyChangedTo: "Mena zmenená na",
	},
	language.Slovenski: {
		ChooseCurrency:    "Izberite želeno valuto za cene",
		CurrencyChangedTo: "Valuta spremenjena v",
	},
	language.Lietùvių: {
		ChooseCurrency:    "Pasirinkite pageidaujamą valiutą kainodarai",
		CurrencyChangedTo: "Valiuta pakeista į",
	},
	language.Latviešu: {
		ChooseCurrency:    "Izvēlieties vēlamo valūtu cenām",
		CurrencyChangedTo: "Valūta mainīta uz",
	},
	language.Eesti: {
		ChooseCurrency:    "Valige eelistatud valuuta hindade jaoks",
		CurrencyChangedTo: "Valuuta muudetud",
	},
	language.Suomi: {
		ChooseCurrency:    "Valitse haluamasi valuutta hinnoittelulle",
		CurrencyChangedTo: "Valuutta muutettu",
	},
	language.Ελληνικά: {
		ChooseCurrency:    "Επιλέξτε το προτιμώμενο νόμισμά σας για τιμολόγηση",
		CurrencyChangedTo: "Νόμισμα άλλαξε σε",
	},
	language.Română: {
		ChooseCurrency:    "Alegeți moneda preferată pentru prețuri",
		CurrencyChangedTo: "Moneda schimbată în",
	},
	language.Magyar: {
		ChooseCurrency:    "Válassza ki a kívánt pénznemet az árakhoz",
		CurrencyChangedTo: "Pénznem megváltoztatva",
	},
	language.Arabic: {
		ChooseCurrency:    "اختر عملتك المفضلة للتسعير",
		CurrencyChangedTo: "تم تغيير العملة إلى",
	},
	language.Farsi: {
		ChooseCurrency:    "ارز مورد نظر خود را برای قیمت‌گذاری انتخاب کنید",
		CurrencyChangedTo: "ارز تغییر یافت به",
	},
	language.Türkçe: {
		ChooseCurrency:    "Fiyatlandırma için tercih ettiğiniz para birimini seçin",
		CurrencyChangedTo: "Para birimi değiştirildi",
	},
	language.Hebrew: {
		ChooseCurrency:    "בחר את המטבע המועדף שלך לתמחור",
		CurrencyChangedTo: "מטבע שונה ל",
	},
	language.ZH中文: {
		ChooseCurrency:    "选择您喜欢的定价货币",
		CurrencyChangedTo: "货币更改为",
	},
	language.JA日本語: {
		ChooseCurrency:    "価格設定の優先通貨を選択してください",
		CurrencyChangedTo: "通貨が変更されました",
	},
	language.KO한국어: {
		ChooseCurrency:    "가격 책정을 위한 선호 통화를 선택하세요",
		CurrencyChangedTo: "통화가 변경됨",
	},
	language.TiếngViệt: {
		ChooseCurrency:    "Chọn đơn vị tiền tệ ưa thích của bạn cho định giá",
		CurrencyChangedTo: "Tiền tệ đã thay đổi thành",
	},
	language.THภาษาไทย: {
		ChooseCurrency:    "เลือกสกุลเงินที่ต้องการสำหรับการกำหนดราคา",
		CurrencyChangedTo: "สกุลเงินเปลี่ยนเป็น",
	},
	language.BahasaIndonesia: {
		ChooseCurrency:    "Pilih mata uang pilihan Anda untuk penetapan harga",
		CurrencyChangedTo: "Mata uang diubah menjadi",
	},
	language.BahasaMelayu: {
		ChooseCurrency:    "Pilih mata wang pilihan anda untuk penetapan harga",
		CurrencyChangedTo: "Mata wang diubah kepada",
	},
	language.Tagalog: {
		ChooseCurrency:    "Piliin ang iyong gustong pera para sa pagpepresyo",
		CurrencyChangedTo: "Pera binago sa",
	},
	language.Hindi: {
		ChooseCurrency:    "मूल्य निर्धारण के लिए अपनी पसंदीदा मुद्रा चुनें",
		CurrencyChangedTo: "मुद्रा बदलकर",
	},
	language.URاردو: {
		ChooseCurrency:    "قیمتوں کے لیے اپنی پسندیدہ کرنسی کا انتخاب کریں",
		CurrencyChangedTo: "کرنسی تبدیل کر دی گئی",
	},
	language.Bengali: {
		ChooseCurrency:    "মূল্য নির্ধারণের জন্য আপনার পছন্দের মুদ্রা নির্বাচন করুন",
		CurrencyChangedTo: "মুদ্রা পরিবর্তন করা হয়েছে",
	},
	language.Tamiḻ: {
		ChooseCurrency:    "விலை நிர்ணயத்திற்கான உங்கள் விருப்பமான நாணயத்தைத் தேர்வுசெய்க",
		CurrencyChangedTo: "நாணயம் மாற்றப்பட்டது",
	},
	language.Telugu: {
		ChooseCurrency:    "ధరలకు మీకు ఇష్టమైన కరెన్సీని ఎంచుకోండి",
		CurrencyChangedTo: "కరెన్సీ మార్చబడింది",
	},
	language.Marathi: {
		ChooseCurrency:    "किंमतसाठी आपली पसंतीची चलन निवडा",
		CurrencyChangedTo: "चलन बदलले",
	},
}
