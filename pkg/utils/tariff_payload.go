package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func GetTariffPayloadData(payload string) (tariffID, days int, err error) {
	tariffID = 1
	for _, part := range strings.Split(payload, ";") {
		key, val, ok := strings.Cut(part, ":")
		if !ok {
			return tariffID, days, fmt.Errorf("invalid payload part: %s", part)
		}

		value, err := strconv.Atoi(val)
		if err != nil {
			return tariffID, days, err
		}

		switch key {
		case "d":
			days = value
		case "t":
			tariffID = value
		default:
			return tariffID, days, fmt.Errorf("unknown key: %s", part)
		}
	}
	return
}
