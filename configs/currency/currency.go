package currency

import "github.com/quickpowered/thebeyond-master/configs"

type Currency string

const (
	XTR Currency = "xtr"
	USD Currency = "usd"
	EUR Currency = "eur"
	GBP Currency = "gbp"
	RUB Currency = "rub"
	UAH Currency = "uah"
	PLN Currency = "pln"
	CZK Currency = "czk"
	RON Currency = "ron"
	HUF Currency = "huf"
	BGN Currency = "bgn"
	RSD Currency = "rsd"
	SEK Currency = "sek"
	NOK Currency = "nok"
	DKK Currency = "dkk"
	INR Currency = "inr"
	CNY Currency = "cny"
	JPY Currency = "jpy"
	KRW Currency = "krw"
	VND Currency = "vnd"
	THB Currency = "thb"
	IDR Currency = "idr"
	MYR Currency = "myr"
	PHP Currency = "php"
	PKR Currency = "pkr"
	BDT Currency = "bdt"
	BRL Currency = "brl"
	CAD Currency = "cad"
	MXN Currency = "mxn"
	SGD Currency = "sgd"
	TRY Currency = "try"
	IRR Currency = "irr"
	ILS Currency = "ils"
	SAR Currency = "sar"
	AED Currency = "aed"
)

var Targets = [][]string{
	{"usd", "eur", "gbp"},
	{"rub", "uah", "pln"},
	{"czk", "ron", "huf"},
	{"bgn", "rsd"},
	{"sek", "nok", "dkk"},
	{"inr", "cny", "jpy"},
	{"krw", "vnd", "thb"},
	{"idr", "myr", "php"},
	{"sgd"},
	{"pkr", "bdt"},
	{"brl", "mxn", "cad"},
	{"try", "irr", "ils"},
	{"sar", "aed"},
}

var Currencies = map[Currency]configs.ItemInfo{
	USD: {Flag: "🇺🇸", Emoji: "$", Name: "US Dollar"},
	EUR: {Flag: "🇪🇺", Emoji: "€", Name: "Euro"},
	GBP: {Flag: "🇬🇧", Emoji: "£", Name: "British Pound"},
	RUB: {Flag: "🇷🇺", Emoji: "₽", Name: "Russian Ruble"},
	UAH: {Flag: "🇺🇦", Emoji: "₴", Name: "Ukrainian Hryvnia"},
	PLN: {Flag: "🇵🇱", Emoji: "zł", Name: "Polish Zloty"},
	CZK: {Flag: "🇨🇿", Emoji: "Kč", Name: "Czech Koruna"},
	RON: {Flag: "🇷🇴", Emoji: "lei", Name: "Romanian Leu"},
	HUF: {Flag: "🇭🇺", Emoji: "Ft", Name: "Hungarian Forint"},
	BGN: {Flag: "🇧🇬", Emoji: "лв", Name: "Bulgarian Lev"},
	RSD: {Flag: "🇷🇸", Emoji: "дин.", Name: "Serbian Dinar"},
	SEK: {Flag: "🇸🇪", Emoji: "kr🇸🇪", Name: "Swedish Krona"},
	NOK: {Flag: "🇳🇴", Emoji: "kr🇳🇴", Name: "Norwegian Krone"},
	DKK: {Flag: "🇩🇰", Emoji: "kr🇩🇰", Name: "Danish Krone"},
	INR: {Flag: "🇮🇳", Emoji: "₹", Name: "Indian Rupee"},
	CNY: {Flag: "🇨🇳", Emoji: "CN¥", Name: "Chinese Yuan"},
	JPY: {Flag: "🇯🇵", Emoji: "¥", Name: "Japanese Yen"},
	KRW: {Flag: "🇰🇷", Emoji: "₩", Name: "South Korean Won"},
	VND: {Flag: "🇻🇳", Emoji: "₫", Name: "Vietnamese Dong"},
	THB: {Flag: "🇹🇭", Emoji: "฿", Name: "Thai Baht"},
	IDR: {Flag: "🇮🇩", Emoji: "Rp", Name: "Indonesian Rupiah"},
	MYR: {Flag: "🇲🇾", Emoji: "RM", Name: "Malaysian Ringgit"},
	PHP: {Flag: "🇵🇭", Emoji: "₱", Name: "Philippine Peso"},
	PKR: {Flag: "🇵🇰", Emoji: "₨", Name: "Pakistani Rupee"},
	BDT: {Flag: "🇧🇩", Emoji: "৳", Name: "Bangladeshi Taka"},
	BRL: {Flag: "🇧🇷", Emoji: "R$", Name: "Brazilian Real"},
	CAD: {Flag: "🇨🇦", Emoji: "CA$", Name: "Canadian Dollar"},
	MXN: {Flag: "🇲🇽", Emoji: "MX$", Name: "Mexican Peso"},
	SGD: {Flag: "🇸🇬", Emoji: "S$", Name: "Singapore Dollar"},
	TRY: {Flag: "🇹🇷", Emoji: "₺", Name: "Turkish Lira"},
	IRR: {Flag: "🇮🇷", Emoji: "﷼🇮🇷", Name: "Iranian Rial"},
	ILS: {Flag: "🇮🇱", Emoji: "₪", Name: "Israeli New Shekel"},
	SAR: {Flag: "🇸🇦", Emoji: "﷼🇸🇦", Name: "Saudi Riyal"},
	AED: {Flag: "🇦🇪", Emoji: "د.إ", Name: "UAE Dirham"},
}

func Get(id Currency) (configs.ItemInfo, bool) {
	currency, ok := Currencies[id]
	return currency, ok
}
