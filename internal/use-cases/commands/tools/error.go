package tools

type CommandError struct {
	Command  string
	Service  string
	Method   string
	ChatID   string
	MemberID int
	Content  string
}

func NewCommandError(cmd, service, method string, chatID string, memberID int, content string) *CommandError {
	return &CommandError{
		Command:  cmd,
		Service:  service,
		Method:   method,
		ChatID:   chatID,
		MemberID: memberID,
		Content:  content,
	}
}

func (e CommandError) Error() string {
	return e.Service + "." + e.Method + ": " + e.Content
}
