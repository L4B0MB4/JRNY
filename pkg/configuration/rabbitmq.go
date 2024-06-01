package configuration

type RabbitMqConfiguration struct {
	Endpoint string
}

func (c *RabbitMqConfiguration) Default() {
	c.Endpoint = "amqp://guest:guest@localhost:5672/"
}
