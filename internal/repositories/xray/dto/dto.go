package dto

type Config struct {
	Inbounds []Inbound `json:"inbounds"`
}

type Inbound struct {
	Tag      string          `json:"tag"`
	Settings InboundSettings `json:"settings"`
}

type InboundSettings struct {
	Clients []Client `json:"clients"`
}

type Client struct {
	Email string `json:"email"`
	Id    string `json:"id"`
	Level int    `json:"level"`
}
