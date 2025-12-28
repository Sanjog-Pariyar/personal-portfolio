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

func (api *Api) Addr() string {
	return fmt.Sprintf("%s:%d", api.host, api.port)
}