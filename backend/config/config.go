package config

import (
	"os"
	"strconv"
	"log"

	"github.com/joho/godotenv"
)

const envConfigJWTsecret = "JWT_SECRET"

type Config struct {
	Api            *Api
	Postgres       *Postgres
	Jwt_secret     string
	Cloudinary     *CloudinaryEnv
	Google_Handler *GoogleHandler
}

func envReadString(envName, defaultValue string) string {
	value := os.Getenv(envName)
	if value == "" {
		value = defaultValue
	}

	return value
}

func envReadNumeric(envName string, defaultValue int) int {
	value, _ := strconv.Atoi(os.Getenv(envName))
	if value == 0 {
		value = defaultValue
	}
	return value
}

func NewConfig() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	apiPort := envReadNumeric(envConfigApiPort, 0)

	config := &Config{
		Api: &Api{
			host: envReadString(envConfigApiHost, ""),
			port: apiPort,
		},
		Postgres: &Postgres{
			host:     envReadString(envConfigPostgresHost, ""),
			port:     envReadNumeric(envConfigPostgresPort, 0),
			password: envReadString(envConfigPostgresPassword, ""),
			user:     envReadString(envConfigPostgresUser, ""),
			dbName:   envReadString(envConfigPostgresDbName, ""),
			sslmode:  envReadString(envConfigPostgresSslMode, ""),
		},
		Jwt_secret: envReadString(envConfigJWTsecret, ""),
		Cloudinary: &CloudinaryEnv{
			envConfigCloudinaryName:      envReadString(envConfigCloudinaryName, ""),
			envConfigApiKey:              envReadString(envConfigApiKey, ""),
			envConfigApiSecret:           envReadString(envConfigApiSecret, ""),
			envConfigApiCloudinaryFolder: envReadString(envConfigApiCloudinaryFolder, ""),
			secure:                       true,
		},
		Google_Handler: &GoogleHandler{
			ClientId:     envReadString(envConfigGoogleClientId, ""),
			ClientSecret: envReadString(envConfigGoogleClientSecret, ""),
		},
	}

	return config

}
