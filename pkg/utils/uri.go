package utils

import (
	"fmt"

	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager/dto"
)

func GenerateVLESSURI(node *dto.Node, keyID, pbk, sid string) string {
	title := fmt.Sprintf("%s ➜ %s", node.Flag, node.Name)
	return "vless://" +
		keyID + "@" +
		node.IP + ":" + fmt.Sprint(node.Port) +
		"?type=xhttp" +
		"&security=reality" +
		"&sni=" + node.SNI +
		"&pbk=" + pbk +
		"&sid=" + sid +
		"&spx=%2F" +
		"&fp=chrome" +
		"&path=%2F" +
		"#" + title
}
