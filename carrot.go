package carrot

import (
	"fmt"

	"github.com/streadway/amqp"
)

type ConnConfig struct {
	User    string
	Pass    string
	Address string
}

type Conn struct {
	Conn          *amqp.Connection
	Ch            *amqp.Channel
	PublishConfig PublishConfig
	QueueConfig   QueueConfig
	ConsumeConfig ConsumeConfig
}

func New(conf ConnConfig) (*Conn, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/", conf.User, conf.Pass, conf.Address))
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Conn{
		Conn: conn,
		Ch:   ch,
	}, nil
}
