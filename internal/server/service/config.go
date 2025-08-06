package service

import "fmt"

type Config struct {
	Host                    string
	Port                    uint16
	PortDoc                 uint16
	CaCert                  string
	Cert                    string
	Key                     string
	SslMutualAuthentication bool
}

func (c *Config) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *Config) AddressDoc() string {
	return fmt.Sprintf("%s:%d", c.Host, c.PortDoc)
}
