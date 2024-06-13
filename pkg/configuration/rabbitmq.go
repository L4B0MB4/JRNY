package configuration

import "os"

type RabbitMqConfiguration struct {
	Endpoint string
}

func (c *RabbitMqConfiguration) Default() {
	url, ok := os.LookupEnv("RABBITMQ_URL")
	if !ok {
		url, ok = os.LookupEnv("RABBIT_MQ__MQTT_SERVICE_SERVICE_HOST")
		port, okPort := os.LookupEnv("RABBIT_MQ__MQTT_SERVICE_SERVICE_PORT")

		if !ok || !okPort {
			c.Endpoint = "amqp://guest:guest@localhost:5672/"
		} else {
			c.Endpoint = "amqp://guest:guest@" + url + ":" + string(port) + "/"
		}

	} else {
		c.Endpoint = url
	}
}
