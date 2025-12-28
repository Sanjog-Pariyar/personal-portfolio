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
			envConfigCloudinaryName:      envReadString(envConfigCloudinaryName, "sanjog-personal"),
			envConfigApiKey:              envReadString(envConfigApiKey, "439893631999215"),
			envConfigApiSecret:           envReadString(envConfigApiSecret, "CJwEHvIWSUaoB8NMaF0yQr249ME"),
			envConfigApiCloudinaryFolder: envReadString(envConfigApiCloudinaryFolder, "test-image"),
			secure:                       true,
		},
		Google_Handler: &GoogleHandler{
			ClientId:     envReadString(envConfigGoogleClientId, "506172879779-ro3m8fqo9ree94ajvc0iaenb950329f2.apps.googleusercontent.com"),
			ClientSecret: envReadString(envConfigGoogleClientSecret, "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"),
		},
	}

	return config

}
