package banner

import (
	"fmt"
	"strings"
)

func Display(env string) {
	if env == "" {
		env = "development"
	}

	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("🌍🚀  THE BEYOND BOT")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("• 🟢 Status   : ONLINE")
	fmt.Println("• 🔀 Mode     : " + strings.ToUpper(string(env[0])) + env[1:])
	fmt.Println("• 📥 Services : Account, Payment")
	fmt.Println("• 🗂️ Modules  : XRay, Tools")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("📡  Listening for incoming messages…")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
}
