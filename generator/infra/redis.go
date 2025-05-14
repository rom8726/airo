package infra

import (
	_ "embed"
	"fmt"

	"github.com/rom8726/airo/config"
)

const (
	redisEnvFormat = "REDIS_HOST=%s\nREDIS_PORT=6379\nREDIS_PASSWORD=password\nREDIS_DB=0"
)

func init() {
	addInfra("redis", InfraInfo{
		Code:      "redis",
		Title:     "Redis",
		Processor: &RedisProcessor{},
		order:     1,
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

func (r *RedisProcessor) ComposeEnv() string {
	host := r.cfg.ProjectName + "-redis"

	return fmt.Sprintf(redisEnvFormat, host)
}

func (r *RedisProcessor) ConfigEnv() string {
	host := "localhost"

	return fmt.Sprintf(redisEnvFormat, host)
}
