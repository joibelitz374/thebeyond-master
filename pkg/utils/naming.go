package utils

func GetUserName(name, surname string) string {
	runeName := []rune(name)
	if len(runeName) == 0 {
		return surname
	}

	return string(runeName[0:1]) + ". " + surname
}
