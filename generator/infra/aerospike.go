package infra

import (
	_ "embed"
	"fmt"

	"github.com/rom8726/airo/config"
)

const (
	aerospikeEnvFormat = `
# Aerospike
AEROSPIKE_HOST=%s
AEROSPIKE_PORT=3000
AEROSPIKE_NAMESPACE=%s
AEROSPIKE_USERNAME=
AEROSPIKE_PASSWORD=`
)

// WithAerospike returns a registry option that adds Aerospike support
func WithAerospike() RegistryOption {
	return WithInfra(
		"aerospike",
		"Aerospike",
		NewAerospikeProcessor(),
		8,
	)
}

//go:embed templates/aerospike.tmpl
var tmplAerospike string

// NewAerospikeProcessor creates a new processor for Aerospike
func NewAerospikeProcessor() Processor {
	return NewDefaultProcessor(tmplAerospike,
		WithImport(func(*config.ProjectConfig) string {
			return `"github.com/aerospike/aerospike-client-go/v6"`
		}),
		WithConfigField("Aerospike Aerospike `envconfig:\"AEROSPIKE\"`"),
		WithConfigFieldName("Aerospike"),
		WithStructField("AerospikeClient *aerospike.Client"),
		WithFillStructField("AerospikeClient: aerospikeClient,"),
		WithComposeEnv(func(cfg *config.ProjectConfig) string {
			return fmt.Sprintf(aerospikeEnvFormat, cfg.ProjectName+"-aerospike", cfg.ProjectName)
		}),
		WithConfigEnv(func() string { return fmt.Sprintf(aerospikeEnvFormat, "localhost", "default") }()),
	)
}
