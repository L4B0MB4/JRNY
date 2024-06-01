package configuration

type HttpServerConfig struct {
	Host string
	Port int
}

func (c *HttpServerConfig) Default() {
	c.Host = ""
	c.Port = 8081
}
