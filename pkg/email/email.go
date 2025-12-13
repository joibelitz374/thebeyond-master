package email

import "fmt"

func NewAccount(id int) string {
	return fmt.Sprintf("id%d@account", id)
}
