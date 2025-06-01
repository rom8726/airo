package infra

import (
	_ "embed"
	"fmt"

	"github.com/rom8726/airo/config"
)

const (
	redisEnvFormat = `
# Redis
REDIS_HOST=%s
REDIS_PORT=6379
REDIS_PASSWORD=password
REDIS_DB=0`
)

// WithRedis returns a registry option that adds Redis support
func WithRedis() RegistryOption {
	return WithInfra(
		"redis",
		"Redis",
		NewRedisProcessor(),
		1,
	)
}

//go:embed templates/redis.tmpl
var tmplRedis string

// NewRedisProcessor creates a new processor for Redis
func NewRedisProcessor() Processor {
	return NewDefaultProcessor(tmplRedis,
		WithImport(func(*config.ProjectConfig) string {
			return `"github.com/redis/go-redis/v9"`
		}),
		WithConfigField("Redis Redis `envconfig:\"REDIS\"`"),
		WithConfigFieldName("Redis"),
		WithStructField("RedisClient *redis.Client"),
		WithFillStructField("RedisClient: redisClient,"),
		WithComposeEnv(func(cfg *config.ProjectConfig) string {
			return fmt.Sprintf(redisEnvFormat, cfg.ProjectName+"-redis")
		}),
		WithConfigEnv(func() string { return fmt.Sprintf(redisEnvFormat, "localhost") }()),
	)
}
