package types

type RichEmbed struct {
	Type        string
	Author      *Author
	Title       string
	Color       int
	Description string
}

type Author struct {
	Name string
	Icon string
}

func NewRichEmbed(title string, description string, color int, opts ...any) *RichEmbed {
	richEmbed := &RichEmbed{
		Type:        "rich",
		Title:       title,
		Color:       color,
		Description: description,
	}

	for _, param := range opts {
		switch val := param.(type) {
		case *Author:
			richEmbed.Author = val
		}
	}

	return richEmbed
}
