package infra

import (
	_ "embed"
	"fmt"

	"github.com/rom8726/airo/config"
)

const mysqlEnvFormat = "MYSQL_HOST=%s\nMYSQL_PORT=3306\nMYSQL_DATABASE=db\nMYSQL_PASSWORD=password\nMYSQL_USER=user"

func init() {
	addDB(config.DBTypeMySQL, DBInfo{
		Code:      config.DBTypeMySQL,
		Title:     "MySQL",
		Processor: &MysqlProcessor{},
		order:     2,
	})
}

//go:embed templates/mysql.tmpl
var tmplMysql string

type MysqlProcessor struct {
	cfg *config.ProjectConfig
}

func (m *MysqlProcessor) SetConfig(cfg *config.ProjectConfig) {
	m.cfg = cfg
}

func (m *MysqlProcessor) Import() string {
	return `"database/sql"
	_ "github.com/go-sql-driver/mysql"`
}

func (m *MysqlProcessor) Config() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: m.cfg.ProjectName,
	}

	return render(tmplMysql, "config", renderData)
}

func (m *MysqlProcessor) ConfigField() string {
	return "MySQL MySQL `envconfig:\"MYSQL\"`"
}

func (m *MysqlProcessor) Constructor() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: m.cfg.ProjectName,
	}

	return render(tmplMysql, "constructor", renderData)
}

func (m *MysqlProcessor) InitInAppConstructor() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: m.cfg.ProjectName,
	}

	return render(tmplMysql, "init_in_app_constructor", renderData)
}

func (m *MysqlProcessor) StructField() string {
	return "MySQLDB *sql.DB"
}

func (m *MysqlProcessor) FillStructField() string {
	return "MySQLDB: mysqlDB,"
}

func (m *MysqlProcessor) Close() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: m.cfg.ProjectName,
	}

	return render(tmplMysql, "close", renderData)
}

func (m *MysqlProcessor) DockerCompose() string {
	renderData := struct {
		ProjectName string
	}{
		ProjectName: m.cfg.ProjectName,
	}

	return render(tmplMysql, "docker_compose", renderData)
}

func (m *MysqlProcessor) ComposeEnv() string {
	host := m.cfg.ProjectName + "-mysql"

	return fmt.Sprintf(mysqlEnvFormat, host)
}

func (m *MysqlProcessor) ConfigEnv() string {
	host := "localhost"

	return fmt.Sprintf(mysqlEnvFormat, host)
}
