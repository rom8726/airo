package config

const (
	DBTypePostgres = "postgresql"
	DBTypeMySQL    = "mysql"
)

type ProjectConfig struct {
	Aborted     bool
	ProjectName string
	ModuleName  string
	OpenAPIPath string
	DB          string
	UseInfra    []string
}
