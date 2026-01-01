package utils

import "fmt"

func ParseAccountID(email string) (int, error) {
	var id int
	if _, err := fmt.Sscanf(email, "id%d@account", &id); err != nil {
		return 0, fmt.Errorf("invalid email format: %w", err)
	}
	return id, nil
}
