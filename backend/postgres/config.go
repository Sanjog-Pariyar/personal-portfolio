package postgres

type Config interface {
	PostgresDSN() string
}