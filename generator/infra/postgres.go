package infra

import (
	_ "embed"

	"github.com/rom8726/airo/config"
)

func init() {
	addDB(config.DBTypePostgres, DBInfo{
		Code:      config.DBTypePostgres,
		Title:     "PostgreSQL",
		Processor: &PostgresProcessor{},
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
