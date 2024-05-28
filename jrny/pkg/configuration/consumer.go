package configuration

type ConsumerConfiguration struct {
	QueueConfig *RabbitMqConfiguration
}

func DefaultConsumerConfiguration() ConsumerConfiguration {
	qC := RabbitMqConfiguration{}
	qC.Default()
	return ConsumerConfiguration{
		QueueConfig: &qC,
	}
}
