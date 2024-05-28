package configuration

type ServerConfiguration struct {
	QueueConfig *RabbitMqConfiguration
	HttpConfig  *HttpServerConfig
}

func DefaultServerConfiguration() ServerConfiguration {
	qC := RabbitMqConfiguration{}
	qC.Default()
	hC := HttpServerConfig{}
	hC.Default()
	return ServerConfiguration{
		QueueConfig: &qC,
		HttpConfig:  &hC,
	}
}
