package schema

import "google.golang.org/protobuf/proto"

type Message struct {
	payload proto.Message
	options *MessageOptions
}

func (m *Message) Payload() proto.Message {
	return m.payload
}

func (m *Message) Options() *MessageOptions {
	return m.options
}

func (m *Message) WithOptions(ff ...MessageOption) *Message {
	if m.options == nil {
		m.options = &MessageOptions{}
	}

	for _, f := range ff {
		m.options = f(m.options)
	}

	return m
}

func NewMessage(payload proto.Message, ff ...MessageOption) *Message {
	options := &MessageOptions{}

	for _, f := range ff {
		options = f(options)
	}

	return &Message{
		payload: payload,
		options: options,
	}
}
