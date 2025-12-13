package domain

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Avatar       string `json:"avatar"`
	Subscription `json:"subscription"`
	Language     string `json:"language"`
	Currency     string `json:"currency"`
}

type Subscription struct {
	ID        string `json:"id"`
	ExpiresAt int    `json:"expires_at"`
}
