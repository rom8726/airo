package runner

import (
	"fmt"
	"os"
)

func NewEnv() Env {
	return Env{
		collidedVars: make(map[string]bool),
		redefinedVars: map[string]string{
			"API_SERVER_ADDR":   ":8080",
			"TECH_SERVER_ADDR":  ":8081",
			"POSTGRES_HOST":     "localhost",
			"POSTGRES_PORT":     "5432",
			"POSTGRES_USER":     "user",
			"POSTGRES_PASSWORD": "password",
			"POSTGRES_DATABASE": "test_db",
			"MIGRATIONS_DIR":    "../migrations",
			"REDIS_HOST":        "localhost",
			"REDIS_PORT":        "6379",
			"REDIS_PASSWORD":    "password",
			"REDIS_DB":          "0",
			"KAFKA_BROKERS":     "localhost:9092",
			"KAFKA_CLIENT_ID":   "app",
			// TODO: add more
		},
	}
}

type Env struct {
	redefinedVars map[string]string
	collidedVars  map[string]bool
}

func (e *Env) SetUp() {
	var err error
	for key, value := range e.redefinedVars {
		if envVar := os.Getenv(key); envVar != "" {
			e.redefinedVars[key] = envVar
			e.collidedVars[key] = true

			continue
		}
		if err = os.Setenv(key, value); err != nil {
			err = fmt.Errorf("can't clear ENV %s: %w", key, err)
			panic(err)
		}
	}
}

func (e *Env) CleanUp() {
	var err error
	for key := range e.redefinedVars {
		if _, ok := e.collidedVars[key]; ok {
			continue
		}
		err = os.Unsetenv(key)
		if err != nil {
			err = fmt.Errorf("can't clear ENV %s: %w", key, err)
			panic(err)
		}
	}
}

func (e *Env) Set(key, val string) {
	e.redefinedVars[key] = val
}

func (e *Env) Get(key string) string {
	return e.redefinedVars[key]
}

func (*Env) SetMock(key, server string) {
	if err := os.Setenv(key, server); err != nil {
		err = fmt.Errorf("can't set mock ENV %s: %w", key, err)
		panic(err)
	}
}
