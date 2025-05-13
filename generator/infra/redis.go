package infra

import (
	_ "embed"

	"github.com/rom8726/airo/config"
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
	cfg *config.ProjectConfig
}

func (r *RedisProcessor) SetConfig(cfg *config.ProjectConfig) {
	r.cfg = cfg
}

func (r *RedisProcessor) Import() string {
	return "\"github.com/redis/go-redis/v9\""
}

func (r *RedisProcessor) ConfigField() string {
	return "Redis Redis `envconfig:\"REDIS\"`"
}

func (r *RedisProcessor) Config() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: r.cfg.ProjectName,
	}

	return render(tmplRedis, "config", renderData)
}

func (r *RedisProcessor) Constructor() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: r.cfg.ProjectName,
	}

	return render(tmplRedis, "constructor", renderData)
}

func (r *RedisProcessor) InitInAppConstructor() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: r.cfg.ProjectName,
	}

	return render(tmplRedis, "init_in_app_constructor", renderData)
}

func (r *RedisProcessor) StructField() string {
	return "RedisClient *redis.Client"
}

func (r *RedisProcessor) FillStructField() string {
	return "RedisClient: redisClient,"
}

func (r *RedisProcessor) Close() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: r.cfg.ProjectName,
	}

	return render(tmplRedis, "close", renderData)
}

func (r *RedisProcessor) DockerCompose() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: r.cfg.ProjectName,
	}

	return render(tmplRedis, "docker_compose", renderData)
}
