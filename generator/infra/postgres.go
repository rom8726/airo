package infra

import (
	_ "embed"
	"fmt"

	"github.com/rom8726/airo/config"
)

const (
	pgEnvFormat = "POSTGRES_HOST=%s\nPOSTGRES_DATABASE=db\nPOSTGRES_PASSWORD=password\nPOSTGRES_PORT=5432\nPOSTGRES_USER=user"
)

func init() {
	addDB(config.DBTypePostgres, DBInfo{
		Code:      config.DBTypePostgres,
		Title:     "PostgreSQL",
		Processor: &PostgresProcessor{},
		order:     1,
	})
}

//go:embed templates/postgres.tmpl
var tmplPostgres string

type PostgresProcessor struct {
	cfg *config.ProjectConfig
}

func (p *PostgresProcessor) SetConfig(cfg *config.ProjectConfig) {
	p.cfg = cfg
}

func (p *PostgresProcessor) Import() string {
	return "\"github.com/jackc/pgx/v5/pgxpool\""
}

func (p *PostgresProcessor) Config() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: p.cfg.ProjectName,
	}

	return render(tmplPostgres, "config", renderData)
}

func (p *PostgresProcessor) ConfigField() string {
	return "Postgres Postgres `envconfig:\"POSTGRES\"`"
}

func (p *PostgresProcessor) Constructor() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: p.cfg.ProjectName,
	}

	return render(tmplPostgres, "constructor", renderData)
}

func (p *PostgresProcessor) InitInAppConstructor() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: p.cfg.ProjectName,
	}

	return render(tmplPostgres, "init_in_app_constructor", renderData)
}

func (p *PostgresProcessor) StructField() string {
	return "PostgresPool *pgxpool.Pool"
}

func (p *PostgresProcessor) FillStructField() string {
	return "PostgresPool: pgPool,"
}

func (p *PostgresProcessor) Close() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: p.cfg.ProjectName,
	}

	return render(tmplPostgres, "close", renderData)
}

func (p *PostgresProcessor) DockerCompose() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: p.cfg.ProjectName,
	}

	return render(tmplPostgres, "docker_compose", renderData)
}

func (p *PostgresProcessor) ComposeEnv() string {
	host := p.cfg.ProjectName + "-postgresql"

	return fmt.Sprintf(pgEnvFormat, host)
}

func (p *PostgresProcessor) ConfigEnv() string {
	host := "localhost"

	return fmt.Sprintf(pgEnvFormat, host)
}
