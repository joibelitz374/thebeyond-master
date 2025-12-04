package utils

const (
	VLESS_SCHEME = "vless://"
	pbk          = "z8Q2qDZKKP2rNNxanKtULLFaEfe4tQQjhvDQeRlkxww"
)

func GenerateVLESSURI(keyID string, ip, sni, sid, location string) string {
	return VLESS_SCHEME +
		keyID + "@" + ip +
		"?type=xhttp" +
		"&security=reality" +
		"&sni=" + sni +
		"&pbk=" + pbk +
		"&sid=" + sid +
		"&spx=%2F" +
		"&fp=chrome" +
		"&path=%2F" +
		"#⚙️🧭🗝🎩⚙️ ➜ " + location
}
