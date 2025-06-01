package infra

import (
	_ "embed"
	"fmt"

	"github.com/rom8726/airo/config"
)

const mysqlEnvFormat = `
# MySQL
MYSQL_HOST=%s
MYSQL_PORT=3306
MYSQL_DATABASE=db
MYSQL_PASSWORD=password
MYSQL_USER=user`

// WithMySQL returns a registry option that adds MySQL support
func WithMySQL() RegistryOption {
	return WithDB(
		config.DBTypeMySQL,
		"MySQL",
		NewMySQLProcessor(),
		2,
	)
}

//go:embed templates/mysql.tmpl
var tmplMysql string

//go:embed files/cmd/server/migrate_mysql_go
var tmplMigrateMySQL []byte

// NewMySQLProcessor creates a new processor for MySQL
func NewMySQLProcessor() Processor {
	return NewDefaultProcessor(tmplMysql,
		WithImport(func(cfg *config.ProjectConfig) string {
			return `"database/sql"
				_ "github.com/go-sql-driver/mysql"`
		}),
		WithConfigField("MySQL MySQL `envconfig:\"MYSQL\"`"),
		WithConfigFieldName("MySQL"),
		WithStructField("MySQLDB *sql.DB"),
		WithFillStructField("MySQLDB: mysqlDB,"),
		WithComposeEnv(func(cfg *config.ProjectConfig) string { return fmt.Sprintf(mysqlEnvFormat, cfg.ProjectName+"-mysql") }),
		WithConfigEnv(func() string { return fmt.Sprintf(mysqlEnvFormat, "localhost") }()),
		WithMigrateFileData(tmplMigrateMySQL),
	)
}
