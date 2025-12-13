package types

import "io"

type Attachments struct {
	Urls  []string
	Files []io.Reader
}

func NewAttachments() *Attachments {
	return &Attachments{}
}

func (a *Attachments) AddURL(url string) *Attachments {
	a.Urls = append(a.Urls, url)
	return a
}

func (a *Attachments) AddFile(file io.Reader) *Attachments {
	a.Files = append(a.Files, file)
	return a
}
