package client

import (
	"github.com/nats-io/nats.go"
	"github.com/opencars/schema"
)

type Client struct {
	conn *nats.Conn
	js   nats.JetStreamContext
}

func New(addr string) (*Client, error) {
	conn, err := nats.Connect(addr)
	if err != nil {
		return nil, err
	}

	js, err := conn.JetStream()
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
		js:   js,
	}, nil
}

func (c *Client) Producer() (*Producer, error) {
	return NewProducer(c.conn)
}

func (c *Client) PullConsumer(subject string, stream string, consumer string, handler schema.ConsumerHandler) (*Consumer, error) {
	return NewPullConsumer(c.conn, subject, stream, consumer, handler)
}
