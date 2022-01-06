package client

import (
	"github.com/nats-io/nats.go"

	"gitlab.jmindsystems.com/infinite/backend/schema"
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

func (c *Client) Producer(topic schema.Topic) (*Producer, error) {
	// p, err := sarama.NewSyncProducerFromClient(c.Client)
	// if err != nil {
	// 	return nil, err
	// }

	// return NewProducer(p, topic), nil

	return NewProducer(c.conn, c.js)
}

func (c *Client) Consumer(pattern string, handler schema.ConsumerHandler) (*Consumer, error) {
	return &Consumer{}
}
