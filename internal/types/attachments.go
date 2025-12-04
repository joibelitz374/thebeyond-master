package types

type Attachments []string

func NewAttachments() *Attachments {
	return &Attachments{}
}

func (a *Attachments) Add(attachment string) *Attachments {
	*a = append(*a, attachment)
	return a
}
