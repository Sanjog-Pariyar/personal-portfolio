package config

import "fmt"

const (
	envConfigApiHost         = "API_HOST"
	envConfigApiPort         = "API_PORT"
)

type Api struct {
	host       string
	port       int
}

func (c *Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.Api.host, c.Api.port)
}