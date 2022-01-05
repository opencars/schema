package schema

type HeaderKey string

var (
	HeaderTraceID HeaderKey = "trace-id"
)

type MessageOptions struct {
	Subject   string
}

type MessageOption func(*MessageOptions) *MessageOptions

func WithSubject(subject Subject) MessageOption {
	return func(m *MessageOptions) *MessageOptions {
		m.Subject = string(subject)

		return m
	}
}