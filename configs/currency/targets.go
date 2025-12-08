package currency

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

var LimitedTargets = [][][]string{
	{
		{"usd", "eur"},
		{"gbp", "rub"},
		{"sek", "nok"},
		{"dkk", "try"},
	},
	{
		{"uah", "pln"},
		{"czk", "huf"},
		{"ron", "bgn"},
		{"rsd", "ils"},
	},
	{
		{"cny", "jpy"},
		{"krw", "inr"},
		{"pkr", "bdt"},
	},
	{
		{"sgd", "thb"},
		{"vnd", "idr"},
		{"myr", "php"},
	},
	{
		{"cad", "brl"},
		{"mxn", "sar"},
		{"aed", "irr"},
	},
}
