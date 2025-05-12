package config

type DBType string

const (
	DBTypeUnknown  DBType = ""
	DBTypePostgres DBType = "postgres"
	DBTypeMySQL    DBType = "mysql"
)

type ProjectConfig struct {
	Aborted     bool
	ProjectName string
	ModuleName  string
	OpenAPIPath string
	DB          DBType
	UseInfra    []string
}
