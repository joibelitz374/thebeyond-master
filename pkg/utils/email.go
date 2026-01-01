package utils

import "fmt"

func NewEmail(id int) string {
	return fmt.Sprintf("id%d@account", id)
}
