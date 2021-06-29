package carrot

import (
	"github.com/streadway/amqp"
)

type PublishConfig struct {
	Exchange    string
	RoutingKey  string
	Mandatory   bool
	Immediate   bool
	ContentType string
	Headers     map[string]interface{}
}

func (r Conn) Produce(body string) error {
	if r.PublishConfig.ContentType == "" {
		r.PublishConfig.ContentType = "text/plain"
	}

	return r.Ch.Publish(
		r.PublishConfig.Exchange,   // exchange
		r.PublishConfig.RoutingKey, // routing key
		r.PublishConfig.Mandatory,  // mandatory
		r.PublishConfig.Immediate,  // immediate
		amqp.Publishing{
			ContentType: r.PublishConfig.ContentType,
			Body:        []byte(body),
			Headers:     r.PublishConfig.Headers,
		})
}
