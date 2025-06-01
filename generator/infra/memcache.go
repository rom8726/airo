package infra

import (
	_ "embed"
	"fmt"

	"github.com/rom8726/airo/config"
)

const (
	memcacheEnvFormat = `
# Memcache
MEMCACHE_HOST=%s
MEMCACHE_PORT=11211`
)

// WithMemcache returns a registry option that adds Memcache support
func WithMemcache() RegistryOption {
	return WithInfra(
		"memcache",
		"Memcache",
		NewMemcacheProcessor(),
		1,
	)
}

//go:embed templates/memcache.tmpl
var tmplMemcache string

// NewMemcacheProcessor creates a new processor for Memcache
func NewMemcacheProcessor() Processor {
	return NewDefaultProcessor(tmplMemcache,
		WithImport(func(*config.ProjectConfig) string {
			return `"github.com/bradfitz/gomemcache/memcache"`
		}),
		WithConfigField("Memcache Memcache `envconfig:\"MEMCACHE\"`"),
		WithConfigFieldName("Memcache"),
		WithStructField("MemcacheClient *memcache.Client"),
		WithFillStructField("MemcacheClient: memcacheClient,"),
		WithComposeEnv(func(cfg *config.ProjectConfig) string {
			return fmt.Sprintf(memcacheEnvFormat, cfg.ProjectName+"-memcache")
		}),
		WithConfigEnv(func() string { return fmt.Sprintf(memcacheEnvFormat, "localhost") }()),
	)
}
