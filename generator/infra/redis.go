package infra

import (
	_ "embed"
	"fmt"
)

const (
	redisEnvFormat = `
# Redis
REDIS_HOST=%s
REDIS_PORT=6379
REDIS_PASSWORD=password
REDIS_DB=0`
)

func WithRedis() Opt {
	return func(reg *Registry) {
		reg.addInfra("redis", &InfraInfo{
			Code:      "redis",
			Title:     "Redis",
			Processor: &RedisProcessor{},
			order:     1,
		})
	}
}

//go:embed templates/redis.tmpl
var tmplRedis string

type RedisProcessor struct {
	BaseProcessor
}

func (r *RedisProcessor) Import() string {
	return "\"github.com/redis/go-redis/v9\""
}

func (r *RedisProcessor) ConfigField() string {
	return "Redis Redis `envconfig:\"REDIS\"`"
}

func (r *RedisProcessor) ConfigFieldName() string {
	return "Redis"
}

func (r *RedisProcessor) Config() string {
	return r.config(tmplRedis)
}

func (r *RedisProcessor) Constructor() string {
	return r.constructor(tmplRedis)
}

func (r *RedisProcessor) InitInAppConstructor() string {
	return r.initInAppConstructor(tmplRedis)
}

func (r *RedisProcessor) StructField() string {
	return "RedisClient *redis.Client"
}

func (r *RedisProcessor) FillStructField() string {
	return "RedisClient: redisClient,"
}

func (r *RedisProcessor) Close() string {
	return r.close(tmplRedis)
}

func (r *RedisProcessor) DockerCompose() string {
	return r.dockerCompose(tmplRedis)
}

func (r *RedisProcessor) ComposeEnv() string {
	host := r.cfg.ProjectName + "-redis"

	return fmt.Sprintf(redisEnvFormat, host)
}

func (r *RedisProcessor) ConfigEnv() string {
	host := "localhost"

	return fmt.Sprintf(redisEnvFormat, host)
}

func (r *RedisProcessor) MigrateFileData() []byte {
	return nil
}
