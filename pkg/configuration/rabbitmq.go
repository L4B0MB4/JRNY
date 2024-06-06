package configuration

import "os"

type RabbitMqConfiguration struct {
	Endpoint string
}

func (c *RabbitMqConfiguration) Default() {
	url, ok := os.LookupEnv("RABBITMQ_URL")
	if !ok {
		c.Endpoint = "amqp://guest:guest@localhost:5672/"

	} else {
		c.Endpoint = url
	}
}
