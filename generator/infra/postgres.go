package infra

import (
	_ "embed"
	"fmt"
	"path/filepath"

	"github.com/rom8726/airo/config"
)

const (
	pgEnvFormat = `
# PostgreSQL
POSTGRES_HOST=%s
POSTGRES_DATABASE=db
POSTGRES_PASSWORD=password
POSTGRES_PORT=5432
POSTGRES_USER=user`
)

func WithPostgres() Opt {
	return func(registry *Registry) {
		registry.addDB(config.DBTypePostgres, &DBInfo{
			Code:      config.DBTypePostgres,
			Title:     "PostgreSQL",
			Processor: &PostgresProcessor{},
			order:     1,
		})
	}
}

//go:embed templates/postgres.tmpl
var tmplPostgres string

//go:embed files/cmd/server/migrate_postgresql_go
var tmplMigratePostgres []byte

type PostgresProcessor struct {
	BaseProcessor
}

func (p *PostgresProcessor) Import() string {
	pkgDBImport := "\"" + filepath.Join(p.cfg.ModuleName, "/pkg/db") + "\""

	return "\"github.com/jackc/pgx/v5/pgxpool\"\n\t" + pkgDBImport
}

func (p *PostgresProcessor) Config() string {
	return p.config(tmplPostgres)
}

func (p *PostgresProcessor) ConfigField() string {
	return "Postgres Postgres `envconfig:\"POSTGRES\"`"
}

func (p *PostgresProcessor) ConfigFieldName() string {
	return "Postgres"
}

func (p *PostgresProcessor) Constructor() string {
	return p.constructor(tmplPostgres)
}

func (p *PostgresProcessor) InitInAppConstructor() string {
	return p.initInAppConstructor(tmplPostgres)
}

func (p *PostgresProcessor) StructField() string {
	return "PostgresPool *pgxpool.Pool"
}

func (p *PostgresProcessor) FillStructField() string {
	return "PostgresPool: pgPool,"
}

func (p *PostgresProcessor) Close() string {
	return p.close(tmplPostgres)
}

func (p *PostgresProcessor) DockerCompose() string {
	return p.dockerCompose(tmplPostgres)
}

func (p *PostgresProcessor) ComposeEnv() string {
	host := p.cfg.ProjectName + "-postgresql"

	return fmt.Sprintf(pgEnvFormat, host)
}

func (p *PostgresProcessor) ConfigEnv() string {
	host := "localhost"

	return fmt.Sprintf(pgEnvFormat, host)
}

func (p *PostgresProcessor) MigrateFileData() []byte {
	return tmplMigratePostgres
}
