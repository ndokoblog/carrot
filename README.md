# carrot

## Usage:
```go
func main() {
    rabbit, err := New(ConnConfig{
		User:    "guest",
		Pass:    "guest",
		Address: "localhost:5672",
	})
	failOnError(err, "Failed to open connection")

	rabbit.PublishConfig = PublishConfig{
		Exchange: "my-exchange",
		Headers: map[string]interface{}{
			"x-delay": 5000,
		},
	}

	go func() {
		body := "Halo Dunia Tipu Tipu"
		err = rabbit.Produce(body)
		failOnError(err, "Failed to produce")
		log.Printf(" [x] Sent: %s", body)
	}()

	rabbit.QueueConfig = QueueConfig{
		Name:    "my-queue",
		Durable: true,
	}

	rabbit.ConsumeConfig = ConsumeConfig{
		Consumer: "consumer-1",
		AutoAck:  true,
	}

	msgs, err := rabbit.Consume()
	failOnError(err, "Failed to consume")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			body := "Bukan Kaleng Kaleng"
			err = rabbit.Produce(body)
			failOnError(err, "Failed to produce")
			log.Printf(" [x] Sent: %s", body)

		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
```