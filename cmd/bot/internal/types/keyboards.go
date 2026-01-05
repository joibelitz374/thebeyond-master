package types

type Keyboard struct {
	ButtonRows [][]Button
}

func NewKeyboard() *Keyboard {
	return &Keyboard{}
}

func (k *Keyboard) NewRow(buttons ...Button) *Keyboard {
	k.ButtonRows = append(k.ButtonRows, buttons)
	return k
}

type Button struct {
	Text, Data, URL string
	Pay             bool
}

func NewPayButton(text string) Button {
	return Button{Text: text, Pay: true}
}

func NewCallbackButton(text, data string) Button {
	return Button{Text: text, Data: data}
}

func NewURLButton(text, url string) Button {
	return Button{Text: text, URL: url}
}
