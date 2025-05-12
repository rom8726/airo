package infra

import (
	_ "embed"
)

func init() {
	addInfra("redis", InfraInfo{
		Code:      "redis",
		Title:     "Redis",
		Processor: &RedisProcessor{},
	})
}

//go:embed templates/redis.tmpl
var tmplRedis string

type RedisProcessor struct {
}

func (r RedisProcessor) Import() string {
	return "\"github.com/redis/go-redis/v9\""
}

func (r RedisProcessor) ConfigField() string {
	return "Redis Redis `envconfig:\"REDIS\"`"
}

func (r RedisProcessor) Config() string {
	return render(tmplRedis, "config", nil)
}

func (r RedisProcessor) Constructor() string {
	return render(tmplRedis, "constructor", nil)
}

func (r RedisProcessor) InitInAppConstructor() string {
	return render(tmplRedis, "init_in_app_constructor", nil)
}

func (r RedisProcessor) StructField() string {
	return "RedisClient *redis.Client"
}

func (r RedisProcessor) FillStructField() string {
	return "RedisClient: redisClient,"
}

func (r RedisProcessor) Close() string {
	return render(tmplRedis, "close", nil)
}
