package client

import (
	"context"
	"time"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"

	"github.com/opencars/schema"
)

type Consumer struct {
	pattern string

	conn *nats.Conn
	js   nats.JetStreamContext

	handler schema.ConsumerHandler
}

func NewConsumer(conn *nats.Conn, pattern string, handler schema.ConsumerHandler) (*Consumer, error) {
	js, err := conn.JetStream()
	if err != nil {
		return nil, err
	}

	return &Consumer{
		pattern: pattern,
		conn:    conn,
		js:      js,
		handler: handler,
	}, nil
}

func (c *Consumer) Consume(ctx context.Context) error {
	sub, err := c.js.SubscribeSync(c.pattern)
	if err != nil {
		return err
	}

	for {
		m, err := sub.NextMsg(time.Minute)
		if err != nil {
			return err
		}

		var event schema.Event
		if err := proto.Unmarshal(m.Data, &event); err != nil {
			return err
		}

		metadata, err := m.Metadata()
		if err != nil {
			return err
		}

		meta := schema.EventMeta{
			Subject: m.Subject,
			Time:    metadata.Timestamp,
		}

		if err := c.handler.Handle(ctx, &event, &meta); err != nil {
			continue
		}

		if err := m.Ack(); err != nil {
			return err
		}
	}
}
