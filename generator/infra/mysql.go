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

func WithMySQL() Opt {
	return func(registry *Registry) {
		registry.addDB(config.DBTypeMySQL, &DBInfo{
			Code:      config.DBTypeMySQL,
			Title:     "MySQL",
			Processor: &MysqlProcessor{},
			order:     2,
		})
	}
}

//go:embed templates/mysql.tmpl
var tmplMysql string

//go:embed files/cmd/server/migrate_mysql_go
var tmplMigrateMySQL []byte

type MysqlProcessor struct {
	BaseProcessor
}

func (m *MysqlProcessor) Import() string {
	return `"database/sql"
	_ "github.com/go-sql-driver/mysql"`
}

func (m *MysqlProcessor) Config() string {
	return m.config(tmplMysql)
}

func (m *MysqlProcessor) ConfigField() string {
	return "MySQL MySQL `envconfig:\"MYSQL\"`"
}

func (m *MysqlProcessor) ConfigFieldName() string {
	return "MySQL"
}

func (m *MysqlProcessor) Constructor() string {
	return m.constructor(tmplMysql)
}

func (m *MysqlProcessor) InitInAppConstructor() string {
	return m.initInAppConstructor(tmplMysql)
}

func (m *MysqlProcessor) StructField() string {
	return "MySQLDB *sql.DB"
}

func (m *MysqlProcessor) FillStructField() string {
	return "MySQLDB: mysqlDB,"
}

func (m *MysqlProcessor) Close() string {
	return m.close(tmplMysql)
}

func (m *MysqlProcessor) DockerCompose() string {
	return m.dockerCompose(tmplMysql)
}

func (m *MysqlProcessor) ComposeEnv() string {
	host := m.cfg.ProjectName + "-mysql"

	return fmt.Sprintf(mysqlEnvFormat, host)
}

func (m *MysqlProcessor) ConfigEnv() string {
	host := "localhost"

	return fmt.Sprintf(mysqlEnvFormat, host)
}

func (m *MysqlProcessor) MigrateFileData() []byte {
	return tmplMigrateMySQL
}
