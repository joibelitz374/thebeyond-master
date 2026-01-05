package i18n

import "github.com/quickpowered/thebeyond-master/configs/language"

type PeriodsLocale struct {
	SelectPeriod string
	Unlimited    string
	Traffic      string
}

var PeriodsMessages = map[language.Language]PeriodsLocale{
	language.English: {
		SelectPeriod: "Select period",
		Unlimited:    "Unlimited",
		Traffic:      "Traffic",
	},
	language.Deutsch: {
		SelectPeriod: "Laufzeit auswählen",
		Unlimited:    "Unbegrenzt",
		Traffic:      "Traffic",
	},
	language.Nederlands: {
		SelectPeriod: "Periode kiezen",
		Unlimited:    "Onbeperkt",
		Traffic:      "Traffic",
	},
	language.Svenska: {
		SelectPeriod: "Välj period",
		Unlimited:    "Obegränsad",
		Traffic:      "Trafik",
	},
	language.Norsk: {
		SelectPeriod: "Velg periode",
		Unlimited:    "Ubegrenset",
		Traffic:      "Trafikk",
	},
	language.Dansk: {
		SelectPeriod: "Vælg periode",
		Unlimited:    "Ubegrænset",
		Traffic:      "Trafik",
	},
	language.Español: {
		SelectPeriod: "Seleccionar período",
		Unlimited:    "Ilimitado",
		Traffic:      "Tráfico",
	},
	language.Français: {
		SelectPeriod: "Sélectionner la période",
		Unlimited:    "Illimité",
		Traffic:      "Trafic",
	},
	language.Português: {
		SelectPeriod: "Selecionar o período",
		Unlimited:    "Ilimitado",
		Traffic:      "Tráfego",
	},
	language.Italiana: { // предположительно опечатка, должно быть Italiano
		SelectPeriod: "Seleziona periodo",
		Unlimited:    "Illimitato",
		Traffic:      "Traffico",
	},
	language.Русский: {
		SelectPeriod: "Выберите срок",
		Unlimited:    "Безлимит",
		Traffic:      "Трафик",
	},
	language.Українська: {
		SelectPeriod: "Виберіть термін",
		Unlimited:    "Безліміт",
		Traffic:      "Трафік",
	},
	language.Polski: {
		SelectPeriod: "Wybierz okres",
		Unlimited:    "Nieograniczony",
		Traffic:      "Transfer", // в польском UI чаще "transfer" для объёма данных
	},
	language.Ceština: {
		SelectPeriod: "Vyberte období",
		Unlimited:    "Neomezený",
		Traffic:      "Traffic",
	},
	language.Български: {
		SelectPeriod: "Изберете период",
		Unlimited:    "Неограничен",
		Traffic:      "Трафик",
	},
	language.Српски: {
		SelectPeriod: "Изаберите рок",
		Unlimited:    "Неограничено",
		Traffic:      "Трафик",
	},
	language.Hrvatski: {
		SelectPeriod: "Odaberite razdoblje",
		Unlimited:    "Neograničeno",
		Traffic:      "Promet",
	},
	language.Slovenčina: {
		SelectPeriod: "Vybrať obdobie",
		Unlimited:    "Neobmedzený",
		Traffic:      "Prevádzka",
	},
	language.Slovenski: {
		SelectPeriod: "Izberite obdobje",
		Unlimited:    "Neomejeno",
		Traffic:      "Promet",
	},
	language.Lietùvių: {
		SelectPeriod: "Pasirinkite laikotarpį",
		Unlimited:    "Neribotas",
		Traffic:      "Trafikas",
	},
	language.Latviešu: {
		SelectPeriod: "Izvēlieties periodu",
		Unlimited:    "Neierobežots",
		Traffic:      "Trafiks",
	},
	language.Eesti: {
		SelectPeriod: "Valige periood",
		Unlimited:    "Piiramatu",
		Traffic:      "Liiklus",
	},
	language.Suomi: {
		SelectPeriod: "Valitse ajanjakso",
		Unlimited:    "Rajoittamaton",
		Traffic:      "Liikenne",
	},
	language.Ελληνικά: {
		SelectPeriod: "Επιλέξτε περίοδο",
		Unlimited:    "Απεριόριστο",
		Traffic:      "Τραφίκ",
	},
	language.Română: {
		SelectPeriod: "Selectați perioada",
		Unlimited:    "Nelimitat",
		Traffic:      "Trafic",
	},
	language.Magyar: {
		SelectPeriod: "Válasszon időszakot",
		Unlimited:    "Korlátlan",
		Traffic:      "Forgalom",
	},
	language.Arabic: {
		SelectPeriod: "اختر الفترة",
		Unlimited:    "غير محدود",
		Traffic:      "الترافيك",
	},
	language.Farsi: {
		SelectPeriod: "دوره را انتخاب کنید",
		Unlimited:    "نامحدود",
		Traffic:      "ترافیک",
	},
	language.Türkçe: {
		SelectPeriod: "Dönem seçin",
		Unlimited:    "Sınırsız",
		Traffic:      "Trafik",
	},
	language.Hebrew: {
		SelectPeriod: "בחר תקופה",
		Unlimited:    "ללא הגבלה",
		Traffic:      "תעבורה",
	},
	language.ZH中文: {
		SelectPeriod: "选择期限",
		Unlimited:    "无限制",
		Traffic:      "流量",
	},
	language.JA日本語: {
		SelectPeriod: "期間を選択",
		Unlimited:    "無制限",
		Traffic:      "トラフィック",
	},
	language.KO한국어: {
		SelectPeriod: "기간 선택",
		Unlimited:    "무제한",
		Traffic:      "트래픽",
	},
	language.TiếngViệt: {
		SelectPeriod: "Chọn kỳ hạn",
		Unlimited:    "Không giới hạn",
		Traffic:      "Lưu lượng",
	},
	language.THภาษาไทย: {
		SelectPeriod: "เลือกช่วงเวลา",
		Unlimited:    "ไม่จำกัด",
		Traffic:      "ทราฟฟิก",
	},
	language.BahasaIndonesia: {
		SelectPeriod: "Pilih periode",
		Unlimited:    "Tanpa batas",
		Traffic:      "Trafik",
	},
	language.BahasaMelayu: {
		SelectPeriod: "Pilih tempoh",
		Unlimited:    "Tanpa had",
		Traffic:      "Trafik",
	},
	language.Tagalog: {
		SelectPeriod: "Pumili ng panahon",
		Unlimited:    "Walang limitasyon",
		Traffic:      "Trapiko",
	},
	language.Hindi: {
		SelectPeriod: "अवधि चुनें",
		Unlimited:    "असीमित",
		Traffic:      "ट्रैफिक",
	},
	language.URاردو: {
		SelectPeriod: "مدت منتخب کریں",
		Unlimited:    "لامحدود",
		Traffic:      "ٹریفک",
	},
	language.Bengali: {
		SelectPeriod: "পিরিয়ড নির্বাচন করুন",
		Unlimited:    "অসীম",
		Traffic:      "ট্রাফিক",
	},
	language.Tamiḻ: {
		SelectPeriod: "காலத்தைத் தேர்ந்தெடுக்கவும்",
		Unlimited:    "வரம்பற்ற",
		Traffic:      "டிராஃபிக்",
	},
	language.Telugu: {
		SelectPeriod: "కాలాన్ని ఎంచుకోండి",
		Unlimited:    "అపరిమిత",
		Traffic:      "ట్రాఫిక్",
	},
	language.Marathi: {
		SelectPeriod: "कालावधी निवडा",
		Unlimited:    "अमर्यादित",
		Traffic:      "ट्रॅफिक",
	},
}
