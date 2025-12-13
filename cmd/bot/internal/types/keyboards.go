package types

type Keyboard struct {
	ButtonRows [][]Button
}

type Button struct {
	Text string
	Data string
	URL  string
}
