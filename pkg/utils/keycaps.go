package utils

import "strings"

var keycaps = map[rune]string{
	'0': "0️⃣",
	'1': "1️⃣",
	'2': "2️⃣",
	'3': "3️⃣",
	'4': "4️⃣",
	'5': "5️⃣",
	'6': "6️⃣",
	'7': "7️⃣",
	'8': "8️⃣",
	'9': "9️⃣",
}

func ToKeycaps(s string) string {
	var b strings.Builder
	for _, r := range s {
		if em, ok := keycaps[r]; ok {
			b.WriteString(em)
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}
