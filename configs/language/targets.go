package language

var Targets = [][]string{
	{"en", "de", "nl"},
	{"sv", "no", "da"},
	{"es", "fr"},
	{"pt", "it"},
	{"ru", "ua", "pl"},
	{"cs", "bg", "sr"},
	{"hr", "sk", "sl"},
	{"lt", "lv"},
	{"et", "fi"},
	{"el", "ro", "hu"},
	{"ar", "fa"},
	{"tr", "he"},
	{"zh", "ja", "ko"},
	{"vi", "th", "id"},
	{"ms", "tl"},
	{"hi", "ur", "bn"},
	{"ta", "te", "mr"},
}

var LimitedTargets = [][][]string{
	{
		{"en", "de"},
		{"nl", "da"},
		{"sv", "no"},
		{"fi", "et"},
	},
	{
		{"fr", "it"},
		{"es", "pt"},
		{"ro", "el"},
	},
	{
		{"ru", "ua"},
		{"pl", "cs"},
		{"sk", "hu"},
		{"lt", "lv"},
	},
	{
		{"sl", "hr"},
		{"sr", "bg"},
		{"tr", "he"},
		{"ar", "fa"},
	},
	{
		{"hi", "ur"},
		{"bn", "mr"},
		{"ta", "te"},
	},
	{
		{"zh", "ja"},
		{"ko", "vi"},
		{"th", "id"},
		{"ms", "tl"},
	},
}
