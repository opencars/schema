package client

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"

	"github.com/opencars/schema"
)

const (
	DefaultBatchSize = 10
)

type Consumer struct {
	subject  string
	stream   string
	consumer string

	conn *nats.Conn
	js   nats.JetStreamContext

	handler schema.ConsumerHandler
}

func NewPullConsumer(conn *nats.Conn, subject string, stream string, consumer string, handler schema.ConsumerHandler) (*Consumer, error) {
	js, err := conn.JetStream()
	if err != nil {
		return nil, err
	}

	return &Consumer{
		stream:   stream,
		consumer: consumer,
		subject:  subject,
		conn:     conn,
		js:       js,
		handler:  handler,
	}, nil
}

func (c *Consumer) Consume(ctx context.Context) error {
	sub, err := c.js.PullSubscribe(string(c.subject), string(c.consumer))
	if err != nil {
		return fmt.Errorf("subscribe: %w", err)
	}

	iter := func() error {
		ctxWithTTL, cancel := context.WithTimeout(ctx, 60*time.Second)
		defer cancel()

		msgs, err := sub.Fetch(DefaultBatchSize, nats.Context(ctxWithTTL))
		if err != nil {
			return fmt.Errorf("fetch: %w", err)
		}

		for _, m := range msgs {
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
				log.Printf("failed to handle event: %s", err)
				continue
			}

			if err := m.Ack(); err != nil {
				return err
			}
		}

		return nil
	}

	for {
		if err := iter(); err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				continue
			}

			if errors.Is(err, context.Canceled) {
				return nil
			}

			return err
		}
	}
}
