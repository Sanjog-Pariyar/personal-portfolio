package config

import "fmt"

const (
	// defaultConfigPostgresPort = 5433

	envConfigPostgresHost     = "POSTGRES_HOST"
	envConfigPostgresPort     = "POSTGRES_PORT"
	envConfigPostgresPassword = "POSTGRES_PASSWORD"
	envConfigPostgresUser     = "POSTGRES_USER"
	envConfigPostgresDbName   = "POSTGRES_DB_NAME"
	envConfigPostgresSslMode  = "SSL_MODE"
)

type Postgres struct {
	host     string
	port     int
	password string
	user     string
	dbName   string
	sslmode  string
}

func (pg *Postgres) PostgresDSN() string {
	return fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%d sslmode=%v",
		pg.host,
		pg.user,
		pg.password,
		pg.dbName,
		pg.port,
		pg.sslmode,
	)
}
