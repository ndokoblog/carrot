package carrot

import "github.com/streadway/amqp"

type QueueConfig struct {
	Name      string
	Durable   bool
	Deleted   bool
	Exclusive bool
	NoWait    bool
	Arguments map[string]interface{}
}

type ConsumeConfig struct {
	Queue     string
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Arguments map[string]interface{}
}

func (r Conn) Consume() (<-chan amqp.Delivery, error) {
	q, err := r.Ch.QueueDeclare(
		r.QueueConfig.Name,      // name
		r.QueueConfig.Durable,   // durable
		r.QueueConfig.Deleted,   // delete when unused
		r.QueueConfig.Exclusive, // exclusive
		r.QueueConfig.NoWait,    // no-wait
		r.QueueConfig.Arguments, // arguments
	)
	if err != nil {
		return nil, err
	}

	return r.Ch.Consume(
		q.Name,                    // queue
		r.ConsumeConfig.Consumer,  // consumer
		r.ConsumeConfig.AutoAck,   // auto-ack
		r.ConsumeConfig.Exclusive, // exclusive
		r.ConsumeConfig.NoLocal,   // no-local
		r.ConsumeConfig.NoWait,    // no-wait
		r.ConsumeConfig.Arguments, // args
	)
}
