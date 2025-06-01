package infra

import (
	_ "embed"
	"fmt"

	"github.com/rom8726/airo/config"
)

const (
	natsEnvFormat = `
# NATS
NATS_URL=nats://%s:4222
NATS_USERNAME=
NATS_PASSWORD=`
)

// WithNats returns a registry option that adds NATS support
func WithNats() RegistryOption {
	return WithInfra(
		"nats",
		"NATS",
		NewNatsProcessor(),
		5,
	)
}

//go:embed templates/nats.tmpl
var tmplNats string

// NewNatsProcessor creates a new processor for NATS
func NewNatsProcessor() Processor {
	return NewDefaultProcessor(tmplNats,
		WithImport(func(*config.ProjectConfig) string {
			return `"github.com/nats-io/nats.go"`
		}),
		WithConfigField("Nats Nats `envconfig:\"NATS\"`"),
		WithConfigFieldName("Nats"),
		WithStructField("NatsClient *nats.Conn"),
		WithFillStructField("NatsClient: natsClient,"),
		WithComposeEnv(func(cfg *config.ProjectConfig) string {
			return fmt.Sprintf(natsEnvFormat, cfg.ProjectName+"-nats")
		}),
		WithConfigEnv(func() string { return fmt.Sprintf(natsEnvFormat, "localhost") }()),
	)
}
