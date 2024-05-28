package configuration

type ServerConfiguration struct {
	QueueConfig *RabbitMqConfiguration
	HttpConfig  *HttpServerConfig
}

func DefaultConfiguration() ServerConfiguration {
	qC := RabbitMqConfiguration{}
	qC.Default()
	hC := HttpServerConfig{}
	hC.Default()
	return ServerConfiguration{
		QueueConfig: &qC,
		HttpConfig:  &hC,
	}
}
