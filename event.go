package schema

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

var (
	ErrPayloadNotProto = errors.New("schema: provided event payload could not be serialized to protobuf")
	ErrEventNotValid   = errors.New("schema: event failed validation")
)

// New creates new generic event - a wrapper for any recognizable payload.
func New(source *Source, payload proto.Message) *Event {
	event := Event{
		Id:     uuid.NewV4().String(),
		Source: source,
	}

	// Such anypb.New marshaling without marshaling options could raise error only in case payload is nil.
	any, _ := anypb.New(payload)
	event.Payload = any

	return &event
}

func (e *Event) Message(ff ...MessageOption) *Message {
	return NewMessage(e, ff...)
}

func NewCUDPayload(typ CUD_Type, version uint64, current, change proto.Message) *CUD {
	cud := CUD{
		Type:    typ,
		Version: 0,
	}

	if current != nil {
		curr, _ := anypb.New(current)
		cud.Current = curr
	}

	if change != nil {
		ch, _ := anypb.New(change)
		cud.Change = ch
	}

	return &cud
}

func (e *CUD) Message(ff ...MessageOption) *Message {
	return NewMessage(e, ff...)
}

// EventMeta contains metadata of event.
type EventMeta struct {
	Key     string
	Topic   string
	Group   string
	Time    time.Time
}
