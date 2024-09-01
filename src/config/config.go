package config

import "fmt"

type Env struct {
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresDB       string `env:"POSTGRES_DB"`
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresPort     string `env:"POSTGRES_PORT"`
}

func (e *Env) GetDBURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		e.PostgresUser,
		e.PostgresPassword,
		e.PostgresHost,
		e.PostgresPort,
		e.PostgresDB,
	)
}
