package update

type Chat struct {
	ID       int
	ThreadID int
}

type ChatInterface interface {
	GetID() int
	GetThreadID() int
}

func (c Chat) GetID() int {
	return c.ID
}

func (c Chat) GetThreadID() int {
	return c.ThreadID
}
