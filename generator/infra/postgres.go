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
}

func (p PostgresProcessor) Import() string {
	return "\"github.com/jackc/pgx/v5/pgxpool\""
}

func (p PostgresProcessor) Config() string {
	return render(tmplPostgres, "config", nil)
}

func (p PostgresProcessor) ConfigField() string {
	return "Postgres Postgres `envconfig:\"POSTGRES\"`"
}

func (p PostgresProcessor) Constructor() string {
	return render(tmplPostgres, "constructor", nil)
}

func (p PostgresProcessor) InitInAppConstructor() string {
	return render(tmplPostgres, "init_in_app_constructor", nil)
}

func (p PostgresProcessor) StructField() string {
	return "PostgresPool *pgxpool.Pool"
}

func (p PostgresProcessor) FillStructField() string {
	return "PostgresPool: pgPool,"
}

func (p PostgresProcessor) Close() string {
	return render(tmplPostgres, "close", nil)
}
