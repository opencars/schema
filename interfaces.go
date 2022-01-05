package schema

import (
	"context"

	"google.golang.org/protobuf/proto"
)

type ProducableMiddleware interface {
	Handle(context.Context, Producable)
}

type ProducableMiddlewareFunc func(context.Context, Producable)

func (f ProducableMiddlewareFunc) Handle(ctx context.Context, p Producable) {
	f(ctx, p)
}

type Producable interface {
	Payload() proto.Message
	Options() *MessageOptions
}

type Producer interface {
	Produce(ctx context.Context, pp ...Producable) error
}

type ConsumerHandler interface {
	Handle(context.Context, *Event, *EventMeta) error
}

type ConsumerHandlerFunc func(context.Context, *Event, *EventMeta) error

func (f ConsumerHandlerFunc) Handle(ctx context.Context, e *Event, em *EventMeta) error {
	return f(ctx, e, em)
}

type ConsumableMessage struct {
	Event *Event
	Meta  *EventMeta
	Ack   func()
}

type BatchConsumerHandler interface {
	Handle(context.Context, []*ConsumableMessage) error
}

type BatchConsumerHandlerFunc func(context.Context, []*ConsumableMessage) error

func (f BatchConsumerHandlerFunc) Handle(ctx context.Context, batch []*ConsumableMessage) error {
	return f(ctx, batch)
}
