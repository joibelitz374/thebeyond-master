package values

type Value struct {
	Flag, Emoji, Name string
}

var currencies = map[string]Value{
	"usd": {"🇺🇸", "💵", "US Dollar"},
	"eur": {"🇪🇺", "💶", "Euro"},
	"gbp": {"🇬🇧", "💷", "British Pound"},
	"rub": {"🇷🇺", "₽", "Russian Ruble"},
	"uah": {"🇺🇦", "₴", "Ukrainian Hryvnia"},
	"pln": {"🇵🇱", "zł", "Polish Zloty"},
	"czk": {"🇨🇿", "Kč", "Czech Koruna"},
	"ron": {"🇷🇴", "lei", "Romanian Leu"},
	"huf": {"🇭🇺", "Ft", "Hungarian Forint"},
	"bgn": {"🇧🇬", "лв", "Bulgarian Lev"},
	"rsd": {"🇷🇸", "дин.", "Serbian Dinar"},
	"sek": {"🇸🇪", "kr🇸🇪", "Swedish Krona"},
	"nok": {"🇳🇴", "kr🇳🇴", "Norwegian Krone"},
	"dkk": {"🇩🇰", "kr🇩🇰", "Danish Krone"},
	"inr": {"🇮🇳", "₹", "Indian Rupee"},
	"cny": {"🇨🇳", "💴", "Chinese Yuan"},
	"jpy": {"🇯🇵", "JP💴", "Japanese Yen"},
	"krw": {"🇰🇷", "₩", "South Korean Won"},
	"vnd": {"🇻🇳", "₫", "Vietnamese Dong"},
	"thb": {"🇹🇭", "฿", "Thai Baht"},
	"idr": {"🇮🇩", "Rp", "Indonesian Rupiah"},
	"myr": {"🇲🇾", "RM", "Malaysian Ringgit"},
	"php": {"🇵🇭", "₱", "Philippine Peso"},
	"pkr": {"🇵🇰", "₨", "Pakistani Rupee"},
	"bdt": {"🇧🇩", "৳", "Bangladeshi Taka"},
	"brl": {"🇧🇷", "R$", "Brazilian Real"},
	"cad": {"🇨🇦", "CA$", "Canadian Dollar"},
	"mxn": {"🇲🇽", "MX$", "Mexican Peso"},
	"sgd": {"🇸🇬", "S$", "Singapore Dollar"},
	"try": {"🇹🇷", "₺", "Turkish Lira"},
	"irr": {"🇮🇷", "﷼🇮🇷", "Iranian Rial"},
	"ils": {"🇮🇱", "₪", "Israeli New Shekel"},
	"sar": {"🇸🇦", "﷼🇸🇦", "Saudi Riyal"},
	"aed": {"🇦🇪", "د.إ", "UAE Dirham"},
}

func GetCurrencies() map[string]Value {
	return currencies
}

func GetCurrency(currency string) (Value, bool) {
	value, ok := currencies[currency]
	return value, ok
}
