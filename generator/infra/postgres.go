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

//go:embed templates/postgres.tmpl
var tmplPostgres string

//go:embed files/cmd/server/migrate_postgresql_go
var tmplMigratePostgres []byte

// WithPostgres returns a registry option that adds PostgreSQL support
func WithPostgres() RegistryOption {
	return WithDB(
		config.DBTypePostgres,
		"PostgreSQL",
		NewPostgresProcessor(),
		1,
	)
}

// NewPostgresProcessor creates a new processor for PostgreSQL
func NewPostgresProcessor() Processor {
	return NewDefaultProcessor(tmplPostgres,
		WithImport(func(cfg *config.ProjectConfig) string {
			pkgDBImport := "\"" + filepath.Join(cfg.ModuleName, "/pkg/db") + "\""
			return "\"github.com/jackc/pgx/v5/pgxpool\"\n\t" + pkgDBImport
		}),
		WithConfigField("Postgres Postgres `envconfig:\"POSTGRES\"`"),
		WithConfigFieldName("Postgres"),
		WithStructField("PostgresPool *pgxpool.Pool"),
		WithFillStructField("PostgresPool: pgPool,"),
		WithComposeEnv(func(cfg *config.ProjectConfig) string { return fmt.Sprintf(pgEnvFormat, cfg.ProjectName+"-postgresql") }),
		WithConfigEnv(func() string { return fmt.Sprintf(pgEnvFormat, "localhost") }()),
		WithMigrateFileData(tmplMigratePostgres),
	)
}
