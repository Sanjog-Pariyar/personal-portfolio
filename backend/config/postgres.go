package config

import (
	"fmt"
	"log"

	"github.com/sanjog-pariyar/user-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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
	db       *gorm.DB
}

func postgresDSN(c *Config) string {
	return fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%d sslmode=%v",
		c.Postgres.host,
		c.Postgres.user,
		c.Postgres.password,
		c.Postgres.dbName,
		c.Postgres.port,
		c.Postgres.sslmode,
	)
}

func (c *Config) NewPostgres() *gorm.DB {
	db, err := gorm.Open(postgres.Open(postgresDSN(c)), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to db: %v", err)
		return nil
	}

	pg := &Postgres{
		db: db,
	}

	autoMigrate(db)
	return pg.db
}

func autoMigrate(database *gorm.DB) {
	database.AutoMigrate(&models.User{})
	fmt.Println("Automigrate complete")
}
