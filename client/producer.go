package client

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"

	"github.com/opencars/schema"
)

type Producer struct {
	conn *nats.Conn
	js   nats.JetStreamContext

	mws []schema.ProducableMiddleware
}

func NewProducer(conn *nats.Conn) (*Producer, error) {
	js, err := conn.JetStream()
	if err != nil {
		return nil, err
	}

	return &Producer{
		conn: conn,
		js:   js,
	}, nil
}

func (p *Producer) Use(mws ...schema.ProducableMiddleware) {
	p.mws = append(p.mws, mws...)
}

func (p *Producer) Produce(ctx context.Context, pp ...schema.Producable) error {
	for _, item := range pp {
		for i := range p.mws {
			p.mws[i].Handle(ctx, item)
		}

		serialized, err := proto.Marshal(item.Payload())
		if err != nil {
			return err
		}

		options := item.Options()

		result, err := p.js.Publish(options.Subject, serialized)
		if err != nil {
			return err
		}

		log.Println("publish: %+v", result)
	}

	return nil
}
