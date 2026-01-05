package utils

import (
	"strconv"
	"strings"
)

var keycapDigits = []string{
	"0️⃣", "1️⃣", "2️⃣", "3️⃣", "4️⃣",
	"5️⃣", "6️⃣", "7️⃣", "8️⃣", "9️⃣",
}

func ToKeycap(n int) string {
	s := strconv.Itoa(n)
	if n < 0 {
		s = "-" + s[1:]
	}

	var builder strings.Builder
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			builder.WriteString(keycapDigits[ch-'0'])
		} else {
			builder.WriteRune(ch)
		}
	}
	return builder.String()
}
