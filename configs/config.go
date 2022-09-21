package configs

import (
	"idnmedia/utils"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type Config struct {
	App      App
	Postgres Postgres
	JWT      JWT
}

type App struct {
	Port int
}

type Postgres struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       string
}

type JWT struct {
	Secret     string
	Expiration int
}

func New() *Config {
	// load default env setup from .env file
	godotenv.Load()

	return &Config{
		App: App{
			Port: utils.StrToInt(os.Getenv("APP_PORT")),
		},
		Postgres: Postgres{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     utils.StrToInt(os.Getenv("POSTGRES_PORT")),
			Username: os.Getenv("POSTGRES_USERNAME"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DB:       os.Getenv("POSTGRES_DB"),
		},
		JWT: JWT{
			Secret:     os.Getenv("JWT_SECRET"),
			Expiration: utils.StrToInt(os.Getenv("JWT_EXPIRATION")),
		},
	}
}
